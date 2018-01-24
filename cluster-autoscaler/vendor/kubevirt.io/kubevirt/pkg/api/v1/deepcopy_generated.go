// +build !ignore_autogenerated

/*
Copyright 2018 The KubeVirt Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1

import (
	core_v1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	types "k8s.io/apimachinery/pkg/types"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Affinity) DeepCopyInto(out *Affinity) {
	*out = *in
	if in.NodeAffinity != nil {
		in, out := &in.NodeAffinity, &out.NodeAffinity
		if *in == nil {
			*out = nil
		} else {
			*out = new(core_v1.NodeAffinity)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Affinity.
func (in *Affinity) DeepCopy() *Affinity {
	if in == nil {
		return nil
	}
	out := new(Affinity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CDRomTarget) DeepCopyInto(out *CDRomTarget) {
	*out = *in
	if in.ReadOnly != nil {
		in, out := &in.ReadOnly, &out.ReadOnly
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CDRomTarget.
func (in *CDRomTarget) DeepCopy() *CDRomTarget {
	if in == nil {
		return nil
	}
	out := new(CDRomTarget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CPU) DeepCopyInto(out *CPU) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CPU.
func (in *CPU) DeepCopy() *CPU {
	if in == nil {
		return nil
	}
	out := new(CPU)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Clock) DeepCopyInto(out *Clock) {
	*out = *in
	in.ClockOffset.DeepCopyInto(&out.ClockOffset)
	if in.Timer != nil {
		in, out := &in.Timer, &out.Timer
		if *in == nil {
			*out = nil
		} else {
			*out = new(Timer)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Clock.
func (in *Clock) DeepCopy() *Clock {
	if in == nil {
		return nil
	}
	out := new(Clock)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClockOffset) DeepCopyInto(out *ClockOffset) {
	*out = *in
	if in.UTC != nil {
		in, out := &in.UTC, &out.UTC
		if *in == nil {
			*out = nil
		} else {
			*out = new(ClockOffsetUTC)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Timezone != nil {
		in, out := &in.Timezone, &out.Timezone
		if *in == nil {
			*out = nil
		} else {
			*out = new(ClockOffsetTimezone)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClockOffset.
func (in *ClockOffset) DeepCopy() *ClockOffset {
	if in == nil {
		return nil
	}
	out := new(ClockOffset)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClockOffsetUTC) DeepCopyInto(out *ClockOffsetUTC) {
	*out = *in
	if in.OffsetSeconds != nil {
		in, out := &in.OffsetSeconds, &out.OffsetSeconds
		if *in == nil {
			*out = nil
		} else {
			*out = new(int)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClockOffsetUTC.
func (in *ClockOffsetUTC) DeepCopy() *ClockOffsetUTC {
	if in == nil {
		return nil
	}
	out := new(ClockOffsetUTC)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudInitNoCloudSource) DeepCopyInto(out *CloudInitNoCloudSource) {
	*out = *in
	if in.UserDataSecretRef != nil {
		in, out := &in.UserDataSecretRef, &out.UserDataSecretRef
		if *in == nil {
			*out = nil
		} else {
			*out = new(core_v1.LocalObjectReference)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudInitNoCloudSource.
func (in *CloudInitNoCloudSource) DeepCopy() *CloudInitNoCloudSource {
	if in == nil {
		return nil
	}
	out := new(CloudInitNoCloudSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Devices) DeepCopyInto(out *Devices) {
	*out = *in
	if in.Disks != nil {
		in, out := &in.Disks, &out.Disks
		*out = make([]Disk, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Watchdog != nil {
		in, out := &in.Watchdog, &out.Watchdog
		if *in == nil {
			*out = nil
		} else {
			*out = new(Watchdog)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Devices.
func (in *Devices) DeepCopy() *Devices {
	if in == nil {
		return nil
	}
	out := new(Devices)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Disk) DeepCopyInto(out *Disk) {
	*out = *in
	in.DiskDevice.DeepCopyInto(&out.DiskDevice)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Disk.
func (in *Disk) DeepCopy() *Disk {
	if in == nil {
		return nil
	}
	out := new(Disk)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskDevice) DeepCopyInto(out *DiskDevice) {
	*out = *in
	if in.Disk != nil {
		in, out := &in.Disk, &out.Disk
		if *in == nil {
			*out = nil
		} else {
			*out = new(DiskTarget)
			**out = **in
		}
	}
	if in.LUN != nil {
		in, out := &in.LUN, &out.LUN
		if *in == nil {
			*out = nil
		} else {
			*out = new(LunTarget)
			**out = **in
		}
	}
	if in.Floppy != nil {
		in, out := &in.Floppy, &out.Floppy
		if *in == nil {
			*out = nil
		} else {
			*out = new(FloppyTarget)
			**out = **in
		}
	}
	if in.CDRom != nil {
		in, out := &in.CDRom, &out.CDRom
		if *in == nil {
			*out = nil
		} else {
			*out = new(CDRomTarget)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskDevice.
func (in *DiskDevice) DeepCopy() *DiskDevice {
	if in == nil {
		return nil
	}
	out := new(DiskDevice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskTarget) DeepCopyInto(out *DiskTarget) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskTarget.
func (in *DiskTarget) DeepCopy() *DiskTarget {
	if in == nil {
		return nil
	}
	out := new(DiskTarget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainSpec) DeepCopyInto(out *DomainSpec) {
	*out = *in
	in.Resources.DeepCopyInto(&out.Resources)
	if in.CPU != nil {
		in, out := &in.CPU, &out.CPU
		if *in == nil {
			*out = nil
		} else {
			*out = new(CPU)
			**out = **in
		}
	}
	if in.Firmware != nil {
		in, out := &in.Firmware, &out.Firmware
		if *in == nil {
			*out = nil
		} else {
			*out = new(Firmware)
			**out = **in
		}
	}
	if in.Clock != nil {
		in, out := &in.Clock, &out.Clock
		if *in == nil {
			*out = nil
		} else {
			*out = new(Clock)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Features != nil {
		in, out := &in.Features, &out.Features
		if *in == nil {
			*out = nil
		} else {
			*out = new(Features)
			(*in).DeepCopyInto(*out)
		}
	}
	in.Devices.DeepCopyInto(&out.Devices)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainSpec.
func (in *DomainSpec) DeepCopy() *DomainSpec {
	if in == nil {
		return nil
	}
	out := new(DomainSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmptyDiskSource) DeepCopyInto(out *EmptyDiskSource) {
	*out = *in
	out.Capacity = in.Capacity.DeepCopy()
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmptyDiskSource.
func (in *EmptyDiskSource) DeepCopy() *EmptyDiskSource {
	if in == nil {
		return nil
	}
	out := new(EmptyDiskSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureAPIC) DeepCopyInto(out *FeatureAPIC) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureAPIC.
func (in *FeatureAPIC) DeepCopy() *FeatureAPIC {
	if in == nil {
		return nil
	}
	out := new(FeatureAPIC)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureHyperv) DeepCopyInto(out *FeatureHyperv) {
	*out = *in
	if in.Relaxed != nil {
		in, out := &in.Relaxed, &out.Relaxed
		if *in == nil {
			*out = nil
		} else {
			*out = new(FeatureState)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.VAPIC != nil {
		in, out := &in.VAPIC, &out.VAPIC
		if *in == nil {
			*out = nil
		} else {
			*out = new(FeatureState)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Spinlocks != nil {
		in, out := &in.Spinlocks, &out.Spinlocks
		if *in == nil {
			*out = nil
		} else {
			*out = new(FeatureSpinlocks)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.VPIndex != nil {
		in, out := &in.VPIndex, &out.VPIndex
		if *in == nil {
			*out = nil
		} else {
			*out = new(FeatureState)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Runtime != nil {
		in, out := &in.Runtime, &out.Runtime
		if *in == nil {
			*out = nil
		} else {
			*out = new(FeatureState)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.SyNIC != nil {
		in, out := &in.SyNIC, &out.SyNIC
		if *in == nil {
			*out = nil
		} else {
			*out = new(FeatureState)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.SyNICTimer != nil {
		in, out := &in.SyNICTimer, &out.SyNICTimer
		if *in == nil {
			*out = nil
		} else {
			*out = new(FeatureState)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Reset != nil {
		in, out := &in.Reset, &out.Reset
		if *in == nil {
			*out = nil
		} else {
			*out = new(FeatureState)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.VendorID != nil {
		in, out := &in.VendorID, &out.VendorID
		if *in == nil {
			*out = nil
		} else {
			*out = new(FeatureVendorID)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureHyperv.
func (in *FeatureHyperv) DeepCopy() *FeatureHyperv {
	if in == nil {
		return nil
	}
	out := new(FeatureHyperv)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureSpinlocks) DeepCopyInto(out *FeatureSpinlocks) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	if in.Retries != nil {
		in, out := &in.Retries, &out.Retries
		if *in == nil {
			*out = nil
		} else {
			*out = new(uint32)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureSpinlocks.
func (in *FeatureSpinlocks) DeepCopy() *FeatureSpinlocks {
	if in == nil {
		return nil
	}
	out := new(FeatureSpinlocks)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureState) DeepCopyInto(out *FeatureState) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureState.
func (in *FeatureState) DeepCopy() *FeatureState {
	if in == nil {
		return nil
	}
	out := new(FeatureState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureVendorID) DeepCopyInto(out *FeatureVendorID) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureVendorID.
func (in *FeatureVendorID) DeepCopy() *FeatureVendorID {
	if in == nil {
		return nil
	}
	out := new(FeatureVendorID)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Features) DeepCopyInto(out *Features) {
	*out = *in
	in.ACPI.DeepCopyInto(&out.ACPI)
	if in.APIC != nil {
		in, out := &in.APIC, &out.APIC
		if *in == nil {
			*out = nil
		} else {
			*out = new(FeatureState)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Hyperv != nil {
		in, out := &in.Hyperv, &out.Hyperv
		if *in == nil {
			*out = nil
		} else {
			*out = new(FeatureHyperv)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Features.
func (in *Features) DeepCopy() *Features {
	if in == nil {
		return nil
	}
	out := new(Features)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Firmware) DeepCopyInto(out *Firmware) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Firmware.
func (in *Firmware) DeepCopy() *Firmware {
	if in == nil {
		return nil
	}
	out := new(Firmware)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FloppyTarget) DeepCopyInto(out *FloppyTarget) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FloppyTarget.
func (in *FloppyTarget) DeepCopy() *FloppyTarget {
	if in == nil {
		return nil
	}
	out := new(FloppyTarget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HPETTimer) DeepCopyInto(out *HPETTimer) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HPETTimer.
func (in *HPETTimer) DeepCopy() *HPETTimer {
	if in == nil {
		return nil
	}
	out := new(HPETTimer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HypervTimer) DeepCopyInto(out *HypervTimer) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HypervTimer.
func (in *HypervTimer) DeepCopy() *HypervTimer {
	if in == nil {
		return nil
	}
	out := new(HypervTimer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *I6300ESBWatchdog) DeepCopyInto(out *I6300ESBWatchdog) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new I6300ESBWatchdog.
func (in *I6300ESBWatchdog) DeepCopy() *I6300ESBWatchdog {
	if in == nil {
		return nil
	}
	out := new(I6300ESBWatchdog)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KVMTimer) DeepCopyInto(out *KVMTimer) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KVMTimer.
func (in *KVMTimer) DeepCopy() *KVMTimer {
	if in == nil {
		return nil
	}
	out := new(KVMTimer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LunTarget) DeepCopyInto(out *LunTarget) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LunTarget.
func (in *LunTarget) DeepCopy() *LunTarget {
	if in == nil {
		return nil
	}
	out := new(LunTarget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Migration) DeepCopyInto(out *Migration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Migration.
func (in *Migration) DeepCopy() *Migration {
	if in == nil {
		return nil
	}
	out := new(Migration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Migration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigrationHostInfo) DeepCopyInto(out *MigrationHostInfo) {
	*out = *in
	if in.Controller != nil {
		in, out := &in.Controller, &out.Controller
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigrationHostInfo.
func (in *MigrationHostInfo) DeepCopy() *MigrationHostInfo {
	if in == nil {
		return nil
	}
	out := new(MigrationHostInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigrationList) DeepCopyInto(out *MigrationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Migration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigrationList.
func (in *MigrationList) DeepCopy() *MigrationList {
	if in == nil {
		return nil
	}
	out := new(MigrationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MigrationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigrationSpec) DeepCopyInto(out *MigrationSpec) {
	*out = *in
	out.Selector = in.Selector
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigrationSpec.
func (in *MigrationSpec) DeepCopy() *MigrationSpec {
	if in == nil {
		return nil
	}
	out := new(MigrationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigrationStatus) DeepCopyInto(out *MigrationStatus) {
	*out = *in
	if in.Instance != nil {
		in, out := &in.Instance, &out.Instance
		if *in == nil {
			*out = nil
		} else {
			*out = new(types.UID)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigrationStatus.
func (in *MigrationStatus) DeepCopy() *MigrationStatus {
	if in == nil {
		return nil
	}
	out := new(MigrationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PITTimer) DeepCopyInto(out *PITTimer) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PITTimer.
func (in *PITTimer) DeepCopy() *PITTimer {
	if in == nil {
		return nil
	}
	out := new(PITTimer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RTCTimer) DeepCopyInto(out *RTCTimer) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RTCTimer.
func (in *RTCTimer) DeepCopy() *RTCTimer {
	if in == nil {
		return nil
	}
	out := new(RTCTimer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryDiskSource) DeepCopyInto(out *RegistryDiskSource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryDiskSource.
func (in *RegistryDiskSource) DeepCopy() *RegistryDiskSource {
	if in == nil {
		return nil
	}
	out := new(RegistryDiskSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceRequirements) DeepCopyInto(out *ResourceRequirements) {
	*out = *in
	if in.Requests != nil {
		in, out := &in.Requests, &out.Requests
		*out = make(core_v1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceRequirements.
func (in *ResourceRequirements) DeepCopy() *ResourceRequirements {
	if in == nil {
		return nil
	}
	out := new(ResourceRequirements)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Spice) DeepCopyInto(out *Spice) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Info = in.Info
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Spice.
func (in *Spice) DeepCopy() *Spice {
	if in == nil {
		return nil
	}
	out := new(Spice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Spice) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpiceInfo) DeepCopyInto(out *SpiceInfo) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpiceInfo.
func (in *SpiceInfo) DeepCopy() *SpiceInfo {
	if in == nil {
		return nil
	}
	out := new(SpiceInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Timer) DeepCopyInto(out *Timer) {
	*out = *in
	if in.HPET != nil {
		in, out := &in.HPET, &out.HPET
		if *in == nil {
			*out = nil
		} else {
			*out = new(HPETTimer)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.KVM != nil {
		in, out := &in.KVM, &out.KVM
		if *in == nil {
			*out = nil
		} else {
			*out = new(KVMTimer)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.PIT != nil {
		in, out := &in.PIT, &out.PIT
		if *in == nil {
			*out = nil
		} else {
			*out = new(PITTimer)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.RTC != nil {
		in, out := &in.RTC, &out.RTC
		if *in == nil {
			*out = nil
		} else {
			*out = new(RTCTimer)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Hyperv != nil {
		in, out := &in.Hyperv, &out.Hyperv
		if *in == nil {
			*out = nil
		} else {
			*out = new(HypervTimer)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Timer.
func (in *Timer) DeepCopy() *Timer {
	if in == nil {
		return nil
	}
	out := new(Timer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VMReplicaSetCondition) DeepCopyInto(out *VMReplicaSetCondition) {
	*out = *in
	in.LastProbeTime.DeepCopyInto(&out.LastProbeTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VMReplicaSetCondition.
func (in *VMReplicaSetCondition) DeepCopy() *VMReplicaSetCondition {
	if in == nil {
		return nil
	}
	out := new(VMReplicaSetCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VMReplicaSetSpec) DeepCopyInto(out *VMReplicaSetSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		if *in == nil {
			*out = nil
		} else {
			*out = new(int32)
			**out = **in
		}
	}
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		if *in == nil {
			*out = nil
		} else {
			*out = new(meta_v1.LabelSelector)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Template != nil {
		in, out := &in.Template, &out.Template
		if *in == nil {
			*out = nil
		} else {
			*out = new(VMTemplateSpec)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VMReplicaSetSpec.
func (in *VMReplicaSetSpec) DeepCopy() *VMReplicaSetSpec {
	if in == nil {
		return nil
	}
	out := new(VMReplicaSetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VMReplicaSetStatus) DeepCopyInto(out *VMReplicaSetStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]VMReplicaSetCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VMReplicaSetStatus.
func (in *VMReplicaSetStatus) DeepCopy() *VMReplicaSetStatus {
	if in == nil {
		return nil
	}
	out := new(VMReplicaSetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VMSelector) DeepCopyInto(out *VMSelector) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VMSelector.
func (in *VMSelector) DeepCopy() *VMSelector {
	if in == nil {
		return nil
	}
	out := new(VMSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VMTemplateSpec) DeepCopyInto(out *VMTemplateSpec) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VMTemplateSpec.
func (in *VMTemplateSpec) DeepCopy() *VMTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(VMTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachine) DeepCopyInto(out *VirtualMachine) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachine.
func (in *VirtualMachine) DeepCopy() *VirtualMachine {
	if in == nil {
		return nil
	}
	out := new(VirtualMachine)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualMachine) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineCondition) DeepCopyInto(out *VirtualMachineCondition) {
	*out = *in
	in.LastProbeTime.DeepCopyInto(&out.LastProbeTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineCondition.
func (in *VirtualMachineCondition) DeepCopy() *VirtualMachineCondition {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineList) DeepCopyInto(out *VirtualMachineList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VirtualMachine, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineList.
func (in *VirtualMachineList) DeepCopy() *VirtualMachineList {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualMachineList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineReplicaSet) DeepCopyInto(out *VirtualMachineReplicaSet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineReplicaSet.
func (in *VirtualMachineReplicaSet) DeepCopy() *VirtualMachineReplicaSet {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineReplicaSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualMachineReplicaSet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineReplicaSetList) DeepCopyInto(out *VirtualMachineReplicaSetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VirtualMachineReplicaSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineReplicaSetList.
func (in *VirtualMachineReplicaSetList) DeepCopy() *VirtualMachineReplicaSetList {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineReplicaSetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualMachineReplicaSetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineSpec) DeepCopyInto(out *VirtualMachineSpec) {
	*out = *in
	in.Domain.DeepCopyInto(&out.Domain)
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		if *in == nil {
			*out = nil
		} else {
			*out = new(Affinity)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.TerminationGracePeriodSeconds != nil {
		in, out := &in.TerminationGracePeriodSeconds, &out.TerminationGracePeriodSeconds
		if *in == nil {
			*out = nil
		} else {
			*out = new(int64)
			**out = **in
		}
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineSpec.
func (in *VirtualMachineSpec) DeepCopy() *VirtualMachineSpec {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineStatus) DeepCopyInto(out *VirtualMachineStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]VirtualMachineCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineStatus.
func (in *VirtualMachineStatus) DeepCopy() *VirtualMachineStatus {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Volume) DeepCopyInto(out *Volume) {
	*out = *in
	in.VolumeSource.DeepCopyInto(&out.VolumeSource)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Volume.
func (in *Volume) DeepCopy() *Volume {
	if in == nil {
		return nil
	}
	out := new(Volume)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeSource) DeepCopyInto(out *VolumeSource) {
	*out = *in
	if in.ISCSI != nil {
		in, out := &in.ISCSI, &out.ISCSI
		if *in == nil {
			*out = nil
		} else {
			*out = new(core_v1.ISCSIVolumeSource)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.PersistentVolumeClaim != nil {
		in, out := &in.PersistentVolumeClaim, &out.PersistentVolumeClaim
		if *in == nil {
			*out = nil
		} else {
			*out = new(core_v1.PersistentVolumeClaimVolumeSource)
			**out = **in
		}
	}
	if in.CloudInitNoCloud != nil {
		in, out := &in.CloudInitNoCloud, &out.CloudInitNoCloud
		if *in == nil {
			*out = nil
		} else {
			*out = new(CloudInitNoCloudSource)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.RegistryDisk != nil {
		in, out := &in.RegistryDisk, &out.RegistryDisk
		if *in == nil {
			*out = nil
		} else {
			*out = new(RegistryDiskSource)
			**out = **in
		}
	}
	if in.EmptyDisk != nil {
		in, out := &in.EmptyDisk, &out.EmptyDisk
		if *in == nil {
			*out = nil
		} else {
			*out = new(EmptyDiskSource)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeSource.
func (in *VolumeSource) DeepCopy() *VolumeSource {
	if in == nil {
		return nil
	}
	out := new(VolumeSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Watchdog) DeepCopyInto(out *Watchdog) {
	*out = *in
	in.WatchdogDevice.DeepCopyInto(&out.WatchdogDevice)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Watchdog.
func (in *Watchdog) DeepCopy() *Watchdog {
	if in == nil {
		return nil
	}
	out := new(Watchdog)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WatchdogDevice) DeepCopyInto(out *WatchdogDevice) {
	*out = *in
	if in.I6300ESB != nil {
		in, out := &in.I6300ESB, &out.I6300ESB
		if *in == nil {
			*out = nil
		} else {
			*out = new(I6300ESBWatchdog)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WatchdogDevice.
func (in *WatchdogDevice) DeepCopy() *WatchdogDevice {
	if in == nil {
		return nil
	}
	out := new(WatchdogDevice)
	in.DeepCopyInto(out)
	return out
}
