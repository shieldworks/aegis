#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                    <aegis.ist>
#     .\_/.
#

# The common version tag assigned to all the things.
VERSION=0.17.1

# tags a release
tag:
	./hack/tag.sh $(VERSION)

#
# ## Aegis Safe ##
#

# Builds “Aegis Safe” into a binary.
safe-build:
	./hack/build.sh "aegis-ist-safe" "./app/safe/cmd/main.go"
# Packages the “Aegis Safe” into a container image.
safe-bundle:
	./hack/bundle.sh "aegis-ist-safe" $(VERSION) "dockerfiles/aegis-ist/safe.Dockerfile"
# Pushes the “Aegis Safe” container to the public registry.
safe-push:
	./hack/push.sh "aegis-ist-safe" $(VERSION) "aegishub/aegis-ist-safe"
# Pushes the “Aegis Safe” container image to the local registry.
safe-push-local:
	./hack/push.sh "aegis-ist-safe" $(VERSION) "localhost:5000/aegis-ist-safe"
# Deploys “Aegis Safe” from the public registry into the cluster.
safe-deploy:
	./hack/safe-deploy.sh
# Deploys t“Aegis Safe” from the local registry into the cluster.
safe-deploy-local:
	./hack/safe-deploy-local.sh

#
# ## Aegis Sentinel ##
#

# Builds “Aegis Sentinel” into a binary.
sentinel-build:
	./hack/build.sh "aegis-ist-sentinel" "./app/sentinel/cmd/main.go"
# Packages the “Aegis Sentinel” binary into a container image.
sentinel-bundle:
	./hack/bundle.sh "aegis-ist-sentinel" $(VERSION) "dockerfiles/aegis-ist/sentinel.Dockerfile"
# Pushes the “Aegis Sentinel” container image the the public registry.
sentinel-push:
	./hack/push.sh "aegis-ist-sentinel" $(VERSION) "aegishub/aegis-ist-sentinel"
# Pushes the “Aegis Sentinel” container image to the local registry.
sentinel-push-local:
	./hack/push.sh "aegis-ist-sentinel" $(VERSION) "localhost:5000/aegis-ist-sentinel"
# Deploys “Aegis Sentinel” from the public registry into the cluster.
sentinel-deploy:
	./hack/sentinel-deploy.sh
# Deploys “Aegis Sentinel” from the local registry into the cluster.
sentinel-deploy-local:
	./hack/sentinel-deploy-local.sh

#
# ## Aegis Init Container ##
#

# Builds “Aegis Init Container” into a binary.
init-container-build:
	./hack/build.sh "aegis-ist-init-container" "./app/init-container/cmd/main.go"
# Packages the “Aegis Init Container” binary into a container image.
init-container-bundle:
	./hack/bundle.sh "aegis-ist-init-container" $(VERSION) "dockerfiles/aegis-ist/init-container.Dockerfile"
# Pushes the “Aegis Init Container” container image to the public registry.
init-container-push:
	./hack/push.sh "aegis-ist-init-container" $(VERSION) "aegishub/aegis-ist-init-container"
# Pushes the “Aegis Init Container” container image to the local registry.
init-container-push-local:
	./hack/push.sh "aegis-ist-init-container" $(VERSION) "localhost:5000/aegis-ist-init-container"

#
# ## Aegis Sidecar ##
#

# Builds “Aegis Sidecar” into a binary.
sidecar-build:
	./hack/build.sh "aegis-ist-sidecar" "./app/sidecar/cmd/main.go"
# Packages the “Aegis Sidecar” binary into a container image.
sidecar-bundle:
	./hack/bundle.sh "aegis-ist-sidecar" $(VERSION) "dockerfiles/aegis-ist/sidecar.Dockerfile"
# Pushes the “Aegis Sidecar” container image to the public registry.
sidecar-push:
	./hack/push.sh "aegis-ist-sidecar" $(VERSION) "aegishub/aegis-ist-sidecar"
# Pushes the “Aegis Sidecar” container image to the local registry.
sidecar-push-local:
	./hack/push.sh "aegis-ist-sidecar" $(VERSION) "localhost:5000/aegis-ist-sidecar"

#
# ## Use Case: Sidecar ##
#

# Builds the “Sidecar” use case into a binary.
example-sidecar-build:
	./hack/build.sh "example-using-sidecar" "./examples/using-sidecar/main.go"
# Packages the “Sidecar” use case binary into a container image.
example-sidecar-bundle:
	./hack/bundle.sh "example-using-sidecar" \
		$(VERSION) "dockerfiles/example/sidecar.Dockerfile"
# Pushes the “Sidecar” use case container image to the public registry.
example-sidecar-push:
	./hack/push.sh "example-using-sidecar" \
		$(VERSION) "aegishub/example-using-sidecar"
# Pushes the “Sidecar” use case container image to the local registry.
example-sidecar-push-local:
	./hack/push.sh "example-using-sidecar" \
		$(VERSION) "localhost:5000/example-using-sidecar"
# Deploys the “Sidecar” use case app from the public registry into the cluster.
example-sidecar-deploy:
	./hack/example-sidecar-deploy.sh
# Deploys the “Sidecar” use case app from the local registry into the cluster.
example-sidecar-deploy-local:
	./hack/example-sidecar-deploy-local.sh

#
# ## Use Case: SDK ##
#

# Builds the “SDK” use case into a binary.
example-sdk-build:
	./hack/build.sh "example-using-sdk" "./examples/using-sdk/main.go"
# Packages the “SDK” use case binary into a container image.
example-sdk-bundle:
	./hack/bundle.sh "example-using-sdk" \
		$(VERSION) "dockerfiles/example/sdk.Dockerfile"
