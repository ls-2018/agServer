package options

import (
	"fmt"
	openapinamer "k8s.io/apiserver/pkg/endpoints/openapi"
	genericapiserver "k8s.io/apiserver/pkg/server"
	genericoptions "k8s.io/apiserver/pkg/server/options"
	"k8s.io/client-go/pkg/version"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"my.domain/guestbook/pkg/scheme"
	"my.domain/guestbook/pkg/server"

	apsv1 "my.domain/guestbook/apis/apps/v1"
	"net"
	"strings"
)

// const RemoteKubeConfig = "./resources/config"
const RemoteKubeConfig = "/Users/acejilam/.kube/koord"
const defaultEtcdPathPrefix = "/registry/myapi.jtthink.com"

type Options struct {
	*genericoptions.RecommendedOptions
	KubeConfig string
}

// NewOptions constructs sources.list new set of default options for metrics-server.
func NewOptions() *Options {

	rc := genericoptions.NewRecommendedOptions(
		defaultEtcdPathPrefix,
		scheme.Codecs.LegacyCodec(apsv1.SchemeGroupVersion), //JSON格式的编码器
	)
	{
		rc.SecureServing.BindPort = 8443
		rc.SecureServing.ServerCert = genericoptions.GeneratableKeyCert{
			CertDirectory: "./certs",
			PairName:      "apiserver",
		}
		//直接写死 省的烦
		rc.Etcd.StorageConfig.Transport.ServerList = []string{}
		//rc.CoreAPI.CoreAPIKubeconfigPath = RemoteKubeConfig
		//rc.Authentication.RemoteKubeConfigFile = RemoteKubeConfig
		//rc.Authorization.RemoteKubeConfigFile = RemoteKubeConfig
	}

	return &Options{
		RecommendedOptions: rc,
		//KubeConfig:         RemoteKubeConfig,
	}
}
func (o Options) ServerConfig() (*server.Config, error) {
	apiServer, err := o.ApiServerConfig()
	if err != nil {
		return nil, err
	}
	restConfig, err := o.restConfig()
	if err != nil {
		return nil, err
	}
	return &server.Config{
		ApiServer: apiServer,
		Rest:      restConfig,
	}, nil
}
func (o Options) ApiServerConfig() (*genericapiserver.Config, error) {
	if err := o.SecureServing.MaybeDefaultWithSelfSignedCerts("0.0.0.0", nil, []net.IP{net.ParseIP("127.0.0.1")}); err != nil {
		return nil, fmt.Errorf("error creating self-signed certificates: %v", err)
	}

	serverConfig := genericapiserver.NewConfig(scheme.Codecs)
	if err := o.SecureServing.ApplyTo(&serverConfig.SecureServing, &serverConfig.LoopbackClientConfig); err != nil {
		return nil, err
	}

	if err := o.Audit.ApplyTo(serverConfig); err != nil {
		return nil, err
	}

	versionGet := version.Get()
	serverConfig.Version = &versionGet
	// enable OpenAPI schemas
	serverConfig.OpenAPIConfig = genericapiserver.DefaultOpenAPIConfig(apsv1.GetOpenAPIDefinitions, openapinamer.NewDefinitionNamer(scheme.Scheme))
	serverConfig.OpenAPIV3Config = genericapiserver.DefaultOpenAPIV3Config(apsv1.GetOpenAPIDefinitions, openapinamer.NewDefinitionNamer(scheme.Scheme))
	serverConfig.OpenAPIConfig.Info.Title = "Kubernetes metrics-server"
	serverConfig.OpenAPIV3Config.Info.Title = "Kubernetes metrics-server"
	serverConfig.OpenAPIConfig.Info.Version = strings.Split(serverConfig.Version.String(), "-")[0] // TODO(directxman12): remove this once autosetting this doesn't require security definitions
	serverConfig.OpenAPIV3Config.Info.Version = strings.Split(serverConfig.Version.String(), "-")[0]

	return serverConfig, nil
}
func (o Options) restConfig() (*rest.Config, error) {
	var config *rest.Config
	var err error
	if len(o.KubeConfig) > 0 {
		loadingRules := &clientcmd.ClientConfigLoadingRules{ExplicitPath: o.KubeConfig}
		loader := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, &clientcmd.ConfigOverrides{})

		config, err = loader.ClientConfig()
	} else {
		config, err = rest.InClusterConfig()
	}
	if err != nil {
		return nil, fmt.Errorf("unable to construct lister client config: %v", err)
	}
	// Use protobufs for communication with apiserver
	config.ContentType = "application/vnd.kubernetes.protobuf"
	err = rest.SetKubernetesDefaults(config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
