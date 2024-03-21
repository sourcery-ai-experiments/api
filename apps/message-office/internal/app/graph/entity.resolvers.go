package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.39

import (
	"context"

	"github.com/kloudlite/api/apps/message-office/internal/app/graph/generated"
	"github.com/kloudlite/api/apps/message-office/internal/app/graph/model"
)

// FindBYOCClusterByMetadataNameAndSpecAccountName is the resolver for the findBYOCClusterByMetadataNameAndSpecAccountName field.
func (r *entityResolver) FindBYOCClusterByMetadataNameAndSpecAccountName(ctx context.Context, metadataName string, specAccountName string) (*model.BYOCCluster, error) {
	return &model.BYOCCluster{
		Metadata: &model.Metadata{
			Name: metadataName,
		},
		Spec: &model.GithubComKloudliteOperatorApisClustersV1BYOCSpec{
			AccountName: specAccountName,
		},
		ClusterToken: "",
	}, nil
}

// FindClusterByMetadataNameAndSpecAccountName is the resolver for the findClusterByMetadataNameAndSpecAccountName field.
func (r *entityResolver) FindClusterByMetadataNameAndSpecAccountName(ctx context.Context, metadataName string, specAccountName string) (*model.Cluster, error) {
	return &model.Cluster{
		Metadata: &model.Metadata{
			Name: metadataName,
		},
		Spec: &model.GithubComKloudliteOperatorApisClustersV1ClusterSpec{
			AccountName: specAccountName,
		},
		ClusterToken: "",
	}, nil
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
