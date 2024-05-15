package scheme

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	apsv1 "my.domain/guestbook/apis/apps/v1"
)

var (
	// Scheme contains the types needed by the resource metrics API.
	Scheme = runtime.NewScheme()
	// Codecs is sources.list codec factory for serving the resource metrics API.
	Codecs = serializer.NewCodecFactory(Scheme)
)

func init() {
	apsv1.AddToScheme(Scheme)
	unversioned := schema.GroupVersion{Group: "", Version: "v1"}
	metav1.AddToGroupVersion(Scheme, unversioned)
}
