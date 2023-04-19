#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                    <aegis.ist>
#     .\_/.
#

# The common version tag assigned to all the things.
VERSION=0.15.6

# tags a release
tag:
	./hack/tag.sh $(VERSION)

#
# ## Aegis Safe ##
#

# Builds “Aegis Safe” into a binary.
safe-build:
	./hack/safe-build.sh "aegis-safe"
# Packages the “Aegis Safe” into a container image.
safe-bundle:
	./hack/bundle.sh "aegis-safe" $(VERSION) "dockerfiles/aegis-safe/Dockerfile"
# Pushes the “Aegis Safe” container to the public registry.
safe-push:
	./hack/push.sh "aegis-safe" $(VERSION) "aegishub/aegis-safe"
# Pushes the “Aegis Safe” container image to the local registry.
safe-push-local:
	./hack/push.sh "aegis-safe" $(VERSION) "localhost:5000/aegis-safe"
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
	./hack/sentinel-build.sh "aegis-sentinel"
# Packages the “Aegis Sentinel” binary into a container image.
sentinel-bundle:
	./hack/bundle.sh "aegis-sentinel" $(VERSION) "dockerfiles/aegis-sentinel/Dockerfile"
# Pushes the “Aegis Sentinel” container image the the public registry.
sentinel-push:
	./hack/push.sh "aegis-sentinel" $(VERSION) "aegishub/aegis-sentinel"
# Pushes the “Aegis Sentinel” container image to the local registry.
sentinel-push-local:
	./hack/push.sh "aegis-sentinel" $(VERSION) "localhost:5000/aegis-sentinel"
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
	./hack/init-container-build.sh "aegis-init-container"
# Packages the “Aegis Init Container” binary into a container image.
init-container-bundle:
	./hack/bundle.sh "aegis-init-container" $(VERSION) "dockerfiles/aegis-init-container/Dockerfile"
# Pushes the “Aegis Init Container” container image to the public registry.
init-container-push:
	./hack/push.sh "aegis-init-container" $(VERSION) "aegishub/aegis-init-container"
# Pushes the “Aegis Init Container” container image to the local registry.
init-container-push-local:
	./hack/push.sh "aegis-init-container" $(VERSION) "localhost:5000/aegis-init-container"

#
# ## Aegis Sidecar ##
#

# Builds “Aegis Sidecar” into a binary.
sidecar-build:
	./hack/sidecar-build.sh "aegis-sidecar"
# Packages the “Aegis Sidecar” binary into a container image.
sidecar-bundle:
	./hack/bundle.sh "aegis-sidecar" $(VERSION) "dockerfiles/aegis-sidecar/Dockerfile"
# Pushes the “Aegis Sidecar” container image to the public registry.
sidecar-push:
	./hack/push.sh "aegis-sidecar" $(VERSION) "aegishub/aegis-sidecar"
# Pushes the “Aegis Sidecar” container image to the local registry.
sidecar-push-local:
	./hack/push.sh "aegis-sidecar" $(VERSION) "localhost:5000/aegis-sidecar"

#
# ## Use Case: Sidecar ##
#

# Builds the “Sidecar” use case into a binary.
example-sidecar-build:
	./hack/example-sidecar-build.sh "aegis-workload-demo-using-sidecar"
# Packages the “Sidecar” use case binary into a container image.
example-sidecar-bundle:
	./hack/bundle.sh "aegis-workload-demo-using-sidecar" \
		$(VERSION) "dockerfiles/example-sidecar/Dockerfile"
# Pushes the “Sidecar” use case container image to the public registry.
example-sidecar-push:
	./hack/push.sh "aegis-workload-demo-using-sidecar" \
		$(VERSION) "aegishub/aegis-workload-demo-using-sidecar"
# Pushes the “Sidecar” use case container image to the local registry.
example-sidecar-push-local:
	./hack/push.sh "aegis-workload-demo-using-sidecar" \
		$(VERSION) "localhost:5000/aegis-workload-demo-using-sidecar"
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
	./hack/example-sdk-build.sh "aegis-workload-demo-using-sdk"
# Packages the “SDK” use case binary into a container image.
example-sdk-bundle:
	./hack/bundle.sh "aegis-workload-demo-using-sdk" \
		$(VERSION) "dockerfiles/example-sdk/Dockerfile"
# Pushes the “SDK” use case container image to the public registry.
example-sdk-push:
	./hack/push.sh "aegis-workload-demo-using-sdk" \
		$(VERSION) "aegishub/aegis-workload-demo-using-sdk"
# Pushes the “SDK” use case container image to the local registry.
example-sdk-push-local:
	./hack/push.sh "aegis-workload-demo-using-sdk" \
		$(VERSION) "localhost:5000/aegis-workload-demo-using-sdk"
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
	./hack/example-multiple-secrets-build.sh "aegis-workload-demo-multiple-secrets"
# Packages the “multiple secrets” use case binary into a container image.
example-multiple-secrets-bundle:
	./hack/bundle.sh "aegis-workload-demo-multiple-secrets" \
		$(VERSION) "dockerfiles/example-multiple-secrets/Dockerfile"
# Pushes the “multiple secrets” use case container image to the public registry.
example-multiple-secrets-push:
	./hack/push.sh "aegis-workload-demo-using-secrets" \
		$(VERSION) "aegishub/aegis-workload-demo-multiple-secrets"
# Pushes the “multiple secrets” use case container image to the local registry.
example-multiple-secrets-push-local:
	./hack/push.sh "aegis-workload-demo-multiple-secrets" \
		$(VERSION) "localhost:5000/aegis-workload-demo-multiple-secrets"
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
	./hack/example-init-container-build.sh "aegis-workload-demo-using-init-container"
# Packages the “Init Container” binary into a container image.
example-init-container-bundle:
	./hack/bundle.sh "aegis-workload-demo-using-init-container" \
		$(VERSION) "dockerfiles/example-init-container/Dockerfile"
# Pushes the “Init Container” container image to the public registry.
example-init-container-push:
	./hack/push.sh "aegis-workload-demo-using-init-container" \
		$(VERSION) "aegishub/aegis-workload-demo-using-init-container"
# Pushes the “Init Container” container image to the local registry.
example-init-container-push-local:
	./hack/push.sh "aegis-workload-demo-using-init-container" \
		$(VERSION) "localhost:5000/aegis-workload-demo-using-init-container"
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
