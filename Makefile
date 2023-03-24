#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                    <aegis.ist>
#     .\_/.
#

# The common version tag assigned to all the things.
VERSION=0.14.5

# Builds “Aegis Safe” into a binary.
safe-build:
	./hack/safe-build.sh "aegis-safe"
# Packages the “Aegis Safe” into a container image.
safe-bundle:
	./hack/bundle.sh "aegis-safe" $(VERSION) "Safe.Dockerfile"
# Pushes the “Aegis Safe” container to the public registry.
safe-push:
	./hack/push.sh "aegis-safe" $(VERSION) "aegishub/aegis-safe"
# Pushes the “Aegis Safe” container image to the local registry.
safe-push-local:
	./hack/push-local.sh "aegis-safe" $(VERSION) "localhost:5000/aegis-safe"
# Deploys “Aegis Safe” from the public registry into the cluster.
safe-deploy:
	./hack/safe-deploy.sh
# Deploys t“Aegis Safe” from the local registry into the cluster.
safe-deploy-local:
	./hack/safe-deploy-local.sh

# Builds “Aegis Sentinel” into a binary.
sentinel-build:
	./hack/sentinel-build.sh "aegis-sentinel"
# Packages the “Aegis Sentinel” binary into a container image.
sentinel-bundle:
	./hack/bundle.sh "aegis-sentinel" $(VERSION) "Sentinel.Dockerfile"
# Pushes the “Aegis Sentinel” container image the the public registry.
sentinel-push:
	./hack/push.sh "aegis-sentinel" $(VERSION) "aegishub/aegis-sentinel"
# Pushes the “Aegis Sentinel” container image to the local registry.
sentinel-push-local:
	./hack/push-local.sh "aegis-sentinel" $(VERSION) "localhost:5000/aegis-sentinel"
# Deploys “Aegis Sentinel” from the public registry into the cluster.
sentinel-deploy:
	./hack/sentinel-deploy.sh
# Deploys “Aegis Sentinel” from the local registry into the cluster.
sentinel-deploy-local:
	./hack/sentinel-deploy-local.sh

# Builds “Aegis Init Container” into a binary.
init-container-build:
	./hack/init-container-build.sh "aegis-init-container"
# Packages the “Aegis Init Container” binary into a container image.
init-container-bundle:
	./hack/bundle.sh "aegis-init-container" $(VERSION) "InitContainer.Dockerfile"
# Pushes the “Aegis Init Container” container image to the public registry.
init-container-push:
	./hack/push.sh "aegis-init-container" $(VERSION) "aegishub/aegis-init-container"
# Pushes the “Aegis Init Container” container image to the local registry.
init-container-push-local:
	./hack/push.sh "aegis-init-container" $(VERSION) "localhost:5000/aegis-init-container"

# Builds “Aegis Sidecar” into a binary.
sidecar-build:
	./hack/sidecar-build.sh "aegis-sidecar"
# Packages the “Aegis Sidecar” binary into a container image.
sidecar-bundle:
	./hack/bundle.sh "aegis-sidecar" $(VERSION) "Sidecar.Dockerfile"
# Pushes the “Aegis Sidecar” container image to the public registry.
sidecar-push:
	./hack/push.sh "aegis-sidecar" $(VERSION) "aegishub/aegis-sidecar"
# Pushes the “Aegis Sidecar” container image to the local registry.
sidecar-push-local:
	./hack/push.sh "aegis-sidecar" $(VERSION) "localhost:5000/aegis-sidecar"

# Builds the “Sidecar” use case into a binary.
example-sidecar-build:
	./hack/sentinel-build.sh "aegis-sentinel"
# Packages the “Sidecar” binary into a container image.
example-sidecar-bundle:
	./hack/bundle.sh "aegis-sentinel" $(VERSION) "Sentinel.Dockerfile"
# Pushes the “Sidecar” use case container image to the public registry.
example-sidecar-push:
	./hack/push.sh "aegis-sentinel" $(VERSION) "aegishub/aegis-sentinel"
# Pushes the “Sidecar” use case container image to the local registry.
example-sidecar-push-local:
	./hack/push-local.sh "aegis-sentinel" $(VERSION) "localhost:5000/aegis-sentinel"
# Deploys the “Sidecar” use case app from the public registry into the cluster.
example-sidecar-deploy:
	./hack/sentinel-deploy.sh
# Deploys the “Sidecar” use case pp from the local registry into the cluster.
example-sidecar-deploy-local:
	./hack/sentinel-deploy-local.sh

