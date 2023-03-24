---
#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                    <aegis.ist>
#     .\_/.
#

layout: default
keywords: Aegis, release, maintenance
title: Releasing a New Version
description: guidelines to maintain, sign, and publish code
micro_nav: true
page_nav:
  prev:
    content: Miscellaneous
    url: '/misc'
---

## Configuring Minikube Local Registry for Linux

Switch to the **Aegis** project folder
Then, delete any existing minikube cluster.

```bash
cd $WORKSPACE/aegis
make k8s-delete
```

Then start the **Minikube** cluster.

```bash 
cd $WORKSPACE/aegis
make k8s-start
```

This will also start the local registry. However, you will need to 
eval some environment variables to be able to use Minikube’s registry insted
of the local Docker registry.

```bash 
cd $WORKSPACE/aegis
eval $(minikube docker-env)

echo $DOCKER_HOST
# example: tcp://192.168.49.2:2376
#
# Any non-empty value to `echo $DOCKER_HOST` means that the environment
# has been set up correctly.
```

> **Add Docker Settings to Your Profile**
> 
> Note that you might need to execute `eval $(minikube docker-env)` whenever 
> you need to use Minikube Docker registry locally. A practical way to do this
> is to add the instruction to your profile file (*i.e., `.bash_profile`, 
> `.profile`, `.zprofile`, etc.*).

## Cutting a Release

Before every release cut, follow the steps outlined below.

### 1. Double Check

Ensure all changes that need to go to a release in all
repositories have been merged to `main`.

```bash
# Checks if there are any uncommitted changes
make status
```

Also check whether you are using the `minikube` version of the Docker registry.
`make help` can help you in that.

```bash 
make help

# Output
#                          ---------------------------------------------------
#             Docker Host: tcp://192.168.49.2:2376
# Minikube Active dockerd: minikube
#                          ---------------------------------------------------
#                    PREP: make delete-k8s;make start-k8s;make clean;make sync;
#                    TEST: make build-local;make deploy-local;make test-local;
#                 RELEASE: make bump;make build;make tag
#                          ---------------------------------------------------
```
If you see `Minikube Active dockerd` as `minikube` then your registry has been
set up properly. If not, execute the following script to set up your regisryy
and then execute `make help` to verify the changes:

```bash 
# shellcheck disable=SC2046
eval $(minikube -p minikube docker-env)
```

### 2. Fetch Assets

```bash
# Uninstall Aegis:
make clean
# Pull recent code:
make pull
# Copy over recent manifests:
make sync
```

### 3. Bump Versions

First, double check all required `Makefile`s and `Deployment.yaml`s have
the same version.

Then edit `./hack/bump-version.sh` to move everything to the new version.

Then execute the following:

```bash
make bump
make sync
```

### 4. Build and Deploy Locally

Before running the following commands, make sure you haven an **insecure**
docker registry running at `localhost:5000`.

For minikube, `minikube addons enable registry` should work.

```bash
make build-local
make deploy-local
```

### 5. Run Integration Tests

```bash
make test-local
```

Ensure that the program succeeds.
It can take several minutes to complete.

### 6. Push the Updated Code

At least, there are new version numbers in the manifests.
They need to be pushed before tagging a release:

```bash
make commit
```

### 7. Tag a New Release

```bash
make build
make tag
```

### 8. All Set 🎉

You’re all set.

Happy releasing.