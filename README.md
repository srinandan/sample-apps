# Sample applications

This repo contains a few sample microservices for experimenting with Apigee.

* [Inventory](./inventory) - A REST service that implements CRUD on `items`
* [Orders](./orders) - A REST service that implements CRUD on `orders`
* [Tracking](./tracking) - A gRPC service that implements CRUD on shipment `tracking`
* [Tracking Client](./tracking) - A REST service that exposes the tracking service as a REST API
* [Client](./client) - A busybox based kubernetes pod for debugging/troubleshooting in kubernetes
* [Orders GraphQL](./orders-gql) - A graphQL implementation of the `orders` and `inventory` service.
* [Load Teest](./load-test) - A [fortio](https://github.com/fortio/fortio) based load test application. For deployment in kubernetes

## Installation

### Prerequisites to build

* kubectl 1.13 or higher
* docker 19.x or higher
* [skaffold](https://skaffold.dev/) 1.1.0 or higher (optional)

### Install steps

These applications can also be installed via [skaffold](https://skaffold.dev/). Modify the [skaffold.yaml](./skaffold.yaml) to set the appropriate project name.

```bash

skaffold run
```

NOTE: `client` and `load-test` are not installed.