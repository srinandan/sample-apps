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

package grpc

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	common "github.com/srinandan/sample-apps/common"
	api "github.com/srinandan/sample-apps/tracking/pkg/api/v1"
	service "github.com/srinandan/sample-apps/tracking/pkg/service/v1"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
)

const maxMsgSize = 1024 * 1024
const timeoutValue = 15 * time.Second

var DISABLE_APISERVER = os.Getenv("DISABLE_APISERVER")

// look for tokens
func authorize(ctx context.Context) (context.Context, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ctx, fmt.Errorf("error getting ctx")
	}

	if len(md["authorization"]) == 0 && len(md["x-api-key"]) == 0 {
		common.Info.Println("no auth used")
	} else {
		if len(md["authorization"]) > 0 {
			if md["authorization"][0] != "" {
				common.Info.Println("Access token is ", md["authorization"][0])
			}
		}

		if len(md["x-api-key"]) > 0 {
			if md["x-api-key"][0] != "" {
				common.Info.Println("api key is ", md["x-api-key"][0])
			}
		}
	}
	return ctx, nil
}

// RunServer runs gRPC and REST endpoints for the tracking service
func RunServer(grpcPort string, restAddress string) error {

	var wait time.Duration
	var apiServer *http.Server

	ctx, cancel := context.WithTimeout(context.Background(), timeoutValue)
	defer cancel()

	common.Info.Println("Starting server on port ", grpcPort)

	listen, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpc_auth.UnaryServerInterceptor(authorize),
		),
		grpc.StreamInterceptor(
			grpc_auth.StreamServerInterceptor(authorize),
		),
	)

	ShipmentServer, err := service.NewShipmentService()
	if err != nil {
		return err
	}

	//register shipment server
	api.RegisterShipmentServer(grpcServer, ShipmentServer)

	//start gRPC server
	go func() {
		if err := grpcServer.Serve(listen); err != nil {
			common.Error.Fatalf("failed to serve: %v", err)
		}
	}()

	if DISABLE_APISERVER != "true" {
		gwmux := runtime.NewServeMux()

		conn, err := grpc.DialContext(
			ctx,
			"0.0.0.0:"+grpcPort,
			grpc.WithBlock(),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(maxMsgSize), grpc.MaxCallSendMsgSize(maxMsgSize)),
		)

		if err != nil {
			common.Error.Println(err)
			return err
		}

		err = api.RegisterShipmentHandler(ctx, gwmux, conn)
		if err != nil {
			common.Error.Println(err)
			return err
		}

		common.Info.Println("Starting API server on port ", common.GetAddress())
		apiServer := &http.Server{
			Addr:         restAddress,
			WriteTimeout: timeoutValue,
			ReadTimeout:  timeoutValue,
			IdleTimeout:  time.Second * 60,
			Handler:      gwmux,
		}
		go func() {
			if err := apiServer.ListenAndServe(); err != nil {
				common.Error.Println(err)
			}
		}()

	}

	// graceful shutdown
	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)
	// Block until we receive our signal.
	<-c
	// Create a deadline to wait for.
	ctxDeadline, cancelDeadline := context.WithTimeout(context.Background(), wait)
	defer cancelDeadline()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	grpcServer.GracefulStop()
	if DISABLE_APISERVER != "true" && apiServer != nil {
		err = apiServer.Shutdown(ctxDeadline)
	}
	common.Info.Println("Shutting down")
	return err
}
