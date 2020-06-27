# websocket-sample

This sample is meant to demonstrate websockets with Apigee hybrid.

## Server

A sample websocket sample is included which responds with `hello`. To install the server, 

### Server Prequisites

* kubectl 1.13 or higher
* docker 19.x or higher
* skaffold 1.1.0 or higher (optional)

### Server Installation

#### Server install via kubectl

1. Build the [docker image](./server/Dockerfile)
2. Deploy to a container register
3. Modify the kubernetes [manifest](./server/websockets.yaml)
4. Deploy the server `kubectl apply -f websockets.yaml`

#### Server install via skaffold

This application can also be installed via [skaffold](https://skaffold.dev/). Modify the [skaffold.yaml](./skaffold.yaml) to set the appropriate project name.

```bash

skaffold run
```

##### Errors in Skaffold

When rerunning/installing the application, you may observe errors like this:

```bash

 - Error from server (Invalid): error when applying patch:
 ...
 ...
 `selector` does not match template `labels`
 ```

There is an open [issue](https://github.com/GoogleContainerTools/skaffold/issues/3133) for this in the skaffold project.

Workaound: first run `skaffold delete` and then `skaffold run`

## Apigee components

An API Proxy is included which verifies an API Key before allowing the websocket connection.  

### Prerequisites

* Apigee hybrid runtime
* GKE 1.13.x
* [apigeecli](https://github.com/srinandan/apigeecli) v1.1 (optional)
* A GCP Service Account with pemrissions to create apis, products ,developers and apps.

### Install Apigee components

```bash

./install-sample-apigee.sh {org-name} {env-name} {path-to-service-account.json}
```

## Client

A sample websocket sample is included which takes an input from the user and sends it to the server. 

### Client Prequisites

* go 1.13 or higher
* goreleaser v0.120.3 (optional)

### Start client steps

Option 1: Download the binary from [releases](https://github.com/srinandan/websocket-sample/releases)
Option 2: Build the binary manually using [./build.sh](./client/build.sh)
Option 3: Build using [goreleaser](https://goreleaser.com/) `goreleaser --snapshot --skip-publish --rm-dist`
NOTE: has been tested with v0.120.3

## Test

Input:

```bash

export API_KEY=xxx
export ENDPOINT=example.com
#use this flag to use ws:// vs. wss://  ; default is false
export TLS=false

./websocket-client
```

Output:

```bash

2019/12/21 21:27:12 connecting to ws://example.com/v1/ws?apikey=xxx
2019/12/21 21:27:13 recv: Replying to: '2019-12-21 21:27:13.403018 -0800 PST m=+1.010449853' with 'hello'
2019/12/21 21:27:14 recv: Replying to: '2019-12-21 21:27:14.400193 -0800 PST m=+2.007614408' with 'hello'
2019/12/21 21:27:15 recv: Replying to: '2019-12-21 21:27:15.403358 -0800 PST m=+3.010769896' with 'hello'
2019/12/21 21:27:16 recv: Replying to: '2019-12-21 21:27:16.400729 -0800 PST m=+4.008131894' with 'hello'
```
