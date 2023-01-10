#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                  <aegis.z2h.dev>
#     .\_/.
#

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
# and all the local changes on them
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
	cd ../aegis-core
	cd ../aegis-core; git stash; git checkout main; git pull;
	cd ../aegis-sdk-go
	cd ../aegis-sdk-go; git stash; git checkout main; git pull;
	cd ../aegis-sentinel
	cd ../aegis-sentinel; git stash; git checkout main; git pull;
	cd ../aegis-sidecar
	cd ../aegis-sidecar; git stash; git checkout main; git pull;
	cd ../aegis-safe
	cd ../aegis-web
	cd ../aegis-web; git stash; git checkout main; git pull;
	cd ../aegis-demo-workload-using-sidecar;
	cd ../aegis-demo-workload-using-sdk;

all:
	@echo ""
	@echo "You can use the following commands based on your needs as follows:"
	@echo ""
	@echo "Clean former installations:"
	@echo "    make clean"
	@echo ""
	@echo "If you do not have SPIRE set up:"
	@echo "    make spire; make aegis"
	@echo ""
	@echo "If you do have SPIRE set up already:"
	@echo "    make aegis"
	@echo ""
	@echo "If you want to run a demo workload to test things out:"
	@echo "    make demo"
	@echo ""
	@echo "If you have dockerhub access:"
	@echo "    make build"
	@echo ""

# A shortcut to install SPIRE, Safe, and Sentinel
install: spire aegis

# SPIRE is required for Workload-to-Safe, Safe-to-Workload, Sentinel-to-Safe
# and Safe-to-Sentinel communication. Better to install it first before
# installing aegis.
.PHONY: spire
spire:
	cd ../aegis-spire && $(MAKE) deploy
	sleep 15 # give some time for SPIRE to bring itself up.

# Installs without rebuilding apps.
aegis: install-safe install-sentinel

# Installs the demo app to play with.
demo: install-demo

install-demo:
	cd demo && $(MAKE) deploy

install-safe:
	cd ../aegis-safe && $(MAKE) deploy

install-sentinel:
	cd ../aegis-sentinel && $(MAKE) deploy

# Fetches the recent changes.
# Then, builds and installs everything.
# You will need dockerhub write access for this task.
# Also note that any uncommitted changes will be stashed.
build: pull spire build-demo build-safe build-sidecar build-sentinel

build-demo:
	cd ../aegis-demo && $(MAKE) all

build-safe:
	cd ../aegis-safe && $(MAKE) all

build-sidecar:
	cd ../aegis-sidecar && $(MAKE) all

build-sentinel:
	cd ../aegis-sentinel && $(MAKE) all
