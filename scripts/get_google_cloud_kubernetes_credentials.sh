#!/usr/bin/env bash

gcloud container clusters get-credentials my-gke-cluster \
    --region="europe-west9-a" \
    --project="interview09052023"
