#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

if ! go list -m k8s.io/code-generator > /dev/null 2>&1; then
    echo "k8s.io/code-generator not found, downloading..."
    go get -u k8s.io/code-generator@v0.28.3
else
    echo "k8s.io/code-generator found"
fi

REPO_ROOT=$(git rev-parse --show-toplevel)
cd "${REPO_ROOT}"

SCRIPT_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
CODEGEN_PKG=$(go list -m -f '{{.Dir}}' k8s.io/code-generator)

${CODEGEN_PKG}/generate-groups.sh "client,informer,lister" \
  github.com/kubebadges/kubebadges/pkg/generated \
  github.com/kubebadges/kubebadges/pkg/apis \
  kubebadges:v1 \
  --output-base "$REPO_ROOT" \
  --go-header-file ${REPO_ROOT}/hack/boilerplate.go.txt

${CODEGEN_PKG}/generate-groups.sh "deepcopy" \
  github.com/kubebadges/kubebadges/pkg/generated \
  github.com/kubebadges/kubebadges/pkg/apis \
  kubebadges:v1 \
  --output-base "$REPO_ROOT" \
  --go-header-file ${REPO_ROOT}/hack/boilerplate.go.txt

cp -r $REPO_ROOT/github.com/kubebadges/kubebadges/pkg/* $REPO_ROOT/pkg/
rm -rf $REPO_ROOT/github.com