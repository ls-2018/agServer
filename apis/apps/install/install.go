package install

import (
	"k8s.io/apimachinery/pkg/runtime"
	util "k8s.io/apimachinery/pkg/util/runtime"
	"my.domain/guestbook/apis/apps/v1alpha1"
)

func Install(scheme *runtime.Scheme) {
	util.Must(v1alpha1.AddToScheme(scheme))
	util.Must(scheme.SetVersionPriority(v1alpha1.SchemeGroupVersion))
}
