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

package ks

import (
	"io/ioutil"
	"os"
	"path"
)

var key  string
var kidBytes, certBytes []byte

func ReadFiles() (err error) {
	var keyPath = os.Getenv("KEY_PATH")

	if keyPath == "" {
		keyPath = "/policy-secret"
	}

	certBytes, err = ioutil.ReadFile(path.Join(keyPath, "remote-service.crt"))
	if err != nil {
		return err
	}

	keyBytes, err := ioutil.ReadFile(path.Join(keyPath, "remote-service.key"))
	if err != nil {
		return err
	}

	key = string(keyBytes)

	kidBytes, err = ioutil.ReadFile(path.Join(keyPath, "remote-service.properties"))
	if err != nil {
		return err
	}

	return nil
}

func GetCert() []byte {
	return certBytes
}

func GetKey() string {
	return key
}

func GetKid() []byte {
	return kidBytes
}
