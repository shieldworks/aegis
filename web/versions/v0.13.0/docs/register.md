---
#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                    <aegis.ist>
#     .\_/.
#

layout: default
keywords: Aegis, tutorial, secrets, secret registration
title: Registering Secrets
description: send those bad boys to your workloads
micro_nav: true
page_nav:
  prev:
    content: quickstart
    url: '/docs/'
  next:
    content: local development
    url: '/docs/contributing'
---

## Introduction

In this tutorial, you will register secrets to Kubernetes workloads
using **Aegis**. 

We will first discuss how to register a secret to a workload using
[**Aegis Sidecar**][sidecar], and then we will cover a more direct approach using
the [**Aegis Go SDK**][sdk-go].

[sidecar]: https://github.com/shieldworks/aegis-sidecar
[sdk-go]: https://github.com/shieldworks/aegis-sdk-go

## Prerequisites

To complete this tutorial, you will need the following:

* A **Kubernetes** cluster that you have sufficient admin rights.
* **Aegis** up and running on that cluster.
* [The `shieldworks/aegis` repository][repo] cloned inside a workspace
  folder (such as `/home/jane-doe/Desktop/WORKSPACE/aegis`)

> **How Do I Set Up Aegis**?
> 
> To set up **Aegis**, [follow the instructions in this quickstart guide][quickstart].

[quickstart]: /docs/
[repo]: https://github.com/shieldworks/aegis

## High-Level Overview

Here is a high-level overview of various components that will interact with 
each other in this demo:

![High-Level Overview](/assets/actors.jpg "High-Level Overview")

On the above diagram:

* **SPIRE** is the identity provider for all intents and purposes.
* **Aegis Safe** is where secrets are stored.
* **Aegis Sentinel** can be considered a bastion host.
* **Demo Workload** is a typical Kubernetes Pod that needs secrets.

> **Want a Deeper Dive**?
> 
> In this tutorial, we cover only the amount of information necessary
> to follow through the steps and make sense of how things tie together
> from a platform operator’s perspective.
> 
> [You can check out this “**Aegis** Deep Dive” article][architecture] 
> to learn more about these components.

[architecture]: /docs/architecture

The **Demo Workload** fetches secrets from **Aegis Safe**. This is either
indirectly done through a **sidecar** or directly by using 
[**Aegis Go SDK**][go-sdk].

Using **Aegis Sentinel**, an admin operator or ar CI/CD pipeline can register
secrets to **Aegis Safe** for the **Demo Workload** to consume.

All the above workload-to-safe and sentinel-to-safe communication are
encrypted through **mTLS** using the **X.509 SVID**s that **SPIRE** 
dispatches to all the actors.

[go-sdk]: https://github.com/shieldworks/aegis-sdk-go

After this high-level overview of your system, let’s create a workload.

## Deploying a Demo Workload With a Sidecar

Here is the deployment manifest of a demo workload that can consume secrets:

```yaml
# k8s/Deployment.yaml

apiVersion: apps/v1
kind: Deployment
metadata:
  name: aegis-workload-demo
  namespace: default
  labels:
    app.kubernetes.io/name: aegis-workload-demo
spec:
  replicas: 1
  selector:
    matchLabels:
      app.kubernetes.io/name: aegis-workload-demo
  template:
    metadata:
      labels:
        app.kubernetes.io/name: aegis-workload-demo
    spec:
      serviceAccountName: aegis-workload-demo
      containers:
      - name: main
        image: aegishub/aegis-workload-demo-using-sidecar:0.9.1
        volumeMounts:
        # `main` shares this volume with `sidecar`.
        - mountPath: /opt/aegis
          name: aegis-secrets-volume
      - name: sidecar
        image: aegishub/aegis-sidecar:0.9.1
        volumeMounts:
        # /opt/aegis/secrets.json is the place the secrets will be at.
        - mountPath: /opt/aegis
          name: aegis-secrets-volume
        # Volume mount for SPIRE unix domain socket.
        - name: spire-agent-socket
          mountPath: /spire-agent-socket
          readOnly: true
      volumes:
      # A memory-backed volume is recommended (but not required) to keep
      # the secrets. The secrets can be stored in any kind of volume.
      - name: aegis-secrets-volume
        emptyDir:
          medium: Memory
      # Using SPIFFE CSI Driver to bind to the SPIRE Agent Socket
      # ref: https://github.com/spiffe/spiffe-csi
      - name: spire-agent-socket
        csi:
          driver: "csi.spiffe.io"
          readOnly: true

---

# This Service Account will be needed later:
apiVersion: v1
kind: ServiceAccount
metadata:
  name: aegis-workload-demo
  namespace: default
automountServiceAccountToken: false
```

