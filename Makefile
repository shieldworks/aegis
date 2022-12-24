#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                  <aegis.z2h.dev>
#     .\_/.
#

prepare:
	kubectl create ns aegis-system

clean:
	kubectl delete ns aegis-system
	kubectl delete deployment aegis-workload-demo -n default

install:
	echo "Not implemented yet!"

clean-prepare-all: clean prepare-all

prepare-all: prepare all-demo all-safe all-sidecar all-sentinel

all: all-demo all-safe all-sidecar all-sentinel

all-demo:
	cd demo && $(MAKE) all

all-safe:
	cd safe && $(MAKE) all

all-sidecar:
	cd sidecar && $(MAKE) all

all-sentinel:
	cd sentinel && $(MAKE) all