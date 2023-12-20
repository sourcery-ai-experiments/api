package domain

import (
	"context"
	"github.com/kloudlite/api/apps/container-registry/internal/domain/entities"
	iamT "github.com/kloudlite/api/apps/iam/types"
	"github.com/kloudlite/api/common"
	"github.com/kloudlite/api/grpc-interfaces/kloudlite.io/rpc/iam"
	"github.com/kloudlite/api/pkg/errors"
	"github.com/kloudlite/api/pkg/repos"
)

func (d *Impl) ListBuildsByCache(ctx RegistryContext, cacheId repos.ID, pagination repos.CursorPagination) (*repos.PaginatedRecord[*entities.Build], error) {
	co, err := d.iamClient.Can(ctx, &iam.CanIn{
		UserId: string(ctx.UserId),
		ResourceRefs: []string{
			iamT.NewResourceRef(ctx.AccountName, iamT.ResourceAccount, ctx.AccountName),
		},
		Action: string(iamT.GetAccount),
	})

	if err != nil {
		return nil, errors.NewE(err)
	}

	if !co.Status {
		return nil, errors.Newf("unauthorized to list builds")
	}

	filter := repos.Filter{"spec.accountName": ctx.AccountName, "spec.cacheKeyName": cacheId}

	return d.buildRepo.FindPaginated(ctx, filter, pagination)
}

func (d *Impl) AddBuild(ctx RegistryContext, build entities.Build) (*entities.Build, error) {
	co, err := d.iamClient.Can(ctx, &iam.CanIn{
		UserId: string(ctx.UserId),
		ResourceRefs: []string{
			iamT.NewResourceRef(ctx.AccountName, iamT.ResourceAccount, ctx.AccountName),
		},
		Action: string(iamT.UpdateAccount),
	})

	if err != nil {
		return nil, errors.NewE(err)
	}

	if !co.Status {
		return nil, errors.Newf("unauthorized to add build")
	}

	if err := validateBuild(build); err != nil {
		return nil, errors.NewE(err)
	}

	var webhookId *int

	if build.Source.Provider == "gitlab" {
		webhookId, err = d.GitlabAddWebhook(ctx, ctx.UserId, d.gitlab.GetRepoId(build.Source.Repository))
		if err != nil {
			return nil, errors.NewE(err)
		}
	}

	build.Spec.AccountName = ctx.AccountName
	return d.buildRepo.Create(ctx, &entities.Build{
		Spec:          build.Spec,
		Name:          build.Name,
		CreatedBy:     common.CreatedOrUpdatedBy{UserId: ctx.UserId, UserName: ctx.UserName, UserEmail: ctx.UserEmail},
		LastUpdatedBy: common.CreatedOrUpdatedBy{},
		Source:        entities.GitSource{Repository: build.Source.Repository, Branch: build.Source.Branch, Provider: build.Source.Provider, WebhookId: webhookId},
		CredUser:      common.CreatedOrUpdatedBy{UserId: ctx.UserId, UserName: ctx.UserName, UserEmail: ctx.UserEmail},
		ErrorMessages: map[string]string{},
		Status:        entities.BuildStatusIdle,
	})
}

func (d *Impl) UpdateBuild(ctx RegistryContext, id repos.ID, build entities.Build) (*entities.Build, error) {
	co, err := d.iamClient.Can(ctx, &iam.CanIn{
		UserId: string(ctx.UserId),
		ResourceRefs: []string{
			iamT.NewResourceRef(ctx.AccountName, iamT.ResourceAccount, ctx.AccountName),
		},
		Action: string(iamT.UpdateAccount),
	})

	if err != nil {
		return nil, errors.NewE(err)
	}

	if !co.Status {
		return nil, errors.Newf("unauthorized to update build")
	}

	if err := validateBuild(build); err != nil {
		return nil, errors.NewE(err)
	}

	return d.buildRepo.UpdateById(ctx, id, &entities.Build{
		Spec:          build.Spec,
		Name:          build.Name,
		CreatedBy:     common.CreatedOrUpdatedBy{},
		LastUpdatedBy: common.CreatedOrUpdatedBy{UserId: ctx.UserId, UserName: ctx.UserName, UserEmail: ctx.UserEmail},
		Source:        build.Source,
		CredUser:      common.CreatedOrUpdatedBy{UserId: ctx.UserId, UserName: ctx.UserName, UserEmail: ctx.UserEmail},
		ErrorMessages: map[string]string{},
		Status:        build.Status,
	})
}

