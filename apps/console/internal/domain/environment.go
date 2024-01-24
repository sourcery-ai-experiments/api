package domain

import (
	"fmt"
	"github.com/kloudlite/api/common/fields"
	crdsv1 "github.com/kloudlite/operator/apis/crds/v1"

	iamT "github.com/kloudlite/api/apps/iam/types"
	"github.com/kloudlite/api/common"
	"github.com/kloudlite/api/grpc-interfaces/kloudlite.io/rpc/iam"
	"github.com/kloudlite/api/pkg/errors"
	"github.com/kloudlite/api/pkg/kv"
	"github.com/kloudlite/operator/operators/resource-watcher/types"

	"github.com/kloudlite/api/constants"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/kloudlite/api/apps/console/internal/entities"
	fc "github.com/kloudlite/api/apps/console/internal/entities/field-constants"
	"github.com/kloudlite/api/pkg/repos"
	t "github.com/kloudlite/api/pkg/types"
)

func (d *domain) findEnvironment(ctx ConsoleContext, projectName string, name string) (*entities.Environment, error) {
	env, err := d.environmentRepo.FindOne(ctx, repos.Filter{
		fields.AccountName:  ctx.AccountName,
		fields.ProjectName:  projectName,
		fields.MetadataName: name,
	})
	if err != nil {
		return nil, errors.NewE(err)
	}
	if env == nil {
		return nil, errors.Newf("no environment with name (%s) and project (%s)", name, projectName)
	}
	return env, nil
}

func (d *domain) envTargetNamespace(ctx ConsoleContext, projectName string, envName string) (string, error) {
	key := fmt.Sprintf("environment-namespace.%s/%s/%s", ctx.AccountName, projectName, envName)
	b, err := d.consoleCacheStore.Get(ctx, key)
	if err != nil {
		if errors.Is(err, kv.ErrKeyNotFound) {
			env, err := d.findEnvironment(ctx, projectName, envName)
			if err != nil {
				return "", err
			}
			defer func() {
				if err := d.consoleCacheStore.Set(ctx, key, []byte(env.Spec.TargetNamespace)); err != nil {
					d.logger.Errorf(err, "while caching environment target namespace")
				}
			}()
			return env.Spec.TargetNamespace, nil
		}
	}

	return string(b), nil
}

func (d *domain) GetEnvironment(ctx ConsoleContext, projectName string, name string) (*entities.Environment, error) {
	if err := d.canReadResourcesInProject(ctx, projectName); err != nil {
		return nil, errors.NewE(err)
	}

	return d.findEnvironment(ctx, projectName, name)
}

func (d *domain) ListEnvironments(ctx ConsoleContext, projectName string, search map[string]repos.MatchFilter, pq repos.CursorPagination) (*repos.PaginatedRecord[*entities.Environment], error) {
	if err := d.canReadResourcesInProject(ctx, projectName); err != nil {
		return nil, errors.NewE(err)
	}

	filter := repos.Filter{
		fields.AccountName:            ctx.AccountName,
		fc.EnvironmentSpecProjectName: projectName,
	}

	return d.environmentRepo.FindPaginated(ctx, d.environmentRepo.MergeMatchFilters(filter, search), pq)
}

func (d *domain) findEnvironmentByTargetNs(ctx ConsoleContext, targetNs string) (*entities.Environment, error) {
	w, err := d.environmentRepo.FindOne(ctx, repos.Filter{
		fields.AccountName:                ctx.AccountName,
		fc.EnvironmentSpecTargetNamespace: targetNs,
	})
	if err != nil {
		return nil, errors.NewE(err)
	}

	if w == nil {
		return nil, errors.Newf("no workspace found for target namespace %q", targetNs)
	}

	return w, nil
}

