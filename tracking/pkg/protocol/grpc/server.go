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
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"

	common "github.com/srinandan/sample-apps/common"
	api "github.com/srinandan/sample-apps/tracking/pkg/api/v1"
	service "github.com/srinandan/sample-apps/tracking/pkg/service/v1"
)

// RunServer runs gRPC service to publish tracking service
func RunServer(port string) error {

	ctx := context.Background()

	common.Info.Println("Starting server on port ", port)

	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	// register service
	server := grpc.NewServer()

	trackingService, err := service.NewTrackingService()
	if err != nil {
		return err
	}

	api.RegisterTrackingServer(server, trackingService)

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
