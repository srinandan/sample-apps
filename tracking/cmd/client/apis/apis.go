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

package apis

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/trace"
	"google.golang.org/grpc"

	"github.com/golang/protobuf/jsonpb"
	"github.com/gorilla/mux"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"

	empty "github.com/golang/protobuf/ptypes/empty"
	common "github.com/srinandan/sample-apps/common"
	v1 "github.com/srinandan/sample-apps/tracking/pkg/api/v1"
)

//endpoint to reach the tracking service
var trackingEndpoint = os.Getenv("TRACKING")

const tokenType = "Bearer"
const authorizationHeader = "Authorization"
const apiKeyHeader = "x-api-key"

var streamTrackingClient v1.ShipmentClient
var streamConn *grpc.ClientConn

type trackingOAuthCreds struct {
	AccessToken string
}

type trackingAPIKeyCreds struct {
	APIKey string
}

func (c *trackingOAuthCreds) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	return map[string]string{
		authorizationHeader: tokenType + " " + c.AccessToken,
	}, nil
}

func (c *trackingAPIKeyCreds) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	return map[string]string{
		apiKeyHeader: c.APIKey,
	}, nil
}

func (c *trackingOAuthCreds) RequireTransportSecurity() bool {
	return false
}

func (c *trackingAPIKeyCreds) RequireTransportSecurity() bool {
	return false
}

func NewTokenFromHeader(jwt string) (credentials.PerRPCCredentials, error) {
	return &trackingOAuthCreds{AccessToken: jwt}, nil
}

func NewKeyFromHeader(key string) (credentials.PerRPCCredentials, error) {
	return &trackingAPIKeyCreds{APIKey: key}, nil
}

func initClient(r *http.Request) (trackingClient v1.ShipmentClient, conn *grpc.ClientConn, err error) {

	credType, cred := getCredential(r)

	// Set up a connection to the server.
	if trackingEndpoint == "" {
		trackingEndpoint = "localhost:50051"
	}
	common.Info.Println("connect to ", trackingEndpoint)

	if credType == "accessToken" {
		creds, _ := NewTokenFromHeader(cred)
		conn, err = grpc.Dial(trackingEndpoint, grpc.WithInsecure(), grpc.WithPerRPCCredentials(creds), grpc.WithStatsHandler(new(ocgrpc.ClientHandler)))
	} else if credType == "apiKey" {
		creds, _ := NewKeyFromHeader(cred)
		conn, err = grpc.Dial(trackingEndpoint, grpc.WithInsecure(), grpc.WithPerRPCCredentials(creds), grpc.WithStatsHandler(new(ocgrpc.ClientHandler)))
	} else {
		conn, err = grpc.Dial(trackingEndpoint, grpc.WithInsecure(), grpc.WithStatsHandler(new(ocgrpc.ClientHandler)))
	}

	if err != nil {
		return nil, nil, fmt.Errorf("did not connect: %v", err)
	}

	trackingClient = v1.NewShipmentClient(conn)

	return trackingClient, conn, nil
}

func initStreamClient(r *http.Request) (err error) {

	if streamConn == nil || streamConn.GetState() != connectivity.Ready {
		credType, cred := getCredential(r)

		// Set up a connection to the server.
		if trackingEndpoint == "" {
			trackingEndpoint = "localhost:50051"
		}
		common.Info.Println("connect to ", trackingEndpoint)
		if credType == "accessToken" {
			creds, _ := NewTokenFromHeader(cred)
			streamConn, err = grpc.Dial(trackingEndpoint, grpc.WithInsecure(), grpc.WithPerRPCCredentials(creds), grpc.WithStatsHandler(new(ocgrpc.ClientHandler)))
		} else if credType == "apiKey" {
			creds, _ := NewKeyFromHeader(cred)
			streamConn, err = grpc.Dial(trackingEndpoint, grpc.WithInsecure(), grpc.WithPerRPCCredentials(creds), grpc.WithStatsHandler(new(ocgrpc.ClientHandler)))
		} else {
			streamConn, err = grpc.Dial(trackingEndpoint, grpc.WithInsecure(), grpc.WithStatsHandler(new(ocgrpc.ClientHandler)))
		}

		if err != nil {
			return fmt.Errorf("did not connect: %v", err)
		}

		streamTrackingClient = v1.NewShipmentClient(streamConn)
		reconnectServer()
	}
	return nil
}

func closeClient(conn *grpc.ClientConn) {
	if conn != nil {
		defer conn.Close()
	}
}

func CloseStreamClient() {
	if streamConn != nil {
		streamConn.Close()
		streamConn = nil
	}
}

