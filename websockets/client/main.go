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
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

// Version
var Version, Git string

func usage() {
	fmt.Println("")
	fmt.Println("websocket-client version ", Version)
	fmt.Println("")
	fmt.Println("Usage: websocker-client --endpoint=<endpoint> --path=<path>")
	fmt.Println("")
	fmt.Println("Options:")
	fmt.Println("endpoint    = hostname or ip of websocket server; default=localhost:3000")
	fmt.Println("path        = api path for the websocket app; default=/v1/ws")
	fmt.Println("token       = OAuth Bearer token")
	fmt.Println("apikey      = API Key")
	fmt.Println("hostHeader  = Set the host header")
	fmt.Println("tls         = Enable or disable tls; default=false")
	fmt.Println("skipVerify  = Skip verification of server cert; default=false")
}

func main() {
	var endpoint, apiKey, hostHeader, apiPath, token string
	var enableTLS, skipVerify, help bool

	flag.StringVar(&endpoint, "endpoint", "0.0.0.0:3000", "Websocket endpoint")
	flag.StringVar(&apiPath, "path", "/v1/ws", "Websocket path")
	flag.StringVar(&token, "token", "", "OAuth Bearer Token")
	flag.StringVar(&apiKey, "apikey", "", "API Key")
	flag.StringVar(&hostHeader, "hostHeader", "", "Set a host header")
	flag.BoolVar(&enableTLS, "tls", false, "TLS Enable/disable")
	flag.BoolVar(&skipVerify, "skipVerify", false, "SKip TLS verification")
	flag.BoolVar(&help, "help", false, "Display usage")

	// Parse commandline parameters
	flag.Parse()

	if help {
		usage()
		os.Exit(0)
	}

	var err error
	var c *websocket.Conn
	var resp *http.Response
	var d websocket.Dialer
	headers := http.Header{}

	if hostHeader != "" {
		headers.Add("Host", hostHeader)
	}

	scheme := "ws"

	if enableTLS {
		scheme = "wss"
		if skipVerify {
			d = websocket.Dialer{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
		} else {
			d = websocket.Dialer{}
		}
	}

	fmt.Println("websocket-client version " + Version + ", Git: " + Git)
	fmt.Println("Endpoint: " + scheme + "://" + endpoint + apiPath)
	fmt.Println("api_key: " + apiKey)
	fmt.Println("token: " + token)
	fmt.Println("skipVerify certificate: ", skipVerify)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: scheme, Host: endpoint, Path: apiPath}

	log.Printf("connecting to %s", u.String())

	if apiKey != "" {
		q := u.Query()
		q.Set("apikey", apiKey)
		u.RawQuery = q.Encode()
		headers.Add("x-api-key", apiKey)
	} else if token != "" {
		headers.Add("Authorization", token)
	}

	c, resp, err = d.Dial(u.String(), headers)

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
