#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                  <aegis.z2h.dev>
#     .\_/.
#

clean:
	@if kubectl get ns | grep aegis-system; then \
		kubectl delete ns aegis-system; \
		kubectl delete deployment aegis-workload-demo -n default; \
	else \
  		echo "Nothing to clean."; \
	fi

configure:
	kubectl create ns aegis-system

install-all: install-demo install-safe install-sidecar install-sentinel

all: all-demo all-safe all-sidecar all-sentinel

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
