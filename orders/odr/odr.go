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

	"github.com/srinandan/sample-apps/common"
	types "github.com/srinandan/sample-apps/common/types"
)

//Orders
var orders = []types.Order{}

func ReadOrdersFile() error {
	orderListBytes, err := Asset("../data/orders.json")
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

func getInventoryEndpoint() string {
	//endpoint to reach the inventory service
	var inventoryEndpoint = os.Getenv("INVENTORY")

	if inventoryEndpoint == "" {
		return "http://inventory.apps.svc.cluster.local:8080"
	}

	return inventoryEndpoint
}

func GetOrderItem(id string) (types.Item, error) {

	var req *http.Request

	item := types.Item{}

	client := &http.Client{}

	itemURL := getInventoryEndpoint() + "/items/" + id
	common.Info.Println("Connecting to: ", itemURL)

	req, err := http.NewRequest("GET", itemURL, nil)
	if err != nil {
		return item, err
	}

	resp, err := client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}

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
