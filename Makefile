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

# Installs without rebuilding apps.
install: install-safe install-sidecar install-sentinel

# Installs the demo app to play with.
demo: install-demo

# Builds and installs everything.
# You will need dockerhub write access for this task.
build-and-install: spire all-demo all-safe all-sidecar all-sentinel

.PHONY: spire
spire:
	cd spire && $(MAKE) deploy
	sleep 15 # give some time for SPIRE to bring itself up.

all-demo:
	cd demo && $(MAKE) all

all-safe:
	cd safe && $(MAKE) all

all-sidecar:
	cd sidecar && $(MAKE) all

all-sentinel:
	cd sentinel && $(MAKE) all

install-demo:
	cd demo && $(MAKE) deploy

install-safe:
	cd safe && $(MAKE) deploy

install-sidecar:
	cd sidecar && echo "no-op"

install-sentinel:
	cd sentinel && $(MAKE) deploy