> **Hint**
> 
> The manifest above is simplified, and most configuration parameters
> have been omitted to use defaults. You can find extensively-commented
> versions of these manifests [inside this installation folder][install-k8s].

[install-k8s]: https://github.com/shieldworks/aegis/tree/main/install/k8s

You’ll see that there are two images in this deployment:

* `aegishub/aegis-workload-demo`: This is the container that has the business logic.
* `aegishub/aegis-sidecar`: This **Aegis**-managed container injects 
  secrets to a place that our demo container can consume.

Here is the source code of the demo container’s app for the sake of completeness:

```go
package main

import (
	"fmt"
	"os"
	"time"
)

func sidecarSecretsPath() string {
	return "/opt/aegis/secrets.json"
}

func main() {
	for {
		dat, err := os.ReadFile(sidecarSecretsPath())
		if err != nil {
			fmt.Println("Will retry in 5 seconds…")
		} else {
			fmt.Println("secret: '", string(dat), "'")
		}

		time.Sleep(5 * time.Second)
	}
}
```

> **Read The Source Luke**
>
> [Check out the **Sidecar Demo** Git repository][git-sidecar]
> to view the source for this use case.

[git-sidecar]: https://github.com/shieldworks/aegis-workload-demo-using-sidecar "Aegis Workload Demo Using Sidecar"

Our demo app tries to read a secret file every 5 seconds forever.

Yet, how do we tell **Aegis** about our app so that it can identify it to
deliver secrets?

For this, there is an identity file that defines a `ClusterSPIFFEID` for
the workload:

```yaml
# k8s/Identity.yaml

{% raw %}apiVersion: spire.spiffe.io/v1alpha1
kind: ClusterSPIFFEID
metadata:
  name: aegis-workload-demo
spec:
  spiffeIDTemplate: "spiffe://aegis.ist\
    /workload/aegis-workload-demo\
    /ns/{{ .PodMeta.Namespace }}\
    /sa/{{ .PodSpec.ServiceAccountName }}\
    /n/{{ .PodMeta.Name }}"
  podSelector:
    matchLabels:
      app.kubernetes.io/name: aegis-workload-demo
  workloadSelectorTemplates:
  - "k8s:ns:{{ .PodMeta.Namespace }}"
  - "k8s:sa:{{ .PodSpec.ServiceAccountName }}"{% endraw %}
```

This identity descriptor, tells **Aegis** that the workload:

* Lives under a certain namespace,
* Is bound to a certain service account,
* And as a certain name.

When the time comes, **Aegis** will read this identity and learn about which
workload is requesting secrets. Then it can decide to deliver
the secrets (*because the workload is registered*) or deny dispatching them
(*because the workload is unknown/unregistered*).

> **ClusterSPIFFEID is an Abstraction**
> 
> Please note that `Identity.yaml` is not a random YAML file:
> It is a binding contract and abstracts a host of operations 
> behind the scenes.
> 
> For every `ClusterSPIFFEID` created this way, 
> `SPIRE` (*Aegis’ identity control plane*) will deliver an **X.509 SVID**
> bundle to the workload.
> 
> Therefore, creating a `ClusterSPIFFEID` is a way to **irrefutably**,
> **securely**, and **cryptographically** identify a workload.

Now that we have these manifests, we can apply them to deploy our workload.

Instead of creating things from scratch, I will use the ones that already exist
inside the `shieldworks/aegis` repo:

```bash
cd $WORKSPACE/aegis
cd ./install/k8s/demo-workload/using-sidecar
kubectl apply -f .
```

Then `kubectl get po` will give you something like this:

```bash 
kubectl get po

NAME                                  READY   STATUS    RESTARTS   AGE
aegis-workload-demo-fd4c8bf8b-6fcq2   2/2     Running   0          9s
```

Let’s check the logs of our pod:

