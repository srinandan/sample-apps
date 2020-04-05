# Sample applications

This repo contains a few sample microservices for experimenting with Apigee.

* [Inventory](./inventory) - A REST service that implements CRUD on `items`
* [Orders](./orders) - A REST service that implements CRUD on `orders`
* [Tracking](./tracking) - A gRPC service that implements CRUD on shipment `tracking`
* [Tracking Client](./tracking) - A REST service that exposes the tracking service as a REST API
* [Client](./client) - A busybox based kubernetes pod for debugging/troubleshooting in kubernetes
* [Orders GraphQL](./orders-gql) - A graphQL implementation of the `orders` and `inventory` service.
* [Customers] - To demonstrate the use of Cloud DLP (https://cloud.google.com/dlp/)
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

### Skaffold Errors

When rerunning/installing the applications, you may observe errors like this:

```bash

 - Error from server (Invalid): error when applying patch:
 - {"metadata":{"annotations":{"kubectl.kubernetes.io/last-applied-configuration":"{\"apiVersion\":\"extensions/v1beta1\",\"kind\":\"Deployment\",\"metadata\":{\"annotations\":{},\"labels\":{\"app.kubernetes.io/managed-by\":\"skaffold-v1.1.0\",\"skaffold.dev/builder\":\"google-cloud-build\",\"skaffold.dev/cleanup\":\"true\",\"skaffold.dev/deployer\":
 ...
 ...
 `selector` does not match template `labels`
 ```

 There is an open [issue](https://github.com/GoogleContainerTools/skaffold/issues/3133) for this in the skaffold project.

 Workaound

 first run

 ```bash

 skaffold delete
 ```

and then

```bash

skaffold run
```
