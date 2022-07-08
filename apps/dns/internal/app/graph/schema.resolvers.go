package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"kloudlite.io/apps/dns/internal/app/graph/generated"
	"kloudlite.io/apps/dns/internal/app/graph/model"
	"kloudlite.io/pkg/repos"
)

func (r *accountResolver) Sites(ctx context.Context, obj *model.Account) ([]*model.Site, error) {
	sitesEntities, err := r.d.GetSites(ctx, string(obj.ID))
	if err != nil {
		return nil, err
	}
	sites := make([]*model.Site, 0)
	for _, e := range sitesEntities {
		sites = append(sites, &model.Site{
			ID:         e.Id,
			AccountID:  e.AccountId,
			IsVerified: e.Verified,
			Domain:     e.Domain,
		})
	}
	return sites, nil
}

func (r *accountResolver) EdgeCname(ctx context.Context, obj *model.Account) (string, error) {
	return r.d.GetAccountEdgeCName(ctx, string(obj.ID))
}

func (r *mutationResolver) DNSCreateSite(ctx context.Context, domain string, accountID repos.ID) (bool, error) {
	err := r.d.CreateSite(ctx, domain, repos.ID(accountID))
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *mutationResolver) DNSDeleteSite(ctx context.Context, siteID repos.ID) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DNSVerifySite(ctx context.Context, siteID repos.ID) (bool, error) {
	err := r.d.VerifySite(ctx, siteId)
	return err == nil, err
}

func (r *queryResolver) DNSGetSite(ctx context.Context, siteID repos.ID) (*model.Site, error) {
	site, err := r.d.GetSite(ctx, string(siteID))
	if err != nil {
		return nil, err
	}
	return &model.Site{
		ID:         site.Id,
		AccountID:  site.AccountId,
		IsVerified: site.Verified,
		Domain:     site.Domain,
	}, nil
}

// Account returns generated.AccountResolver implementation.
func (r *Resolver) Account() generated.AccountResolver { return &accountResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type accountResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
