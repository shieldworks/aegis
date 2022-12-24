![Aegis](assets/aegis-banner.png "Aegis")

## Status of This Software

This project is a work in progress and has yet to be ready for production 
consumption. Use it at your own risk.

## Before Anything… How Do I Pronounce “Aegis”?

“Aegis” is a word of Greek origin and is pronounced `EE-jiss`. 

[Here’s a YouTube pronunciation guide][pronounce].

[pronounce]: http://www.youtube.com/watch?v=x4bUgXWdNfM

**Aegis** has two definitions:

1. (Classical Mythology) The shield or breastplate of Zeus or Athena, bearing 
   at its center the head of the Gorgon. 
2. Protection; support.

Here’s an image of an aegis (*shield*) as depicted in Greek mythology:

![aegis](assets/aegis-shield.jpg "aegis")

## A Video Is Worth a Lot of Thousands of Words

[Here is a six-minute video introduction to Aegis][aegis-demo-video].

There is also [aegis.z2h.dev][aegis-website], the project’s website.

[aegis-website]: https://aegis.z2h.dev/ "Aegis webiste"
[aegis-demo-video]: https://vimeo.com/v0lkan/aegis "Introducing Aegis: A Cloud Native Solution to Secure Your Sensitive Data"

## About Aegis

**Aegis** is perfect for storing arbitrary configuration information at a 
central location and securely dispatching it to workloads.

**Aegis** is a lightweight secrets management solution that keeps your secrets
secret. With **Aegis**, you can rest assured that your sensitive data is 
always **secure** and **protected**. 

**Aegis** ensures that your secrets are only accessible to authorized workloads, 
helping you safeguard your business and protect against data breaches.

If you haven’t watched [this six-minute introductory video yet][aegis-demo-video],
now might be a good time 🙂.

Keeping *Aegis* **slim**, **secure**, and **boringly-easy** to install and 
operate are the three pillars of the project. 

We follow the **guidelines** outlined in the next few sections to achieve these.

Since **Aegis** is still in development, some of these goals discussed in the 
following sections may still need to be fully implemented. Regardless, 
they are the **guiding principles** we steer towards while shaping the future 
of **Aegis**.

### Be Cloud Native

**Aegis** is designed to run on Kubernetes and **only** on Kubernetes. 
That helps us leverage Kubernetes concepts like *Operators*, *Custom Resources*, 
and *Controllers* to our advantage to simplify workflow and state management. 

If you are looking for a solution that runs outside Kubernetes or as a 
standalone binary, then Aegis is not the Secrets Store you’re looking for.

### Do One Thing Well

At a 5000-feet level, **Aegis** is a secure Key-Value store. It can securely 
store arbitrary values that you, as an administrator, associate with keys. It 
does that, and it does that well.

If you are searching for a solution to create and store X.509 certificates, create
dynamic secrets, automate your PKI infrastructure, federate your identities,
use as an OTP generator, policy manager, in short, anything other than a
secure key-value store, then Aegis is likely not the solution you are looking 
for.

### Have a Minimal and Intuitive API

As an administrator, there is a limited set of API endpoints that you can 
interact with **Aegis**. This makes **Aegis** easy to manage. In addition,
a minimal set of APIs means a smaller attack surface, a smaller footprint, and
a codebase that is easy to understand, test, audit, and develop; all good things.

### Be Practically Secure

Corollary: Do not be annoyingly secure. Provide a delightful user experience
while taking security seriously.

**Aegis** is a secure solution, yet still delightful to operate. 
You won’t have to jump over the hoops or wake up in the middle of the night
to keep it up and running. Instead, **Aegis** will work seamlessly, as if it 
doesn’t exist at all.

## Secure By Default

**Aegis** stores your sensitive data in memory by default. Yes, that brings 
up resource limitations since you cannot store a gorilla holding a banana and
the entire jungle in your store; however, a couple of gigabytes of RAM can store
a lot of plain text secrets in it, so it’s good enough for most practical 
purposes. 

More importantly, almost all modern instruction set architectures and 
operating systems implement [*memory protection*][memory-protection]. The primary 
purpose of memory protection is to prevent a process from accessing memory that 
has not been allocated to it. This prevents a bug or malware within a process 
from affecting other processes or the operating system itself.

[memory-protection]: https://en.wikipedia.org/wiki/Memory_protection "Memory Protection (Wikipedia)"

Therefore, reading a variable’s value from a process’s memory is practically 
impossible unless you attach a debugger to it.

So, until we implement ways to securely store a backup of the state data
encrypted on disk, all the secrets **Aegis** has will be held in memory
only: Never persisted to disk and never written or streamed to log files.

Other related works are in progress to make **Aegis**’s system 
architecture secure by default. We are slowly and steadily getting there.
You can check out [aegis.txt](aegis.txt) for the overall progress of 
what has been done and what is in progress.

## Where NOT To Use Aegis

Aegis is **not** a Database, nor is it a distributed caching layer. Of course,
you may tweak it to act like one if you try hard enough, yet, that is 
generally not a good use of the tool.

Aegis is suitable for storing secrets and dispatching them; however, it
is a *terrible* idea to use it as a centralized database to store everything
but the kitchen sink.

Use Aegis to store service keys, database credentials, access tokens, 
etc. However, **do not** use Aegis to store the username and
passwords of your 1000 customers: That’s what a database is for (*where you
hopefully hash and salt the passwords before you store them*).

## Hey, Where Are the GitHub Issues?

Right now, [there is a single text file](aegis.txt) that lists all the issues
in [todo.txt format][todo-txt]. 

This is the fastest way to bootstrap a project, and it is also the best way for 
me to manage things as I am the only developer working on the project with a 
**very** tight time budget. I literally do not have time to triage and label 
GitHub issues.

Once the project matures enough and I have a more maintainable development
burden, I’ll move items on the [aegis.txt](aegis.txt)
file to GitHub issues and GitHub projects.

[todo-txt]: https://github.com/todotxt "todo.txt"

## Project Timeline

Check [aegis.txt](aegis.txt) for task prioritization and timeline information.

For a quick guideline to parse that text file.

Let’s say you see the following line in `aegis.txt`:

```text 
(A) a mini operations manual on README.md due:2022-12-24 +aegis @▶️
```

It would mean:

* This is very important (`(A)`: *sorted alphabetically: 
  `A` is the most important, `Z` is least important*)
* It is associated with the `+aegis` projects.
* It is a work in progress (`@▶️`).
* It is guesstimated to be done by `2022-12-24`, with no promises.

## Installation

You need a **Kubernetes** cluster and sufficient admin rights on that cluster to
install **Aegis**.

As of now, the only installation option is to clone the project and install
it using `make` as follows:

```bash 
git clone https://github.com/zerotohero-dev/aegis.git
cd aegis
make install
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

## Architecture Details

[Check out this document](ARCHITECTURE.md) for detailed information about
**Aegis**’s project folder structure, system design, sequence diagrams, 
workflows, and internal operating principles.

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
* **Notary** refresh interval (*see [architecture details](ARCHITECTURE.md)*)
* **Sidecar** poll frequency (*see [architecture details](ARCHITECTURE.md)*)

We recommend you benchmark with a realistic production-like 
cluster and allocate your resources accordingly.

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
