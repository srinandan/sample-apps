module github.com/srinandan/websocket-sample

go 1.21

require internal/common v1.0.0

require (
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/gorilla/mux v1.8.1 // indirect
)

replace internal/common => ../../internal/common

require github.com/gorilla/websocket v1.5.0
