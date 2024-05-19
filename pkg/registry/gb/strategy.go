package gb

import (
	"context"
	"fmt"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/apiserver/pkg/storage"
	"k8s.io/apiserver/pkg/storage/names"
	"my.domain/guestbook/apis/apps/v1alpha1"
)

type Strategy interface {
	rest.RESTCreateStrategy
	rest.RESTUpdateStrategy
}

type GuestBookStrategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}

// PrepareForCreate RESTCreateStrategy
func (g GuestBookStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
	//TODO implement me
	panic("implement me")
}

// Validate RESTCreateStrategy
func (g GuestBookStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	fmt.Println("Validate")
	return nil
}

// WarningsOnCreate RESTCreateStrategy
func (g GuestBookStrategy) WarningsOnCreate(ctx context.Context, obj runtime.Object) []string {
	fmt.Println("WarningsOnCreate")
	return nil
}

// NamespaceScoped RESTUpdateStrategy
func (g GuestBookStrategy) NamespaceScoped() bool {
	return false
}

// AllowCreateOnUpdate RESTUpdateStrategy
func (g GuestBookStrategy) AllowCreateOnUpdate() bool {
	//TODO implement me
	panic("implement me")
}

// PrepareForUpdate RESTUpdateStrategy
func (g GuestBookStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
	//TODO implement me
	panic("implement me")
}

// ValidateUpdate RESTUpdateStrategy
func (g GuestBookStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	//TODO implement me
	panic("implement me")
}

// WarningsOnUpdate RESTUpdateStrategy
func (g GuestBookStrategy) WarningsOnUpdate(ctx context.Context, obj, old runtime.Object) []string {
	//TODO implement me
	panic("implement me")
}

// Canonicalize RESTUpdateStrategy
func (g GuestBookStrategy) Canonicalize(obj runtime.Object) {
	//TODO implement me
	panic("implement me")
}

// AllowUnconditionalUpdate RESTUpdateStrategy
func (g GuestBookStrategy) AllowUnconditionalUpdate() bool {
	//TODO implement me
	panic("implement me")
}

func NewStrategy(typer runtime.ObjectTyper) Strategy {
	return &GuestBookStrategy{typer, names.SimpleNameGenerator}
}

func GetAttrs(obj runtime.Object) (labels.Set, fields.Set, error) {
	object, ok := obj.(*v1alpha1.GuestBook)
	if !ok {
		return nil, nil, fmt.Errorf("the object isn't a GuestBook")
	}
	fs := generic.ObjectMetaFieldsSet(&object.ObjectMeta, true)
	return object.ObjectMeta.Labels, fs, nil
}

func MatchGuestBook(label labels.Selector, field fields.Selector) storage.SelectionPredicate {
	return storage.SelectionPredicate{
		Label:    label,
		Field:    field,
		GetAttrs: GetAttrs,
	}
}
