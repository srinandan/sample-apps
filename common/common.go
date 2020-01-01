package common

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

//log levels, default is error
var (
	Info  *log.Logger
	Error *log.Logger
)

//ErrorMessage hold the return value when there is an error
type ErrorMessage struct {
	StatusCode int    `json:"status_code,omitempty"`
	Message    string `json:"message,omitempty"`
}

var errorMessage = ErrorMessage{StatusCode: http.StatusInternalServerError}

//Address to start server
const address = "0.0.0.0:"

const defaultPort = "8080"

const defaultgRPCPort = "50051"

//InitLog function initializes the logger objects
func InitLog() {
	var infoHandle = ioutil.Discard

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

//GetAddress returns the REST API port for the server to listen to
func GetAddress() string {
	port := os.Getenv("PORT")
	if port == "" {
		return address + port
	}
	return address + defaultPort
}

//GetgRPCPort returns the gRPC port for the server to listen to
func GetgRPCPort() string {
	port := os.Getenv("GRPC_PORT")
	if port == "" {
		return defaultgRPCPort
	}
	return port
}

//HealthHandler handles kubernetes healthchecks
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

func ErrorHandler(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusInternalServerError)

	errorMessage.Message = err.Error()

	if err := json.NewEncoder(w).Encode(errorMessage); err != nil {
		Error.Println(err)
	}
}

//NotFoundHandler returns a 404 when an entity is not found
func NotFoundHandler(w http.ResponseWriter, msg string) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)

	errorMessage.Message = msg
	errorMessage.StatusCode = http.StatusNotFound

	if err := json.NewEncoder(w).Encode(errorMessage); err != nil {
		Error.Println(err)
	}
}

//BadRequestHandler returns a 400 when the client sends an incorrect payload
func BadRequestHandler(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)

	errorMessage.Message = err.Error()
	errorMessage.StatusCode = http.StatusBadRequest

	if err := json.NewEncoder(w).Encode(errorMessage); err != nil {
		Error.Println(err)
	}
}

//ResponseHandler returns a 200 when the response is successful
func ResponseHandler(w http.ResponseWriter, response interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		Error.Println(err)
	}
}