func (d *domain) CreateEnvironment(ctx ConsoleContext, projectName string, env entities.Environment) (*entities.Environment, error) {
	project, err := d.findProject(ctx, projectName)
	if err != nil {
		return nil, errors.NewE(err)
	}

	if err := d.canMutateResourcesInProject(ctx, project.Name); err != nil {
		return nil, errors.NewE(err)
	}

	env.ProjectName = project.Name
	env.Namespace = project.Spec.TargetNamespace

	env.EnsureGVK()
	if err := d.k8sClient.ValidateObject(ctx, &env.Environment); err != nil {
		return nil, errors.NewE(err)
	}

	env.IncrementRecordVersion()

	if env.Spec.TargetNamespace == "" {
		env.Spec.TargetNamespace = fmt.Sprintf("env-%s", env.Name)
	}

	if env.Spec.Routing == nil {
		env.Spec.Routing = &crdsv1.EnvironmentRouting{
			Mode: crdsv1.EnvironmentRoutingModePrivate,
		}
	}

	env.CreatedBy = common.CreatedOrUpdatedBy{
		UserId:    ctx.UserId,
		UserName:  ctx.UserName,
		UserEmail: ctx.UserEmail,
	}
	env.LastUpdatedBy = env.CreatedBy

	env.AccountName = ctx.AccountName
	env.SyncStatus = t.GenSyncStatus(t.SyncActionApply, env.RecordVersion)

	if _, err := d.upsertEnvironmentResourceMapping(ResourceContext{ConsoleContext: ctx, ProjectName: env.ProjectName, EnvironmentName: env.Name}, &env); err != nil {
		return nil, errors.NewE(err)
	}

	nenv, err := d.environmentRepo.Create(ctx, &env)
	if err != nil {
		if d.environmentRepo.ErrAlreadyExists(err) {
			// TODO: better insights into error, when it is being caused by duplicated indexes
			return nil, errors.NewE(err)
		}
		return nil, errors.NewE(err)
	}

	d.resourceEventPublisher.PublishConsoleEvent(ctx, entities.ResourceTypeEnvironment, nenv.Name, PublishAdd)

	if _, err := d.iamClient.AddMembership(ctx, &iam.AddMembershipIn{
		UserId:       string(ctx.UserId),
		ResourceType: string(iamT.ResourceEnvironment),
		ResourceRef:  iamT.NewResourceRef(ctx.AccountName, iamT.ResourceEnvironment, nenv.Name),
		Role:         string(iamT.RoleResourceOwner),
	}); err != nil {
		d.logger.Errorf(err, "error while adding membership")
	}

	if err := d.applyK8sResource(ctx, nenv.ProjectName, &nenv.Environment, nenv.RecordVersion); err != nil {
		return nil, errors.NewE(err)
	}

	return nenv, nil
}

