## Project Folder Structure

```text
TODO: this has changed and needs an update.
Also, there are new(er) sequence diagrams, and the current diagrams
here are outdated too.
```

**Aegis** is a monorepo. Here’s a brief overview of essential files and folders:

* `./Makefile`: This is the file to install and test things.
* `README.md`: The very file that you are reading.
* `CONTRIBUTING.md`: Instructions about how to contribute to the project.
* `CODE_OF_CONDUCT.md`: The document that tells everyone to be nice human beings.
* `./demo`: A demo workload that can be used to test **Aegis**’s functionality.
* `./notary`: Source code of **Notary** (`aegis-notary`) which is a Kubernetes
  controller that acts as the mediator between **Safe** and workloads.
* `./safe`: Source code of **Safe** (`aegis-safe`). **Safe** is where all the
  secrets are stored, so you better keep it extra safe with proper RBAC. That is
  true for all **Aegis** components, but extra-true for **Safe**.
* `./sentinel`: Source code for **Sentinel**. **Sentinel** is a utility pod
  that you can diagnose the system and do administrative tasks.
* `./sidecar`: Source code of **Sidecar** (`aegis-sidecar`), a sidecar that’s
  injected to workloads to fetch secrets from **Safe**/

Each folder also has their associated `README.md` files to provide further
details about each child project.

## Components of Aegis

**Aegis**, as a system, has the following components.

### **Safe** (`aegis-safe`)

**Safe** stores secrets and dispatches them to workloads.

### **Sidecar** (`aegis-sidecar`)

`aegis-sidecar` is a sidecar that facilitates delivering secrets to workloads.

### **Notary** (`aegis-notary`)

A [Kubernetes Controller][k8s-controller] that lets **Safe** and the **Sidecar**s
securely communicate with each other.

### **Sentinel** (`aegis-sentinel`)

**Sentinel** is a pod you can shell in and do administrative tasks such as
registering secrets for workloads.

**Sentinel** is a Swiss army knife that should **NOT** run on production. If you
have to diagnose the production system using *Sentinel*, it is strongly
recommended that you *delete* Sentinel when you no longer need it—you have been
warned.

[k8s-controller]: https://kubernetes.io/docs/concepts/architecture/controller/ "Kubernetes Controller"

## Safe’s Bootstrapping Process

Bootstrapping is when **Notary** talks to **Safe** to deliver secure and
randomly generated admin token and notary token.

![Aegis Bootstrapping](assets/aegis-bootstrap.png "Bootstrapping")

**Notary** exchanges its token (*which is safe, but still known*)
with a more secure, unknown, randomly-generated token.

Bootstrapping is automatically done by **Notary**. It does not need any manual
intervention. Without successful bootstrapping, almost nothing works in the
system.

Bootstrapping is only done **once**. Executing the bootstrap flow on **Safe**
more than once is a **no-op**.

It is important to note that the notary id is **never** used anywhere in the
system after a successful bootstrap, and the bootstrapping typically happens
blazing fast.

## Dispatching Workload Ids and Secrets

**Notary** is also responsible for dispatching the workload ids and secrets to
workloads and notifying **Safe** about those ids and secrets, so that
workloads can communicate with **Safe** safely.

After a successful bootstrap, here’s how **Notary** dispatches ids and secrets
at a high level:

![Notary Dispatch](assets/aegis-notary-dispatch.png "Notary Dispatch")

The **id** provided to the workload is the **id** that is the value of
`aegis-workload-id` annotation in its deployment template. **Notary** will not
dispatch secrets to pods that don’t have this annotation.

## Workload Fetching Secrets

After receiving their **id** and **secret**, the **Sidecar** will
periodically call **Safe** to fetch and update the Pod’s secrets (*see the
bottom two line of the above sequence diagram*).
