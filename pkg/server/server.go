package server

import (
	"fmt"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	rest2 "k8s.io/apiserver/pkg/registry/rest"
	genericapiserver "k8s.io/apiserver/pkg/server"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	apsv1 "my.domain/guestbook/apis/apps/v1"
	"my.domain/guestbook/pkg/scheme"
	"my.domain/guestbook/pkg/storage"
	"net/http"
)

type Config struct {
	ApiServer    *genericapiserver.Config
	Rest         *rest.Config
	NodeSelector string
}

func NewServer(
	nodes cache.Controller,
	pods cache.Controller,
	apiServer *genericapiserver.GenericAPIServer) *Server {
	return &Server{
		nodes:            nodes,
		pods:             pods,
		GenericAPIServer: apiServer,
	}
}

func (c Config) Complete() (*Server, error) {

	podInformerFactory, err := runningPodMetadataInformer(c.Rest)
	if err != nil {
		return nil, err
	}
	podInformer := podInformerFactory.ForResource(corev1.SchemeGroupVersion.WithResource("pods"))
	informer, err := informerFactory(c.Rest)
	if err != nil {
		return nil, err
	}
	nodes := informer.Core().V1().Nodes()

	// Disable default metrics handler and create custom one
	c.ApiServer.EnableMetrics = false
	genericServer, err := c.ApiServer.Complete(nil).New("metrics-server", genericapiserver.NewEmptyDelegate())
	if err != nil {
		return nil, err
	}
	genericServer.Handler.NonGoRestfulMux.HandleFunc("/metrics", func(writer http.ResponseWriter, request *http.Request) {

	})

	{
		agi := genericapiserver.NewDefaultAPIGroupInfo(apsv1.SchemeGroupVersion.Group, scheme.Scheme, metav1.ParameterCodec, scheme.Codecs)

		agi.VersionedResourcesStorageMap[apsv1.SchemeGroupVersion.Version] = map[string]rest2.Storage{
			"guestbook": storage.NewStorage("guestbook"),
		}
		genericServer.InstallAPIGroup(&agi)

	}
	fmt.Println(genericServer.Handler.ListedPaths())
	s := NewServer(
		nodes.Informer(),
		podInformer.Informer(),
		genericServer,
	)
	return s, nil
}

// Server scrapes metrics and serves then using k8s api.
type Server struct {
	*genericapiserver.GenericAPIServer

	pods  cache.Controller
	nodes cache.Controller

	storage storage.Storage
}

func (s Server) RunUntil(stopCh <-chan struct{}) error {

	// Start informers
	go s.nodes.Run(stopCh)
	go s.pods.Run(stopCh)

	ok := cache.WaitForCacheSync(stopCh, s.nodes.HasSynced)
	if !ok {
		return nil
	}
	ok = cache.WaitForCacheSync(stopCh, s.pods.HasSynced)
	if !ok {
		return nil
	}

	return s.GenericAPIServer.PrepareRun().Run(stopCh)
}

// 海龙张
// jtthink
