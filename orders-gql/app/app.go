package app

import (
	"context"
	"github.com/srinandan/sample-apps/orders-gql/models"
)

func GetOrder(ctx context.Context, id string) (*models.Order, error) {
	return &models.Order{
		ID:         "1",
		ShipmentID: getStringPtr("123456"),
		LineItems:  nil,
		Carrier:    getStringPtr("FedEx"),
		TrackingID: getStringPtr("200001"),
	}, nil
}

func ListOrders(ctx context.Context) ([]*models.Order, error) {
	return nil, nil
}

func getStringPtr(str string) *string {
	return &str
}
