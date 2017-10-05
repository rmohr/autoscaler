/*
 * This file is part of the KubeVirt project
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * Copyright 2017 Red Hat, Inc.
 *
 */

package kubevirt

import (
	"encoding/json"

	k8sv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/apimachinery/announced"
	"k8s.io/apimachinery/pkg/apimachinery/registered"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes/scheme"

	"github.com/mohae/deepcopy"
)

// GroupName is the group name use in this package
const GroupName = "kubevirt.io"

// GroupVersion is group version used to register these objects
var GroupVersion = schema.GroupVersion{Group: GroupName, Version: "v1alpha1"}

// GroupVersionKind
var VirtualMachineGroupVersionKind = schema.GroupVersionKind{Group: GroupName, Version: GroupVersion.Version, Kind: "VirtualMachine"}

var VMReplicaSetGroupVersionKind = schema.GroupVersionKind{Group: GroupName, Version: GroupVersion.Version, Kind: "VirtualMachineReplicaSet"}

var VirtualMachineReplicaSetResource = "virtualmachinereplicasets"

var (
	groupFactoryRegistry = make(announced.APIGroupFactoryRegistry)
	registry             = registered.NewOrDie(GroupVersion.String())
)

// Adds the list of known types to api.Scheme.
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(GroupVersion,
		&VirtualMachine{},
		&VirtualMachineList{},
		&metav1.ListOptions{},
		&metav1.DeleteOptions{},
		&VirtualMachineReplicaSet{},
		&VirtualMachineReplicaSetList{},
		&metav1.GetOptions{},
	)
	return nil
}

func init() {
	SchemeBuilder := runtime.NewSchemeBuilder(addKnownTypes)
	if err := announced.NewGroupMetaFactory(
		&announced.GroupMetaFactoryArgs{
			GroupName:              GroupName,
			VersionPreferenceOrder: []string{GroupVersion.Version},
		},
		announced.VersionToSchemeFunc{
			GroupVersion.Version: SchemeBuilder.AddToScheme,
		},
	).Announce(groupFactoryRegistry).RegisterAndEnable(registry, scheme.Scheme); err != nil {
		panic(err)
	}
}

// VirtualMachine is *the* VM Definition. It represents a virtual machine in the runtime environment of kubernetes.
type VirtualMachine struct {
	metav1.TypeMeta `json:",inline"`
	ObjectMeta      metav1.ObjectMeta `json:"metadata,omitempty"`
	// VM Spec contains the VM specification.
	Spec VMSpec `json:"spec,omitempty" valid:"required"`
	// Status is the high level overview of how the VM is doing. It contains information available to controllers and users.
	Status VMStatus `json:"status"`
}

func (in *VirtualMachine) DeepCopyInto(out *VirtualMachine) {
	v := deepcopy.Copy(in)
	out = v.(*VirtualMachine)
	return
}

func (in *VirtualMachine) DeepCopy() *VirtualMachine {
	if in == nil {
		return nil
	}
	out := new(VirtualMachine)
	in.DeepCopyInto(out)
	return out
}

func (in *VirtualMachine) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// VirtualMachineList is a list of VirtualMachines
type VirtualMachineList struct {
	metav1.TypeMeta `json:",inline"`
	ListMeta        metav1.ListMeta  `json:"metadata,omitempty"`
	Items           []VirtualMachine `json:"items"`
}

func (in *VirtualMachineList) DeepCopyInto(out *VirtualMachineList) {
	v := deepcopy.Copy(in)
	out = v.(*VirtualMachineList)
	return
}

func (in *VirtualMachineList) DeepCopy() *VirtualMachineList {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineList)
	in.DeepCopyInto(out)
	return out
}

