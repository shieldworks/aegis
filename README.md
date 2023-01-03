![Aegis](assets/aegis-banner.png "Aegis")

## Status of This Software

This project is a **work in progress** (*WIP*) and has yet to be ready for 
production consumption. Use it at your own risk.

```text
TODO: some of the content in this file needs an update.
Especially thing around **notary** since we don’t use notary anymore.
```

## A Video Is Worth a Lot of Thousands of Words

[Here is a six-minute video introduction to Aegis][aegis-demo-video].

There is also [aegis.z2h.dev][aegis-website], the project’s website.

[aegis-website]: https://aegis.z2h.dev/ "Aegis webiste"
[aegis-demo-video]: https://vimeo.com/v0lkan/aegis "Introducing Aegis: A Cloud Native Solution to Secure Your Sensitive Data"

## About Aegis

**Aegis** is a lightweight secrets management solution that keeps your secrets
secret. With **Aegis**, you can rest assured that your sensitive data is 
always **secure** and **protected**. 

**Aegis** is perfect for securely storing arbitrary configuration information at a 
central location and securely dispatching it to workloads.

**Aegis** ensures that your secrets are only accessible to authorized workloads, 
helping you safeguard your business and protect against data breaches.

If you haven’t watched [this six-minute introductory video][aegis-demo-video] yet,
now might be a good time 🙂.

## System Requirements

**Aegis** has been recently tested with the following Kubernetes version:

```text
Client Version: v1.26.0
Kustomize Version: v4.5.7
Server Version: v1.25.3
```

Although not explicitly tested, any recent Kubernetes installation will work
just fine.

As in any secrets management solution, your compute and memory requirements
will depend on several factors, such as:

* The number of workloads in the cluster
* The number of secrets **Safe** has to manage (*see [architecture details](ARCHITECTURE.md)*)
* **Sidecar** poll frequency (*see [architecture details](ARCHITECTURE.md)*)

We recommend you benchmark with a realistic production-like
cluster and allocate your resources accordingly.

## Installation

You need a **Kubernetes** cluster and sufficient admin rights on that cluster to
install **Aegis**.

As of now, the only installation option is to clone the project and install
it using `make` as follows:

```bash 
git clone https://github.com/zerotohero-dev/aegis.git
cd aegis
make clean 
make install-all
```

To verify installation check out the `aegis-system` namespace:

```bash
kubectl get deployment -n aegis-system
# Output:
# NAME             READY   UP-TO-DATE   AVAILABLE
# aegis-safe       1/1     1            1
# aegis-sentinel   1/1     1            1
# aegis-notary     1/1     1            1
```

## Registering a Secret to a Workload

[Here is a sample Kubernetes deployment descriptor](demo/k8s/Deployment.yaml) 
for a workload that **Aegis** can inject secrets:

Assuming we have **Aegis** up and running, and the above workload deployed, 
to register a secret, you can execute the following API call from 
the **Sentinel** (*or any other place that has a network route to **Safe**).

```bash
http PUT http://aegis-safe.aegis-system.svc.cluster.local:8017/v1/secret \
  token=$ADMIN_TOKEN \
  key=aegis-workload-demo \
  value='{"username": "me@volkan.io", "password": "ToppyTopSecret"}'
```

Where `$ADMIN_TOKEN` is the token that you get using the **Sentinel**.

You can also do the same using **Sentinel**:

```bash 
# replace aegis-sentinel-aabbccdd-11223344
# with the pod name you have on the system.
kubectl exec -it aegis-sentinel-aabbccdd-11223344 -n aegis-system -- /bin/zsh
cd /tests
vim ./create.sh
./create.sh
```

## How Do I Get the Admin Token?

Unfortunately, this is a work in progress at the moment ☹️.

As a temporary workaround, the **Safe** Pod displays the admin token in its logs
during the bootstrapping process. Note that this is pretty insecure, and we
will remove the log line once we establish a secure way to deliver the admin 
token.

## Where Do I Store the Admin Token?

Keep the admin token safe; **do not** store it in source control; **do not** 
store it on disk as plain text. An ideal place to store it is a password manager 
or an encrypted file that only the administrators know how to decrypt.

## Design Decisions

Keeping *Aegis* **slim**, **secure**, and **boringly-easy** to install and
operate are the three pillars of the project.

[Check out the **Design Decisions** document](DESIGN_DECISIONS.md) for a 
deeper discussion about how we maintain the architectural balance in **Aegis**.

## Where **NOT** To Use Aegis

Aegis is **not** a Database, nor is it a distributed caching layer. Of course,
you may tweak it to act like one if you try hard enough, yet, that is
generally not a good use of the tool.

Aegis is suitable for storing secrets and dispatching them; however, it
is a *terrible* idea to use it as a centralized database to store everything
but the kitchen sink.

Use Aegis to store service keys, database credentials, access tokens,
etc. 

## Architecture Details

[Check out this **Architecture** document](ARCHITECTURE.md) for detailed
information about **Aegis**’s project folder structure, system design, sequence 
diagrams, workflows, and internal operating principles.

## Umm… How Do I Pronounce “Aegis”?

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

You can see the project’s progress [in this **Aegis MDP** board][mdp].

The board outlines what are the current outstanding work items, and what is
currently being worked on.

There is also [this **Aegis v1.0.0** board][v100] that contains longer-term
goals that we’ll start once the MDP board is mostly done.

[mdp]: https://github.com/orgs/zerotohero-dev/projects/2/views/2
[v100]: https://github.com/orgs/zerotohero-dev/projects/3/views/2

There is also [a text file](aegis.txt) that is a more free-form list of
issues. I sometimes jot things down there before creating more detailed GitHub
issues.

[todo-txt]: https://github.com/todotxt "todo.txt"

## Code Of Conduct

[Be a nice citizen](CODE_OF_CONDUCT.md).

## Contributing

It’s a bit chaotic around here, yet if you want to lend a hand,
[here are the contributing guidelines](CONTRIBUTING.md).

## Maintainers

As of now, I, [Volkan Özçelik][me], am the sole maintainer of **Aegis**.

[me]: https://github.com/v0lkan "Volkan Özçelik"

Please send your feedback, suggestions, recommendations, and comments to
[me@volkan.io](mailto:me@volkan.io). I’d love to have them.
