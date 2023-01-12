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

# Cleans the former Aegis deployment.
clean:
	./hack/uninstall.sh

# Clones all satellite repos into the workspace.
clone:
	./hack/clone.sh

# Destructively and irreversibly removes all the satellite repos
# and all the local changes on them.
rimraf:
	./hack/rimraf.sh

# Switches to the `main` branches and pulls everything.
pull:
	./hack/pull.sh

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