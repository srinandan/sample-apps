package app

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/srinandan/sample-apps/common"
	"github.com/srinandan/sample-apps/common/types"
	"github.com/srinandan/sample-apps/orders-gql/models"

	"io/ioutil"
	"net/http"
	"os"
)

func getOrderEndpoint() string {
	//endpoint to reach the order service
	var orderHostname = os.Getenv("ORDERS")

	if orderHostname == "" {
		return "http://orders.apps.svc.cluster.local:8080"
	}

	return orderHostname
}

func callOrderEndpoint(uri string) (respBody []byte, err error) {
	var req *http.Request

	orderURL := getOrderEndpoint() + uri

	client := &http.Client{}
	common.Info.Println("Connecting to: ", orderURL)

	req, err = http.NewRequest("GET", orderURL, nil)
	if err != nil {
		common.Error.Println("error in client: ", err)
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		common.Error.Println("error connecting: ", err)
		return nil, err
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	respBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		common.Error.Println("error in response: ", err)
		return nil, err
	} else if resp.StatusCode > 299 {
		common.Error.Println("error in response: ", string(respBody))
		return nil, fmt.Errorf("error in response")
	}

	return respBody, nil
}

func getStringPtr(str string) *string {
	return &str
}

func getIntPtr(i int) *int {
	return &i
}

func GetOrder(ctx context.Context, id string) (*models.Order, error) {
	
	respBody, err := callOrderEndpoint("/orders/" + id)
	if err != nil {
		return nil, err		
	}
	
	order := &models.Order{}

	if err = json.Unmarshal(respBody, order); err != nil {
		common.Error.Println("error unmarshaling: ", err)
		return nil, err
	}

	respBody, err = callOrderEndpoint("/orders/" + id + "/items")
	if err != nil {
		return nil, err		
	}

	items := []types.Item{}
	if err = json.Unmarshal(respBody, &items); err != nil {
		common.Error.Println("error unmarshaling: ", err)
		return nil, err
	}
	
	for i, item := range items {
		lineItem := &models.LineItem{}
		lineItem.Item = &models.Item{}
		lineItem.Item.Inventory = &models.Inventory{}
		lineItem.Item.Inventory.SalePrice = &models.Price{}
		lineItem.Item.Inventory.SalePrice.Currency = getStringPtr(item.Inventory.SalePrice.Currency) 
		lineItem.Item.Inventory.SalePrice.Value = getStringPtr(item.Inventory.SalePrice.Value)
		lineItem.Item.Inventory.SellOnGoogleQuantity = getStringPtr(item.Inventory.SellOnGoogleQuantity)
		lineItem.Item.Inventory.Availability = getIntPtr(item.Inventory.Availability)
		lineItem.Item.Inventory.Kind = getStringPtr(item.Inventory.Kind)
		lineItem.Item.Inventory.Price = &models.Price{}
		lineItem.Item.Inventory.Price.Currency = getStringPtr(item.Inventory.Price.Currency) 
		lineItem.Item.Inventory.Price.Value = getStringPtr(item.Inventory.Price.Value)
		lineItem.Item.ProductID = getStringPtr(item.ProductId)
		lineItem.Item.StoreCode = getStringPtr(item.StoreCode)
		lineItem.Quantity = order.LineItems[i].Quantity
		order.LineItems[i] = lineItem
	}

	return order, nil
}

func ListOrders(ctx context.Context) ([]*models.Order, error) {
	return nil, nil
}