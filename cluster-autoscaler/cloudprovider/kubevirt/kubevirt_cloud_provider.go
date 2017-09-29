package kubevirt

import (
		apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider"
	"k8s.io/autoscaler/cluster-autoscaler/utils/errors"
	"k8s.io/client-go/rest"
)

type KubeVirtCloudProvider struct {
	nodeGroups []string
	client *rest.RESTClient
}

// Name returns name of the cloud provider.
func (*KubeVirtCloudProvider) Name() string {
	return "KubeVirt"

}

// NodeGroups returns all node groups configured for this cloud provider.
func (*KubeVirtCloudProvider) NodeGroups() []cloudprovider.NodeGroup {
	return nil
}

// NodeGroupForNode returns the node group for the given node, nil if the node
// should not be processed by cluster autoscaler, or non-nil error if such
// occurred. Must be implemented.
func (*KubeVirtCloudProvider) NodeGroupForNode(*apiv1.Node) (cloudprovider.NodeGroup, error) {
	// TODO ControllerRef can help us here, right now we can do a label check
	return nil, nil
}

// Pricing returns pricing model for this cloud provider or error if not available.
// Implementation optional.
func (*KubeVirtCloudProvider) Pricing() (cloudprovider.PricingModel, errors.AutoscalerError) {
	return nil, cloudprovider.ErrNotImplemented
}

// GetAvailableMachineTypes get all machine types that can be requested from the cloud provider.
// Implementation optional.
func (*KubeVirtCloudProvider) GetAvailableMachineTypes() ([]string, error) {
	return nil, nil
}

// NewNodeGroup builds a theoretical node group based on the node definition provided. The node group is not automatically
// created on the cloud provider side. The node group is not returned by NodeGroups() until it is created.
// Implementation optional.
func (*KubeVirtCloudProvider) NewNodeGroup(machineType string, labels map[string]string, extraResources map[string]resource.Quantity) (cloudprovider.NodeGroup, error) {
	return nil, cloudprovider.ErrNotImplemented
}
