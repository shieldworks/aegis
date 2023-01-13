![Aegis](assets/aegis-banner.png "Aegis")

## Introduction

This documents outlines **Aegis**’ architecture details and project structure.

It is especially useful if you want to have an in-depth understanding of **Aegis**
to contribute to the project.

## Components of Aegis

**Aegis**, as a system, has the following components.

### SPIRE

**SPIRE** is what makes communication within **Aegis** components and workloads
possible. It dispatches x.509 SVID certificates to the required parties to make
secure mTLS communication possible.

### **Safe** (`aegis-safe`)

[`aegis-safe`][safe] stores secrets and dispatches them to workloads.

### **Sidecar** (`aegis-sidecar`)

[`aegis-sidecar`][sidecar] is a sidecar that facilitates delivering secrets to workloads.

### **Sentinel** (`aegis-sentinel`)

[`aegis-sentinel`][sentinel] is a pod you can shell in and do administrative tasks such as
registering secrets for workloads.

[safe]: https://github.com/zerotohero-dev/aegis-safe
[sidecar]: https://github.com/zerotohero-dev/aegis-sidecar
[sentinel]: https://github.com/zerotohero-dev/aegis-sentinel

## Projects

Also, check out [the list of projects][projects] that **Aegis** is composed of.

[projects]: https://github.com/zerotohero-dev/aegis/blob/main/README.md#projects

## High Level Architecture

### Dispatching Identities

**SPIRE** is responsible for delivering short-lived X.509 SVIDs to **Aegis**
components and consumer workloads.

**Sidecar** periodically talks to **Safe** to check if there is a new secret
to be updated.

![Aegis High Level Architecture](assets/aegis-hla.png "Aegis High Level Architecture")

### Creating Secrets

**Sentinel** is the only place that secrets can be created and registered
to **Safe**.

![Creating Secrets](assets/aegis-create-secrets.png "Creating Secrets")

### Component and Workload SVID Schemas

SPIFFE ID format wor workloads is as follows:

```text
spiffe://aegis.z2h.dev/workload/$workloadName/ns/{{ .PodMeta.Namespace }}
/sa/{{ .PodSpec.ServiceAccountName }}/n/{{ .PodMeta.Name }}
```

For the non-aegis-system workloads that **Safe** injects secrets,
`$workloadName` is determined by the workload’s `ClusterSPIFFEID` CRD.

For `aegis-system` components we use `aegis-safe` and `aegis-sentinel` 
for the `$workloadName`:

```text
spiffe://aegis.z2h.dev/workload/aegis-safe/ns/{{ .PodMeta.Namespace }}
/sa/{{ .PodSpec.ServiceAccountName }}/n/{{ .PodMeta.Name }}

```text
spiffe://aegis.z2h.dev/workload/aegis-sentinel/ns/{{ .PodMeta.Namespace }}
/sa/{{ .PodSpec.ServiceAccountName }}/n/{{ .PodMeta.Name }}
```
