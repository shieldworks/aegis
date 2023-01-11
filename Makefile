#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                  <aegis.z2h.dev>
#     .\_/.
#

# Bump versions (to push new container images)
bump:
	./hack/bump-version.sh

# Tag a new release when you are sure everything works.
tag:
	./hack/tag.sh

# Cleans the former deployment.
clean:
	@if kubectl get ns | grep aegis-system; then \
		kubectl delete ns spire-system; \
		kubectl delete ns aegis-system; \
		kubectl delete deployment aegis-workload-demo -n default; \
		kubectl delete ClusterSPIFFEID aegis-workload-demo; \
		kubectl delete ClusterSPIFFEID aegis-sentinel; \
		kubectl delete ClusterSPIFFEID aegis-safe; \
	else \
  		echo "Nothing to clean."; \
	fi

# Clones all satellite repos into the workspace.
clone:
	cd ..; git clone git@github.com:zerotohero-dev/aegis-spire.git
	cd ..; git clone git@github.com:zerotohero-dev/aegis-core.git
	cd ..; git clone git@github.com:zerotohero-dev/aegis-sdk-go.git
	cd ..; git clone git@github.com:zerotohero-dev/aegis-safe.git
	cd ..; git clone git@github.com:zerotohero-dev/aegis-sentinel.git
	cd ..; git clone git@github.com:zerotohero-dev/aegis-sidecar.git
	cd ..; git clone git@github.com:zerotohero-dev/aegis-workload-demo-using-sidecar.git
	cd ..; git clone git@github.com:zerotohero-dev/aegis-workload-demo-using-sdk.git
	cd ..; git clone git@github.com:zerotohero-dev/aegis-web.git

# Destructively and irreversibly removes all the satellite repos
# and all the local changes on them.
rimraf:
	cd ..; rm -rf aegis-spire
	cd ..; rm -rf aegis-core
	cd ..; rm -rf aegis-sdk-go
	cd ..; rm -rf aegis-safe
	cd ..; rm -rf aegis-sentinel
	cd ..; rm -rf aegis-sidecar
	cd ..; rm -rf aegis-workload-demo-using-sidecar
	cd ..; rm -rf aegis-workload-demo-using-sdk
	cd ..; rm -rf aegis-web

pull:
	cd ../aegis-spire;
	cd ../aegis-spire; git stash; git checkout main; git pull;
	cd ../aegis-core
	cd ../aegis-core; git stash; git checkout main; git pull;
	cd ../aegis-sdk-go
	cd ../aegis-sdk-go; git stash; git checkout main; git pull;
	cd ../aegis-sentinel
	cd ../aegis-sentinel; git stash; git checkout main; git pull;
	cd ../aegis-sidecar
	cd ../aegis-sidecar; git stash; git checkout main; git pull;
	cd ../aegis-safe
	cd ../aegis-safe; git stash; git checkout main; git pull;
	cd ../aegis-web
	cd ../aegis-web; git stash; git checkout main; git pull;
	cd ../aegis-workload-demo-using-sidecar;
	cd ../aegis-workload-demo-using-sidecar; git stash; git checkout main; git pull;
	cd ../aegis-workload-demo-using-sdk;
	cd ../aegis-workload-demo-using-sdk; git stash; git checkout main; git pull;

# For repo-admin-use only.
build: build-demo-sidecar build-demo-sdk build-safe build-sidecar build-sentinel

build-demo-sidecar:
	cd ../aegis-workload-demo-using-sidecar && $(MAKE) build && $(MAKE) bundle && $(MAKE) push

build-demo-sdk:
	cd ../aegis-workload-demo-using-sdk && $(MAKE) build && $(MAKE) bundle && $(MAKE) push

build-safe:
	cd ../aegis-safe && $(MAKE) build && $(MAKE) bundle && $(MAKE) push

build-sidecar:
	cd ../aegis-sidecar && $(MAKE) build && $(MAKE) bundle && $(MAKE) push

build-sentinel:
	cd ../aegis-sentinel && $(MAKE) build && $(MAKE) bundle && $(MAKE) push

# Deploys Aegis to the cluster.
deploy: spire safe sentinel

# SPIRE is required for Workload-to-Safe, Safe-to-Workload, Sentinel-to-Safe
# and Safe-to-Sentinel communication. Better to install it first before
# installing aegis.
spire:
	cd ../aegis-spire && $(MAKE) deploy
	sleep 15 # give some time for SPIRE to bring itself up.

# Sentinel acts as a bastion.
sentinel:
	cd ../aegis-sentinel && $(MAKE) deploy

# Safe is the secrets store.
safe:
	cd ../aegis-safe && $(MAKE) deploy

# Installs the demo app to play with.
demo-sidecar:
	cd ../aegis-workload-demo-using-sidecar && $(MAKE) deploy

# Installs the demo app to play with.
demo-sdk:
	cd ../aegis-workload-demo-using-sdk && $(MAKE) deploy