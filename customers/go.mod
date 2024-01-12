module github.com/srinandan/sample-apps/customers

go 1.21

require internal/common v1.0.0

replace internal/common => ../internal/common

require internal/datatypes v1.0.0

require github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect

replace internal/datatypes => ../internal/datatypes

require customersdata v1.0.0

replace customersdata => ../customers/data

require github.com/gorilla/mux v1.8.1
