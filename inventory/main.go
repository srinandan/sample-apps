package main

import (
	"context"
	"log"

	"github.com/gorilla/mux"
	common "github.com/srinandan/sample-apps/common"
	apis "github.com/srinandan/sample-apps/inventory/apis"
	app "github.com/srinandan/sample-apps/inventory/app"

	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	var wait time.Duration

	app.Initialize()

	r := mux.NewRouter()
	r.HandleFunc("/healthz", common.HealthHandler).
		Methods("GET")
	r.HandleFunc("/items", apis.ListInventoryHandler).
		Methods("GET")
	r.HandleFunc("/items", apis.CreateInventoryHandler).
		Methods("POST")
	r.HandleFunc("/items/{id}", apis.GetInventoryHandler).
		Methods("GET")
	r.HandleFunc("/items/{id}", apis.DeleteInventoryHandler).
		Methods("DELETE")

	common.Info.Println("Starting server - ", common.Address)

	//the following code is from gorilla mux samples
	srv := &http.Server{
		Addr:         common.Address,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
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
