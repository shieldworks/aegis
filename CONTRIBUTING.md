# Contributing

When contributing to this repository, please first discuss the change you wish 
to make via issue, email, or any other method with the owners of this repository 
before making a change. That will save time for everyone.

Please note [we have a code of conduct](CODE_OF_CONDUCT.md), please follow it 
in all your interactions with the project.

## Pull Request Process

1. Ensure that all components build and function properly on a local Kubernetes
   cluster (*such as minikube*).
2. Update necessary `README.md` documents.
3. Keep the pull request as granular as possible, since reviewing large amount
   of code can be error-prone (*not to mention time-consuming for the reviewers*).
4. Follow the discussion under the pull request and proceed accordingly.

## Building Aegis for Development

> **Caveat**
> 
> This build workflow is still a work in progress.
> 
> You will likely need to hack things around with a local repository.

First, create a workspace folder and `cd` into it:

```bash 
mkdir $HOME/WORKSPACE
cd $HOME/WORKSPACE
```

Then clone `aegis` and its friends:

```bash 
# Switch to the workspace folder.
cd $HOME/WORKSPACE

# Clone the main Aegis repo.
git clone git@github.com:zerotohero-dev/aegis.git

# Clone other aegis-* repositories.
# The exact command will differ if you are working on a forked repository.
git clone git@github.com:zerotohero-dev/aegis-spire.git
git clone git@github.com:zerotohero-dev/aegis-core.git
git clone git@github.com:zerotohero-dev/aegis-sdk-go.git
git clone git@github.com:zerotohero-dev/aegis-safe.git
git clone git@github.com:zerotohero-dev/aegis-sentinel.git
git clone git@github.com:zerotohero-dev/aegis-sidecar.git
git clone git@github.com:zerotohero-dev/aegis-workload-demo-using-sidecar.git
git clone git@github.com:zerotohero-dev/aegis-workload-demo-using-sdk.git
git clone git@github.com:zerotohero-dev/aegis-web.git
```

Alternatively, if you have contributor access, you can execute `make pull`
to pull all the satellite repositories.

```bash 
cd $HOME/WORKSPACE/aegis
make pull
```

If you have dockerhub access, you can bump versions and publish new artifacts:

```bash 
cd $HOME/WORKSPACE/aegis
./hack/bump-version.sh
make build
```

## For Organization Admins

When in doubt, [follow the runbook](runbook.txt).
