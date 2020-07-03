# cloud-datastore sample

## Prerequisite

* the user should have installed the google-auth-sidecar service on Kubernetes
* the user should have deployed the OAuth [shared flow](../oauth-sharedflow)

The sample invoke the [Cloud Datastore API](https://cloud.google.com/datastore/docs/). Setup Cloud Datastore to create a kind (the same uses a kind called `sample`). The data is initialized with `id` and `name`.

### Deploying the proxy

Be sure to change the Target url to point to your GCP project
`https://datastore.googleapis.com/v1/projects/{project-id}:runQuery`

### Invoking the sample

```bash

curl https://api.examples.com/data
```

Expected result:

```bash

{
  "batch": {
    "entityResultType": "FULL",
    "entityResults": [
      {
        "entity": {
          "key": {
            "partitionId": {
              "projectId": "xx"
            },
            "path": [
              {
                "kind": "sample",
                "id": "5629499534213120"
              }
            ]
          },
          "properties": {
            "name": {
              "stringValue": "test"
            }
          }
        },
        "cursor": "xx",
        "version": "xx"
      }
    ],
    "endCursor": "xx",
    "moreResults": "NO_MORE_RESULTS"
  }
}
```