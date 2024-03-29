#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                    <aegis.ist>
#     .\_/.
#

# Packages the “Aegis Sidecar” binary into a container image.
sidecar-bundle-ist:
	./hack/bundle.sh "aegis-ist-sidecar" \
		$(VERSION) "dockerfiles/aegis-ist/sidecar.Dockerfile"

# Packages the “Aegis Sidecar” binary into a container image for FIPS.
sidecar-bundle-ist-fips:
	./hack/bundle.sh "aegis-ist-fips-sidecar" \
		$(VERSION) "dockerfiles/aegis-ist-fips/sidecar.Dockerfile"

# Packages the “Aegis Sidecar” binary into a container image for Photon OS.
sidecar-bundle-photon:
	./hack/bundle.sh "aegis-photon-sidecar" \
		$(VERSION) "dockerfiles/aegis-photon/sidecar.Dockerfile"

# Packages the “Aegis Sidecar” binary into a container image for Photon OS and FIPS.
sidecar-bundle-photon-fips:
	./hack/bundle.sh "aegis-photon-fips-sidecar" \
		$(VERSION) "dockerfiles/aegis-photon-fips/sidecar.Dockerfile"

# Pushes the “Aegis Sidecar” container image to the public registry.
sidecar-push-ist:
	./hack/push.sh "aegis-ist-sidecar" \
		$(VERSION) "aegishub/aegis-ist-sidecar"

# Pushes the “Aegis Sidecar” (FIPS) container image to the public registry.
sidecar-push-ist-fips:
	./hack/push.sh "aegis-ist-fips-sidecar" \
		$(VERSION) "aegishub/aegis-ist-fips-sidecar"

# Pushes the “Aegis Sidecar” (Photon OS) container image to the public registry.
sidecar-push-photon:
	./hack/push.sh "aegis-photon-sidecar" \
		$(VERSION) "aegishub/aegis-photon-sidecar"

# Pushes the “Aegis Sidecar” (Photon OS and FIPS) container image to the public registry.
sidecar-push-photon-fips:
	./hack/push.sh "aegis-photon-fips-sidecar" \
		$(VERSION) "aegishub/aegis-photon-fips-sidecar"

# Pushes the “Aegis Sidecar” container image to the local registry.
sidecar-push-ist-local:
	./hack/push.sh "aegis-ist-sidecar" \
		$(VERSION) "localhost:5000/aegis-ist-sidecar"

# Pushes the “Aegis Sidecar” (FIPS) container image to the local registry.
sidecar-push-ist-fips-local:
	./hack/push.sh "aegis-ist-fips-sidecar" \
		$(VERSION) "localhost:5000/aegis-ist-fips-sidecar"

# Pushes the “Aegis Sidecar” (Photon OS) container image to the local registry.
sidecar-push-photon-local:
	./hack/push.sh "aegis-photon-sidecar" \
		$(VERSION) "localhost:5000/aegis-photon-sidecar"

# Pushes the “Aegis Sidecar” (Photon OS and FIPS) container image to the local registry.
sidecar-push-photon-fips-local:
	./hack/push.sh "aegis-photon-fips-sidecar" \
		$(VERSION) "localhost:5000/aegis-photon-fips-sidecar"