```bash 
kubectl get logs aegis-workload-demo-fd4c8bf8b-6fcq2

Failed to read the secrets file. Will retry in 5 seconds…
Failed to read the secrets file. Will retry in 5 seconds…
Failed to read the secrets file. Will retry in 5 seconds…
secret: '  '
secret: '  '
secret: '  '
secret: '  '
…
```

What we see here that our workload checks for the secrets file and cannot
find it for a while, and displays a failure message. And once the sidecar
creates the secrets file that the workload pod was trying to parse, it
shows an empty string.

At this point, we have an empty secrets file because we haven’t registered
any secrets to this workload yet. 

Later, **Aegis Safe** will identify and acknowledge this workload
and deliver it an empty secrets file.

Next, we will add some secrets to that file using **Aegis Sentinel**:

> **What Is Aegis Sentinel**?
> 
> For all practical purposes, you can think of **Aegis Sentinel** as the
> “*bastion host*” you log in and execute sensitive operations.
> 
> In our case, we will register secrets to workloads using it.

## Registering Secrets to a Workload

For that, let’s first find where our sentinel is:

```bash
kubectl get po -n aegis-system

NAME                             READY   STATUS    RESTARTS   AGE
aegis-safe-c85c5c7d9-9k8dq       1/1     Running   0          11m
aegis-sentinel-b55f8bff5-7m7n7   1/1     Running   0          11m
```

Let’s execute a command to register a secret to our demo workload.

If you remember from the beginning of this tutorial, our demo workload had
a SPIFFE ID that matched the following template:

```text
{% raw %}spiffe://aegis.ist/workload/aegis-workload-demo
/ns/{{ .PodMeta.Namespace }}
/sa/{{ .PodSpec.ServiceAccountName }}
/n/{{ .PodMeta.Name }}{% endraw %}
```

The `aegis-workload-demo` part from that template is the **name** that **Aegis**
will identify this workload as.

Since we know the name of our workload, adding secrets to it will be a single
command that we’ll execute on **Aegis Sentinel**:

```bash 
kubectl exec -it aegis-sentinel-b55f8bff5-7m7n7 \
-n aegis-system \
-- aegis \
-w aegis-workload-demo \
-s '{"username":"Aegis", "password": "KeepYourSecrets"}'

OK

```

Once you do that, wait a few moments, and check the logs of our workload pod,
you can see the updated secret displayed on the console:

```text
kubectl logs aegis-workload-demo-fd4c8bf8b-6fcq2 -f

…
secret: '  '
secret: '  '
secret: ' {"username":"Aegis", "password": "KeepYourSecrets"} '
secret: ' {"username":"Aegis", "password": "KeepYourSecrets"} '
…
```

> **Aegis Sentinel Commands**
>
> You can execute
> `kubectl exec -it $sentinelPod -n aegis-sytem -- aegis --help`
> for a list of all available commands and command-line flags
> that **Aegis Sentinel** has.

## BONUS: Setting Aegis Safe’s Log Level

A relatively-hidden (*and subject to change*) feature of **Aegis** is you
can use **Aegis Sentinel** to set secrets and update the behavior of 
**Aegis** system components.

Here is how you increase the log verbosity of **Aegis Safe**, for example:

```bash 
kubectl exec aegis-safe-c85c5c7d9-9k8dq -it \
-n aegis-system -- \
aegis -w aegis-safe 
-s '{"logLevel": 6}'

OK

```

You can set the log level to any number from `1` to `6`.

```text 
# 1: logs are off, 6: highest verbosity.
# Off = 1, 
# Error = 2, 
# Warn = 3, 
# Info = 4, 
# Debug = 5, 
# Trace = 6
```

## Deploying a Demo Workload Without a Sidecar

You can also programmatically consume the **Aegis Secrets** API from your 
workload. That way, you will have more control over how you consume and cache 
your secrets, and you will not need to add a sidecar to your pod.

The advantage of this approach is: you are in charge.
The downside of it is: Well, you are in charge 🙂. 

But, jokes aside, your application will have to be 
more tightly bound to **Aegis** without a sidecar.

However, when you use a sidecar, your application does not have any idea of 
**Aegis**’s existence. From its perspective, it is merely reading from a file
that something magically updates every once in a while. This 
“*separation of concerns*” can make your application architecture more 
adaptable to changes.

