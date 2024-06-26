package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const GroupName = "my.domain/guestbook"

var SchemeGroupVersion = schema.GroupVersion{Group: GroupName, Version: "v1alpha"}

func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

func Kind(kind string) schema.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

// var的定义和internal version的register中基本类似，
// 只是创建Builder时多了一个中间产物local scheme builder，local builder会在包括生成代码的init中去使用
var (
	SchemeBuilder      runtime.SchemeBuilder
	localSchemeBuilder = &SchemeBuilder
	AddToScheme        = localSchemeBuilder.AddToScheme
)

func init() {
	//这里去注册本version的类型，以及它们向internal version的转换函数
	localSchemeBuilder.Register(addKnownTypes)
}

// 被SchemeBuilder调用，从而把自己知道的Object（Type）注册到scheme中
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(
		SchemeGroupVersion,
		&GuestBook{},
		&GuestBookList{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
