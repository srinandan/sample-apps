#!/bin/bash

# build images
skaffold build -p gcb --default-repo=us-docker.pkg.dev/$PROJECT_ID/images --quiet > build_result.json

# deploy to GKE cluster
gcloud deploy apply --file=clouddeploy.yaml --region=$LOCATION --project=$PROJECT_ID

gcloud deploy releases create sample-apps-$((RANDOM % 900 + 100)) --project=$PROJECT_ID --build-artifacts=build_result.json --region=$LOCATION --delivery-pipeline=sample-apps-pipeline

rm build_result.json