As in anything, there is no one true way to do it. Your approach will depend
on your project’s requirements.

> **Read The Source Luke**
>
> [Check out the **SDK Demo** Git repository][git-sdk]
> to view the source for this use case.

[git-sdk]: https://github.com/shieldworks/aegis-workload-demo-using-sdk "Aegis Workload Demo Using SDK"

That part taken care of; let’s deploy a workload that does not use a sidecar.

Here is the deployment manifest for our workload:

```yaml
# k8s/Deployment.yaml

apiVersion: apps/v1
kind: Deployment
metadata:
  name: aegis-workload-demo
  namespace: default
  labels:
    app.kubernetes.io/name: aegis-workload-demo
spec:
  replicas: 1
  selector:
    matchLabels:
      app.kubernetes.io/name: aegis-workload-demo
  template:
    metadata:
      labels:
        app.kubernetes.io/name: aegis-workload-demo
    spec:
      serviceAccountName: aegis-workload-demo
      containers:
      - name: main
        image: aegishub/aegis-workload-demo-using-sdk:0.9.8
        volumeMounts:
        - name: spire-agent-socket
          mountPath: /spire-agent-socket
          readOnly: true
      volumes:
      - name: spire-agent-socket
        csi:
          driver: "csi.spiffe.io"
          readOnly: true
```

The `Identity.yaml` and `Service.yaml` will be the same as the demo workload
with a sidecar. And, as a reminder, you can find those files
[inside this **GitHub** folder][install-k8s] as well.

Here’s how the source code of `aegishub/aegis-workload-demo-using-sdk` looks like:

```go 
package main

import (
	"fmt"
	"github.com/shieldworks/aegis-sdk-go/sentry"
	"log"
	"time"
)

func main() {
	for {
		d, err := sentry.Fetch()

		if err != nil {
			fmt.Println("Will retry in 5 seconds…")
			fmt.Println(err.Error())
			time.Sleep(5 * time.Second)
			continue
		}

		if d.Data == "" {
			fmt.Println("no secret yet… will check again later.")
			time.Sleep(5 * time.Second)
			continue
		}

		fmt.Printf(
			"secret: updated: %s, created: %s, value: %s\n",
			d.Updated, d.Created, d.Data,
		)
		time.Sleep(5 * time.Second)
	}
}
```

Where all the heavy lifting is done by 
`github.com/shieldworks/aegis-sdk-go/sentry`.

The `sentry.Fetch()` operation will fetch the most recent secret from 
**Aegis Safe** and returns a go `struct` that our workload application 
can directly consume.

Having seen the source code, let’s deploy our workload and see if it
can fetch the secret that we registered a while ago:

```bash
cd $WORKSPACE/aegis 
cd ./install/k8s/demo-workload/using-sdk
kubectl apply -f .
```

Then let’s check our workload:

```bash
kubectl get po

NAME                                   READY   STATUS    RESTARTS   AGE
aegis-workload-demo-544dd799d8-rpzqc   1/1     Running   0          5s
```

It looks healthy. Let us view its logs too

```bash 
{% raw %}kubectl logs 

# lines are wrapped to fit on the web page
…
secret: updated: "Sun Jan 22 18:18:14 +0000 2023", 
created: "Sun Jan 22 18:09:47 +0000 2023", 
value: {"username":"Aegis", "password": "KeepYourSecrets"}
2023/01/22 19:14:38 fetch
secret: updated: "Sun Jan 22 18:18:14 +0000 2023", 
created: "Sun Jan 22 18:09:47 +0000 2023", 
value: {"username":"Aegis", "password": "KeepYourSecrets"}
2023/01/22 19:14:43 fetch
secret: updated: "Sun Jan 22 18:18:14 +0000 2023", 
created: "Sun Jan 22 18:09:47 +0000 2023", 
value: {"username":"Aegis", "password": "KeepYourSecrets"}{% endraw %}
…
```

It looks like our workload was able to receive its secret too. 

In addition,
we were able to fetch important meta-information about the secret,
such as the creation and update time stamps.

## Deploying a Demo Workload With an Init Container

In certain situations you might not have full control over the source code
of your workloads. For example, your workload can be a containerized third
party binary executable that you don’t have the source code of. It might
be consuming Kubernetes `Secret`s through injected environment variables, 

