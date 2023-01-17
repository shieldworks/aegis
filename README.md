![Aegis](assets/aegis-banner.png "Aegis")

[spire]: https://spiffe.io/ "SPIFFE: Secure Production Identity Framework for Everyone"

## Status of This Software

This project is a **work in progress** (*WIP*) and has yet to be ready for 
production consumption. 

**Use it at your own risk**.

## A Video Is Worth a Lot of Thousands of Words

[![Watch the video](https://github.com/zerotohero-dev/aegis/blob/main/assets/capture.png)](https://vimeo.com/v0lkan/secrets)

There is also [aegis.z2h.dev][aegis-website], the project’s website.

[aegis-website]: https://aegis.z2h.dev/ "Aegis webiste"
[aegis-demo-video]: https://vimeo.com/v0lkan/secrets "Aegis: Keep your secrets… Secret." 

## About Aegis

**Aegis** is a delightfully-secure Kubernetes-native secrets store.

**Aegis** keeps your secrets secret. 

With **Aegis**, you can rest assured that your 
sensitive data is always **secure** and **protected**. 

**Aegis** is perfect for securely storing arbitrary configuration information at a 
central location and securely dispatching it to workloads.

**Aegis** ensures that your secrets are only accessible to authorized workloads, 
helping you safeguard your business and protect against data breaches.

If you haven’t watched [this six-minute introductory video][aegis-demo-video] yet,
now might be a good time 🙂.

## Projects

**Aegis** consists of the following sister projects:

* [**Aegis SPIRE**][aegis-spire]: **Aegis** uses [SPIRE][spire] as its Identity Control Plane.
* [**Aegis Safe**][aegis-safe]: **Safe** is the **secrets store** of **Aegis**.
* [**Aegis Sentinel**][aegis-sentinel]: **Sentinel** acts as a bastion that an operator (or a CI) can register secrets to workloads.
* [**Aegis Sidecar**][aegis-sidecar]: **Sidecar** is a utility that can help workloads retrieve secrets dynamically at runtime.
* [**Aegis Go SDK**][aegis-sdk-go]: **Go SDK** is a library that workloads can use to directly talk to **Safe** (instead of using the **Sidecar**).
* [**Aegis Core**][aegis-core]: Common modules that other projects share.
* [**Aegis Demo Workload (using Go SDK)**][aegis-workload-demo-using-sdk]: A demo workload that uses the **Go SDK** to talk to **Safe**.
* [**Aegis Demo Workload (using Aegis Sidecar)**][aegis-workload-demo-using-sidecar]: A demo workload dynamically injects secrets to itself using an **Aegis Sidecar**.
* [**Aegis Web**][aegis-web]: The sourcecode of <https://aegis.z2h.dev>.

[aegis-core]: https://github.com/zerotohero-dev/aegis-core
[aegis-safe]: https://github.com/zerotohero-dev/aegis-safe
[aegis-sdk-go]: https://github.com/zerotohero-dev/aegis-sdk-go
[aegis-sentinel]: https://github.com/zerotohero-dev/aegis-sentinel
[aegis-sidecar]: https://github.com/zerotohero-dev/aegis-sidecar
[aegis-spire]: https://github.com/zerotohero-dev/aegis-spire
[aegis-web]: https://github.com/zerotohero-dev/aegis-web
[aegis-workload-demo-using-sdk]: https://github.com/zerotohero-dev/aegis-workload-demo-using-sdk
[aegis-workload-demo-using-sidecar]: https://github.com/zerotohero-dev/aegis-workload-demo-using-sidecar

## Community

If you are a security enthusiast, [**join Aegis’ Slack Workspace**][slack-invite]
and let us change the world together 🤘.

[slack-invite]: https://join.slack.com/t/aegis-6n41813/shared_invite/zt-1myzqdi6t-jTvuRd1zDLbHX0gN8VkCqg "Join aegis.slack.com"

## Installation

You need a **Kubernetes** cluster and sufficient admin rights on that cluster to
install **Aegis**.

```bash 
./hack/install.sh
```

You can also deploy a demo workload to experiment with it:

```bash
# Demo workload that uses `aegis-sidecar` 
./hack/install-workload-using-sidecar.sh

# Demo workload that directly talks to `aegis-safe` using Aegis Go SDK
./hack/install-workload-using-sdk.sh
```

To verify installation check out the `aegis-system` namespace:

```bash
kubectl get deployment -n aegis-system

# Output:
#
# NAME             READY   UP-TO-DATE   AVAILABLE
# aegis-safe       1/1     1            1
# aegis-sentinel   1/1     1            1
```

## Registering a Secret to a Workload

You can use **Sentinel** to add a secret to a workload:

```bash
# Change `aegis-sentinel-aabbccdd11223344` with the name of the Sentinel
# pod.
kubectl exec -it aegis-sentinel-aabbccdd11223344 -n aegis-system \
-- /bin/sentinel \ 
-w demo-workload # Name of the workload \
-s '{"username":"root@admin-db", \
"password":"KeepYourSecrets!."}' # The secret to bind to the workload.
```

**Sentinel** is the only entry point that an operator can register secrets
to the system.

## Configuring Aegis Safe Using Sentinel

Starting with `v0.9.1`, it is possible to use the Aegis API to dynamically 
configure Aegis at runtime without having to restart the workloads:

```bash 
# Change the log level from WARN (3) default to ERROR (2):
kubectl exec $sentinelPodName -it -- aegis -w "aegis-safe" -s '{"logLevel":2}'

# Here are the possible values for `logLevel`:
#   Off   = 1
#   Error = 2
#   Warn  = 3
#   Info  = 4
#   Debug = 5
#   Trace = 6
# The higher the level is, the more verbose will the logs be.
```

## How Do I Get the Root Token? Where Do I Store It?

Unlike some other secret vaults, you do not need an admin token
to operate Aegis 🙂.

Benefits of this approach is: It helps the Ops team `#sleepmore`, since 
everything is automated, and you won’t have to manually unlock Aegis upon 
a system crash, for example.

However, there’s no free lunch, and as the operator of a production system, 
your homework is to secure access to **Sentinel**.

**Aegis** leverages Kubernetes security primitives and modern cryptography 
to secure access to secrets. And **Sentinel** is the **only** system part that 
has direct write access to the secrets store. Therefore, once you secure your 
access to **Sentinel** with proper RBAC and policies, you secure your access 
to your secrets.

We believe that this approach is **Kubernetes-native**, convenient, simpler, 
and delightfully secure (*as opposed to being “annoyingly secure”*).

Here are the resource allocation reported by `kubectl top` for a demo setup
on a single-node minikube cluster to give an idea:

```text 
NAMESPACE     WORKLOAD            CPU(cores) MEMORY(bytes)
aegis-system  aegis-safe          1m         9Mi
aegis-system  aegis-sentinel      1m         3Mi
default       aegis-workload-demo 2m         7Mi
spire-system  spire-agent         4m         35Mi
spire-system  spire-server        6m         41Mi
```

Note that 1000m is 1 full CPU core.

## System Requirements

**Aegis** has been recently tested with the following Kubernetes version:

```text
Client Version: v1.26.0
Kustomize Version: v4.5.7
Server Version: v1.25.3
```

Although not explicitly tested, any recent Kubernetes installation will
likely work just fine.

As in any secrets management solution, your compute and memory requirements
will depend on several factors, such as:

* The number of workloads in the cluster
* The number of secrets **Safe** (*Aegis’ Secrets Store*) has to manage
  (*see [architecture details](ARCHITECTURE.md)*)
* The amount of workloads interacting with **Safe**
  (*see [architecture details](ARCHITECTURE.md)*)
* **Sidecar** poll frequency (*see [architecture details](ARCHITECTURE.md)*)
* etc.

We recommend you benchmark with a realistic production-like
cluster and allocate your resources accordingly.

## Design Decisions

Keeping *Aegis*, **Kubernetes-native**, **slim**, **secure**, and 
**boringly-easy** to install and operate are the pillars of the project.

[Check out the **Design Decisions** document](DESIGN_DECISIONS.md) for a 
deeper discussion about how we maintain the architectural balance in **Aegis**.

## Where **Not** to Use Aegis

Aegis is **not** a Database, nor is it a distributed caching layer. Of course,
you may tweak it to act like one if you try hard enough, yet, that is
generally not a good use of the tool.

Aegis is suitable for storing secrets and dispatching them; however, it
is a *terrible* idea to use it as a centralized database to store everything
but the kitchen sink.

Use Aegis to store service keys, database credentials, access tokens,
etc. 

## Technologies Used

Without these technologies, implementing **Aegis** would have been a very 
hard, time-consuming, and error-prone endeavor. 

* [SPIFFE and SPIRE][spire] for establishing an Identity Control Plane.
* [Mozilla Sops][sops] (*in design phase*) to enable integration with cloud 
  secrets stores, such as AWS KMS, GCP KMS, Azure KeyVault, and even HashiCorp 
  Vault.
* [Age Encryption][age] (*in design phase*) to enable out-of-memory encrypted 
  backup of the secrets store for disaster recovery.

[spire]: https://spiffe.io/ "SPIFFE: Secure Production Identity Framework for Everyone"
[sops]: https://github.com/mozilla/sops "Sops: Simple and flexible tool for managing secrets"
[age]: https://github.com/FiloSottile/age "Age: A secure and modern encryption tool"

## Architecture Details

[Check out this **Architecture** document](ARCHITECTURE.md) for detailed
information about **Aegis**’s project folder structure, system design, sequence 
diagrams, workflows, and internal operating principles.

## Uninstallation

To uninstall Aegis from your cluster, execute the following script:

```bash
./hack/uninstall.sh
```

## One More Thing… How Do I Pronounce “Aegis”?

“Aegis” is a word of Greek origin and is pronounced `EE-jiss`.

[Here’s a YouTube pronunciation guide][pronounce].

[pronounce]: http://www.youtube.com/watch?v=x4bUgXWdNfM

**Aegis** has two definitions:

1. (Classical Mythology) The shield or breastplate of Zeus or Athena, bearing
   at its center the head of the Gorgon.
2. Protection; support.

Here’s an image of an aegis (*shield*) as depicted in Greek mythology:

![aegis](assets/aegis-shield.jpg "aegis")

## What’s Coming Up Next?

You can see the project’s progress [in these **Aegis** boards][mdp].

The board outlines what are the current outstanding work items, and what is
currently being worked on.

[mdp]: https://github.com/zerotohero-dev/aegis/projects?query=is%3Aopen

There is also [a text file](aegis.txt) that is a more free-form list of
issues. I sometimes jot things down there before creating more detailed GitHub
issues.

[todo-txt]: https://github.com/todotxt "todo.txt"

## Code Of Conduct

[Be a nice citizen](CODE_OF_CONDUCT.md).

## Contributing

It’s a bit chaotic around here, yet if you want to lend a hand,
[here are the contributing guidelines](CONTRIBUTING.md).

## Security

We take security **very** seriously.

[Follow the instructions here to report any security issues for **Aegis**
or any of its sister projects](SECURITY.md).

## Maintainers

As of now, I, [Volkan Özçelik][me], am the sole maintainer of **Aegis**.

[me]: https://github.com/v0lkan "Volkan Özçelik"

Please send your feedback, suggestions, recommendations, and comments to
[me@volkan.io](mailto:me@volkan.io). I’d love to have them.
