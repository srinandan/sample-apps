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
	"math/rand"
	"net/http"
	"strconv"
	"time"

	types "internal/datatypes"

	odr "ordersdata"

	"github.com/gorilla/mux"

	common "internal/common"
)

func ListOrdersHandler(w http.ResponseWriter, r *http.Request) {
	orders := odr.ListOrders()

	common.ResponseHandler(w, orders, false)
}

func GetOrderHandler(w http.ResponseWriter, r *http.Request) {
	// read path variables
	vars := mux.Vars(r)
	order, pos := odr.GetOrder(vars["id"])

	if pos != -1 {
		common.ResponseHandler(w, order, false)
	} else {
		common.NotFoundHandler(w, "order not found")
	}
}

func CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	// read the body
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

	common.ResponseHandler(w, order, false)
}

func DeleteOrderHandler(w http.ResponseWriter, r *http.Request) {
	// read path variables
	vars := mux.Vars(r)
	err := odr.DeleteOrder(vars["id"])

	if err != nil {
		common.NotFoundHandler(w, "order not found")
		return
	} else {
		common.ResponseHandler(w, map[string]string{"msg": vars["key"] + " is deleted"}, false)
	}
}

func GetOrderItemsHandler(w http.ResponseWriter, r *http.Request) {
	// read path variables
	vars := mux.Vars(r)
	order, pos := odr.GetOrder(vars["id"])

	if pos != -1 {
		items := []types.Item{}
		for _, lineItem := range order.LineItems {
			item, _ := odr.GetOrderItem(lineItem.LineItemId)
			items = append(items, item)
		}
		common.ResponseHandler(w, items, false)
	} else {
		common.NotFoundHandler(w, "order not found")
	}
}

func GetOrderDelayHandler(w http.ResponseWriter, r *http.Request) {
	// read path variables
	vars := mux.Vars(r)
	interval, err := strconv.ParseInt(vars["interval"], 10, 64)
	if err != nil {
		common.BadRequestHandler(w, err)
		return
	}

	stop := make(chan bool, 1)
	go func() {
		time.Sleep(time.Duration(interval) * time.Second)
		order, pos := odr.GetOrder(vars["id"])
		if pos != -1 {
			common.ResponseHandler(w, order, false)
		} else {
			common.NotFoundHandler(w, "order not found")
		}
		stop <- true
	}()

	<-stop
}

func GetOrderRandomizedDelayHandler(w http.ResponseWriter, r *http.Request) {
	// read path variables
	vars := mux.Vars(r)

	min := 50
	max := 75

	interval := rand.Intn(max-min) + min
	stop := make(chan bool, 1)

	go func() {
		for {
			select {
			case <-stop:
				return
			case <-time.After(time.Duration(interval) * time.Millisecond):
				order, pos := odr.GetOrder(vars["id"])
				if pos != -1 {
					common.ResponseHandler(w, order, false)
				} else {
					common.NotFoundHandler(w, "order not found")
				}
				stop <- true
			}
		}
	}()
	<-stop
}
