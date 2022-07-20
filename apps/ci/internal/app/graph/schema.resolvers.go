package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"errors"

	"kloudlite.io/apps/ci/internal/app/graph/generated"
	"kloudlite.io/apps/ci/internal/app/graph/model"
	"kloudlite.io/apps/ci/internal/domain"
	"kloudlite.io/common"
	fn "kloudlite.io/pkg/functions"
	httpServer "kloudlite.io/pkg/http-server"
	"kloudlite.io/pkg/repos"
	"kloudlite.io/pkg/types"
)

func (r *appResolver) Pipelines(ctx context.Context, obj *model.App) ([]*model.GitPipeline, error) {
	pipelines, err := r.Domain.GetAppPipelines(ctx, obj.ID)
	if err != nil {
		return nil, err
	}
	var res []*model.GitPipeline
	for _, pipeline := range pipelines {
		res = append(
			res, &model.GitPipeline{
				ID:          pipeline.Id,
				Name:        pipeline.Name,
				GitProvider: pipeline.GitProvider,
				GitRepoURL:  pipeline.GitRepoUrl,
				GitBranch:   pipeline.GitBranch,
				Build: func() *model.GitPipelineBuild {
					if pipeline.Build == nil {
						return nil
					}
					return &model.GitPipelineBuild{
						BaseImage: &pipeline.Build.BaseImage,
						Cmd:       pipeline.Build.Cmd,
					}
				}(),
				Run: func() *model.GitPipelineRun {
					if pipeline.Run == nil {
						return nil
					}
					return &model.GitPipelineRun{
						BaseImage: &pipeline.Run.BaseImage,
						Cmd:       pipeline.Run.Cmd,
					}
				}(),
				Metadata: pipeline.Metadata,
			},
		)
	}
	return res, nil
}

func (r *appResolver) CiCreateDockerPipeLine(ctx context.Context, obj *model.App, containerName string, in model.GitDockerPipelineIn) (map[string]interface{}, error) {
	session := httpServer.GetSession[*common.AuthSession](ctx)
	if session == nil {
		return nil, errors.New("not authorized")
	}
	var pipeline, err = r.Domain.CreatePipeline(
		ctx, session.UserId, domain.Pipeline{
			Name:          in.Name,
			ProjectName:   in.ProjectName,
			ProjectId:     in.ProjectID,
			AccountId:     in.AccountID,
			AppId:         string(obj.ID),
			ContainerName: containerName,
			GitProvider:   in.GitProvider,
			GitRepoUrl:    in.GitRepoURL,
			GitBranch:     in.GitBranch,
			DockerBuildInput: &domain.DockerBuildInput{
				DockerFile: in.DockerFile,
				ContextDir: in.ContextDir,
				BuildArgs:  in.BuildArgs,
			},
			ArtifactRef: domain.ArtifactRef{
				DockerImageName: fn.DefaultIfNil(in.ArtifactRef.DockerImageName),
				DockerImageTag:  fn.DefaultIfNil(in.ArtifactRef.DockerImageTag),
			},
		},
	)
	if err != nil {
		return nil, err
	}
	marshal, err := json.Marshal(pipeline)
	if err != nil {
		return nil, err
	}
	x := make(map[string]any)
	err = json.Unmarshal(marshal, &x)
	if err != nil {
		return nil, err
	}
	return x, err
}

func (r *appResolver) CiCreatePipeLine(ctx context.Context, obj *model.App, containerName string, in model.GitPipelineIn) (map[string]interface{}, error) {
	session := httpServer.GetSession[*common.AuthSession](ctx)
	if session == nil {
		return nil, errors.New("not authorized")
	}
	var pipeline, err = r.Domain.CreatePipeline(
		ctx, session.UserId, domain.Pipeline{
			Name:          in.Name,
			ProjectId:     in.ProjectID,
			AppId:         string(obj.ID),
			ContainerName: containerName,
			GitProvider:   in.GitProvider,
			GitRepoUrl:    in.GitRepoURL,
			GitBranch:     in.GitBranch,
			Build: &domain.ContainerImageBuild{
				BaseImage: in.Build.BaseImage,
				Cmd:       in.Build.Cmd,
			},
			Run: &domain.ContainerImageRun{
				BaseImage: fn.DefaultIfNil(in.Run.BaseImage),
				Cmd:       in.Run.Cmd,
			},
			ArtifactRef: domain.ArtifactRef{
				DockerImageName: fn.DefaultIfNil(in.ArtifactRef.DockerImageName),
				DockerImageTag:  fn.DefaultIfNil(in.ArtifactRef.DockerImageTag),
			},
		},
	)
	if err != nil {
		return nil, err
	}
	marshal, err := json.Marshal(pipeline)
	if err != nil {
		return nil, err
	}
	x := make(map[string]any)
	err = json.Unmarshal(marshal, &x)
	if err != nil {
		return nil, err
	}
	return x, err
}

