package kubevirt

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider"
	"io"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/apimachinery/pkg/runtime/serializer"
)

const (
	ProviderName = "kubevirt"
	NodeGroupMaxSize = 1000
	NodeGroupMinSize = 0
)

type KubeVirtManager interface {
	GetClient() *rest.RESTClient
}

type kubeVirtManger struct {
	Client *rest.RESTClient
}

func CreateManager(configReader io.Reader, clusterName string, rest *rest.RESTClient) (KubeVirtManager, error) {
	return &kubeVirtManger{rest}, nil
}

func BuildKubeVirtProvider(manager KubeVirtManager, discoveryOpts cloudprovider.NodeGroupDiscoveryOptions) (*KubeVirtCloudProvider, error) {
	return &KubeVirtCloudProvider{[]string{"testreplicaset"}, manager.GetClient()}, nil
}


func GetKubevirtClientFromFlags(master string, kubeconfig string) (*rest.RESTClient, error) {
	config, err := clientcmd.BuildConfigFromFlags(master, kubeconfig)
	if err != nil {
		return nil, err
	}

	config.GroupVersion = &GroupVersion
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}
	config.APIPath = "/apis"
	config.ContentType = runtime.ContentTypeJSON

	return rest.RESTClientFor(config)
}

func (k *kubeVirtManger) GetClient() *rest.RESTClient {
	return k.Client
}
