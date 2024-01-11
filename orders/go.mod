module github.com/srinandan/sample-apps/orders

go 1.21

require internal/common v1.0.0

replace internal/common => ../internal/common

require internal/datatypes v1.0.0

replace internal/datatypes => ../internal/datatypes

require ordersdata v1.0.0

replace ordersdata => ../orders/data

require github.com/gorilla/mux v1.8.1

require github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
