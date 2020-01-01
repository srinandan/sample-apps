package resolvers

import (
	"context"

	"github.com/srinandan/sample-apps/orders-gql/gql"
	"github.com/srinandan/sample-apps/orders-gql/models"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Query() gql.QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) GetOrder(ctx context.Context, id string) (*models.Order, error) {
	panic("not implemented")
}
func (r *queryResolver) ListOrders(ctx context.Context) ([]*models.Order, error) {
	panic("not implemented")
}
