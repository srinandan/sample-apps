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

package inventory

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"time"

	types "internal/datatypes"
)

// Items
var items = []types.Item{}

func ReadInventoryFile() error {
	inventoryBytes, err := ioutil.ReadFile("items.json")
	if err != nil {
		return err
	}

	if err = json.Unmarshal(inventoryBytes, &items); err != nil {
		return err
	}

	return nil
}

func ListItems() []types.Item {
	return items
}

func GetItem(id string) (types.Item, int) {
	item := types.Item{}
	for pos, item := range items {
		if item.ProductId == id {
			return item, pos
		}
	}
	return item, -1
}

func getId() string {
	rand.Seed(time.Now().UnixNano())
	min := 900000
	max := 999996

	return strconv.Itoa(rand.Intn(max-min+1) + min)
}

func CreateItem(item types.Item) types.Item {
	item.ProductId = getId()
	items = append(items, item)
	return item
}

func DeleteItem(id string) error {
	_, pos := GetItem(id)

	if pos == -1 {
		return fmt.Errorf("item not found")
	}

	itemLength := len(items)

	if itemLength > 1 {
		items[pos] = items[itemLength-1]
		items = items[:itemLength-1]
	} else {
		items = []types.Item{}
	}

	return nil
}
