---
#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                    <aegis.ist>
#     .\_/.
#

layout: default
keywords: Aegis, installation, deployment, production, reference architecture
title: Production Deployment
description: things to pay attention to
micro_nav: true
page_nav:
  prev:
    content: Miscellaneous
    url: '/misc'
---

## Introduction

There are certain aspects that you need to pay attention to and certain
parts of the system that you’d need to harden for a **production** *Aegis* setup.
This article will overview them.

## Version Compatibility

By the time of this writing,
**Aegis** has been tested with the following Kubernetes version:

```text
Client Version: v1.26.0
Kustomize Version: v4.5.7
Server Version: v1.25.3
```

Any Kubernetes setup that has components with version greater
than or equal to the ones above will likely work just fine.

## Back Up Your Cluster Regularly

**Aegis** is designed to be resilient; however, it is still a possibility
to lose access to your sensitive data by inadvertently deleting a Kubernetes
`Secret` that you are not supposed to delete. Or, your backing store that
contains the secrets can get corrupted for any reason. 

Cloud Native, or not, you rely on hardware which—intrinsically—is unreliable.

Things happen. Make sure you back up your cluster 
[using a tool like **Velero**][velero], so when things do happen, you 
can revert to the last known good state of your cluster.

Set up your backups from day zero.

[velero]: https://velero.io/ "Velero"

## Restrict Access To `aegis-age-key`

The `aegis-age-key` secret that **Aegis Safe** stores in the `aegis-system` 
namespace contains the keys to encrypt and decrypt secret data on the data
volume of **Aegis Safe**.

While reading the secret alone is not enough to plant an attack on the secrets
(*because the attacker also needs to access the Aegis Safe Pod or the `/data`
volume in that Pod*) it is still **crucial** to follow the **principle of least
privilege** guideline and do not allow anyone on the cluster read or write
to the `aegis-age-key` secret.

The only entity allowed to have read/write (*but not delete*) access to
`aegis-safe-key` is the **Aegis Safe** Pod inside the `aegis-system` namespace
with an `aegis-safe` service account.

## Restrict Access to Aegis Sentinel

To be extra secure, all **Aegis** images are based on [distroless][distroless]
containers. Thus, an operator cannot execute a shell on the pod to try
a privilege escalation or container escape attack. However, this does not mean
you can leave the `aegis-system` namespace like an open buffet.

Always take a **principle of least privilege** stance. Do not let anyone who
does not need to fiddle with the `aegis-system` namespace see and use the 
resources there.

This stance is especially important for the **Aegis Sentinel** Pod since an
attacker who has access to that pod can override (*but not read*) secrets on
workloads.

**Aegis** leverages Kubernetes security primitives and modern cryptography
to secure access to secrets. And **Aegis Sentinel** is the **only** system 
part that has direct write access to the **Aegis Safe** secrets store. Therefore, 
o nce you secure your access to **Aegis Sentinel** with proper RBAC and policies, 
you secure your access to your secrets.

[distroless]: https://github.com/GoogleContainerTools/distroless

## Volume Selection for Aegis Safe Backing Store

[**Aegis Safe** default deployment descriptor][aegis-safe-deployment-yaml]
uses [`HostPath`][k8s-pv] to store encrypted backups for secrets.

It is highly recommended to make sure that the backing store that **Aegis Safe**
uses is durable, performant, and reliable.

It is a best practice to avoid `HostPath` volumes for production deployments.
For production setups, you are strongly encouraged to [choose a `PersistentVolume`
that suits your needs][k8s-pv].

[aegis-safe-deployment-yaml]: https://github.com/shieldworks/aegis/blob/main/install/k8s/safe/Deployment.yaml
[k8s-pv]: https://kubernetes.io/docs/concepts/storage/volumes/

## Set CPU and Memory Limits

Benchmark your system usage and set **CPU** and **Memory** limits to the
**Aegis Safe** pod.

It is recommended to…

* Set a memory **request** and **limit**,
* Set a CPU **request**; but **not** set a CPU limit (*i.e., 
  the **Aegis Safe** pod will ask for a baseline CPU
  and burst for more upon need*).

As in any secrets management solution, your compute and memory requirements
will depend on several factors, such as:

* The number of workloads in the cluster
* The number of secrets **Safe** (*Aegis’ Secrets Store*) has to manage
  (*see [architecture details][architecture] for more context*)
* The number of workloads interacting with **Safe**
  (*see [architecture details][architecture] for more context*)
* **Sidecar** poll frequency (*see [architecture details][architecture] for more context*)
* etc.

[architecture]: /docs/architecture 

We recommend you benchmark with a realistic production-like
cluster and allocate your resources accordingly.

That being said, here are the resource allocation reported by `kubectl top`
for a demo setup on a single-node minikube cluster to give an idea:

```text 
NAMESPACE     WORKLOAD            CPU(cores) MEMORY(bytes)
aegis-system  aegis-safe          1m         9Mi
aegis-system  aegis-sentinel      1m         3Mi
default       aegis-workload-demo 2m         7Mi
spire-system  spire-agent         4m         35Mi
spire-system  spire-server        6m         41Mi
```

Note that 1000m is 1 full CPU core.

Based on these findings, the following resource and limit allocations can be
a starting point for **Aegis**-managed containers:

```text
  # Resource allocation will highly-depend on the system.
  # Benchmark your deployment, monitor your resource utilization,
  # and adjust these values accordingly.
  resources:
    requests:
      memory: "128Mi"
      cpu: "250m"
    limits:
      memory: "128Mi"
      # We recommend “NOT” setting a CPU limit.
      # As long as you have configured your CPU “requests”
      # correctly, everything would work fine.
```

## Conclusion

Since **Aegis** is a *Kubernetes-native* framework, its security is strongly
related to how you secure your cluster. You should be safe if you keep your 
cluster and the`aegis-system` namespace secure and follow 
“*the principle of least privilege*” as a guideline.

**Aegis** is a lightweight secrets manager; however, that does not mean it
runs on water: It needs CPU and Memory resources. The amount of resources you
need will depend on the criteria outlined in the former sections. You can either
benchmark your system and set your resources accordingly. Or set generous-enough
limits and adjust your settings as time goes by.

Also, you are strongly encouraged **not** to set a limit on **Aegis** Pods’ CPU
usage. Instead, it is recommended to let **Aegis Safe** burst the CPU when 
it needs.

On the same topic, you are encouraged to set a **request** for **Aegis Safe**
to guarantee a baseline compute allocation.
