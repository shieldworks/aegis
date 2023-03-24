---
#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                    <aegis.ist>
#     .\_/.
#

layout: default
keywords: Aegis, sentinel, cli, bastion
title: Aegis Sentinel CLI
description: <strong>Aegis Sentinel</strong> command line interface
micro_nav: true
page_nav:
  prev:
    content: <strong>Aegis</strong> go SDK
    url: '/docs/sdk'
  next:
    content: <strong>Aegis</strong> deep dive
    url: '/docs/architecture'
---

## Introduction

This section contains usage examples and documentation for **Aegis Sentinel**.

## Finding **Aegis Sentinel**

First, find which pod belongs to **Aegis System**:

```bash 
kubetctl get po -n aeis-system
```

The response to the above command will be similar to the following:

```text 
NAME                              READY
aegis-safe-5f6948c84c-vkrdh       1/1
aegis-sentinel-5998b5dbfc-lvw44   1/1
```

There, `aegis-sentinel-5998b5dbfc-lvw44` is the name of the Pod you’d need.

You can also execute a script similar to the following to save the Pod’s name
into an environment variable:

```bash 
SENTINEL=$(kubectl get po -n aegis-system \
  | grep "aegis-sentinel-" | awk '{print $1}')
```

In the following examples, we’ll use `$SENTINEL` in lieu of the **Aegis Sentinel**’s
Pod name.

## Displaying Help Information

```bash 
kubectl exec $SENTINEL -n aegis-system -- aegis --help
```

Output:

```text 
usage: aegis [-h|--help] [-l|--list] [-k|--use-k8s] [-n|--namespace "<value>"]
             [-b|--store "<value>"] [-w|--workload "<value>"] [-s|--secret
             "<value>"] [-t|--template "<value>"] [-f|--format "<value>"]
             [-e|--encrypt]

             Assigns secrets to workloads.

Arguments:

  -h  --help       Print help information
  -l  --list       lists all registered workloads.. Default: false
  -k  --use-k8s    update an associated Kubernetes secret upon save. Overrides
                   AEGIS_SAFE_USE_KUBERNETES_SECRETS.. Default: false
  -n  --namespace  the namespace of the Kubernetes Secret to create.. Default:
                   aegis-system
  -b  --store      backing store type (file|memory|cluster). Overrides
                   AEGIS_SAFE_BACKING_STORE.
  -w  --workload   name of the workload (i.e. the '$name' segment of its
                   ClusterSPIFFEID
                   ('spiffe://trustDomain/workload/$name/…')).
  -s  --secret     the secret to store for the workload.
  -t  --template   The template transformation to use. 
  -f  --format     The format to transform the secrets (yaml|json|none; 
                   default: none)
```

## Registering a Secret for a Workload

Given our workload has the SPIFFE ID 
`"spiffe://aegis.ist/workload/billing/: …[trunacted]"`

```bash
kubectl exec $SENTINEL -n aegis-system -- aegis \
-w billing
-s "very secret value"
```

will register the secret `"very secret value"` to `billing`.

## Choosing a Backing Store

The registered secrets will be encrypted and backed up to 
**Aegis Safe**’s Kubernetes volume by default. This behavior can be configured
by changing the `AEGIS_SAFE_BACKING_STORE` environment variable that
**Aegis Safe** sees. In addition, this behavior can be overridden on a per-secret
basis too.

The following commands stores the secret to the backing volume 
(*default behavior*):

```bash
kubectl exec $SENTINEL -n aegis-system -- aegis \
-w billing
-s "very secret value"
-b file
```

This one, will **not** store the secret on file; the secret will only be 
persisted in memory, and will be lost if **Aegis Sentinel** needs to restart:

```bash
kubectl exec $SENTINEL -n aegis-system -- aegis \
-w billing
-s "very secret value"
-b memory
```

The following will stored the secret on the cluster itself as a Kubernetes
`Secret`. The value of the secret will be **encrypted** with the public key
of **Aegis Safe** before storing it on the Secret. 

```bash
kubectl exec $SENTINEL -n aegis-system -- aegis \
-w billing
-s "very secret value"
-b cluster
```

## Template Transformations

You can transform how the stored secret is displayed to the consuming workloads:

```bash 
{% raw %}kubectl exec "$SENTINEL" -n aegis-system -- aegis \
-w "aegis-workload-demo" \
-s '{"username": "root", "password": "SuperSecret", "value": "AegisRocks"}' \
-t '{"USER":"{{.username}}", "PASS":"{{.password}}", "VALUE": "{{.value}}"}'{% endraw %}
```

When the workload fetches the secret through the workload API, this is what
it sees as the value:

