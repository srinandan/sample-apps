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
	items "github.com/srinandan/sample-apps/inventory/items"
)

func ListInventoryHandler(w http.ResponseWriter, r *http.Request) {
	items := items.ListItems()

	common.ResponseHandler(w, items)
}

func GetInventoryHandler(w http.ResponseWriter, r *http.Request) {
	//read path variables
	vars := mux.Vars(r)
	item, pos := items.GetItem(vars["id"])

	if pos != -1 {
		common.ResponseHandler(w, item)
	} else {
		common.NotFoundHandler(w, "item not found")
	}
}

func CreateInventoryHandler(w http.ResponseWriter, r *http.Request) {
	//read the body
	itemBytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		common.ErrorHandler(w, err)
		return
	}

	item := types.Item{}
	err = json.Unmarshal(itemBytes, &item)
	if err != nil {
		common.ErrorHandler(w, err)
		return
	}

	item = items.CreateItem(item)

	common.ResponseHandler(w, item)
}

func DeleteInventoryHandler(w http.ResponseWriter, r *http.Request) {
	//read path variables
	vars := mux.Vars(r)
	err := items.DeleteItem(vars["id"])

	if err != nil {
		common.NotFoundHandler(w, "item not found")
		return
	} else {
		common.ResponseHandler(w, "{\"msg\":\""+vars["id"]+" is deleted\"}")
	}
}