func getCredential(r *http.Request) (credType string, cred string) {
	//get access token
	if authHeaderValue := r.Header.Get(authorizationHeader); authHeaderValue != "" {
		splitToken := strings.Split(authHeaderValue, tokenType)
		if len(splitToken) > 1 {
			accessToken := splitToken[1]
			common.Info.Println("Using access token ", accessToken)
			return "accessToken", strings.TrimSpace(accessToken)
		}
		return "", ""
	} else if apiKeyHeaderValue := r.Header.Get(apiKeyHeader); apiKeyHeaderValue != "" {
		common.Info.Println("Using api key ", apiKeyHeaderValue)
		return "apiKey", apiKeyHeaderValue
	} else {
		common.Info.Println("no auth used")
	}
	return "", ""
}

func ListTrackingDetailsHandler(w http.ResponseWriter, r *http.Request) {

	ctx, rootspan := trace.StartSpan(context.Background(), "ListTrackingDetailsHandler")
	defer rootspan.End()

	// create child span for backend call
	_, childspan := trace.StartSpan(ctx, "call to tracking server")
	defer childspan.End()

	trackingClient, conn, err := initClient(r)

	if err != nil {
		common.ErrorHandler(w, err)
		return
	}

	defer closeClient(conn)

	childCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Call Create
	resp, err := trackingClient.ListTracking(childCtx, &empty.Empty{})
	if err != nil {
		e, _ := status.FromError(err)
		if e.Code() == codes.Unavailable {
			common.ErrorHandler(w, err)
		} else if e.Code() == codes.PermissionDenied || e.Code() == codes.Unauthenticated {
			common.PermissionDeniedHandler(w, err)
		} else {
			common.NotFoundHandler(w, err.Error())
		}
		return
	}

	m := &jsonpb.Marshaler{}
	trackingListResponse, err := m.MarshalToString(resp)
	if err != nil {
		common.ErrorHandler(w, err)
		return
	}

	fmt.Fprintln(w, string(trackingListResponse))
}

func GetTrackingDetailsHandler(w http.ResponseWriter, r *http.Request) {

	ctx, rootspan := trace.StartSpan(context.Background(), "GetTrackingDetailsHandler")
	defer rootspan.End()

	// create child span for backend call
	_, childspan := trace.StartSpan(ctx, "call to tracking server")
	defer childspan.End()

	trackingClient, conn, err := initClient(r)

	if err != nil {
		common.ErrorHandler(w, err)
		return
	}

	defer closeClient(conn)

	//read path variables
	vars := mux.Vars(r)
	trackingId := vars["id"]

	childCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Call Create
	resp, err := trackingClient.GetTracking(childCtx, &v1.GetTrackingRequest{
		TrackingId: trackingId,
	})
	if err != nil {
		e, _ := status.FromError(err)
		if e.Code() == codes.Unavailable {
			common.ErrorHandler(w, err)
		} else if e.Code() == codes.PermissionDenied || e.Code() == codes.Unauthenticated {
			common.PermissionDeniedHandler(w, err)
		} else {
			common.NotFoundHandler(w, err.Error())
		}
		return
	}

	m := &jsonpb.Marshaler{}
	trackingResponse, err := m.MarshalToString(resp)
	if err != nil {
		common.ErrorHandler(w, err)
		return
	}

	fmt.Fprintln(w, string(trackingResponse))
}

func NotifyTrackingDetailsHandler(w http.ResponseWriter, r *http.Request) {

	err := initStreamClient(r)
	if err != nil {
		common.ErrorHandler(w, err)
		return
	}

	//read path variables
	vars := mux.Vars(r)
	trackingId := vars["id"]

	stream, err := streamTrackingClient.NotifyTracking(context.Background())
	if err != nil {
		common.ErrorHandler(w, err)
		return
	}

	done := make(chan bool)

	//go routine to send request
	go func() {
		req := &v1.GetTrackingRequest{
			TrackingId: trackingId,
		}
		if err := stream.Send(req); err != nil {
			common.ErrorHandler(w, err)
			return
		}
		if err := stream.CloseSend(); err != nil {
			common.ErrorHandler(w, err)
			return
		}
	}()

	//go routine to receive requests
	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				close(done)
				return
			}
			if err != nil {
				common.ErrorHandler(w, err)
				return
			}
			m := &jsonpb.Marshaler{}
			trackingResponse, err := m.MarshalToString(resp)
			if err != nil {
				common.ErrorHandler(w, err)
				return
			}
			fmt.Fprintln(w, string(trackingResponse))
		}
	}()

	<-done
}

func reconnectServer() {
	ticker := time.NewTicker(30 * time.Second)

	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				common.Info.Println("closing connection ", t)
				CloseStreamClient()
			}
		}
	}()
}