```text 
{"USER": "root", "PASS": "SuperSecret", "VALUE": "AegisRocks"}
```

Instead of this default transformation, you can output it as `yaml` too:

```bash 
{% raw %}kubectl exec "$SENTINEL" -n aegis-system -- aegis \
-w "aegis-workload-demo" \
-s '{"username": "root", "password": "SuperSecret", "value": "AegisRocks"}' \
-t '{"USER":"{{.username}}", "PASS":"{{.password}}", "VALUE": "{{.value}}"}' \
-f yaml{% endraw %}
```

The above command will result in the following secret value to the workload 
that receives it:

```text 
USER: root
PASS: SuperSecret
VALUE: AegisRocks
```

You can create a YAML secret value without the `-t` flag too. In that case
**Aegis Safe** will assume an identity transformation:

```bash 
kubectl exec "$SENTINEL" -n aegis-system -- aegis \
-w "aegis-workload-demo" \
-s '{"username": "root", "password": "SuperSecret", "value": "AegisRocks"}' \
-f yaml
```

The above command will result in the following secret value to the workload
that receives it:

```text 
username: root
password: SuperSecret
value: AegisRocks
```

If you provide `-f json` as the format argument, the secret will have to be
a strict JSON, or an error will occur.

For `-f none` (*default*), any arbitrary transformation is possible:

```bash
{% raw %}kubectl exec "$SENTINEL" -n aegis-system -- aegis \
-w "aegis-workload-demo" \
-s '{"username": "root", "password": "SuperSecret", "value": "AegisRocks"}' \
-t 'USER»»»{{.username}}' \
-f none{% endraw %}
```

The above command will provide the following secret value to the workload
(*which is neither YAML, nor JSON; it is a free form text*):

```text 
USER»»»root
```

> **Gotcha**
> 
> If `-t` is given, the `-s` argument will have to be a valid JSON regardless
> of what is chosen for `-f`.

The following is also possible:

```bash 
{% raw %}kubectl exec "$SENTINEL" -n aegis-system -- aegis \
-w "aegis-workload-demo" \
-s 'USER»»»{{.username}}'{% endraw %}
```

and will result in the following as the secret value for the workload:

```text 
USER»»»root
```

Or, equivalently:

```bash 
{% raw %}kubectl exec "$SENTINEL" -n aegis-system -- aegis \
-w "aegis-workload-demo" \
-s 'USER»»»{{.username}}' \
-f none{% endraw %}
```

Will give:

```text 
USER»»»root
```

However, the following will raise an error:

```bash 
{% raw %}kubectl exec "$SENTINEL" -n aegis-system -- aegis \
-w "aegis-workload-demo" \
-s 'USER»»»{{.username}}' \
-f yaml{% endraw %}
```

Or, similarly, this will raise an error:

```bash 
{% raw %}kubectl exec "$SENTINEL" -n aegis-system -- aegis \
-w "aegis-workload-demo" \
-s 'USER»»»{{.username}}' \
-f json{% endraw %}
```

To transform the value to *YAML*, or *JSON*, `-s` has to be a **valid** *JSON*.

## Creating Kubernetes Secrets

**Aegis Safe**-managed secrets can be interpolated onto Kubernetes secrets if
a template transformation is given.

Let’s take the following as an example:

```bash
{% raw %}kubectl exec "$SENTINEL" -n aegis-system -- aegis \
-w "billing" \
-n "finance" \
-s '{"username": "root", "password": "SuperSecret"}' \
-t '{"USERNAME":"{{.username}}", "PASSWORD":"{{.password}}"' \
-k{% endraw %}
```

The `-k` flag hints **Aegis Safe** that the secret will be synced with a 
Kubernetes `Secret`. `-n` tells that the namespace of the secret is `"finance"`.

Before running this command, a secret with the following manifest should exist
in the cluster:

```yaml 
apiVersion: v1
kind: Secret
metadata:
  name: aegis-secret-billing
  namespace: finance
type: Opaque
```

The `aegis-secret-` prefix is required for **Aegis Safe** to locate the 
`Secret`. Also `metadata.namespace` attribute should match the namespace 
that’s provided with the `-n` flag to **Aegis Sentinel**.

After executing the command the secret will contain the new values in its
`data:` section.

```bash 
kubectl describe secret aegis-secret-billing -n finance

# Output:
#   Name:         aegis-secret-billing
#   Namespace:    finance
#   Labels:       <none>
#   Annotations:  <none>
#
#   Type:  Opaque
#
#   Data
#   ====
#   USERNAME:  137 bytes
#   PASSWORD:  196 bytes
```