func (in *VirtualMachineList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// VMSpec is a description of a VM. Not to be confused with api.DomainSpec in virt-handler.
// It is expected that v1.DomainSpec will be merged into this structure.
type VMSpec struct {
	// Domain is the actual libvirt domain.
	Domain *DomainSpec `json:"domain,omitempty"`
	// If labels are specified, only nodes marked with all of these labels are considered when scheduling the VM.
	NodeSelector map[string]string `json:"nodeSelector,omitempty"`
}

// VMStatus represents information about the status of a VM. Status may trail the actual
// state of a system.
type VMStatus struct {
	// NodeName is the name where the VM is currently running.
	NodeName string `json:"nodeName,omitempty"`
	// MigrationNodeName is the node where the VM is live migrating to.
	MigrationNodeName string `json:"migrationNodeName,omitempty"`
	// Conditions are specific points in VM's pod runtime.
	Conditions []VMCondition `json:"conditions,omitempty"`
	// Phase is the status of the VM in kubernetes world. It is not the VM status, but partially correlates to it.
	Phase VMPhase `json:"phase"`
	// Graphics represent the details of available graphical consoles.
	Graphics []VMGraphics `json:"graphics"`
}

type VMGraphics struct {
	Type string `json:"type"`
	Host string `json:"host"`
	Port int32  `json:"port"`
}

// Required to satisfy Object interface
func (v *VirtualMachine) GetObjectKind() schema.ObjectKind {
	return &v.TypeMeta
}

// Required to satisfy ObjectMetaAccessor interface
func (v *VirtualMachine) GetObjectMeta() metav1.Object {
	return &v.ObjectMeta
}

func (v *VirtualMachine) IsReady() bool {
	// TODO once we support a ready condition, use it instead
	return v.IsRunning()
}

func (v *VirtualMachine) IsRunning() bool {
	return v.Status.Phase == Running || v.Status.Phase == Migrating
}

func (v *VirtualMachine) IsFinal() bool {
	return v.Status.Phase == Failed || v.Status.Phase == Succeeded
}

// Required to satisfy Object interface
func (vl *VirtualMachineList) GetObjectKind() schema.ObjectKind {
	return &vl.TypeMeta
}

// Required to satisfy ListMetaAccessor interface
func (vl *VirtualMachineList) GetListMeta() meta.List {
	return &vl.ListMeta
}

func (v *VirtualMachine) UnmarshalJSON(data []byte) error {
	type VMCopy VirtualMachine
	tmp := VMCopy{}
	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}
	tmp2 := VirtualMachine(tmp)
	*v = tmp2
	return nil
}

func (vl *VirtualMachineList) UnmarshalJSON(data []byte) error {
	type VMListCopy VirtualMachineList
	tmp := VMListCopy{}
	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}
	tmp2 := VirtualMachineList(tmp)
	*vl = tmp2
	return nil
}

type VMConditionType string

// These are valid conditions of VMs.
const (
	// PodCreated means that the VM request was translated into a Pod which can be scheduled and started by
	// Kubernetes.
	PodCreated VMConditionType = "PodCreated"
	// VMReady means the pod is able to service requests and should be added to the
	// load balancing pools of all matching services.
	VMReady VMConditionType = "Ready"
)

type VMCondition struct {
	Type               VMConditionType       `json:"type"`
	Status             k8sv1.ConditionStatus `json:"status"`
	LastProbeTime      metav1.Time           `json:"lastProbeTime,omitempty"`
	LastTransitionTime metav1.Time           `json:"lastTransitionTime,omitempty"`
	Reason             string                `json:"reason,omitempty"`
	Message            string                `json:"message,omitempty"`
}

// VMPhase is a label for the condition of a VM at the current time.
type VMPhase string

