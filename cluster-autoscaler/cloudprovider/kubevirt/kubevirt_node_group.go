package kubevirt

import (
	"fmt"
	"github.com/golang/glog"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider"
	"k8s.io/client-go/rest"
	"k8s.io/kubernetes/plugin/pkg/scheduler/schedulercache"
	virtv1 "kubevirt.io/kubevirt/pkg/api/v1"
)

var VirtualMachineReplicaSetResource = "virtualmachinereplicasets"

type ReplicaSetNodeGroup struct {
	name      string
	namespace string
	maxSize   int
	minSize   int
	client    *rest.RESTClient
	selector  labels.Selector
	template  *virtv1.VirtualMachineReplicaSet
}

func (r *ReplicaSetNodeGroup) getReplicaSet() (*virtv1.VirtualMachineReplicaSet, error) {
	rs := &virtv1.VirtualMachineReplicaSet{}
	err := r.client.Get().Namespace(r.namespace).Name(r.name).Resource(VirtualMachineReplicaSetResource).Do().Into(rs)
	return rs, err
}

func (r *ReplicaSetNodeGroup) scaleReplicaSet(replicas int) error {
	requestBody := fmt.Sprintf("[{\"op\":\"replace\",\"path\":\"/spec/replicas\",\"value\":%d}]", replicas)
	_, err := r.client.Patch(types.JSONPatchType).Namespace(r.namespace).Name(r.name).Resource(VirtualMachineReplicaSetResource).Body([]byte(requestBody)).Do().Get()
	return err
}

func (r *ReplicaSetNodeGroup) PauseReplicaSet() error {
	requestBody := "[{\"op\":\"replace\",\"path\":\"/spec/pause\",\"value\":\"true\"}]"
	_, err := r.client.Patch(types.JSONPatchType).Namespace(r.namespace).Name(r.name).Resource(VirtualMachineReplicaSetResource).Body([]byte(requestBody)).Do().Get()
	return err
}

func (r *ReplicaSetNodeGroup) ResumeReplicaSet() error {
	requestBody := "[{\"op\":\"replace\",\"path\":\"/spec/pause\",\"value\":\"false\"}]"
	_, err := r.client.Patch(types.JSONPatchType).Namespace(r.namespace).Name(r.name).Resource(VirtualMachineReplicaSetResource).Body([]byte(requestBody)).Do().Get()
	return err
}

func (r *ReplicaSetNodeGroup) deleteNode(node string) error {
	return r.client.Delete().Namespace(r.namespace).Namespace(node).Do().Error()
}

func (r *ReplicaSetNodeGroup) IsVirtualMachineOwned(vm *virtv1.VirtualMachine) bool {
	return r.selector.Matches(labels.Set(vm.ObjectMeta.Labels))
}

func (r *ReplicaSetNodeGroup) IsNodeOwned(node *v1.Node) (bool, error) {
	vm := &virtv1.VirtualMachine{}
	err := r.client.Get().Resource("virtualmachines").Name(node.ObjectMeta.Name).Namespace(r.namespace).Do().Into(vm)
	if err != nil {
		return false, err
	}

	return r.IsVirtualMachineOwned(vm), nil
}

// MaxSize returns maximum size of the node group.
func (r *ReplicaSetNodeGroup) MaxSize() int {
	return r.maxSize
}

// MinSize returns minimum size of the node group.
func (r *ReplicaSetNodeGroup) MinSize() int {
	return r.minSize
}

// TargetSize returns the current target size of the node group. It is possible that the
// number of nodes in Kubernetes is different at the moment but should be equal
// to Size() once everything stabilizes (new nodes finish startup and registration or
// removed nodes are deleted completely). Implementation required.
func (r *ReplicaSetNodeGroup) TargetSize() (int, error) {
	rs, err := r.getReplicaSet()
	if err != nil {
		return 0, err
	}
	return int(*rs.Spec.Replicas), nil
}

// IncreaseSize increases the size of the node group. To delete a node you need
// to explicitly name it and use DeleteNode. This function should wait until
// node group size is updated. Implementation required.
func (r *ReplicaSetNodeGroup) IncreaseSize(delta int) error {
	rs, err := r.getReplicaSet()
	if err != nil {
		return err
	}
	replicas := int(*rs.Spec.Replicas) + delta
	return r.scaleReplicaSet(replicas)
}

// DeleteNodes deletes nodes from this node group. Error is returned either on
// failure or if the given node doesn't belong to this node group. This function
// should wait until node group size is updated. Implementation required.
func (r *ReplicaSetNodeGroup) DeleteNodes(nodes []*v1.Node) error {
	rs, err := r.getReplicaSet()
	if err != nil {
		return err
	}
	for _, node := range nodes {

		owned, err := r.IsNodeOwned(node)

		if err != nil {
			return err
		}

		if !owned {
			return fmt.Errorf("Node %s is not owned by node group %s", node.ObjectMeta.Name, r.name)
		}
	}

	// Pause the vmrs
	err = r.PauseReplicaSet()
	if err != nil {
		return err
	}

	// Make sure that we resume the nodegroup. If we don't manage to resume it, panic
	defer func() {
		var err error
		for x := 0; x < 3; x++ {
			err = r.ResumeReplicaSet()
			if err != nil {
				glog.Errorf("Failed to resume paused nodegroup: %v", err)
				continue
			}
			break
		}
		if err != nil {
			glog.Fatalf("Failed to  resume paused nodegroup 3 times: %v", err)
		}
	}()

	// TODO: Wait for the controller to report that it is paused, to avoid race conditions wit node deletions

	for _, node := range nodes {
		err := r.deleteNode(node.ObjectMeta.Name)
		if err != nil {
			return err
		}
	}

	err = r.scaleReplicaSet(int(*rs.Spec.Replicas) - len(nodes))
	if err != nil {
		return err
	}

	return nil
}

