package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.39

import (
	"context"
	"github.com/kloudlite/api/pkg/errors"
	"time"

	"github.com/kloudlite/api/apps/console/internal/app/graph/generated"
	"github.com/kloudlite/api/apps/console/internal/app/graph/model"
	"github.com/kloudlite/api/apps/console/internal/entities"
	fn "github.com/kloudlite/api/pkg/functions"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CreationTime is the resolver for the creationTime field.
func (r *routerResolver) CreationTime(ctx context.Context, obj *entities.Router) (string, error) {
	if obj == nil {
		return "", errNilRouter
	}
	return obj.BaseEntity.CreationTime.Format(time.RFC3339), nil
}

// Spec is the resolver for the spec field.
func (r *routerResolver) Spec(ctx context.Context, obj *entities.Router) (*model.GithubComKloudliteOperatorApisCrdsV1RouterSpec, error) {
	if obj == nil {
		return nil, errNilRouter
	}

	m := &model.GithubComKloudliteOperatorApisCrdsV1RouterSpec{}
	if err := fn.JsonConversion(obj.Spec, &m); err != nil {
		return nil, errors.NewE(err)
	}
	return m, nil
}

// UpdateTime is the resolver for the updateTime field.
func (r *routerResolver) UpdateTime(ctx context.Context, obj *entities.Router) (string, error) {
	if obj == nil {
		return "", errNilRouter
	}
	return obj.BaseEntity.UpdateTime.Format(time.RFC3339), nil
}

// Metadata is the resolver for the metadata field.
func (r *routerInResolver) Metadata(ctx context.Context, obj *entities.Router, data *v1.ObjectMeta) error {
	if obj == nil {
		return errNilRouter
	}

	if data != nil {
		obj.ObjectMeta = *data
	}

	return nil
}

// Spec is the resolver for the spec field.
func (r *routerInResolver) Spec(ctx context.Context, obj *entities.Router, data *model.GithubComKloudliteOperatorApisCrdsV1RouterSpecIn) error {
	if obj == nil {
		return errNilRouter
	}
	return fn.JsonConversion(data, &obj.Spec)
}

// Router returns generated.RouterResolver implementation.
func (r *Resolver) Router() generated.RouterResolver { return &routerResolver{r} }

// RouterIn returns generated.RouterInResolver implementation.
func (r *Resolver) RouterIn() generated.RouterInResolver { return &routerInResolver{r} }

type routerResolver struct{ *Resolver }
type routerInResolver struct{ *Resolver }
