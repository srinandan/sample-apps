package resolvers

import (
	"context"

	"github.com/srinandan/sample-apps/orders-gql/app"
	"github.com/srinandan/sample-apps/orders-gql/gql"
	"github.com/srinandan/sample-apps/orders-gql/models"
)

type Resolver struct{}

func (r *Resolver) Query() gql.QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) GetOrder(ctx context.Context, id string) (*models.Order, error) {
	order, err := app.GetOrder(ctx, id)
	return order, err
}

func (r *queryResolver) ListOrders(ctx context.Context) ([]*models.Order, error) {
	orders, err := app.ListOrders(ctx)
	return orders, err
}
