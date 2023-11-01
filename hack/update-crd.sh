#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

if ! command -v controller-gen &> /dev/null; then
    echo "controller-gen not found, installing..."
    go install sigs.k8s.io/controller-tools/cmd/controller-gen@v0.13.0
else
    echo "controller-gen found"
fi

controller-gen crd paths=./pkg/apis/... output:crd:dir=./manifests/crd