# Pushes the “SDK” use case container image to the public registry.
example-sdk-push:
	./hack/push.sh "example-using-sdk" \
		$(VERSION) "aegishub/example-using-sdk"
# Pushes the “SDK” use case container image to the local registry.
example-sdk-push-local:
	./hack/push.sh "example-using-sdk" \
		$(VERSION) "localhost:5000/example-using-sdk"
# Deploys the “SDK” use case app from the public registry into the cluster.
example-sdk-deploy:
	./hack/example-sdk-deploy.sh
# Deploys the “SDK” use case app from the local registry into the cluster.
example-sdk-deploy-local:
	./hack/example-sdk-deploy-local.sh

#
# ## Use Case: Multiple Secrets ##
#

# Builds the “multiple secrets” use case into a binary.
example-multiple-secrets-build:
	./hack/build.sh "example-multiple-secrets" "./examples/multiple-secrets/main.go"
# Packages the “multiple secrets” use case binary into a container image.
example-multiple-secrets-bundle:
	./hack/bundle.sh "example-multiple-secrets" \
		$(VERSION) "dockerfiles/example/multiple-secrets.Dockerfile"
# Pushes the “multiple secrets” use case container image to the public registry.
example-multiple-secrets-push:
	./hack/push.sh "example-multiple-secrets" \
		$(VERSION) "aegishub/example-multiple-secrets"
# Pushes the “multiple secrets” use case container image to the local registry.
example-multiple-secrets-push-local:
	./hack/push.sh "example-multiple-secrets" \
		$(VERSION) "localhost:5000/example-multiple-secrets"
# Deploys the “multiple secrets” use case app from the public registry into the cluster.
example-multiple-secrets-deploy:
	./hack/example-multiple-secrets-deploy.sh
# Deploys the “multiple secrets” use case app from the local registry into the cluster.
example-multiple-secrets-deploy-local:
	./hack/example-multiple-secrets-deploy-local.sh

#
# ## Use Case: Init Container ##
#

# Builds the “Init Container” use case into a binary.
example-init-container-build:
	./hack/build.sh "example-using-init-container" "./examples/using-init-container/main.go"
# Packages the “Init Container” binary into a container image.
example-init-container-bundle:
	./hack/bundle.sh "example-using-init-container" \
		$(VERSION) "dockerfiles/example/init-container.Dockerfile"
# Pushes the “Init Container” container image to the public registry.
example-init-container-push:
	./hack/push.sh "example-using-init-container" \
		$(VERSION) "aegishub/example-using-init-container"
# Pushes the “Init Container” container image to the local registry.
example-init-container-push-local:
	./hack/push.sh "example-using-init-container" \
		$(VERSION) "localhost:5000/example-using-init-container"
# Deploys the “Init Container” app from the public registry into the cluster.
example-init-container-deploy:
	./hack/example-init-container-deploy.sh
# Deploys the “Init Container” app from the local registry into the cluster.
example-init-container-deploy-local:
	./hack/example-init-container-deploy-local.sh

#
# ## Lifecycle ##
#

# Removes the former Aegis deployment without entirely destroying the cluster.
clean:
	./hack/uninstall.sh

# Completely removes the Minikube cluster.
k8s-delete:
	./hack/minikube-delete.sh
# Brings up a fresh Minikube cluster.
k8s-start:
	./hack/minikube-start.sh

# Deploys Aegis to the cluster.
deploy:
	./hack/deploy.sh
deploy-local:
	./hack/deploy-local.sh

#
# ## Tests ##
#

# Integration tests.
test:
	./hack/test.sh "remote"
test-local:
	./hack/test.sh

#
# ## Build and Push ##
#

# Builds everything and pushes to registries.
build: \
	example-sidecar-build \
	example-sidecar-bundle \
	example-sidecar-push \
	example-sdk-build \
	example-sdk-bundle \
	example-sdk-push \
	example-multiple-secrets-build \
	example-multiple-secrets-bundle \
	example-multiple-secrets-push \
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
	example-multiple-secrets-build \
	example-multiple-secrets-bundle \
	example-multiple-secrets-push-local \
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

#
# ## Help ##
#

help:
	@echo ""
	@echo "                         ---------------------------------------------------"
	@echo "                         eval $$ (minikube -p minikube docker-env)"
	@echo "            Docker Host: ${DOCKER_HOST}"
	@echo "Minikube Active dockerd: ${MINIKUBE_ACTIVE_DOCKERD}"
	@echo "                         ---------------------------------------------------"
	@echo "                   PREP: make k8s-delete;make k8s-start;\n\
                   TEST: make build-local;make deploy-local;make test-local;\n\
                         ---------------------------------------------------\n\
      EXAMPLE (SIDECAR): make example-sidecar-deploy-local |\n\
                         make example-sidecar-deploy\n\
                         ---------------------------------------------------\n\
          EXAMPLE (SDK): make example-sdk-deploy-local |\n\
                         make example-sdk-deploy\n\
                         ---------------------------------------------------\n\
    EXAMPLE (N SECRETS): make example-multiple-secrets-deploy-local |\n\
                         make example-multiple-secrets-deploy\n\
                         ---------------------------------------------------\n\
       EXAMPLE (INIT C): make example-init-container-deploy-local |\n\
                         make example-init-container-deploy\n\
                         ---------------------------------------------------\n\
                CLEANUP: make clean\n\
                         ---------------------------------------------------\n\
                RELEASE: make k8s-delete;make bump;make build;\n\
         TEST (release): make k8s-start;make deploy;make test;\n\
                    TAG: make tag\n\
                         ---------------------------------------------------\n"
	@echo ""
