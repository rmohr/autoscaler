package kubevirt

import (
	"fmt"
	"github.com/golang/glog"
	"gopkg.in/gcfg.v1"
	"io"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"strings"
	virtv1 "kubevirt.io/kubevirt/pkg/api/v1"
)

const (
	ProviderName     = "kubevirt"
	NodeGroupMaxSize = 1000
	NodeGroupMinSize = 0
)

type KubeVirtManager interface {
	GetClient() *rest.RESTClient
	GetAllReplicaSets(labels *v1.LabelSelector) ([]virtv1.VirtualMachineReplicaSet, error)
}

type kubeVirtManger struct {
	Client *rest.RESTClient
}

type KubeVirtConfig struct {
	Cluster struct {
		Kubeconfig string
		Master     string
	}
}

func CreateManager(configReader io.Reader, _ string) (KubeVirtManager, error) {
	var client *rest.RESTClient
	config := KubeVirtConfig{}
	var err error
	if configReader == nil {

	} else {
		err = gcfg.ReadInto(&config, configReader)
		if err != nil {
			return nil, fmt.Errorf("Failed to read provided configuration: %v", err)
		}
	}
	client, err = GetKubevirtClientFromFlags(config.Cluster.Master, config.Cluster.Kubeconfig)
	if err != nil {
		glog.Fatalf("Failed to create KubeVirt REST client: %v", err)
	}
	return &kubeVirtManger{client}, nil
}

func BuildKubeVirtProvider(manager KubeVirtManager, discoveryOpts cloudprovider.NodeGroupDiscoveryOptions) (*KubeVirtCloudProvider, error) {
	if err := discoveryOpts.Validate(); err != nil {
		return nil, err
	}
	nodeGroups := []*ReplicaSetNodeGroup{}

	// autodiscovery: kubevirt:labels=kubevirt.io/autoscaler
	if discoveryOpts.AutoDiscoverySpecified() {
		selector, err := v1.ParseToLabelSelector(strings.TrimPrefix(discoveryOpts.NodeGroupAutoDiscoverySpec, "kubevirt:labels="))
		if err != nil {
			return nil, fmt.Errorf("Error parsing auto-discovery labels: %v", err)
		}

		rsList, err := manager.GetAllReplicaSets(selector)
		for _, rs := range rsList {
			glog.Infof("Detected node group %s/%s", rs.ObjectMeta.Namespace, rs.ObjectMeta.Name)
			nodeGroup, err := NodeGroupFromReplicaSet(&rs, manager.GetClient())
			if err != nil {
				return nil, fmt.Errorf("Creating a NodeGroup out of a vmrs failed: %v", err)
			}
			// Make sure that on a fresh start we unpause the vmrs.
			if rs.Spec.Paused {
				err = nodeGroup.ResumeReplicaSet()
				if err != nil {
					return nil, fmt.Errorf("Resuming node group %s failed: %v ", nodeGroup.Id(), err)
				}
			}

			nodeGroups = append(nodeGroups, nodeGroup)
		}
	} else {
		return nil, fmt.Errorf("No discovery option specified.")
	}

	return &KubeVirtCloudProvider{nodeGroups}, nil
}

func GetKubevirtClientFromFlags(master string, kubeconfig string) (*rest.RESTClient, error) {
	config, err := clientcmd.BuildConfigFromFlags(master, kubeconfig)
	if err != nil {
		return nil, err
	}

	config.GroupVersion = &virtv1.GroupVersion
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}
	config.APIPath = "/apis"
	config.ContentType = runtime.ContentTypeJSON

	return rest.RESTClientFor(config)
}

func (k *kubeVirtManger) GetClient() *rest.RESTClient {
	return k.Client
}

func (k *kubeVirtManger) GetAllReplicaSets(labels *v1.LabelSelector) ([]virtv1.VirtualMachineReplicaSet, error) {
	list := &virtv1.VirtualMachineReplicaSetList{}
	err := k.Client.Get().Resource(VirtualMachineReplicaSetResource).Param("label-selector", labels.String()).Do().Into(list)
	if err != nil {
		return nil, err
	}
	return list.Items, nil
}