# Builds the “SDK” use case into a binary.
example-sdk-build:
	./hack/sentinel-build.sh "aegis-sentinel"
# Packages the “SDK” binary into a container image.
example-sdk-bundle:
	./hack/bundle.sh "aegis-sentinel" $(VERSION) "Sentinel.Dockerfile"
# Pushes the “SDK” container image to the public registry.
example-sdk-push:
	./hack/push.sh "aegis-sentinel" $(VERSION) "aegishub/aegis-sentinel"
# Pushes the “SDK” container image to the local registry.
example-sdk-push-local:
	./hack/push-local.sh "aegis-sentinel" $(VERSION) "localhost:5000/aegis-sentinel"
# Deploys the “SDK” app from the public registry into the cluster.
example-sdk-deploy:
	./hack/sentinel-deploy.sh
# Deploys the “SDK” app from the local registry into the cluster.
example-sdk-deploy-local:
	./hack/sentinel-deploy-local.sh

# Builds the “Init Container” use case into a binary.
example-init-container-build:
	./hack/sentinel-build.sh "aegis-sentinel"
# Packages the “Init Container” binary into a container image.
example-init-container-bundle:
	./hack/bundle.sh "aegis-sentinel" $(VERSION) "Sentinel.Dockerfile"
# Pushes the “Init Container” container image to the public registry.
example-init-container-push:
	./hack/push.sh "aegis-sentinel" $(VERSION) "aegishub/aegis-sentinel"
# Pushes the “Init Container” container image to the local registry.
example-init-container-push-local:
	./hack/push-local.sh "aegis-sentinel" $(VERSION) "localhost:5000/aegis-sentinel"
# Deploys the “Init Container” app from the public registry into the cluster.
example-init-container-deploy:
	./hack/sentinel-deploy.sh
# Deploys the “Init Container” app from the local registry into the cluster.
example-init-container-deploy-local:
	./hack/sentinel-deploy-local.sh

# Removes the former Aegis deployment without entirely destroying the cluster.
clean:
	./hack/uninstall.sh

# Completely removes the Minikube cluster.
delete-k8s:
	./hack/minikube-start.sh
# Brings up a fresh Minikube cluster.
start-k8s:
	./hack/minikube-delete.sh

# Deploys Aegis to the cluster.
deploy:
	./hack/deploy.sh
deploy-local:
	./hack/deploy-local.sh

# Integration tests.
test:
	./hack/test.sh
test-local:
	./hack/test-local.sh

# Builds everything and pushes to registries.
build: \
	example-sidecar-build \
	example-sidecar-bundle \
	example-sidecar-push \
	example-sdk-build \
	example-sdk-bundle \
	example-sdk-push \
	example-init-container-build \
	example-init-container-bundle \
	example-init-container-push \
	safe-build \
	safe-bundle \
	safe-push \
	sidecar-build \
	sidecar-bundle \
	sidecar-push \
	sentinel-build \
	sentinel-bundle \
	sentinel-push \
	init-container-build \
	init-container-bundle \
	init-container-push
build-local: \
	example-sidecar-build \
	example-sidecar-bundle \
	example-sidecar-push-local \
	example-sdk-build \
	example-sdk-bundle \
	example-sdk-push-local \
	example-init-container-build \
	example-init-container-bundle \
	example-init-container-push-local \
	safe-build \
	safe-bundle \
	safe-push-local \
	sidecar-build \
	sidecar-bundle \
	sidecar-push-local \
	sentinel-build \
	sentinel-bundle \
	sentinel-push-local \
	init-container-build \
	init-container-bundle \
	init-container-push-local

help:
	@echo ""
	@echo "                         ---------------------------------------------------"
	@echo "            Docker Host: ${DOCKER_HOST}"
	@echo "Minikube Active dockerd: ${MINIKUBE_ACTIVE_DOCKERD}"
	@echo "                         ---------------------------------------------------"
	@echo "                   PREP: make delete-k8s;make start-k8s;\n\
                   TEST: make build-local;make deploy-local;make test-local;\n\
 TEST (docker/aegishub): make build;make deploy;make test\n\
                RELEASE: make bump;make build;make tag\n\
      EXAMPLE (SIDECAR): make example-sidecar-deploy-local |\
                         make example-sidecar-deploy\n\
          EXAMPLE (SDK): make example-sdk-deploy-local |\
                         make example-sdk-deploy\n\
       EXAMPLE (INIT C): make example-init-container-deploy-local |\
                         make example-init-container-deploy\n\
                CLEANUP: make clean\n"
	@echo "                         ---------------------------------------------------"
	@echo ""