// These are the valid statuses of pods.
const (
	//When a VM Object is first initialized and no phase, or Pending is present.
	VmPhaseUnset VMPhase = ""
	// Pending means the VM has been accepted by the system.
	Pending VMPhase = "Pending"
	// Either a target pod does not yet exist or a target Pod exists but is not yet scheduled and in running state.
	Scheduling VMPhase = "Scheduling"
	// A target pod was scheduled and the system saw that Pod in runnig state.
	// Here is where the responsibility of virt-controller ends and virt-handler takes over.
	Scheduled VMPhase = "Scheduled"
	// VMRunning means the pod has been bound to a node and the VM is started.
	Running VMPhase = "Running"
	// VMMigrating means the VM is currently migrated by a controller.
	Migrating VMPhase = "Migrating"
	// VMSucceeded means that the VM stopped voluntarily, e.g. reacted to SIGTERM or shutdown was invoked from
	// inside the VM.
	Succeeded VMPhase = "Succeeded"
	// VMFailed means that associated Pod is in failure state (exited with a non-zero exit code or was stopped by
	// the system).
	Failed VMPhase = "Failed"
	// VMUnknown means that for some reason the state of the VM could not be obtained, typically due
	// to an error in communicating with the host of the VM.
	Unknown VMPhase = "Unknown"
)

const (
	AppLabel          string = "kubevirt.io/app"
	DomainLabel       string = "kubevirt.io/domain"
	VMUIDLabel        string = "kubevirt.io/vmUID"
	NodeNameLabel     string = "kubevirt.io/nodeName"
	MigrationUIDLabel string = "kubevirt.io/migrationUID"
	MigrationLabel    string = "kubevirt.io/migration"
)

func NewVM(name string, uid types.UID) *VirtualMachine {
	return &VirtualMachine{
		Spec: VMSpec{},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			UID:       uid,
			Namespace: k8sv1.NamespaceDefault,
		},
		Status: VMStatus{},
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       VirtualMachineGroupVersionKind.Kind,
		},
	}
}

type SyncEvent string

const (
	Created    SyncEvent = "Created"
	Deleted    SyncEvent = "Deleted"
	Started    SyncEvent = "Started"
	Stopped    SyncEvent = "Stopped"
	SyncFailed SyncEvent = "SyncFailed"
	Resumed    SyncEvent = "Resumed"
)

func (s SyncEvent) String() string {
	return string(s)
}

type VMSelector struct {
	// Name of the VM to migrate
	Name string `json:"name" valid:"required"`
}

// VM is *the* VM Definition. It represents a virtual machine in the runtime environment of kubernetes.
type VirtualMachineReplicaSet struct {
	metav1.TypeMeta `json:",inline"`
	ObjectMeta      metav1.ObjectMeta `json:"metadata,omitempty"`
	// VM Spec contains the VM specification.
	Spec VMReplicaSetSpec `json:"spec,omitempty" valid:"required"`
	// Status is the high level overview of how the VM is doing. It contains information available to controllers and users.
	Status VMReplicaSetStatus `json:"status"`
}

// VMList is a list of VMs
type VirtualMachineReplicaSetList struct {
	metav1.TypeMeta `json:",inline"`
	ListMeta        metav1.ListMeta            `json:"metadata,omitempty"`
	Items           []VirtualMachineReplicaSet `json:"items"`
}

type VMReplicaSetSpec struct {
	// Number of desired pods. This is a pointer to distinguish between explicit
	// zero and not specified. Defaults to 1.
	// +optional
	Replicas *int32 `json:"replicas,omitempty"`

	// Label selector for pods. Existing ReplicaSets whose pods are
	// selected by this will be the ones affected by this deployment.
	// +optional
	Selector *metav1.LabelSelector `json:"selector,omitempty" valid:"required"`

	// Template describes the pods that will be created.
	Template *VMTemplateSpec `json:"template" valid:"required"`

	// Indicates that the replica set is paused.
	// +optional
	Paused bool `json:"paused,omitempty" protobuf:"varint,7,opt,name=paused"`
}

type VMReplicaSetStatus struct {
	// Total number of non-terminated pods targeted by this deployment (their labels match the selector).
	// +optional
	Replicas int32 `json:"replicas,omitempty" protobuf:"varint,2,opt,name=replicas"`

	// The number of ready replicas for this replica set.
	// +optional
	ReadyReplicas int32 `json:"readyReplicas,omitempty" protobuf:"varint,4,opt,name=readyReplicas"`

	Conditions []VMReplicaSetCondition `json:"conditions"`
}

