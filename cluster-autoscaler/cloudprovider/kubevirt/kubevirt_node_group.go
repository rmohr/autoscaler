package kubevirt

import (
	"k8s.io/api/core/v1"
	"k8s.io/client-go/rest"
	"k8s.io/kubernetes/plugin/pkg/scheduler/schedulercache"
	"fmt"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/api/resource"
	"strings"
)

type ReplicaSetNodeGroup struct {
	name      string
	namespace string
	maxSize   int
	minSize   int
	client    *rest.RESTClient
	selector  labels.Selector
	template *VMReplicaSetSpec
}

func (r *ReplicaSetNodeGroup) getReplicaSet() (*VirtualMachineReplicaSet, error) {
	var rs *VirtualMachineReplicaSet
	err := r.client.Get().Namespace(r.namespace).Name(r.name).Resource(VirtualMachineReplicaSetResource).Do().Into(rs)
	return rs, err
}

func (r *ReplicaSetNodeGroup) scaleReplicaSet(replicas int) error {
	requestBody := fmt.Sprintf("[{\"op\":\"replace\",\"path\":\"/spec/replicas\",\"value\":%d}]", replicas)
	_, err := r.client.Patch(types.JSONPatchType).Namespace(r.namespace).Name(r.name).Resource(VirtualMachineReplicaSetResource).Body(requestBody).Do().Get()
	return err
}

func (r *ReplicaSetNodeGroup) pauseReplicaSet() error {
	requestBody := "[{\"op\":\"replace\",\"path\":\"/spec/pause\",\"value\":\"true\"}]"
	_, err := r.client.Patch(types.JSONPatchType).Namespace(r.namespace).Name(r.name).Resource(VirtualMachineReplicaSetResource).Body(requestBody).Do().Get()
	return err
}

func (r *ReplicaSetNodeGroup) resumeReplicaSet() error {
	requestBody := "[{\"op\":\"replace\",\"path\":\"/spec/pause\",\"value\":\"false\"}]"
	_, err := r.client.Patch(types.JSONPatchType).Namespace(r.namespace).Name(r.name).Resource(VirtualMachineReplicaSetResource).Body(requestBody).Do().Get()
	return err
}

func (r *ReplicaSetNodeGroup) deleteNode(node string) error {
	return r.client.Delete().Namespace(r.namespace).Namespace(node).Do().Error()
}

func (r *ReplicaSetNodeGroup) IsNodeOwned(node *v1.Node) (bool, error) {
	vm := &VirtualMachine{}
	err := r.client.Get().Resource("virtualmachines").Name(node.ObjectMeta.Name).Namespace(r.namespace).Do().Into(vm)
	if err != nil {
		return false, err
	}

	return r.selector.Matches(labels.Set(node.ObjectMeta.Labels)), nil
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
	return int(rs.Spec.Replicas), nil
}

// IncreaseSize increases the size of the node group. To delete a node you need
// to explicitly name it and use DeleteNode. This function should wait until
// node group size is updated. Implementation required.
func (r *ReplicaSetNodeGroup) IncreaseSize(delta int) error {
	rs, err := r.getReplicaSet()
	if err != nil {
		return err
	}
	replicas := int(rs.Spec.Replicas) + delta
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
	for _, node := range  nodes {

		owned, err := r.IsNodeOwned(node)

		if err != nil {
			return err
		}

		if !owned {
			return fmt.Errorf("Node %s is not owned by node group %s", node.ObjectMeta.Name, r.name)
		}
	}

	err = r.pauseReplicaSet()
	if err != nil {
		return err
	}

	for _, node := range  nodes {
		err := r.deleteNode(node.ObjectMeta.Name)
		if err != nil {
			return err
		}
	}

	err = r.scaleReplicaSet(int(*rs.Spec.Replicas) - len(nodes))
	if err != nil {
		return err
	}

	// TODO what if that fails? We need to always resume. Maybe try a few times and then panic?
	return r.resumeReplicaSet()
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
	replicas := int(rs.Spec.Replicas) + delta
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
	l := &VirtualMachineList{}
	err := r.client.Get().Namespace(r.namespace).Resource("virtualmachines").Param("label-selector", r.template.Selector.String()).Do().Into(l)
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

	node := &v1.Node{}

	// TODO, don't hardcode the values
	node.Status.Allocatable[v1.ResourceCPU] = *resource.NewQuantity(800, "m")
	unit := r.template.Template.Spec.Domain.Memory.Unit
	unit = strings.TrimSuffix(unit, "b")
	unit = strings.TrimSuffix(unit, "bytes")
	unit = strings.TrimSuffix(unit, "byte")
	unit = strings.TrimSuffix(unit, "B")
	value := r.template.Template.Spec.Domain.Memory.Value
	node.Status.Allocatable[v1.ResourceMemory] = resource.MustParse(fmt.Sprint("%v%s", value, unit))
	node.Status.Allocatable[v1.ResourceStorage] = *resource.NewQuantity(10, "Gi")

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
