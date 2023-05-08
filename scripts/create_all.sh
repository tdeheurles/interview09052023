#!/usr/bin/env bash
set -euo pipefail

(
    cd infrastructure
    terraform apply -auto-approve -var-file=config.dev.tfvars
)

./scripts/kubectl_apply.sh
