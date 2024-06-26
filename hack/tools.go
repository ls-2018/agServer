//go:build tools
// +build tools

// This package imports things required by build scripts, to force `go mod` to see them as dependencies
package hack

import (
	_ "github.com/go-bindata/go-bindata/go-bindata"
	_ "k8s.io/code-generator"
	_ "k8s.io/kube-openapi/cmd/openapi-gen"
)
