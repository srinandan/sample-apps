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

// go run cmd/server/main.go
package main

import (
	"fmt"
	"os"

	"github.com/srinandan/sample-apps/common"
	grpc "github.com/srinandan/sample-apps/tracking/pkg/protocol/grpc"
)

func main() {
	//initialize logging
	common.InitLog()

	if err := grpc.RunServer(common.GetgRPCPort(), common.GetAddress()); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
