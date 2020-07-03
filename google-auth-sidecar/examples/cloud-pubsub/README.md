# cloud-pubsub sample

## Prerequisite

* the user should have installed the google-auth-sidecar service on Kubernetes
* the user should have deployed the OAuth [shared flow](../oauth-sharedflow)

The sample invoke the [Cloud PubSub API](https://cloud.google.com/dlp/docs). The same masks emails and phone numbers (PII data)

### Deploying the proxy

Be sure to change the Target url to point to your GCP project
`hhttps://pubsub.googleapis.com/v1/projects/{project-id}/topics/{topicname}:publish`

### Invoking the sample

```bash

curl https://api.examples.com/pubsub -d 'this is a sample'
```

Expected result:

```bash

{
  "messageIds": [
    "902025832274037"
  ]
}
```