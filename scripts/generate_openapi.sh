#!/usr/bin/env bash

go mod vendor
retVal=$?
if [ $retVal -ne 0 ]; then
    exit $retVal
fi

set -ex

TMP_DIR=$(mktemp -d)
mkdir -p "${TMP_DIR}"/src/my.domain/guestbook
cp -r ./{apis,hack,vendor} "${TMP_DIR}"/src/my.domain/guestbook/

(cd "${TMP_DIR}"/src/my.domain/guestbook; \
    GOPATH=${TMP_DIR} GO111MODULE=off go run vendor/k8s.io/kube-openapi/cmd/openapi-gen/openapi-gen.go \
    -O zz_generated.openapi \
    -i ./apis/apps/v1 \
    -i k8s.io/metrics/pkg/apis/metrics/v1beta1,k8s.io/apimachinery/pkg/apis/meta/v1,k8s.io/apimachinery/pkg/api/resource,k8s.io/apimachinery/pkg/version \
    -i k8s.io/apimachinery/pkg/runtime \
    -p my.domain/guestbook/apis/apps/v1 \
    -h ./hack/boilerplate.go.txt \
    --report-filename ./violation_exceptions.list)

cp -f "${TMP_DIR}"/src/my.domain/guestbook/apis/apps/v1/zz_generated.openapi.go ./apis/apps/v1
