# Sample applications

This repo contains a few sample microservices for experimenting with Apigee.

## Installation

### Prerequisites to build

* kubectl 1.13 or higher
* docker 19.x or higher
* skaffold 1.1.0 or higher (optional)

### Install steps

This application can also be installed via [skaffold](https://skaffold.dev/). Modify the [skaffold.yaml](./skaffold.yaml) to set the appropriate project name.

```bash

skaffold run
```