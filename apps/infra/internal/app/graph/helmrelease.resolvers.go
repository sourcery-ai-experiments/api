package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.39

import (
	"context"
	"github.com/kloudlite/api/pkg/errors"
	"time"

	"github.com/kloudlite/api/apps/infra/internal/app/graph/generated"
	"github.com/kloudlite/api/apps/infra/internal/app/graph/model"
	"github.com/kloudlite/api/apps/infra/internal/entities"
	fn "github.com/kloudlite/api/pkg/functions"
	"github.com/kloudlite/api/pkg/repos"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CreationTime is the resolver for the creationTime field.
func (r *helmReleaseResolver) CreationTime(ctx context.Context, obj *entities.HelmRelease) (string, error) {
	if obj == nil {
		return "", errors.Newf("helmRelease obj is nil")
	}

	return obj.CreationTime.Format(time.RFC3339), nil
}

// ID is the resolver for the id field.
func (r *helmReleaseResolver) ID(ctx context.Context, obj *entities.HelmRelease) (repos.ID, error) {
	if obj == nil {
		return "", errors.Newf("helmRelease obj is nil")
	}

	return obj.Id, nil
}

// Spec is the resolver for the spec field.
func (r *helmReleaseResolver) Spec(ctx context.Context, obj *entities.HelmRelease) (*model.GithubComKloudliteOperatorApisCrdsV1HelmChartSpec, error) {
	if obj == nil {
		return nil, errors.Newf("helmRelease is nil")
	}

	var spec model.GithubComKloudliteOperatorApisCrdsV1HelmChartSpec
	if err := fn.JsonConversion(&obj.Spec, &spec); err != nil {
		return nil, errors.NewE(err)
	}
	return &spec, nil
}

// Status is the resolver for the status field.
func (r *helmReleaseResolver) Status(ctx context.Context, obj *entities.HelmRelease) (*model.GithubComKloudliteOperatorApisCrdsV1HelmChartStatus, error) {
	if obj == nil {
		return nil, errors.Newf("helmRelease is nil")
	}

	var status model.GithubComKloudliteOperatorApisCrdsV1HelmChartStatus
	if err := fn.JsonConversion(&obj.Status, &status); err != nil {
		return nil, errors.NewE(err)
	}
	return &status, nil
}

// UpdateTime is the resolver for the updateTime field.
func (r *helmReleaseResolver) UpdateTime(ctx context.Context, obj *entities.HelmRelease) (string, error) {
	if obj == nil {
		return "", errors.Newf("helmRelease is nil")
	}
	return obj.UpdateTime.Format(time.RFC3339), nil
}

// Metadata is the resolver for the metadata field.
func (r *helmReleaseInResolver) Metadata(ctx context.Context, obj *entities.HelmRelease, data *v1.ObjectMeta) error {
	if obj == nil {
		return errors.Newf("helmRelease is nil")
	}

	return fn.JsonConversion(data, &obj.ObjectMeta)
}

// Spec is the resolver for the spec field.
func (r *helmReleaseInResolver) Spec(ctx context.Context, obj *entities.HelmRelease, data *model.GithubComKloudliteOperatorApisCrdsV1HelmChartSpecIn) error {
	if obj == nil {
		return errors.Newf("helmRelease is nil")
	}

	return fn.JsonConversion(data, &obj.Spec)
}

// HelmRelease returns generated.HelmReleaseResolver implementation.
func (r *Resolver) HelmRelease() generated.HelmReleaseResolver { return &helmReleaseResolver{r} }

// HelmReleaseIn returns generated.HelmReleaseInResolver implementation.
func (r *Resolver) HelmReleaseIn() generated.HelmReleaseInResolver { return &helmReleaseInResolver{r} }

type helmReleaseResolver struct{ *Resolver }
type helmReleaseInResolver struct{ *Resolver }