func (r *mutationResolver) CiDeletePipeline(ctx context.Context, pipelineID repos.ID) (bool, error) {
	session := httpServer.GetSession[*common.AuthSession](ctx)
	if err := r.Domain.DeletePipeline(ctx, session.UserId, pipelineID); err != nil {
		return false, err
	}
	return true, nil
}

func (r *mutationResolver) CiCreatePipeline(ctx context.Context, in model.GitPipelineIn) (map[string]interface{}, error) {
	session := httpServer.GetSession[*common.AuthSession](ctx)
	if session == nil {
		return nil, errors.New("not authorized")
	}
	var pipeline, err = r.Domain.CreatePipeline(
		ctx, session.UserId, domain.Pipeline{
			Name:        in.Name,
			ProjectName: in.ProjectName,
			AccountId:   in.AccountID,
			ProjectId:   in.ProjectID,
			AppId:       in.AppID,
			GitProvider: in.GitProvider,
			GitRepoUrl:  in.GitRepoURL,
			GitBranch:   in.GitBranch,
			Build: &domain.ContainerImageBuild{
				BaseImage: in.Build.BaseImage,
				Cmd:       in.Build.Cmd,
				OutputDir: fn.DefaultIfNil(in.Build.OutputDir),
			},
			Run: &domain.ContainerImageRun{
				BaseImage: fn.DefaultIfNil(in.Run.BaseImage),
				Cmd:       in.Run.Cmd,
			},
			ArtifactRef: domain.ArtifactRef{
				DockerImageName: fn.DefaultIfNil(in.ArtifactRef.DockerImageName),
				DockerImageTag:  fn.DefaultIfNil(in.ArtifactRef.DockerImageTag),
			},
		},
	)
	if err != nil {
		return nil, err
	}
	marshal, err := json.Marshal(pipeline)
	if err != nil {
		return nil, err
	}
	x := make(map[string]any)
	err = json.Unmarshal(marshal, &x)
	if err != nil {
		return nil, err
	}
	return x, err
}

func (r *queryResolver) CiGithubInstallations(ctx context.Context, pagination *types.Pagination) (interface{}, error) {
	session := httpServer.GetSession[*common.AuthSession](ctx)
	if session == nil {
		return nil, errors.New("not authenticated")
	}
	return r.Domain.GithubListInstallations(ctx, session.UserId, pagination)
}

func (r *queryResolver) CiGithubInstallationToken(ctx context.Context, repoURL string) (interface{}, error) {
	return r.Domain.GithubInstallationToken(ctx, repoURL)
}

func (r *queryResolver) CiGithubRepos(ctx context.Context, installationID int, pagination *types.Pagination) (interface{}, error) {
	session := httpServer.GetSession[*common.AuthSession](ctx)
	if session == nil {
		return nil, errors.New("not authenticated")
	}
	return r.Domain.GithubListRepos(ctx, session.UserId, int64(installationID), pagination)
}

func (r *queryResolver) CiGithubRepoBranches(ctx context.Context, repoURL string, pagination *types.Pagination) (interface{}, error) {
	session := httpServer.GetSession[*common.AuthSession](ctx)
	if session == nil {
		return nil, errors.New("not authenticated")
	}
	branches, err := r.Domain.GithubListBranches(ctx, session.UserId, repoURL, pagination)
	return branches, err
}

