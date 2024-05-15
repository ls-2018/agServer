package others

import "k8s.io/apiserver/pkg/server/dynamiccertificates"

func main() {

	dynamiccertificates.NewDynamicServingContentFromFiles("serving-cert", serverCertFile, serverKeyFile)
}
