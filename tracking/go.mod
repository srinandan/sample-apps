module github.com/srinandan/sample-apps/tracking

go 1.21

require internal/common v1.0.0

replace internal/common => ../internal/common

replace internal/datatypes => ../internal/datatypes

require (
	github.com/golang/protobuf v1.5.3
	github.com/gorilla/mux v1.8.1
	github.com/grpc-ecosystem/go-grpc-middleware v1.4.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.15.2
	go.opencensus.io v0.24.0
	google.golang.org/genproto v0.0.0-20230410155749-daa745c078e1
	google.golang.org/grpc v1.54.0
	google.golang.org/protobuf v1.30.0
)

require (
	cloud.google.com/go/compute v1.19.1 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/stretchr/testify v1.8.2 // indirect
	golang.org/x/net v0.9.0 // indirect
	golang.org/x/oauth2 v0.7.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/text v0.9.0 // indirect
)
