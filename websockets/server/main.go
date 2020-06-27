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

package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/v1/ws", func(w http.ResponseWriter, r *http.Request) {
		var conn, _ = upgrader.Upgrade(w, r, nil)
		go func(conn *websocket.Conn) {
			for {
				mType, msg, err := conn.ReadMessage()
				if err != nil {
					fmt.Println(err)
					conn.Close()
					break
				} else {
					fmt.Println("Received message: ", string(msg))
					reply := "Replying to: '" + string(msg) + "' with 'hello'"
					conn.WriteMessage(mType, []byte(reply))
				}
			}
		}(conn)
	})

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("Error starting: ", err)
	}
}
