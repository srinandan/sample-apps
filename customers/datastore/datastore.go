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

package datastore

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"time"

	types "github.com/srinandan/sample-apps/common/types"
)

//Customers
var customers = []types.Customer{}

func ReadFile() error {
	customerBytes, err := ioutil.ReadFile("customers.json")
	if err != nil {
		return err
	}

	if err = json.Unmarshal(customerBytes, &customers); err != nil {
		return err
	}

	return nil
}

func ListCustomers() []types.Customer {
	return customers
}

func GetCustomer(id string) (types.Customer, int) {
	customer := types.Customer{}
	for pos, customer := range customers {
		if customer.CustomerID == id {
			return customer, pos
		}
	}
	return customer, -1
}

func getId() string {
	rand.Seed(time.Now().UnixNano())
	min := 600000
	max := 699996

	return strconv.Itoa(rand.Intn(max-min+1) + min)
}

func CreateCustomer(customer types.Customer) types.Customer {
	customer.CustomerID = getId()
	customers = append(customers, customer)
	return customer
}

func DeleteCustomer(id string) error {
	_, pos := GetCustomer(id)

	if pos == -1 {
		return fmt.Errorf("customer not found")
	}

	customersLength := len(customers)

	if customersLength > 1 {
		customers[pos] = customers[customersLength-1]
		customers = customers[:customersLength-1]
	} else {
		customers = []types.Customer{}
	}

	return nil
}
