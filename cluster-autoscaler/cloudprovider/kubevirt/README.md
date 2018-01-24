# KubeVirt Cloud Provider

[VirtualMachineReplicaSet](https://kubevirt.gitbooks.io/user-guide/replicaset.html)
based backend for [KubeVirt](https://github.com/kubevirt/kubevirt).

## Auto Discovery

A `LabelSelector` can be used to select all `VirtualMachineReplicaSets` which
represent a `NodeGroup`. When setting `--node-group-auto-discovery` to
`kubevirt:labels=kubevirt.io/autodiscover`, all sets with the label
`kubevirt.io/autodiscovery` will be managed by the autoscaler.

## Static discovery

Not implemented.

## Config

```ini
[cluster]
kubeconfig=/config/to/kubevirt-k8s.cfg
server=https://apiserver.override:8664
```

If no config is provided, the autoscaler will try to connect to kubernetes via
the in-cluster config lookup.

## Example

To scale nodes of a nested Kubernetes cluster (can be called `overcloud`) from a KubeVirt
enabled Kubernetes (can be called `undercloud`, or `basecloud`), take the config from above and save it in
`undercloud.ini`. Create a `VirtualMachineReplicaSet` with a cloud-init secret
and the `kubevirt.io/autoscaler` label. Then run

```bash
./autoscaler --kubeconfig /config/to/overcloud.cfg  --cloud-provider --cloud-config undercloud.ini --node-group-auto-discovery "kubevirt:labels=kubevirt.io/autoscaler"
```

## KubeVirt Dependencies

TODO: Properly vendor the kubevirt API.
