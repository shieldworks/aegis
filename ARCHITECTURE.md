![Aegis](assets/aegis-banner.png "Aegis")

## Introduction

This documents outlines **Aegis**’ architecture details and project structure.

It is especially useful if you want to have an in-depth understanding of **Aegis**
to contribute to the project.

## Project Folder Structure

**Aegis** is a monorepo. Here’s a brief overview of essential files and folders:

* `./Makefile`: This is the file to install and test things.
* `README.md`: The very file that you are reading.
* `CONTRIBUTING.md`: Instructions about how to contribute to the project.
* `CODE_OF_CONDUCT.md`: The document that tells everyone to be nice human beings.
* `./demo`: A demo workload that can be used to test **Aegis**’s functionality.
* `./spire`: Contains a [SPIFFE/SPIRE][spire] installation to be used as an
  identity control plane.
* `./safe`: Source code of **Safe** (`aegis-safe`). **Safe** is where all the
  secrets are stored, so you better keep it extra safe with proper RBAC. That is
  true for all **Aegis** components, but extra-true for **Safe**.
* `./sentinel`: Source code for **Sentinel**. **Sentinel** is a utility pod
  that you can diagnose the system and do administrative tasks.
* `./sidecar`: Source code of **Sidecar** (`aegis-sidecar`), a sidecar that’s
  injected to workloads to fetch secrets from **Safe**/

[spire]: https://spiffe.io/ "SPIFFE: Secure Production Identity Framework for Everyone"

Each folder also has their associated `README.md` files to provide further
details about each child project.

## Components of Aegis

**Aegis**, as a system, has the following components.

### SPIRE

**SPIRE** is what makes communication within **Aegis** components and workloads
possible. It dispatches x.509 SVID certificates to the required parties to make
secure mTLS communication possible.

### **Safe** (`aegis-safe`)

**Safe** stores secrets and dispatches them to workloads.

### **Sidecar** (`aegis-sidecar`)

`aegis-sidecar` is a sidecar that facilitates delivering secrets to workloads.

### **Sentinel** (`aegis-sentinel`)

**Sentinel** is a pod you can shell in and do administrative tasks such as
registering secrets for workloads.

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
