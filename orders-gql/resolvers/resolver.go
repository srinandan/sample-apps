// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
