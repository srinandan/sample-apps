package main

import (
	"net/http"

	"github.com/99designs/gqlgen/handler"
	common "github.com/srinandan/sample-apps/common"
	gql "github.com/srinandan/sample-apps/orders-gql/gql"
	resolvers "github.com/srinandan/sample-apps/orders-gql/resolvers"
)

func main() {
	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(gql.NewExecutableSchema(gql.Config{Resolvers: &resolvers.Resolver{}})))
	
	common.InitLog()
	
	common.Info.Printf("connect to http://%s/ for GraphQL playground", common.GetAddress())
	common.Error.Fatal(http.ListenAndServe(common.GetAddress(), nil))
}
