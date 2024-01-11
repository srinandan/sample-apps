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
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	common "internal/common"

	"github.com/gorilla/mux"
	apis "github.com/srinandan/sample-apps/inventory/apis"
	app "github.com/srinandan/sample-apps/inventory/app"
)

func main() {
	var wait time.Duration

	app.Initialize()

	r := mux.NewRouter()
	r.Use(common.Middleware())

	r.HandleFunc("/items", apis.ListInventoryHandler).
		Methods("GET")
	r.HandleFunc("/items", apis.CreateInventoryHandler).
		Methods("POST")
	r.HandleFunc("/items/{id}", apis.GetInventoryHandler).
		Methods("GET")
	r.HandleFunc("/items/{id}", apis.DeleteInventoryHandler).
		Methods("DELETE")

	common.Info.Println("Starting server - ", common.GetAddress())

	// the following code is from gorilla mux samples
	srv := &http.Server{
		Addr:         common.GetAddress(),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			common.Error.Println(err)
		}
	}()

	rHealth := mux.NewRouter()
	rHealth.HandleFunc("/healthz", common.HealthHandler).
		Methods("GET")

	healthSvr := &http.Server{
		Addr:         common.GetHealthAddress(),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      rHealth,
	}

	common.Info.Println("Starting healthcheck - ", common.GetHealthAddress())

	go func() {
		if err := healthSvr.ListenAndServe(); err != nil {
			common.Error.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)

	common.Info.Println("Shutting down")

	os.Exit(0)
}
