// Copyright 2018 The Kubernetes Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain sources.list copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0

//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package storage

import (
	"context"
	"fmt"
	"k8s.io/apimachinery/pkg/apis/meta/internalversion"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/apiserver/pkg/storage/names"
	apsv1 "my.domain/guestbook/apis/apps/v1"
	"my.domain/guestbook/pkg/watcher"
	"net/http"
	"sync"
)

var Ng = names.SimpleNameGenerator

type storage struct {
	rest.TableConvertor
	SingularName string
	isNamespaced bool
	muWatchers   sync.RWMutex
	newFunc      func() runtime.Object
	newListFunc  func() runtime.Object
}

func (s *storage) Create(ctx context.Context, name string, obj runtime.Object, createValidation rest.ValidateObjectFunc, options *metav1.CreateOptions) (runtime.Object, error) {
	return &apsv1.GuestBook{}, nil
}

func (s *storage) Get(ctx context.Context, name string, options runtime.Object) (runtime.Object, error) {
	return &apsv1.GuestBook{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Spec: apsv1.GuestBookSpec{
			Name: name,
		},
	}, nil
}

func (s *storage) NewGetOptions() (runtime.Object, bool, string) {
	//TODO implement me
	panic("implement me")
}

func (s *storage) Connect(ctx context.Context, id string, options runtime.Object, r rest.Responder) (http.Handler, error) {
	//TODO implement me
	panic("implement me")
}

func (s *storage) NewConnectOptions() (runtime.Object, bool, string) {
	//TODO implement me
	panic("implement me")
}

func (s *storage) ConnectMethods() []string {
	//TODO implement me
	panic("implement me")
}

func (s *storage) StorageVersion() runtime.GroupVersioner {
	//TODO implement me
	panic("implement me")
}

func (s *storage) AcceptsGroupVersion(gv schema.GroupVersion) bool {
	//TODO implement me
	panic("implement me")
}

func (s *storage) ProducesMIMETypes(verb string) []string {
	//TODO implement me
	panic("implement me")
}

func (s *storage) ProducesObject(verb string) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *storage) GetSingularName() string {
	return s.SingularName
}

func (s *storage) GenerateName(base string) string {
	fmt.Println("GenerateName")
	return Ng.GenerateName(base)
}

func (s *storage) ObjectKinds(object runtime.Object) ([]schema.GroupVersionKind, bool, error) {
	fmt.Println("ObjectKinds")
	return nil, false, nil
}

func (s *storage) Recognizes(gvk schema.GroupVersionKind) bool {
	fmt.Println("Recognizes")
	return true
}

func (s *storage) NamespaceScoped() bool {
	fmt.Println("NamespaceScoped")
	return false
}

// Canonicalize 规范化，可用于修改对象，或对象类型检查
// 调用是在 验证已成功，但在对象被持久化之前
func (s *storage) Canonicalize(obj runtime.Object) {
	fmt.Println("Canonicalize")
}

func (s *storage) AllowUnconditionalUpdate() bool {
	fmt.Println("AllowUnconditionalUpdate")
	return true
}

func (s *storage) NewList() runtime.Object {
	fmt.Println("NewList")
	return &apsv1.GuestBookList{}
}

func (s *storage) List(ctx context.Context, options *internalversion.ListOptions) (runtime.Object, error) {
	return &apsv1.GuestBookList{Items: make([]apsv1.GuestBook, 0)}, nil
}

func (s *storage) New() runtime.Object {
	return &apsv1.GuestBook{}
}

func (s *storage) Update(ctx context.Context, name string, objInfo rest.UpdatedObjectInfo, createValidation rest.ValidateObjectFunc, updateValidation rest.ValidateObjectUpdateFunc, forceAllowCreate bool, options *metav1.UpdateOptions) (runtime.Object, bool, error) {
	return &apsv1.GuestBook{ObjectMeta: metav1.ObjectMeta{
		Name: name,
	}}, true, nil
}

func (s *storage) Delete(ctx context.Context, name string, deleteValidation rest.ValidateObjectFunc, options *metav1.DeleteOptions) (runtime.Object, bool, error) {
	return nil, true, nil
}

func (s *storage) DeleteCollection(ctx context.Context, deleteValidation rest.ValidateObjectFunc, options *metav1.DeleteOptions, listOptions *internalversion.ListOptions) (runtime.Object, error) {
	fmt.Println("DeleteCollection")
	return &apsv1.GuestBookList{}, nil
}

func (s *storage) Watch(ctx context.Context, options *internalversion.ListOptions) (watch.Interface, error) {
	return &watcher.Watch{}, nil
}

func (s *storage) Destroy() {
	fmt.Println("Destroy")
}
func (s *storage) ConvertToTable(ctx context.Context, object runtime.Object, tableOptions runtime.Object) (*metav1.Table, error) {
	return rest.NewDefaultTableConvertor(apsv1.Resource("guestbook")).ConvertToTable(ctx, object, tableOptions)
}

var _ Storage = &storage{}

func NewStorage(SingularName string) Storage {
	return &storage{SingularName: SingularName}
}
