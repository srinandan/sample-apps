package odr

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	types "github.com/srinandan/sample-apps/common/types"
)

//Orders
var orders = []types.Order{}

//endpoint to reach the inventory service
var inventoryEndpoint = os.Getenv("INVENTORY")

func ReadOrdersFile() error {
	orderListBytes, err := ioutil.ReadFile("orders.json")
	if err != nil {
		return err
	}

	if err = json.Unmarshal(orderListBytes, &orders); err != nil {
		return err
	}

	return nil
}

func ListOrders() []types.Order {
	return orders
}

func GetOrder(id string) (types.Order, int) {
	order := types.Order{}
	for pos, order := range orders {
		if order.OperationId == id {
			return order, pos
		}
	}
	return order, -1
}

func getId() string {
	rand.Seed(time.Now().UnixNano())
	min := 100003
	max := 199999

	return strconv.Itoa(rand.Intn(max-min+1) + min)
}

func GetOrderItem(id string) (types.Item, error) {

	var req *http.Request

	item := types.Item{}

	client := &http.Client{}

	itemUrl := inventoryEndpoint + "/items/" + id

	req, err := http.NewRequest("GET", itemUrl, nil)
	if err != nil {
		return item, err
	}

	resp, err := client.Do(req)
	defer resp.Body.Close()

	if err != nil {
		return item, err
	}	

	if resp.StatusCode != 200 {
		return item, fmt.Errorf("item was not found")
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return item, err
	}	

	err = json.Unmarshal(respBody, &item)
	if err != nil {
		return item, err
	}

	return item, nil
}

func CreateOrder(order types.Order) (types.Order, error) {
	order.OperationId = getId()

	for _, lineItem := range order.LineItems {
		_, err := GetOrderItem(lineItem.LineItemId)
		if err != nil {
			return order, err
		}
	}
	orders = append(orders, order)
	return order, nil
}

func DeleteOrder(id string) error {
	_, pos := GetOrder(id)

	if pos == -1 {
		return fmt.Errorf("order not found")
	}

	orderLength := len(orders)

	if orderLength > 1 {
		orders[pos] = orders[orderLength-1]
		orders = orders[:orderLength-1]
	} else {
		orders = []types.Order{}
	}

	return nil
}