func (d *domain) CloneEnvironment(ctx ConsoleContext, projectName string, sourceEnvName string, destinationEnvName string, displayName string, environmentRoutingMode crdsv1.EnvironmentRoutingMode) (*entities.Environment, error) {
	if err := d.canMutateResourcesInProject(ctx, projectName); err != nil {
		return nil, errors.NewE(err)
	}

	env, err := d.findEnvironment(ctx, projectName, sourceEnvName)
	if err != nil {
		return nil, errors.NewE(err)
	}

	env.ProjectName = projectName
	env.DisplayName = destinationEnvName
	env.Name = destinationEnvName
	env.Id = d.environmentRepo.NewId()
	env.PrimitiveId = ""
	env.DisplayName = displayName

	if env.Spec.Routing == nil {
		env.Spec.Routing = &crdsv1.EnvironmentRouting{}
	}
	env.Spec.Routing.Mode = environmentRoutingMode

	env.EnsureGVK()
	if err := d.k8sClient.ValidateObject(ctx, &env.Environment); err != nil {
		return nil, errors.NewE(err)
	}

	env.IncrementRecordVersion()

	env.Spec.TargetNamespace = fmt.Sprintf("env-%s", destinationEnvName)

	env.SyncStatus = t.GenSyncStatus(t.SyncActionApply, env.RecordVersion)

	if _, err := d.upsertEnvironmentResourceMapping(ResourceContext{ConsoleContext: ctx, ProjectName: env.ProjectName, EnvironmentName: env.Name}, env); err != nil {
		return nil, errors.NewE(err)
	}

	nenv, err := d.environmentRepo.Create(ctx, env)
	if err != nil {
		if d.environmentRepo.ErrAlreadyExists(err) {
			// TODO: better insights into error, when it is being caused by duplicated indexes
			return nil, errors.NewE(err)
		}
		return nil, errors.NewE(err)
	}
	resContext := ResourceContext{
		ConsoleContext:  ctx,
		ProjectName:     env.ProjectName,
		EnvironmentName: nenv.Name,
	}
	filters := repos.Filter{
		fields.AccountName:            resContext.AccountName,
		fc.EnvironmentSpecProjectName: resContext.ProjectName,
		fields.EnvironmentName:        sourceEnvName,
	}

	apps, err := d.appRepo.Find(ctx, repos.Query{
		Filter: filters,
		Sort:   nil,
	})
	if err != nil {
		return nil, errors.NewE(err)
	}

	secrets, err := d.secretRepo.Find(ctx, repos.Query{
		Filter: filters,
		Sort:   nil,
	})
	if err != nil {
		return nil, errors.NewE(err)
	}
	configs, err := d.configRepo.Find(ctx, repos.Query{
		Filter: filters,
		Sort:   nil,
	})
	if err != nil {
		return nil, errors.NewE(err)
	}
	routers, err := d.routerRepo.Find(ctx, repos.Query{
		Filter: filters,
		Sort:   nil,
	})
	if err != nil {
		return nil, errors.NewE(err)
	}
	managedResources, err := d.mresRepo.Find(ctx, repos.Query{
		Filter: filters,
		Sort:   nil,
	})
	if err != nil {
		return nil, errors.NewE(err)
	}

	for _, app := range apps {
		app.Namespace = env.Spec.TargetNamespace
		app.EnvironmentName = destinationEnvName
		app.Id = d.appRepo.NewId()
		app.PrimitiveId = ""
		app.Spec.Intercept = nil
		err := cloneResource(resContext, d, d.appRepo, app, &app.App)
		if err != nil {
			return nil, errors.NewE(err)
		}
	}
	for _, secret := range secrets {
		secret.Namespace = env.Spec.TargetNamespace
		secret.EnvironmentName = destinationEnvName
		secret.Id = d.secretRepo.NewId()
		secret.PrimitiveId = ""
		err := cloneResource(resContext, d, d.secretRepo, secret, &secret.Secret)
		if err != nil {
			return nil, errors.NewE(err)
		}
	}
	for _, config := range configs {
		config.Namespace = env.Spec.TargetNamespace
		config.EnvironmentName = destinationEnvName
		config.Id = d.configRepo.NewId()
		config.PrimitiveId = ""
		err := cloneResource(resContext, d, d.configRepo, config, &config.ConfigMap)
		if err != nil {
			return nil, errors.NewE(err)
		}
	}
	for _, router := range routers {
		router.Namespace = env.Spec.TargetNamespace
		router.EnvironmentName = destinationEnvName
		router.Id = d.routerRepo.NewId()
		router.PrimitiveId = ""
		err := cloneResource(resContext, d, d.routerRepo, router, &router.Router)
		if err != nil {
			return nil, errors.NewE(err)
		}
	}
	for _, managedResource := range managedResources {
		managedResource.Namespace = env.Spec.TargetNamespace
		managedResource.EnvironmentName = destinationEnvName
		managedResource.Id = d.mresRepo.NewId()
		managedResource.PrimitiveId = ""
		managedResource.Spec.ResourceName = fmt.Sprintf("env-%s-%s", destinationEnvName, managedResource.Name)
		err := cloneResource(resContext, d, d.mresRepo, managedResource, &managedResource.ManagedResource)
		if err != nil {
			return nil, errors.NewE(err)
		}
	}

	if _, err := d.iamClient.AddMembership(ctx, &iam.AddMembershipIn{
		UserId:       string(ctx.UserId),
		ResourceType: string(iamT.ResourceEnvironment),
		ResourceRef:  iamT.NewResourceRef(ctx.AccountName, iamT.ResourceEnvironment, nenv.Name),
		Role:         string(iamT.RoleResourceOwner),
	}); err != nil {
		d.logger.Errorf(err, "error while adding membership")
	}

	if err := d.applyK8sResource(ctx, env.ProjectName, &nenv.Environment, nenv.RecordVersion); err != nil {
		return nil, errors.NewE(err)
	}

	return nenv, nil
}

func (d *domain) UpdateEnvironment(ctx ConsoleContext, projectName string, env entities.Environment) (*entities.Environment, error) {
	if err := d.canMutateResourcesInProject(ctx, projectName); err != nil {
		return nil, errors.NewE(err)
	}

	env.Namespace = "trest"

	env.EnsureGVK()
	if err := d.k8sClient.ValidateObject(ctx, &env.Environment); err != nil {
		return nil, errors.NewE(err)
	}

	patchForUpdate := common.PatchForUpdate(
		ctx,
		&env,
		common.PatchOpts{
			XPatch: repos.Document{
				fc.EnvironmentSpec: env.Spec,
			},
		},
	)

	upEnv, err := d.environmentRepo.Patch(
		ctx,
		repos.Filter{
			fields.AccountName:  ctx.AccountName,
			fields.MetadataName: env.Name,
		},
		patchForUpdate,
	)
	if err != nil {
		return nil, errors.NewE(err)
	}
	d.resourceEventPublisher.PublishConsoleEvent(ctx, entities.ResourceTypeEnvironment, upEnv.Name, PublishUpdate)

	if err := d.applyK8sResource(ctx, upEnv.ProjectName, &upEnv.Environment, upEnv.RecordVersion); err != nil {
		return nil, errors.NewE(err)
	}

	return upEnv, nil
}

