#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                    <aegis.ist>
#     .\_/.
#

# Packages the “Aegis Safe” into a container image.
safe-bundle:
	./hack/bundle.sh "aegis-ist-safe" \
		$(VERSION) "dockerfiles/aegis-ist/safe.Dockerfile"

# Packages the “Aegis Safe” into a container image for Photon OS.
safe-bundle-photon:
	./hack/bundle.sh "aegis-photon-safe" \
		$(VERSION) "dockerfiles/aegis-photon/safe.Dockerfile"

# Pushes the “Aegis Safe” container to the public registry.
safe-push:
	./hack/push.sh "aegis-ist-safe" $(VERSION) "aegishub/aegis-ist-safe"

# Pushes the “Aegis Safe” (Photon OS) container to the public registry.
safe-push-photon:
	./hack/push.sh "aegis-photon-safe" \
		$(VERSION) "aegishub/aegis-photon-safe"

# Pushes the “Aegis Safe” container image to the local registry.
safe-push-local:
	./hack/push.sh "aegis-ist-safe" $(VERSION) "localhost:5000/aegis-ist-safe"

# Pushes the “Aegis Safe” (Photon OS) container image to the local registry.
safe-push-photon-local:
	./hack/push.sh "aegis-photon-safe" \
		$(VERSION) "localhost:5000/aegis-photon-safe"

# Deploys “Aegis Safe” from the public registry into the cluster.
safe-deploy:
	./hack/safe-deploy.sh

# Deploys “Aegis Safe” (Photon OS) from the public registry into the cluster.
safe-deploy-photon:
	./hack/safe-deploy-photon.sh

# Deploys “Aegis Safe” from the local registry into the cluster.
safe-deploy-local:
	./hack/safe-deploy-local.sh

# Deploys “Aegis Safe” (Photon OS) from the local registry into the cluster.
safe-deploy-photon-local:
	./hack/safe-deploy-photon.sh
