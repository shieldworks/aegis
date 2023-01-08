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
	cd spire && $(MAKE) deploy
	sleep 15 # give some time for SPIRE to bring itself up.

# Installs without rebuilding apps.
aegis: install-safe install-sentinel

# Installs the demo app to play with.
demo: install-demo

install-demo:
	cd demo && $(MAKE) deploy

install-safe:
	cd safe && $(MAKE) deploy

install-sentinel:
	cd sentinel && $(MAKE) deploy

# Builds and installs everything.
# You will need dockerhub write access for this task.
build: spire build-demo build-safe build-sidecar build-sentinel

build-demo:
	cd demo && $(MAKE) all

build-safe:
	cd safe && $(MAKE) all

build-sidecar:
	cd sidecar && $(MAKE) all

build-sentinel:
	cd sentinel && $(MAKE) all
