package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.28

import (
	"context"
	"fmt"

	"github.com/kloudlite/operator/pkg/operator"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"kloudlite.io/apps/container-registry/internal/app/graph/generated"
	"kloudlite.io/pkg/types"
)

// Labels is the resolver for the labels field.
func (r *metadataResolver) Labels(ctx context.Context, obj *v1.ObjectMeta) (map[string]interface{}, error) {
	panic(fmt.Errorf("not implemented: Labels - labels"))
}

// Annotations is the resolver for the annotations field.
func (r *metadataResolver) Annotations(ctx context.Context, obj *v1.ObjectMeta) (map[string]interface{}, error) {
	panic(fmt.Errorf("not implemented: Annotations - annotations"))
}

// CreationTimestamp is the resolver for the creationTimestamp field.
func (r *metadataResolver) CreationTimestamp(ctx context.Context, obj *v1.ObjectMeta) (string, error) {
	panic(fmt.Errorf("not implemented: CreationTimestamp - creationTimestamp"))
}

// DeletionTimestamp is the resolver for the deletionTimestamp field.
func (r *metadataResolver) DeletionTimestamp(ctx context.Context, obj *v1.ObjectMeta) (*string, error) {
	panic(fmt.Errorf("not implemented: DeletionTimestamp - deletionTimestamp"))
}

// Checks is the resolver for the checks field.
func (r *statusResolver) Checks(ctx context.Context, obj *operator.Status) (map[string]interface{}, error) {
	panic(fmt.Errorf("not implemented: Checks - checks"))
}

// DisplayVars is the resolver for the displayVars field.
func (r *statusResolver) DisplayVars(ctx context.Context, obj *operator.Status) (map[string]interface{}, error) {
	panic(fmt.Errorf("not implemented: DisplayVars - displayVars"))
}

// SyncScheduledAt is the resolver for the syncScheduledAt field.
func (r *syncStatusResolver) SyncScheduledAt(ctx context.Context, obj *types.SyncStatus) (string, error) {
	panic(fmt.Errorf("not implemented: SyncScheduledAt - syncScheduledAt"))
}

// LastSyncedAt is the resolver for the lastSyncedAt field.
func (r *syncStatusResolver) LastSyncedAt(ctx context.Context, obj *types.SyncStatus) (*string, error) {
	panic(fmt.Errorf("not implemented: LastSyncedAt - lastSyncedAt"))
}

// Labels is the resolver for the labels field.
func (r *metadataInResolver) Labels(ctx context.Context, obj *v1.ObjectMeta, data map[string]interface{}) error {
	panic(fmt.Errorf("not implemented: Labels - labels"))
}

// Annotations is the resolver for the annotations field.
func (r *metadataInResolver) Annotations(ctx context.Context, obj *v1.ObjectMeta, data map[string]interface{}) error {
	panic(fmt.Errorf("not implemented: Annotations - annotations"))
}

// Metadata returns generated.MetadataResolver implementation.
func (r *Resolver) Metadata() generated.MetadataResolver { return &metadataResolver{r} }

// Status returns generated.StatusResolver implementation.
func (r *Resolver) Status() generated.StatusResolver { return &statusResolver{r} }

// SyncStatus returns generated.SyncStatusResolver implementation.
func (r *Resolver) SyncStatus() generated.SyncStatusResolver { return &syncStatusResolver{r} }

// MetadataIn returns generated.MetadataInResolver implementation.
func (r *Resolver) MetadataIn() generated.MetadataInResolver { return &metadataInResolver{r} }

type metadataResolver struct{ *Resolver }
type statusResolver struct{ *Resolver }
type syncStatusResolver struct{ *Resolver }
type metadataInResolver struct{ *Resolver }
