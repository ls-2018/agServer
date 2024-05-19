package gb

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/generic"
	gRegistry "k8s.io/apiserver/pkg/registry/generic/registry"
	v1alpha1 "my.domain/guestbook/apis/apps/v1alpha1"
	"my.domain/guestbook/pkg/registry"
)

func NewREST(scheme *runtime.Scheme, optsGetter generic.RESTOptionsGetter) (*registry.REST, error) {
	strategy := NewStrategy(scheme)

	store := &gRegistry.Store{
		NewFunc:                   func() runtime.Object { return &v1alpha1.GuestBook{} },
		NewListFunc:               func() runtime.Object { return &v1alpha1.GuestBookList{} },
		PredicateFunc:             MatchGuestBook,
		DefaultQualifiedResource:  v1alpha1.Resource("guestbook"),
		SingularQualifiedResource: v1alpha1.Resource("guestbook"),

		CreateStrategy: strategy,
		UpdateStrategy: strategy,
		DeleteStrategy: strategy,

		//TableConvertor: rest.NewDefaultTableConvertor(v1alpha1.Resource("guestbook")),
		TableConvertor: &xx{},
	}
	options := &generic.StoreOptions{RESTOptions: optsGetter, AttrFunc: GetAttrs}
	if err := store.CompleteWithOptions(options); err != nil {
		return nil, err
	}
	return &registry.REST{Store: store}, nil
}

var (
	TableListColumns = []string{"Name", "Namespace", "Path", "Host"}
)

type xx struct {
}

func (x xx) ConvertToTable(ctx context.Context, obj runtime.Object, tableOptions runtime.Object) (*metav1.Table, error) {
	t := &metav1.Table{}
	t.Kind = "Table"
	t.APIVersion = "meta.k8s.io/v1"

	//设置表头
	th := make([]metav1.TableColumnDefinition, len(TableListColumns))
	for i, h := range TableListColumns {
		th[i] = metav1.TableColumnDefinition{Name: h, Type: "string"}
	}
	t.ColumnDefinitions = th                        //设置表头
	if v, ok := obj.(*v1alpha1.GuestBookList); ok { //代表取列表
		//设置 行  数据
		rows := make([]metav1.TableRow, len(v.Items))
		for i, item := range v.Items {
			rows[i] = metav1.TableRow{
				Cells: []interface{}{item.Name, item.Namespace, item.Spec.InstanceAmount},
			}
		}
		t.Rows = rows
	}
	if v, ok := obj.(*v1alpha1.GuestBook); ok {
		rows := make([]metav1.TableRow, 1)
		rows[0] = metav1.TableRow{
			Cells: []interface{}{v.Name, v.Namespace, v.Spec.InstanceAmount},
		}
		t.Rows = rows
	}
	return t, nil
}
