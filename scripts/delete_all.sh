#!/usr/bin/env bash
set -euo pipefail

kubectl delete -f ./kubernetes/web.yaml --ignore-not-found
(
    cd infrastructure
    terraform destroy -auto-approve -var-file=config.dev.tfvars
)
