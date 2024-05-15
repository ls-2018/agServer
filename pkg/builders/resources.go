package builders

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apsv1 "my.domain/guestbook/apis/apps/v1"
)

const applyKey = "kubectl.kubernetes.io/last-applied-configuration"

func ApiResourceList() metav1.APIResourceList {
	apiList := metav1.APIResourceList{
		GroupVersion: apsv1.SchemeGroupVersion.String(),
		APIResources: []metav1.APIResource{
			{
				Name:         "",
				SingularName: "myingress",
				Kind:         "MyIngress",
				ShortNames:   []string{"mi"},
				Namespaced:   true,
				Verbs:        []string{"get", "list", "create", "watch"},
			},
		},
	}
	apiList.APIVersion = "v1"
	apiList.Kind = "APIResourceList"
	return apiList
}