type VMReplicaSetCondition struct {
	Type               VMReplicaSetConditionType `json:"type"`
	Status             k8sv1.ConditionStatus     `json:"status"`
	LastProbeTime      metav1.Time               `json:"lastProbeTime,omitempty"`
	LastTransitionTime metav1.Time               `json:"lastTransitionTime,omitempty"`
	Reason             string                    `json:"reason,omitempty"`
	Message            string                    `json:"message,omitempty"`
}

type VMReplicaSetConditionType string

const (
	// VMReplicaSetReplicaFailure is added in a replica set when one of its vms
	// fails to be created due to insufficient quota, limit ranges, pod security policy, node selectors,
	// etc. or deleted due to kubelet being down or finalizers are failing.
	VMReplicaSetReplicaFailure VMReplicaSetConditionType = "ReplicaFailure"

	// VMReplicaSetReplicaPaused is added in a replica set when the replica set got paused by the controller.
	// After this condition was added, it is safe to remove or add vms by hand and adjust the replica count by hand.
	VMReplicaSetReplicaPaused VMReplicaSetConditionType = "ReplicaPaused"
)

type VMTemplateSpec struct {
	ObjectMeta metav1.ObjectMeta `json:"metadata,omitempty"`
	// VM Spec contains the VM specification.
	Spec VMSpec `json:"spec,omitempty" valid:"required"`
}

// Required to satisfy Object interface
func (v *VirtualMachineReplicaSet) GetObjectKind() schema.ObjectKind {
	return &v.TypeMeta
}

// Required to satisfy ObjectMetaAccessor interface
func (v *VirtualMachineReplicaSet) GetObjectMeta() metav1.Object {
	return &v.ObjectMeta
}

func (v *VirtualMachineReplicaSet) UnmarshalJSON(data []byte) error {
	type VMReplicaSetCopy VirtualMachineReplicaSet
	tmp := VMReplicaSetCopy{}
	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}
	tmp2 := VirtualMachineReplicaSet(tmp)
	*v = tmp2
	return nil
}

func (vl *VirtualMachineReplicaSetList) UnmarshalJSON(data []byte) error {
	type VMReplicaSetListCopy VirtualMachineReplicaSetList
	tmp := VMReplicaSetListCopy{}
	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}
	tmp2 := VirtualMachineReplicaSetList(tmp)
	*vl = tmp2
	return nil
}

// Required to satisfy Object interface
func (vl *VirtualMachineReplicaSetList) GetObjectKind() schema.ObjectKind {
	return &vl.TypeMeta
}

// Required to satisfy ListMetaAccessor interface
func (vl *VirtualMachineReplicaSetList) GetListMeta() meta.List {
	return &vl.ListMeta
}

func (in *VirtualMachineReplicaSet) DeepCopyInto(out *VirtualMachineReplicaSet) {
	v := deepcopy.Copy(in)
	out = v.(*VirtualMachineReplicaSet)
	return
}

func (in *VirtualMachineReplicaSet) DeepCopy() *VirtualMachineReplicaSet {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineReplicaSet)
	in.DeepCopyInto(out)
	return out
}

func (in *VirtualMachineReplicaSet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

func (in *VirtualMachineReplicaSetList) DeepCopyInto(out *VirtualMachineReplicaSetList) {
	v := deepcopy.Copy(in)
	out = v.(*VirtualMachineReplicaSetList)
	return
}

func (in *VirtualMachineReplicaSetList) DeepCopy() *VirtualMachineReplicaSetList {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineReplicaSetList)
	in.DeepCopyInto(out)
	return out
}

func (in *VirtualMachineReplicaSetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

type DomainSpec struct {
	Memory Memory `json:"memory"`
}

type Memory struct {
	Value uint   `json:"value"`
	Unit  string `json:"unit"`
}
