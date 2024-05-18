#!/usr/bin/env bash

# Copyright 2017 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit
set -o nounset
set -o pipefail
set -ex
PROJECT="my.domain/guestbook"
find .|grep generated|grep -v vendor|xargs -I F rm -rf F  || echo "clean success"
PROJECT_ROOT="$(realpath $(dirname $0)/..)"

TMP_DIR=$(mktemp -d)
echo "tmp_dir: ${TMP_DIR}"

mkdir -p "${TMP_DIR}"/src/${PROJECT}
cp -r ./{apis,hack,vendor,go.mod,.git} "${TMP_DIR}"/src/${PROJECT}
cp -r ./vendor/* "${TMP_DIR}"/src/

export GO111MODULE=off

SCRIPT_ROOT="${TMP_DIR}"/src/${PROJECT}
(

cd ${SCRIPT_ROOT}

CODEGEN_PKG=${CODEGEN_PKG:-$(cd "${SCRIPT_ROOT}"; ls -d -1 ./vendor/k8s.io/code-generator 2>/dev/null || echo ../code-generator)}

source "${CODEGEN_PKG}/kube_codegen.sh"

export GOPATH=${TMP_DIR}

export KUBE_VERBOSE=9

kube::codegen::gen_helpers \
    --boilerplate "${SCRIPT_ROOT}/hack/boilerplate.go.txt" \
   "./apis"

kube::codegen::gen_openapi \
   --output-dir "./pkg/client/openapi" \
   --output-pkg "openapi" \
   --report-filename "${PROJECT_ROOT}/report/api_violations.report" \
   --update-report \
   --boilerplate "./hack/boilerplate.go.txt" \
   "${SCRIPT_ROOT}/apis"

kube::codegen::gen_client \
    --with-watch \
    --output-dir "${SCRIPT_ROOT}/pkg/client" \
    --output-pkg "${PROJECT}/pkg/client" \
    --boilerplate "${SCRIPT_ROOT}/hack/boilerplate.go.txt" \
    "./apis"
)

rm -rf ./pkg/client/{clientset,informers,listers,openapi}
cp -rf "${TMP_DIR}"/src/${PROJECT}/apis/* ${PROJECT_ROOT}/apis
cp -rf "${TMP_DIR}"/src/${PROJECT}/pkg/client/* ${PROJECT_ROOT}/pkg/client