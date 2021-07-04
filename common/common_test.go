// Copyright 2021 Google LLC
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

package common

import (
	"context"
	"testing"

	"github.com/gorilla/mux"

	"net/http"
	"time"

	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/plugin/ochttp/propagation/tracecontext"
)

func TestCommon(t *testing.T) {
	var wait time.Duration

	InitLog()

	r := mux.NewRouter()
	r.Use(Middleware())

	r.HandleFunc("/healthz", HealthHandler).
		Methods("GET")

	och := &ochttp.Handler{
		Handler:     r,
		Propagation: &tracecontext.HTTPFormat{},
	}

	//the following code is from gorilla mux samples
	srv := &http.Server{
		Addr:         GetAddress(),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      och, //r,
	}

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	_ = srv.Shutdown(ctx)

	Info.Println("Shutting down")

}
