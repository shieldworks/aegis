#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                    <aegis.ist>
#     .\_/.
#

# Packages the “Aegis Init Container” binary into a container image.
init-container-bundle:
	./hack/bundle.sh "aegis-ist-init-container" \
		$(VERSION) "dockerfiles/aegis-ist/init-container.Dockerfile"

# Packages the “Aegis Init Container” binary into a container image for Photon OS.
init-container-bundle-photon:
	./hack/bundle.sh "aegis-photon-init-container" \
		$(VERSION) "dockerfiles/aegis-photon/init-container.Dockerfile"

# Pushes the “Aegis Init Container” container image to the public registry.
init-container-push:
	./hack/push.sh "aegis-ist-init-container" \
		$(VERSION) "aegishub/aegis-ist-init-container"

# Pushes the “Aegis Init Container” (Photon OS) container image to the public registry.
init-container-push-photon:
	./hack/push.sh "aegis-photon-init-container" \
		$(VERSION) "aegishub/aegis-photon-init-container"

# Pushes the “Aegis Init Container” container image to the local registry.
init-container-push-local:
	./hack/push.sh "aegis-ist-init-container" $(VERSION) \
		"localhost:5000/aegis-ist-init-container"

# Pushes the “Aegis Init Container” (Photon OS) container image to the local registry.
init-container-push-photon-local:
	./hack/push.sh "aegis-photon-init-container" $(VERSION) \
		"localhost:5000/aegis-photon-init-container"
