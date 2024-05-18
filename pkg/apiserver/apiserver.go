package apiserver

import (
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	clientgorest "k8s.io/client-go/rest"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"
	"my.domain/guestbook/apis/apps/v1alpha1"
	jsclient "my.domain/guestbook/pkg/client/clientset/versioned"
	informerfactory "my.domain/guestbook/pkg/client/informers/externalversions"
	"my.domain/guestbook/pkg/controller/guestbook"
	cicdregistry "my.domain/guestbook/pkg/registry"
	jsstorage "my.domain/guestbook/pkg/registry/gb"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/apimachinery/pkg/version"
	"k8s.io/apiserver/pkg/registry/rest"
	genericapiserver "k8s.io/apiserver/pkg/server"
	"my.domain/guestbook/apis/apps/install"
)

var (
	Scheme = runtime.NewScheme()
	Codecs = serializer.NewCodecFactory(Scheme)
)

// 如下方法需要更新至相应phase，开始漏掉了
func init() {
	install.Install(Scheme)
	metav1.AddToGroupVersion(Scheme, schema.GroupVersion{Version: "v1"})
	unversioned := schema.GroupVersion{Group: "", Version: "v1"}
	Scheme.AddUnversionedTypes(
		unversioned,
		&metav1.Status{},
		&metav1.APIVersions{},
		&metav1.APIGroupList{},
		&metav1.APIGroup{},
		&metav1.APIResourceList{},
	)
}

// 如下环节制作Server的Config
type Config struct {
	GenericConfig *genericapiserver.RecommendedConfig
	KubeConfig    string

	// ExtraConfig   ExtraConfig // 如果有自己需要的config的话，可以扩展field
}

type completedConfig struct {
	GenericConfig genericapiserver.CompletedConfig
	KubeConfig    string
}

// 完善后的config
type CompletedConfig struct {
	*completedConfig
}

type CicdServer struct {
	GenericAPIServer *genericapiserver.GenericAPIServer
}

// 完善初始的config
func (cfg *Config) Complete() CompletedConfig {
	config := completedConfig{
		GenericConfig: cfg.GenericConfig.Complete(),
		KubeConfig:    cfg.KubeConfig,
	}
	config.GenericConfig.Version = &version.Info{
		Major: "1",
		Minor: "0",
	}
	return CompletedConfig{&config}
}

// 有了这个方法，完善后的config就可以制作server的instance了
func (ccfg completedConfig) NewServer() (*CicdServer, error) {
	genericServer, err := ccfg.GenericConfig.New("my.domain/guestbook", genericapiserver.NewEmptyDelegate())
	if err != nil {
		return nil, err
	}

	server := &CicdServer{
		GenericAPIServer: genericServer,
	}

	//重点是把我们各个版本的api object都注入到server中去，开始
	apiGroupInfo := genericapiserver.NewDefaultAPIGroupInfo(
		v1alpha1.GroupName,
		Scheme,
		metav1.ParameterCodec,
		Codecs,
	)
	v1alphastorage := map[string]rest.Storage{}
	v1alphastorage["guestbook"] = cicdregistry.RESTWithErrorHandler(jsstorage.NewREST(Scheme, ccfg.GenericConfig.RESTOptionsGetter))
	apiGroupInfo.VersionedResourcesStorageMap[v1alpha1.SchemeGroupVersion.Version] = v1alphastorage

	if err = server.GenericAPIServer.InstallAPIGroup(&apiGroupInfo); err != nil {
		return nil, err
	}
	var config *restclient.Config
	kubeConfig := ccfg.KubeConfig
	// fallback to kubeConfig
	if envVar := os.Getenv("KUBECONFIG"); len(envVar) > 0 {
		kubeConfig = envVar
	}
	if kubeConfig != "" {
		config, err = clientcmd.BuildConfigFromFlags("", kubeConfig)
	} else {
		config, err = clientgorest.InClusterConfig()

	}
	// 创建相关控制器
	if err != nil {
		klog.ErrorS(err, "The kubeConfig cannot be loaded: %v\n")
		panic(err)
	}
	coreAPIClientSet, err := kubernetes.NewForConfig(config)

	client, err := jsclient.NewForConfig(genericServer.LoopbackClientConfig)
	if err != nil {
		klog.Error("Can't create client set for CICD API Server during creating server")
	}
	jsInformerFactory := informerfactory.NewSharedInformerFactory(client, 0)
	controller := guestbook.NewGuestBookController(jsInformerFactory.My().V1alpha1().GuestBooks(), coreAPIClientSet)

	// 向Server启动钩子中注入控制器启动函数
	genericServer.AddPostStartHookOrDie("my.domain/guestbook-controller", func(ctx genericapiserver.PostStartHookContext) error {
		ctxjs := wait.ContextForChannel(ctx.StopCh)
		go func() {
			controller.Run(ctxjs, 2)
		}()
		return nil
	})
	genericServer.AddPostStartHookOrDie("my.domain/guestbook-informer", func(context genericapiserver.PostStartHookContext) error {
		jsInformerFactory.Start(context.StopCh)
		return nil
	})
	return server, nil
}
