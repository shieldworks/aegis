#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                    <aegis.ist>
#     .\_/.
#

# The common version tag assigned to all the things.
VERSION=0.17.3

# Utils
include ./AegisMacOs.mk
include ./AegisDeploy.mk
## Aegis
include ./AegisSafe.mk
include ./AegisSentinel.mk
include ./AegisInitContainer.mk
include ./AegisSidecar.mk
## Examples
include ./AegisExampleSidecar.mk
include ./AegisExampleSdk.mk
include ./AegisExampleMultipleSecrets.mk
include ./AegisExampleInitContainer.mk

## Build
include ./AegisBuild.mk

help:
	@echo "--------------------------------------------------------------------"
	@echo "          🛡️ Aegis: Keep your secrets… secret."
	@echo "          🛡️ https://aegis.ist"
	@echo "--------------------------------------------------------------------"
	@echo "        ℹ️ This Makefile assumes you use Minikube and Docker"
	@echo "        ℹ️ for most operations."
	@echo "--------------------------------------------------------------------"

	@if [ "`uname`" = "Darwin" ]; then \
		if type docker > /dev/null 2>&1; then \
			echo "  Using Docker for Mac?"; \
			echo "        ➡ 'make mac-tunnel' to proxy to the internal registry."; \
		else \
			echo "  Docker is not installed on this Mac."; \
		fi; \
	fi

	@echo ""

	@if [ -z "$(DOCKER_HOST)" -o -z "$(MINIKUBE_ACTIVE_DOCKERD)" ]; then \
		echo "  Using Minikube? If DOCKER_HOST and MINIKUBE_ACTIVE_DOCKERD are"; \
		echo '  not set, then run: eval $$(minikube -p minikube docker-env)'; \
		echo "        ➡ \$$DOCKER_HOST            : ${DOCKER_HOST}"; \
		echo "        ➡ \$$MINIKUBE_ACTIVE_DOCKERD: ${MINIKUBE_ACTIVE_DOCKERD}"; \
	else \
	    echo "  Make sure DOCKER_HOST and MINIKUBE_ACTIVE_DOCKERD are current:"; \
		echo '          eval $$(minikube -p minikube docker-env)'; \
	    echo "          (they may change if you reinstall Minikube)"; \
		echo "        ➡ \$$DOCKER_HOST            : ${DOCKER_HOST}"; \
		echo "        ➡ \$$MINIKUBE_ACTIVE_DOCKERD: ${MINIKUBE_ACTIVE_DOCKERD}"; \
	fi

	@echo "--------------------------------------------------------------------"
	@echo "  Prep/Cleanup:"
	@echo "        ˃ make k8s-delete;make k8s-start;"
	@echo "        ˃ make clean;"
	@echo "--------------------------------------------------------------------"
	@echo "  Testing:"
	@echo "    ⦿ Istanbul images:"
	@echo "        ˃ make build-local;make deploy-local;make test-local;"
	@echo "    ⦿ Photon images:"
	@echo "        ˃ make build-local;make deploy-photon-local;make test-local;"
	@echo "    ⦿ Istanbul (remote) images:"
	@echo "        ˃ make build;make deploy;make test-remote;"
	@echo "    ⦿ Photon (remote) images:"
	@echo "        ˃ make build;make deploy-photon;make test-remote"
	@echo "--------------------------------------------------------------------"
	@echo "  Tagging:"
	@echo "        ˃ make tag;"
	@echo "--------------------------------------------------------------------"
	@echo "  Example Use Cases:"
	@echo "        ˃ make example-sidecar-deploy(-local);"
	@echo "        ˃ make example-sdk-deploy(-local);"
	@echo "        ˃ make example-multiple-secrets-deploy(-local);"
	@echo "--------------------------------------------------------------------"
