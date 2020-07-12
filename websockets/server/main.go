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
	"net/http"
	"os"
	"time"

	"github.com/gorilla/websocket"
	"github.com/srinandan/sample-apps/common"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/plugin/ochttp/propagation/tracecontext"
)

var upgrader = websocket.Upgrader{}

func main() {
	//init logging
	common.InitLog()
	//init tracing
	if os.Getenv("DISABLE_TRACING") == "" {
		common.Info.Println("Tracing enabled.")
		go common.InitTracing("orders")
	} else {
		common.Info.Println("Tracing disabled.")
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/ws", wsResponse)
	mux.HandleFunc("/healthz", common.HealthHandler)

	och := &ochttp.Handler{
		Handler:     mux,
		Propagation: &tracecontext.HTTPFormat{},
	}

	//the following code is from gorilla mux samples
	srv := &http.Server{
		Addr:         "0.0.0.0:3000",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      och,
	}

	err := srv.ListenAndServe()
	if err != nil {
		common.Error.Println("Error starting: ", err)
	}
}

func wsResponse(w http.ResponseWriter, r *http.Request) {
	var conn, _ = upgrader.Upgrade(w, r, nil)
	go func(conn *websocket.Conn) {
		for {
			mType, msg, err := conn.ReadMessage()
			if err != nil {
				common.Error.Println(err)
				conn.Close()
				break
			} else {
				common.Info.Println("Received message: ", string(msg))
				reply := "Replying to: '" + string(msg) + "' with 'hello'"
				conn.WriteMessage(mType, []byte(reply))
			}
		}
	}(conn)
}
