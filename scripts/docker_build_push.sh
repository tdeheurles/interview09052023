#!/usr/bin/env bash

set -euo pipefail

version="${1}"
image="tdeheurles/interview09052023-web:${version}"

(
    cd web
    docker build -t "${image}" .
)

docker push "${image}"
