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

package common

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// log levels, default is error
var (
	Info  *log.Logger
	Error *log.Logger
)

// ErrorMessage hold the return value when there is an error
type ErrorMessage struct {
	StatusCode int    `json:"status_code,omitempty"`
	Message    string `json:"message,omitempty"`
}

var errorMessage = ErrorMessage{StatusCode: http.StatusInternalServerError}

// Address to start server
const address = "0.0.0.0:"

const (
	defaultPort       = "8080"
	defaultHealthPort = "8090"
)

const (
	defaultgRPCPort       = "50051"
	defaultgRPCHealthPort = "5000"
)

// InitLog function initializes the logger objects
func InitLog() {
	infoHandle := ioutil.Discard

	debug, _ := strconv.ParseBool(os.Getenv("DEBUG"))

	if debug {
		infoHandle = os.Stdout
	}

	errorHandle := os.Stdout

	Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

// GetHealthAddress returns the healthcheck port
func GetHealthAddress() string {
	port := os.Getenv("HEALTH_PORT")
	if port == "" {
		return address + defaultHealthPort
	}
	return address + port
}

// GetAddress returns the REST API port for the server to listen to
func GetAddress() string {
	port := os.Getenv("PORT")
	if port == "" {
		return address + defaultPort
	}
	return address + port
}

// GetgRPCPort returns the gRPC port for the server to listen to
func GetgRPCPort() string {
	port := os.Getenv("GRPC_PORT")
	if port == "" {
		return defaultgRPCPort
	}
	return port
}

// GetgRPCHealthPort returns the gRPC port for the server to listen to
func GetgRPCHealthPort() string {
	port := os.Getenv("GRPC_HEALTH_PORT")
	if port == "" {
		return defaultgRPCHealthPort
	}
	return port
}

// HealthHandler handles kubernetes healthchecks
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	_ = json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

func ErrorHandler(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusInternalServerError)

	errorMessage.Message = err.Error()

	if err := json.NewEncoder(w).Encode(errorMessage); err != nil {
		Error.Println(err)
	}
}

// NotFoundHandler returns a 404 when an entity is not found
func NotFoundHandler(w http.ResponseWriter, msg string) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)

	errorMessage.Message = msg
	errorMessage.StatusCode = http.StatusNotFound

	if err := json.NewEncoder(w).Encode(errorMessage); err != nil {
		Error.Println(err)
	}
}

// BadRequestHandler returns a 400 when the client sends an incorrect payload
func BadRequestHandler(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusBadRequest)

	errorMessage.Message = err.Error()
	errorMessage.StatusCode = http.StatusBadRequest

	if err := json.NewEncoder(w).Encode(errorMessage); err != nil {
		Error.Println(err)
	}
}

// ResponseHandler returns a 200 when the response is successful
func ResponseHandler(w http.ResponseWriter, response interface{}, text bool) {
	if !text {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(response); err != nil {
			Error.Println(err)
		}
	} else {
		w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
		if str, ok := response.(string); ok {
			w.Write([]byte(str))
		}
	}
	// w.(http.Flusher).Flush()
}

func PermissionDeniedHandler(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusForbidden)

	errorMessage.Message = err.Error()
	errorMessage.StatusCode = http.StatusForbidden

	if err := json.NewEncoder(w).Encode(errorMessage); err != nil {
		Error.Println(err)
	}
}

func Middleware() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			printHeaders(req)
			next.ServeHTTP(w, req)
		})
	}
}

func getRouteName(route *mux.Route, req *http.Request) string {
	name := route.GetName()
	if name == "" {
		name, _ = route.GetPathTemplate()
		if name != "" {
			name = req.Method + " " + name
		}
	}

	return name
}

func printHeaders(req *http.Request) {
	headerBytes, _ := json.Marshal(req.Header)
	// don't print kubernetes probe
	if !strings.Contains(req.Header.Get("User-Agent"), "kube-probe") {
		Info.Println(string(headerBytes))
	}
}
