# google-auth-sidecar Samples

Install samples included in this repo.

## Prerequisites

* Apigee hybrid runtime
* GKE 1.13.x
* [apigeecli](https://github.com/srinandan/apigeecli) v1.1
* A GCP Service Account with pemrissions to create apis, products ,developers and apps.

## Install the Google OAuth sharedflow

This step is a prerequisite for any of the samples. This steps deploys a sharedflow to get a Google OAuth token and stores the token in the `accessToken` variable.

```bash
./install-oauth-sharedflow.sh {org} {env} {path-to-service-account.json}
```

## Install the Google Pub Sub sample

This sample install a sample API Proxy (basePath `/pubsub`) and a second sharedflow to publish messages to a Google Pub Sub topic. The topic name is set [here](https://github.com/srinandan/google-auth-sidecar/blob/master/examples/cloud-pubsub/sharedflowbundle/policies/Publish-Message.xml#L21). The default topic name is `mytopic`.

 ```bash

 ./install-pubsub-sample.sh {org} {env} {path-to-service-account.json}
 ```
