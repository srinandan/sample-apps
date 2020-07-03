# cloud-dlp sample

## Prerequisite

* the user should have installed the google-auth-sidecar service on Kubernetes
* the user should have deployed the OAuth [shared flow](../oauth-sharedflow)

The sample invoke the [Cloud DLP API](https://cloud.google.com/dlp/docs). The same masks emails and phone numbers (PII data)

### Deploying the proxy

Be sure to change the Target url to point to your GCP project
`https://dlp.googleapis.com/v2/projects/{project-id}/content:deidentify`

### Invoking the sample

```bash

curl https://api.examples.com/dlp -d 'my email is test@email.com and my phone is 123-456-7890'
```

Expected result:

```bash

my email is ****************** and my phone is ************
```