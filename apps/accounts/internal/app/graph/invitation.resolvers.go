package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.28

import (
	"context"
	"fmt"
	"time"

	"kloudlite.io/apps/accounts/internal/app/graph/generated"
	"kloudlite.io/apps/accounts/internal/entities"
)

// CreationTime is the resolver for the creationTime field.
func (r *invitationResolver) CreationTime(ctx context.Context, obj *entities.Invitation) (string, error) {
	if obj == nil {
		return "", fmt.Errorf("invitation obj is nil")
	}
	return obj.CreationTime.Format(time.RFC3339), nil
}

// ID is the resolver for the id field.
func (r *invitationResolver) ID(ctx context.Context, obj *entities.Invitation) (string, error) {
	if obj == nil {
		return "", fmt.Errorf("invitation obj is nil")
	}
	return string(obj.Id), nil
}

// UpdateTime is the resolver for the updateTime field.
func (r *invitationResolver) UpdateTime(ctx context.Context, obj *entities.Invitation) (string, error) {
	if obj == nil {
		return "", fmt.Errorf("invitation obj is nil")
	}
	return obj.UpdateTime.Format(time.RFC3339), nil
}

// Invitation returns generated.InvitationResolver implementation.
func (r *Resolver) Invitation() generated.InvitationResolver { return &invitationResolver{r} }

type invitationResolver struct{ *Resolver }
