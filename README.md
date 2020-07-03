# Sample applications

This repo contains a few sample microservices for experimenting with Apigee.

* [Inventory](./inventory) - A REST service that implements CRUD on `items`
* [Orders](./orders) - A REST service that implements CRUD on `orders`
* [Tracking](./tracking) - A gRPC service that implements CRUD on shipment `tracking`
* [Tracking Client](./tracking) - A REST service that exposes the tracking gRPC service as a REST API
* [Client](./client) - A busybox based kubernetes pod for debugging/troubleshooting in kubernetes
* [Orders GraphQL](./orders-gql) - A graphQL implementation of the `orders` and `inventory` service.
* [Customers](./customers) - To demonstrate the use of Cloud DLP (https://cloud.google.com/dlp/)
* [Websockets](./websockets) - To demostrate websockets via Apigee runtime
* [Google Auth Service](./google-auth-sidecar) - Obtain a Google OAuth token to call Google APIs
* [Load Test](./load-test) - A [fortio](https://github.com/fortio/fortio) based load test application. For deployment in kubernetes

## Installation

### Prerequisites to build

* kubectl 1.15 or higher
* [skaffold](https://skaffold.dev/) 1.12.0
* docker 19.x or higher (optional)

### Create a Kubernetes Secret with Google Service Account

```bash
./create-secret.sh
```

### Install steps

* Option 1 - Use Cloud Build

PREREQUISITES: Cloud Build service account must have role: "Kubernetes Engine Developer"

```bash
export GOOGLE_APPLICATION_CREDENTIALS=my-service-account.json
skaffold run -p gcb --default-repo=gcr.io/$PROJECT_ID
```

* Option 2 - local docker

```bash
skaffold run --default-repo=gcr.io/$PROJECT_ID
```

NOTE: `client` and `load-test` are not installed.

#### Installing a single application

```bash
skaffold run -p gcb --default-repo=gcr.io/$PROJECT_ID -f skaffold-{appname}.yaml
```

where {appname} can be orders, inventory, tracking, customers or websockets

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