func (r *queryResolver) CiSearchGithubRepos(ctx context.Context, search *string, org string, pagination *types.Pagination) (interface{}, error) {
	session := httpServer.GetSession[*common.AuthSession](ctx)
	if session == nil {
		return nil, errors.New("not authenticated")
	}
	return r.Domain.GithubSearchRepos(ctx, session.UserId, *search, org, pagination)
}

func (r *queryResolver) CiGitlabGroups(ctx context.Context, query *string, pagination *types.Pagination) (interface{}, error) {
	session := httpServer.GetSession[*common.AuthSession](ctx)
	return r.Domain.GitlabListGroups(ctx, session.UserId, query, pagination)
}

func (r *queryResolver) CiGitlabRepos(ctx context.Context, groupID string, search *string, pagination *types.Pagination) (interface{}, error) {
	session := httpServer.GetSession[*common.AuthSession](ctx)
	return r.Domain.GitlabListRepos(ctx, session.UserId, groupID, search, pagination)
}

func (r *queryResolver) CiGitlabRepoBranches(ctx context.Context, repoID string, search *string, pagination *types.Pagination) (interface{}, error) {
	session := httpServer.GetSession[*common.AuthSession](ctx)
	return r.Domain.GitlabListBranches(ctx, session.UserId, repoID, search, pagination)
}

func (r *queryResolver) CiGetPipelines(ctx context.Context, projectID repos.ID) ([]*model.GitPipeline, error) {
	pipelineEntities, err := r.Domain.GetPipelines(ctx, projectID)
	if err != nil {
		return nil, err
	}
	pipelines := make([]*model.GitPipeline, len(pipelineEntities))
	for i, pItem := range pipelineEntities {
		pipelines[i] = &model.GitPipeline{
			ID:          pItem.Id,
			Name:        pItem.Name,
			GitProvider: pItem.GitProvider,
			GitRepoURL:  pItem.GitRepoUrl,
			GitBranch:   pItem.GitBranch,
			Metadata:    pItem.Metadata,
		}

		if pItem.Build != nil {
			pipelines[i].Build = &model.GitPipelineBuild{
				BaseImage: &pItem.Build.BaseImage,
				Cmd:       pItem.Build.Cmd,
			}
		}

		if pItem.Run != nil {
			pipelines[i].Run = &model.GitPipelineRun{
				BaseImage: &pItem.Run.BaseImage,
				Cmd:       pItem.Run.Cmd,
			}
		}
	}
	return pipelines, nil
}

func (r *queryResolver) CiGetPipeline(ctx context.Context, pipelineID repos.ID) (*model.GitPipeline, error) {
	pipeline, err := r.Domain.GetPipeline(ctx, pipelineID)
	if err != nil {
		return nil, err
	}

	pRecord := model.GitPipeline{
		ID:          pipeline.Id,
		Name:        pipeline.Name,
		GitProvider: pipeline.GitProvider,
		GitRepoURL:  pipeline.GitRepoUrl,
		Build: &model.GitPipelineBuild{
			BaseImage: &pipeline.Build.BaseImage,
			Cmd:       pipeline.Build.Cmd,
		},
		Run: &model.GitPipelineRun{
			BaseImage: &pipeline.Run.BaseImage,
			Cmd:       pipeline.Run.Cmd,
		},
		Metadata: pipeline.Metadata,
	}

	if pipeline.Build != nil {
		pRecord.Build = &model.GitPipelineBuild{
			BaseImage: &pipeline.Build.BaseImage,
			Cmd:       pRecord.Build.Cmd,
		}
	}

	if pipeline.Run != nil {
		pRecord.Run = &model.GitPipelineRun{
			BaseImage: &pipeline.Run.BaseImage,
			Cmd:       pipeline.Run.Cmd,
		}
	}

	return &pRecord, nil
}

func (r *queryResolver) CiTriggerPipeline(ctx context.Context, pipelineID repos.ID) (*bool, error) {
	session := httpServer.GetSession[*common.AuthSession](ctx)
	if err := r.Domain.TriggerPipeline(ctx, session.UserId, pipelineID); err != nil {
		return fn.New(false), err
	}
	return fn.New(true), nil
}

// App returns generated.AppResolver implementation.
func (r *Resolver) App() generated.AppResolver { return &appResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type appResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
