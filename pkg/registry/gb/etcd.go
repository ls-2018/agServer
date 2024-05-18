package gb

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/generic"
	gRegistry "k8s.io/apiserver/pkg/registry/generic/registry"
	"k8s.io/apiserver/pkg/registry/rest"
	v1alpha1 "my.domain/guestbook/apis/apps/v1alpha1"
	"my.domain/guestbook/pkg/registry"
)

func NewREST(scheme *runtime.Scheme, optsGetter generic.RESTOptionsGetter) (*registry.REST, error) {
	strategy := NewStrategy(scheme)

	store := &gRegistry.Store{
		NewFunc:                  func() runtime.Object { return &v1alpha1.GuestBook{} },
		NewListFunc:              func() runtime.Object { return &v1alpha1.GuestBookList{} },
		PredicateFunc:            MatchGuestBook,
		DefaultQualifiedResource: v1alpha1.Resource("guestbook"),

		CreateStrategy: strategy,
		UpdateStrategy: strategy,
		DeleteStrategy: strategy,

		TableConvertor: rest.NewDefaultTableConvertor(v1alpha1.Resource("guestbook")),
	}
	options := &generic.StoreOptions{RESTOptions: optsGetter, AttrFunc: GetAttrs}
	if err := store.CompleteWithOptions(options); err != nil {
		return nil, err
	}
	return &registry.REST{Store: store}, nil
}
