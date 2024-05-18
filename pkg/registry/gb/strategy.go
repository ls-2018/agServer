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

func (GuestBookStrategy) AllowCreateOnUpdate() bool {
	return false
}
func (GuestBookStrategy) Canonicalize(obj runtime.Object) {

}
func (GuestBookStrategy) NamespaceScoped() bool {
	return true
}
func (GuestBookStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {

}

func (GuestBookStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	errs := field.ErrorList{} //承载发现的错误

	js := obj.(*v1alpha1.GuestBook)
	if js.Spec.InstanceAmount > 10 {
		errs = append(errs, field.TooMany(field.NewPath("spec").Key("instanceamount"), js.Spec.InstanceAmount, 10))
	}
	if len(errs) > 0 {
		return errs
	} else {
		return nil
	}
}
func (GuestBookStrategy) WarningsOnCreate(ctx context.Context, obj runtime.Object) []string {
	return []string{}
}

func (GuestBookStrategy) AllowUnconditionalUpdate() bool {
	return false
}
func (GuestBookStrategy) PrepareForUpdate(ctx context.Context, obj runtime.Object, old runtime.Object) {

}
func (GuestBookStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	return field.ErrorList{}
}
func (GuestBookStrategy) WarningsOnUpdate(ctx context.Context, obj runtime.Object, old runtime.Object) []string {
	return []string{}
}
