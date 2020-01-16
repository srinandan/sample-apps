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

package apis

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"

	common "github.com/srinandan/sample-apps/common"
	types "github.com/srinandan/sample-apps/common/types"
	odr "github.com/srinandan/sample-apps/orders/odr"
)

func ListOrdersHandler(w http.ResponseWriter, r *http.Request) {
	orders := odr.ListOrders()

	common.ResponseHandler(w, orders)
}

func GetOrderHandler(w http.ResponseWriter, r *http.Request) {
	//read path variables
	vars := mux.Vars(r)
	order, pos := odr.GetOrder(vars["id"])

	if pos != -1 {
		common.ResponseHandler(w, order)
	} else {
		common.NotFoundHandler(w,"order not found")
	}
}

func CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	//read the body
	orderBytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		common.ErrorHandler(w, err)
		return
	}

	order := types.Order{}
	err = json.Unmarshal(orderBytes, &order)
	if err != nil {
		common.ErrorHandler(w, err)
		return
	}

	order, err = odr.CreateOrder(order)
	if err != nil {
		common.BadRequestHandler(w, err)
		return
	}

	common.ResponseHandler(w, order)
}

func DeleteOrderHandler(w http.ResponseWriter, r *http.Request) {
	//read path variables
	vars := mux.Vars(r)
	err := odr.DeleteOrder(vars["id"])

	if err != nil {
		common.NotFoundHandler(w, "order not found")
		return
	} else {
		common.ResponseHandler(w, map[string]string{"msg": vars["key"] + " is deleted"})
	}
}

func GetOrderItemsHandler(w http.ResponseWriter, r *http.Request) {
	//read path variables
	vars := mux.Vars(r)
	order, pos := odr.GetOrder(vars["id"])

	if pos != -1 {
		items := []types.Item{}
		for _, lineItem := range order.LineItems {
			item, _ := odr.GetOrderItem(lineItem.LineItemId)
			items = append(items, item)
		}
		common.ResponseHandler(w, items)
	} else {
		common.NotFoundHandler(w, "order not found")
	}
}
