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
	"os"
	"os/signal"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	common "github.com/srinandan/sample-apps/common"
	api "github.com/srinandan/sample-apps/tracking/pkg/api/v1"
	service "github.com/srinandan/sample-apps/tracking/pkg/service/v1"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
)

//look for tokens
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

// RunServer runs gRPC service to publish tracking service
func RunServer(port string) error {

	ctx := context.Background()

	common.Info.Println("Starting server on port ", port)

	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	server := grpc.NewServer(
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

	api.RegisterShipmentServer(server, ShipmentServer)

	if err := server.Serve(listen); err != nil {
		common.Error.Fatalf("failed to serve: %v", err)
	}

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			common.Info.Println("shutting down gRPC server...")

			server.GracefulStop()

			<-ctx.Done()
		}
	}()

	// start gRPC server
	common.Info.Println("starting gRPC server...")
	return server.Serve(listen)
}
