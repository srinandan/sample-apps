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
	"net/http"

	"github.com/99designs/gqlgen/handler"
	common "github.com/srinandan/sample-apps/common"
	gql "github.com/srinandan/sample-apps/orders-gql/gql"
	"github.com/srinandan/sample-apps/orders-gql/odr"
	resolvers "github.com/srinandan/sample-apps/orders-gql/resolvers"
)

func main() {
	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(gql.NewExecutableSchema(gql.Config{Resolvers: &resolvers.Resolver{}})))

	common.InitLog()

	// ReadOrdersFile
	if err := odr.ReadOrdersFile(); err != nil {
		common.Error.Println("error reading orders file ", err)
	}

	common.Info.Printf("connect to http://%s/ for GraphQL playground", common.GetAddress())
	common.Error.Fatal(http.ListenAndServe(common.GetAddress(), nil))
}
