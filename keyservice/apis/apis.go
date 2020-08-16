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
	"net/http"

	common "github.com/srinandan/sample-apps/common"
	ks "github.com/srinandan/sample-apps/keyservice/ks"
)

func GetKeyHandler(w http.ResponseWriter, r *http.Request) {
	key := ks.GetKey()
	if key != "" {
		common.ResponseHandler(w, key, true)
	} else {
		common.NotFoundHandler(w, "key not found")
	}
}

func GetCertHandler(w http.ResponseWriter, r *http.Request) {
	certBytes := ks.GetCert()
	var jsonMap map[string]interface{}

	if err := json.Unmarshal(certBytes, &jsonMap); err != nil {
		common.ErrorHandler(w, err)
	}

	if len(certBytes) != 0 {
		common.ResponseHandler(w, jsonMap, false)
	} else {
		common.NotFoundHandler(w, "cert not found")
	}
}

func GetKidHandler(w http.ResponseWriter, r *http.Request) {
	kidBytes := ks.GetKid()
	var jsonMap map[string]interface{}

	if err := json.Unmarshal(kidBytes, &jsonMap); err != nil {
		common.ErrorHandler(w, err)
	}

	if len(kidBytes) != 0 {
		common.ResponseHandler(w, jsonMap, false)
	} else {
		common.NotFoundHandler(w, "kid not found")
	}
}
