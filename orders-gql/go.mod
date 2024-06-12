module github.com/srinandan/sample-apps/orders-gql

go 1.21

require internal/common v1.0.0

replace internal/common => ../internal/common

require internal/datatypes v1.0.0

replace internal/datatypes => ../internal/datatypes

require (
	github.com/99designs/gqlgen v0.17.5
	github.com/vektah/gqlparser/v2 v2.5.14
)

require (
	github.com/agnivade/levenshtein v1.1.1 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/gorilla/mux v1.8.1 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
)