// DecreaseTargetSize decreases the target size of the node group. This function
// doesn't permit to delete any existing node and can be used only to reduce the
// request for new nodes that have not been yet fulfilled. Delta should be negative.
// It is assumed that cloud provider will not delete the existing nodes when there
// is an option to just decrease the target. Implementation required.
func (r *ReplicaSetNodeGroup) DecreaseTargetSize(delta int) error {
	rs, err := r.getReplicaSet()
	if err != nil {
		return err
	}
	replicas := int(*rs.Spec.Replicas) + delta
	// TODO add available nodes to the replicaset, to add an additional sort parameter for removing not-so-long ready nodes
	if replicas < int(rs.Status.ReadyReplicas) {
		return fmt.Errorf("Can't decrease the amount of requested virtual machines, since it would remove ready nodes.")
	}
	return r.scaleReplicaSet(replicas)
}

// Id returns an unique identifier of the node group.
func (r *ReplicaSetNodeGroup) Id() string {
	return r.namespace + "/" + r.name
}

// Debug returns a string containing all information regarding this node group.
func (r *ReplicaSetNodeGroup) Debug() string {
	return ""
}

// Nodes returns a list of all nodes that belong to this node group.
func (r *ReplicaSetNodeGroup) Nodes() ([]string, error) {
	l := &virtv1.VirtualMachineList{}
	err := r.client.Get().Namespace(r.namespace).Resource("virtualmachines").Param("label-selector", r.template.Spec.Selector.String()).Do().Into(l)
	if err != nil {
		return nil, err
	}
	nodes := []string{}
	for _, obj := range l.Items {
		nodes = append(nodes, obj.ObjectMeta.Name)
	}
	return nodes, nil
}

// TemplateNodeInfo returns a schedulercache.NodeInfo structure of an empty
// (as if just started) node. This will be used in scale-up simulations to
// predict what would a new node look like if a node group was expanded. The returned
// NodeInfo is expected to have a fully populated Node object, with all of the labels,
// capacity and allocatable information as well as all pods that are started on
// the node by default, using manifest (most likely only kube-proxy). Implementation optional.
func (r *ReplicaSetNodeGroup) TemplateNodeInfo() (*schedulercache.NodeInfo, error) {

	node := &v1.Node{Status: v1.NodeStatus{Capacity: v1.ResourceList{}}}

	var cpuCores uint32 = 1
	if r.template.Spec.Template.Spec.Domain.CPU != nil {
		cpuCores = r.template.Spec.Template.Spec.Domain.CPU.Cores
	}

	node.Status.Capacity[v1.ResourceCPU] = *resource.NewQuantity(int64(cpuCores), "m")
	node.Status.Capacity[v1.ResourceMemory] = r.template.Spec.Template.Spec.Domain.Resources.Requests[v1.ResourceMemory]

	var err error
	if val, ok := r.template.Annotations["kubevirt.io/resourceStorage"]; ok {
		node.Status.Capacity[v1.ResourceStorage], err = resource.ParseQuantity(val)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("No 'kubevirt.io/resourceStorage' annotation found.")
	}

	if val, ok := r.template.Annotations["kubevirt.io/resourcePods"]; ok {
		node.Status.Capacity[v1.ResourcePods], err = resource.ParseQuantity(val)
		if err != nil {
			return nil, err
		}
	} else {
		node.Status.Capacity[v1.ResourcePods] = resource.MustParse("110")
	}

	node.Status.Conditions = cloudprovider.BuildReadyConditions()

	// TODO Make capacity and allocatable configurable
	node.Status.Allocatable = node.Status.Capacity

	nodeInfo := schedulercache.NewNodeInfo()
	nodeInfo.SetNode(node)

	return nodeInfo, nil
}

// Exist checks if the node group really exists on the cloud provider side. Allows to tell the
// theoretical node group from the real one. Implementation required.
func (r *ReplicaSetNodeGroup) Exist() bool {
	return true
}

// Create creates the node group on the cloud provider side. Implementation optional.
func (*ReplicaSetNodeGroup) Create() error {
	return nil
}

// Delete deletes the node group on the cloud provider side.
// This will be executed only for autoprovisioned node groups, once their size drops to 0.
// Implementation optional.
func (*ReplicaSetNodeGroup) Delete() error {
	return nil
}

// Autoprovisioned returns true if the node group is autoprovisioned. An autoprovisioned group
// was created by CA and can be deleted when scaled to 0.
func (*ReplicaSetNodeGroup) Autoprovisioned() bool {
	return false
}