Luckily, with **Aegis Init Container** you can interpolate secrets stored in 
**Aegis Safe** to the `Data` section of Kubernetes `Secret`s at runtime to 
be consumed by the workloads.

> **Read The Source Luke**
>
> [Check out the **Init Container Demo** Git repository][git-init-container]
> to view the source for this use case.

Here is a sample deployment descriptor for your workload that uses
**Aegis Init Container**:

```yaml
# Deployment.yaml

apiVersion: apps/v1
kind: Deployment
metadata:
  name: aegis-workload-demo
  namespace: default
  labels:
    app.kubernetes.io/name: aegis-workload-demo
spec:
  replicas: 1
  selector:
    matchLabels:
      app.kubernetes.io/name: aegis-workload-demo
  template:
    metadata:
      labels:
        app.kubernetes.io/name: aegis-workload-demo
    spec:
      serviceAccountName: aegis-workload-demo
      containers:
      - name: main
        image: aegishub/aegis-workload-demo-using-init-container:0.13.0
        
        # These environment variables are interpolated dynamically at runtime.
        env:
        - name: SECRET
          valueFrom:
            secretKeyRef:
              name: aegis-secret-aegis-workload-demo
              key: VALUE
        - name: USERNAME
          valueFrom:
            secretKeyRef:
              name: aegis-secret-aegis-workload-demo
              key: USERNAME
        - name: PASSWORD
          valueFrom:
            secretKeyRef:
              name: aegis-secret-aegis-workload-demo
              key: PASSWORD

      initContainers:
      # Using Aegis Init Container.
      - name: init-container
        image: aegishub/aegis-init-container:0.13.0
        volumeMounts:
        # Volume mount for SPIRE unix domain socket.
        - name: spire-agent-socket
          mountPath: /spire-agent-socket
          readOnly: true
      volumes:
      - name: spire-agent-socket
        csi:
          driver: "csi.spiffe.io"
          readOnly: true
```

Then you can execute the following code to inject the secrets that the
container needs.

```bash
{% raw %}# ./hack/register.sh

# Find a Sentinel node.
SENTINEL=$(kubectl get po -n aegis-system \
  | grep "aegis-sentinel-" | awk '{print $1}')

# Execute the command needed to interpolate the secret.
kubectl exec "$SENTINEL" -n aegis-system -- aegis \
-w "aegis-workload-demo" \
-n "default" \
-s '{"username": "root", "password": "SuperSecret", "value": "AegisRocks"}' \
-t '{"USERNAME":"{{.username}}", "PASSWORD":"{{.password}}", "VALUE": "{{.value}}"}' \
-k

# Sit back and relax.{% endraw %}
```

The `Pod` that your `Deployment` manages will not initialize until you register
secrets to your workload.

Once you register secrets using the above command, **Aegis Init Container** will
exit with a success status code and let the main container initialize with the
updated Kubernetes `Secret`.

Here is a sequence diagram of how the secret is transformed (*open the image
in a new tab for a larger version*):

![Transforming Secrets](/assets/secret-transformation.png "Transforming Secrets")

You can also watch a demo video that implements the above flow. The video
visually explains the above concepts in greater detail:

[![Watch the video](/doks-theme/assets/images/capture.png)](https://vimeo.com/v0lkan/aegis-secrets)

[git-init-container]: https://github.com/shieldworks/aegis-workload-demo-using-init-container "Aegis Workload Demo Using Init Container"

## Conclusion

In this tutorial, you have seen how to register secrets to workloads using
**Aegis Sentinel**. First, we have used a **sidecar** to streamline the process
and keep the workload oblivious to the existence of **Aegis**. Then we used
[**Aegis** Go SDK][go-sdk] to skip the workload and directly consume secrets
from **Aegis Safe**. Finally, we looked into a use case where we can dynamically
create Kubernetes `Secret`s and bind them to workloads with the help of 
**Aegis Sentinel**, **Aegis Safe**, and **Aegis Init Container**.

For the interested, [the following section][next-section] covers the **Aegis** 
Go SDK’s methods in more detail; and the section after that is about
[**Aegis** Sentinel Command Line Interface Usage](/docs/sentinel).

[next-section]: /docs/sdk