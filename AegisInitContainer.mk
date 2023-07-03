#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                    <aegis.ist>
#     .\_/.
#

# Packages the “Aegis Init Container” binary into a container image.
init-container-bundle-ist:
	./hack/bundle.sh "aegis-ist-init-container" \
		$(VERSION) "dockerfiles/aegis-ist/init-container.Dockerfile"

# Packages the “Aegis Init Container” binary into a container image for FIPS.
init-container-bundle-ist-fips:
	./hack/bundle.sh "aegis-ist-fips-init-container" \
		$(VERSION) "dockerfiles/aegis-ist-fips/init-container.Dockerfile"

# Packages the “Aegis Init Container” binary into a container image for Photon OS.
init-container-bundle-photon:
	./hack/bundle.sh "aegis-photon-init-container" \
		$(VERSION) "dockerfiles/aegis-photon/init-container.Dockerfile"

# Packages the “Aegis Init Container” binary into a container image for Photon OS and FIPS.
init-container-bundle-photon-fips:
	./hack/bundle.sh "aegis-photon-fips-init-container" \
		$(VERSION) "dockerfiles/aegis-photon-fips/init-container.Dockerfile"

# Pushes the “Aegis Init Container” container image to the public registry.
init-container-push-ist:
	./hack/push.sh "aegis-ist-init-container" \
		$(VERSION) "aegishub/aegis-ist-init-container"

# Pushes the “Aegis Init Container” (FIPS) container image to the public registry.
init-container-push-ist-fips:
	./hack/push.sh "aegis-ist-fips-init-container" \
		$(VERSION) "aegishub/aegis-ist-fips-init-container"

# Pushes the “Aegis Init Container” (Photon OS) container image to the public registry.
init-container-push-photon:
	./hack/push.sh "aegis-photon-init-container" \
		$(VERSION) "aegishub/aegis-photon-init-container"

# Pushes the “Aegis Init Container” (Photon OS and FIPS) container image to the public registry.
init-container-push-photon-fips:
	./hack/push.sh "aegis-photon-fips-init-container" \
		$(VERSION) "aegishub/aegis-photon-fips-init-container"

# Pushes the “Aegis Init Container” container image to the local registry.
init-container-push-ist-local:
	./hack/push.sh "aegis-ist-init-container" $(VERSION) \
		"localhost:5000/aegis-ist-init-container"

init-container-push-ist-fips-local:
	./hack/push.sh "aegis-ist-fips-init-container" $(VERSION) \
		"localhost:5000/aegis-ist-fips-init-container"

# Pushes the “Aegis Init Container” (Photon OS) container image to the local registry.
init-container-push-photon-local:
	./hack/push.sh "aegis-photon-init-container" $(VERSION) \
		"localhost:5000/aegis-photon-init-container"

# Pushes the “Aegis Init Container” (Photon OS and FIPS) container image to the local registry.
init-container-push-photon-fips-local:
	./hack/push.sh "aegis-photon-fips-init-container" $(VERSION) \
		"localhost:5000/aegis-photon-fips-init-container"
