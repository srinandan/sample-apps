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
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"

	"net/http"
)

//Version
var Version, Git string

func main() {

	var err error
	var c *websocket.Conn
	var resp *http.Response

	endpoint := os.Getenv("ENDPOINT")
	if endpoint == "" {
		endpoint = "localhost:3000"
	}

	apiKey := os.Getenv("API_KEY")

	apiPath := os.Getenv("API_PATH")
	if apiPath == "" {
		apiPath = "/v1/ws"
	}

	token := os.Getenv("TOKEN")
	if token != "" {
		token = "Bearer " + token
	}

	enableTLS := os.Getenv("TLS")
	scheme := "ws"
	if enableTLS != "" {
		scheme = "wss"
	}

	fmt.Println("websocket-client version " + Version + ", Git: " + Git)
	fmt.Println("Endpoint: " + scheme + "://" + endpoint + apiPath)
	fmt.Println("api_key: " + apiKey)
	fmt.Println("token: " + token)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: scheme, Host: endpoint, Path: apiPath}

	log.Printf("connecting to %s", u.String())

	if apiKey != "" {
		q := u.Query()
		q.Set("apikey", apiKey)
		u.RawQuery = q.Encode()
		c, resp, err = websocket.DefaultDialer.Dial(u.String(), http.Header{"x-api-key": {apiKey}})
	} else if token != "" {
		c, resp, err = websocket.DefaultDialer.Dial(u.String(), http.Header{"Authorization": {token}})
	} else {
		c, resp, err = websocket.DefaultDialer.Dial(u.String(), nil)
	}

	if err != nil {
		if resp != nil {
			log.Fatalf("handshake failed with status %d and message %v", resp.StatusCode, err)
		} else {
			log.Fatalf("handshake failed with status message %v", err)
		}
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
			if err != nil {
				log.Println("write:", err)
				return
			}
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