func (d *Impl) UpdateBuildInternal(ctx context.Context, build *entities.Build) (*entities.Build, error) {
	return d.buildRepo.UpdateById(ctx, build.Id, build)
}

func (d *Impl) ListBuildsByGit(ctx context.Context, repoUrl, branch, provider string) ([]*entities.Build, error) {
	filter := repos.Filter{
		"source.repository": repoUrl,
		"source.branch":     branch,
		"source.provider":   provider,
	}

	b, err := d.buildRepo.Find(ctx, repos.Query{
		Filter: filter,
	})
	if err != nil {
		return nil, errors.NewE(err)
	}

	return b, nil
}

func (d *Impl) ListBuilds(ctx RegistryContext, repoName string, search map[string]repos.MatchFilter, pagination repos.CursorPagination) (*repos.PaginatedRecord[*entities.Build], error) {
	co, err := d.iamClient.Can(ctx, &iam.CanIn{
		UserId: string(ctx.UserId),
		ResourceRefs: []string{
			iamT.NewResourceRef(ctx.AccountName, iamT.ResourceAccount, ctx.AccountName),
		},
		Action: string(iamT.GetAccount),
	})

	if err != nil {
		return nil, errors.NewE(err)
	}

	if !co.Status {
		return nil, errors.Newf("unauthorized to list builds")
	}

	filter := repos.Filter{"spec.accountName": ctx.AccountName, "spec.registry.repo.name": repoName}

	return d.buildRepo.FindPaginated(ctx, d.buildRepo.MergeMatchFilters(filter, search), pagination)
}

func (d *Impl) GetBuild(ctx RegistryContext, buildId repos.ID) (*entities.Build, error) {
	co, err := d.iamClient.Can(ctx, &iam.CanIn{
		UserId: string(ctx.UserId),
		ResourceRefs: []string{
			iamT.NewResourceRef(ctx.AccountName, iamT.ResourceAccount, ctx.AccountName),
		},
		Action: string(iamT.GetAccount),
	})

	if err != nil {
		return nil, errors.NewE(err)
	}

	if !co.Status {
		return nil, errors.Newf("unauthorized to get build")
	}

	b, err := d.buildRepo.FindOne(ctx, repos.Filter{"spec.accountName": ctx.AccountName, "id": buildId})
	if err != nil {
		return nil, errors.NewE(err)
	}

	if b == nil {
		return nil, errors.Newf("build not found")
	}

	return b, nil
}

func (d *Impl) DeleteBuild(ctx RegistryContext, buildId repos.ID) error {
	co, err := d.iamClient.Can(ctx, &iam.CanIn{
		UserId: string(ctx.UserId),
		ResourceRefs: []string{
			iamT.NewResourceRef(ctx.AccountName, iamT.ResourceAccount, ctx.AccountName),
		},
		Action: string(iamT.UpdateAccount),
	})

	if err != nil {
		return errors.NewE(err)
	}

	if !co.Status {
		return errors.Newf("unauthorized to delete build")
	}

	b, err := d.buildRepo.FindById(ctx, buildId)
	if err != nil {
		return errors.NewE(err)
	}

	if err = d.buildRepo.DeleteOne(ctx, repos.Filter{"spec.accountName": ctx.AccountName, "id": buildId}); err != nil {
		return errors.NewE(err)
	}

	if b.Source.Provider == "gitlab" {
		at, err := d.getAccessTokenByUserId(ctx, "gitlab", ctx.UserId)
		if err != nil {
			return nil
		}

		d.gitlab.DeleteWebhook(ctx, at, string(b.Source.Repository), entities.GitlabWebhookId(*b.Source.WebhookId))
	}

	return nil
}

func (d *Impl) TriggerBuild(ctx RegistryContext, buildId repos.ID) error {
	panic("implement me")
}
