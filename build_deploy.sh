#!/bin/bash

# build images
skaffold build -p gcb --default-repo=us-docker.pkg.dev/$PROJECT_ID/images --quiet > build_result.json

# deploy to GKE cluster
gcloud deploy apply --file=clouddeploy.yaml --region=$LOCATION --project=$PROJECT_ID

# openssl rand -base64 25 | tr -dc 'a-z0-9' | cut -c1-4
gcloud deploy releases create sample-apps-$(base64 < /dev/urandom | tr -dc 'a-z0-9' | fold -w 4 | head -n 1) --project=$PROJECT_ID --build-artifacts=build_result.json --region=$LOCATION --delivery-pipeline=sample-apps-pipeline

rm build_result.json