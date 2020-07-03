# google-auth-sidecar

[![Go Report Card](https://goreportcard.com/badge/github.com/srinandan/google-auth-sidecar)](https://goreportcard.com/report/github.com/srinandan/google-auth-sidecar)

This service is meant to run as a sidecar to the Apigee Runtime (also known as Message Processor). The service uses a Google Cloud Service Account generates an OAuth Access Token.

## Use Case

This service is meant to be used with the Apigee hybrid [API Runtime](https://docs.apigee.com/hybrid). When developing API Proxies on Apigee, a developer may want to interact with Google Cloud Services. Google Cloud Services are available on `*.googleapis.com` and require an OAuth access token to access the services through APIs.

A typical pattern/example would be:

* Instantiate the `google-auth-sidecar` services as a service or as a sidecar to the Message Processor
* Use a [Service Callout policy](https://docs.apigee.com/api-platform/reference/policies/service-callout-policy) to invoke the `google-auth-sidecar` service. The response contains an access token
* Use the access token to make a subsequent API call to any of the Google APIs
* For performance optimization, use the Apigee [cache policy](https://docs.apigee.com/api-platform/reference/policies/populate-cache-policy) and cache the token for a few seconds less than the token expiry

## Samples

A [shared flow](./examples/oauth-sharedflow) that generates and caches the token. See [here](./examples) for instructions on how to install the samples.

Three samples that makes use of the shared flow:

* [Query data from Cloud Datastore](./examples/cloud-datastore)
* [Mask sensitive data using Cloud DLP](./examples/cloud-dlp)
* [Publish to a topic using Cloud PubSub](./examples/cloud-pubsub)
* [Log messages to Stackdriver Logging](./examples/stackdriver-logging)

## Using the service via sharedflows

This [sample](./exmaples/oauth-sharedflow) implements a service callout to invoke the OAuth service.

## Using the service directly

Input:

```bash

curl http://google-auth-sidecar.apps.svc.cluster.local:8080/token
```

Output:

```bash

> User-Agent: curl/7.64.1
> Accept: */*
>
< HTTP/1.1 200 OK
< Content-Type: application/json; charset=UTF-8
< Date: xx, 00 Xxx 20xx 00:00:00 GMT
< Content-Length: 195
<
{"access_token":"xxx","expires_in":3600,"token_type":"Bearer"}
```
