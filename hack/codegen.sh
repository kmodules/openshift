#!/bin/bash

set -x

GOPATH=$(go env GOPATH)
PACKAGE_NAME=kmodules.xyz/openshift
REPO_ROOT="$GOPATH/src/$PACKAGE_NAME"
DOCKER_REPO_ROOT="/go/src/$PACKAGE_NAME"
DOCKER_CODEGEN_PKG="/go/src/k8s.io/code-generator"
apiGroups=(apps/v1 security/v1)

pushd $REPO_ROOT

docker run --rm -ti -u $(id -u):$(id -g) \
  -v "$REPO_ROOT":"$DOCKER_REPO_ROOT" \
  -w "$DOCKER_REPO_ROOT" \
  appscode/gengo:release-1.14 "$DOCKER_CODEGEN_PKG"/generate-groups.sh all \
  kmodules.xyz/openshift/client \
  kmodules.xyz/openshift/apis \
  "apps:v1 security:v1" \
  --go-header-file "$DOCKER_REPO_ROOT/hack/empty.txt"

popd
