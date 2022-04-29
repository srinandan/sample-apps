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

package app

import (
	"context"

	"github.com/srinandan/sample-apps/orders-gql/models"
	"github.com/srinandan/sample-apps/orders-gql/odr"
)

func getStringPtr(str string) *string {
	return &str
}

func getIntPtr(i int) *int {
	return &i
}

func GetOrder(ctx context.Context, id string) (*models.Order, error) {

	o, _ := odr.GetOrder(id)

	order := &models.Order{}

	order.OperationID = o.OperationId
	order.ShipmentID = getStringPtr(o.ShipmentId)
	order.Carrier = getStringPtr(o.Carrier)
	order.TrackingID = getStringPtr(o.TrackingID)

	return order, nil
}

func ListOrders(ctx context.Context) ([]*models.Order, error) {
	orders := []*models.Order{}
	os := odr.ListOrders()

	for _, o := range os {
		order := &models.Order{}

		order.OperationID = o.OperationId
		order.ShipmentID = getStringPtr(o.ShipmentId)
		order.Carrier = getStringPtr(o.Carrier)
		order.TrackingID = getStringPtr(o.TrackingID)

		orders = append(orders, order)
	}
	return orders, nil
}
