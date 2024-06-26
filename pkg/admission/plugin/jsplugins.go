package plugin

import (
	"context"
	"fmt"
	"io"
	"my.domain/guestbook/apis/apps/v1alpha1"

	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apiserver/pkg/admission"
	informers "my.domain/guestbook/pkg/client/informers/externalversions"
	listers "my.domain/guestbook/pkg/client/listers/apps/v1alpha1"
)

type GuestBookPlugin struct {
	*admission.Handler
	jsLister listers.GuestBookLister
}

func Register(plugin *admission.Plugins) {
	plugin.Register("GuestBook", func(config io.Reader) (admission.Interface, error) {
		return New()
	})
}

func New() (*GuestBookPlugin, error) {
	var _ admission.Interface = GuestBookPlugin{}
	return &GuestBookPlugin{
		Handler: admission.NewHandler(admission.Create),
	}, nil
}

// Validate 有了validate方法就实现了admission.ValidationInterface，从而在validating阶段被调用
func (jsp *GuestBookPlugin) Validate(ctx context.Context, a admission.Attributes, _ admission.ObjectInterfaces) error {
	if a.GetKind().GroupKind() != v1alpha1.Kind("GuestBook") { //所有object的valid都会进来，所以我们要验一下是不是该关心的
		return nil
	}

	if !jsp.WaitForReady() { // 例如informer还没有把远程信息sync到本地
		return admission.NewForbidden(a, fmt.Errorf("the plugin isn't ready for handling request"))
	}

	// 下面就可以进行我们期望的校验了
	// 区别于registry部分strategy中的valid strategy，此处的校验更多是多实体之间的关联正确性，而不是单个jenkins service内容的正确
	// 例如，我们规定整个系统中只能存在10 个GuestBook对象，多了不行，就可以在这里做检查
	existedGuestBooks, err := jsp.jsLister.List(labels.Everything())
	if err != nil {
		return admission.NewForbidden(a, fmt.Errorf("the plugin encounter internal error during retrieve jenkins service objects from api server"))
	}
	if len(existedGuestBooks) >= 10 {
		return admission.NewForbidden(a, fmt.Errorf("too many service instances exist, %d exist but max is 10", len(existedGuestBooks)))
	}

	return nil

}

// SetInformerFactory 有了这个方法，plugin就实现了WantsCicdInformerFactory接口，可以获取到cicd informer了
func (jsp *GuestBookPlugin) SetInformerFactory(f informers.SharedInformerFactory) {
	jsp.jsLister = f.My().V1alpha1().GuestBooks().Lister()
	jsp.SetReadyFunc(f.My().V1alpha1().GuestBooks().Informer().HasSynced)
}
