package admission

import (
	"k8s.io/apiserver/pkg/admission"
	informers "my.domain/guestbook/pkg/client/informers/externalversions"
)

// 需要admission plugin去实现这个接口，从而保证可以接收informerfactory；
type WantsCicdInformerFactory interface {
	SetInformerFactory(informers.SharedInformerFactory)
}

type CicdInformerPluginInitializer struct {
	informers informers.SharedInformerFactory
}

func (i CicdInformerPluginInitializer) Initialize(plugin admission.Interface) {
	if wants, ok := plugin.(WantsCicdInformerFactory); ok { //如果目标plugin通过实现接口，声明需要cicd informer，那么我们就给它
		wants.SetInformerFactory(i.informers)
	}
}

// NewGBInformerPluginInitializer server启动时在config阶段被调用，从而把informer交给plugin
func NewGBInformerPluginInitializer(informers informers.SharedInformerFactory) CicdInformerPluginInitializer {
	return CicdInformerPluginInitializer{
		informers: informers,
	}
}
