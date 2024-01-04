package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.28

import (
	"context"
	"github.com/kloudlite/api/pkg/errors"
	"time"

	"github.com/kloudlite/api/apps/infra/internal/app/graph/generated"
	"github.com/kloudlite/api/apps/infra/internal/app/graph/model"
	"github.com/kloudlite/api/apps/infra/internal/entities"
	fn "github.com/kloudlite/api/pkg/functions"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CreationTime is the resolver for the creationTime field.
func (r *persistentVolumeResolver) CreationTime(ctx context.Context, obj *entities.PersistentVolume) (string, error) {
	if obj == nil {
		return "", errors.Newf("persistent-volume/creation-time is nil")
	}
	return obj.CreationTime.Format(time.RFC3339), nil
}

// ID is the resolver for the id field.
func (r *persistentVolumeResolver) ID(ctx context.Context, obj *entities.PersistentVolume) (string, error) {
	if obj == nil {
		return "", errors.Newf("persitent volume is nil")
	}
	return string(obj.Id), nil
}

// Spec is the resolver for the spec field.
func (r *persistentVolumeResolver) Spec(ctx context.Context, obj *entities.PersistentVolume) (*model.K8sIoAPICoreV1PersistentVolumeSpec, error) {
	var m model.K8sIoAPICoreV1PersistentVolumeSpec
	if err := fn.JsonConversion(obj.Spec, &m); err != nil {
		return nil, errors.NewE(err)
	}
	return &m, nil
}

// Status is the resolver for the status field.
func (r *persistentVolumeResolver) Status(ctx context.Context, obj *entities.PersistentVolume) (*model.K8sIoAPICoreV1PersistentVolumeStatus, error) {
	var m model.K8sIoAPICoreV1PersistentVolumeStatus
	if err := fn.JsonConversion(obj.Status, &m); err != nil {
		return nil, errors.NewE(err)
	}
	return &m, nil
}

// UpdateTime is the resolver for the updateTime field.
func (r *persistentVolumeResolver) UpdateTime(ctx context.Context, obj *entities.PersistentVolume) (string, error) {
	if obj == nil || obj.UpdateTime.IsZero() {
		return "", errors.Newf("persistent-volume/update-time is nil")
	}
	return obj.UpdateTime.Format(time.RFC3339), nil
}

// Metadata is the resolver for the metadata field.
func (r *persistentVolumeInResolver) Metadata(ctx context.Context, obj *entities.PersistentVolume, data *v1.ObjectMeta) error {
	if obj == nil {
		return errors.Newf("persistance volume is nil")
	}
	return fn.JsonConversion(data, &obj.ObjectMeta)
}

// Spec is the resolver for the spec field.
func (r *persistentVolumeInResolver) Spec(ctx context.Context, obj *entities.PersistentVolume, data *model.K8sIoAPICoreV1PersistentVolumeSpecIn) error {
	if obj == nil {
		return errors.Newf("persistance volume is nil")
	}
	return fn.JsonConversion(data, &obj.Spec)
}

// Status is the resolver for the status field.
func (r *persistentVolumeInResolver) Status(ctx context.Context, obj *entities.PersistentVolume, data *model.K8sIoAPICoreV1PersistentVolumeStatusIn) error {
	if obj == nil {
		return errors.Newf("persistance volume is nil")
	}
	return fn.JsonConversion(data, &obj.Status)
}

// PersistentVolume returns generated.PersistentVolumeResolver implementation.
func (r *Resolver) PersistentVolume() generated.PersistentVolumeResolver {
	return &persistentVolumeResolver{r}
}

// PersistentVolumeIn returns generated.PersistentVolumeInResolver implementation.
func (r *Resolver) PersistentVolumeIn() generated.PersistentVolumeInResolver {
	return &persistentVolumeInResolver{r}
}

type persistentVolumeResolver struct{ *Resolver }
type persistentVolumeInResolver struct{ *Resolver }
