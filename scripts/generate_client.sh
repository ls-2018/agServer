#!/usr/bin/env bash

go mod vendor
retVal=$?
if [ $retVal -ne 0 ]; then
    exit $retVal
fi

set -x
TMP_DIR=$(mktemp -d)
mkdir -p "${TMP_DIR}"/src/my.domain/guestbook/pkg/client
cp -r ./{apis,hack,vendor,go.mod,.git} "${TMP_DIR}"/src/my.domain/guestbook/

chmod +x "${TMP_DIR}"/src/my.domain/guestbook/vendor/k8s.io/code-generator/generate-internal-groups.sh
echo "tmp_dir: ${TMP_DIR}"

(cd "${TMP_DIR}"/src/my.domain/guestbook; \
    GOPATH=${TMP_DIR} GO111MODULE=off /bin/bash vendor/k8s.io/code-generator/generate-groups.sh client,deepcopy,informer,lister \
    my.domain/guestbook/pkg/client my.domain/guestbook/apis "apps:v1" -h ./hack/boilerplate.go.txt)

rm -rf ./pkg/client/{clientset,informers,listers}
mv "${TMP_DIR}"/src/my.domain/guestbook/pkg/client/* ./pkg/client

rm -rf vendor