func (d *domain) DeleteEnvironment(ctx ConsoleContext, projectName string, name string) error {
	if err := d.canMutateResourcesInProject(ctx, projectName); err != nil {
		return errors.NewE(err)
	}

	uenv, err := d.environmentRepo.Patch(
		ctx,
		repos.Filter{
			fields.AccountName:  ctx.AccountName,
			fields.MetadataName: name,
		},
		common.PatchForMarkDeletion(),
	)

	if err != nil {
		return errors.NewE(err)
	}

	d.resourceEventPublisher.PublishConsoleEvent(ctx, entities.ResourceTypeEnvironment, uenv.Name, PublishUpdate)

	if err := d.deleteK8sResource(ctx, uenv.ProjectName, &uenv.Environment); err != nil {
		if errors.Is(err, ErrNoClusterAttached) {
			return d.appRepo.DeleteById(ctx, uenv.Id)
		}
		return errors.NewE(err)
	}

	return nil
}

func (d *domain) OnEnvironmentApplyError(ctx ConsoleContext, errMsg, namespace, name string, opts UpdateAndDeleteOpts) error {
	uenv, err := d.environmentRepo.Patch(
		ctx,
		repos.Filter{
			fields.AccountName:  ctx.AccountName,
			fields.MetadataName: name,
		},
		common.PatchForErrorFromAgent(
			errMsg,
			common.PatchOpts{
				MessageTimestamp: opts.MessageTimestamp,
			},
		),
	)

	if err != nil {
		return errors.NewE(err)
	}

	d.resourceEventPublisher.PublishConsoleEvent(ctx, entities.ResourceTypeEnvironment, uenv.Name, PublishDelete)

	return errors.NewE(err)
}

func (d *domain) OnEnvironmentDeleteMessage(ctx ConsoleContext, env entities.Environment) error {
	err := d.environmentRepo.DeleteOne(
		ctx,
		repos.Filter{
			fields.AccountName:  ctx.AccountName,
			fields.MetadataName: env.Name,
		},
	)
	if err != nil {
		return errors.NewE(err)
	}

	if _, err = d.iamClient.RemoveMembership(ctx, &iam.RemoveMembershipIn{
		UserId:      string(ctx.UserId),
		ResourceRef: iamT.NewResourceRef(ctx.AccountName, iamT.ResourceEnvironment, env.Name),
	}); err != nil {
		d.logger.Errorf(err, "error while removing membership")
	}

	d.resourceEventPublisher.PublishConsoleEvent(ctx, entities.ResourceTypeEnvironment, env.Name, PublishDelete)
	return nil
}

func (d *domain) OnEnvironmentUpdateMessage(ctx ConsoleContext, env entities.Environment, status types.ResourceStatus, opts UpdateAndDeleteOpts) error {
	xenv, err := d.findEnvironment(ctx, env.Spec.ProjectName, env.Name)
	if err != nil {
		return errors.NewE(err)
	}

	if xenv == nil {
		return errors.Newf("no environment found")
	}

	recordVersion, err := d.MatchRecordVersion(env.Annotations, xenv.RecordVersion)
	if err != nil {
		return d.resyncK8sResource(ctx, xenv.ProjectName, xenv.SyncStatus.Action, &xenv.Environment, xenv.RecordVersion)
	}

	uenv, err := d.environmentRepo.PatchById(
		ctx,
		xenv.Id,
		common.PatchForSyncFromAgent(
			&env,
			recordVersion,
			status,
			common.PatchOpts{
				MessageTimestamp: opts.MessageTimestamp,
			}))

	if err != nil {
		return err
	}

	d.resourceEventPublisher.PublishConsoleEvent(ctx, entities.ResourceTypeEnvironment, uenv.Name, PublishUpdate)
	return nil
}

func (d *domain) ResyncEnvironment(ctx ConsoleContext, projectName string, name string) error {
	if err := d.canMutateResourcesInProject(ctx, projectName); err != nil {
		return errors.NewE(err)
	}

	e, err := d.findEnvironment(ctx, projectName, name)
	if err != nil {
		return errors.NewE(err)
	}

	if err := d.resyncK8sResource(ctx, e.ProjectName, t.SyncActionApply, &corev1.Namespace{
		TypeMeta: metav1.TypeMeta{APIVersion: "v1", Kind: "Namespace"},
		ObjectMeta: metav1.ObjectMeta{
			Name: e.Spec.TargetNamespace,
			Labels: map[string]string{
				constants.EnvNameKey: e.Name,
			},
		},
	}, e.RecordVersion); err != nil {
		return errors.NewE(err)
	}

	return d.resyncK8sResource(ctx, e.ProjectName, e.SyncStatus.Action, &e.Environment, e.RecordVersion)
}
