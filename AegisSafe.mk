#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                    <aegis.ist>
#     .\_/.
#

# Packages the “Aegis Safe” into a container image.
safe-bundle-ist:
	./hack/bundle.sh "aegis-ist-safe" \
		$(VERSION) "dockerfiles/aegis-ist/safe.Dockerfile"

# Packages the “Aegis Safe” into a container image for FIPS.
safe-bundle-ist-fips:
	./hack/bundle.sh "aegis-ist-fips-safe" \
		$(VERSION) "dockerfiles/aegis-ist-fips/safe.Dockerfile"

# Packages the “Aegis Safe” into a container image for Photon OS.
safe-bundle-photon:
	./hack/bundle.sh "aegis-photon-safe" \
		$(VERSION) "dockerfiles/aegis-photon/safe.Dockerfile"

# Packages the “Aegis Safe” into a container image for Photon OS and FIPS.
safe-bundle-photon-fips:
	./hack/bundle.sh "aegis-photon-fips-safe" \
		$(VERSION) "dockerfiles/aegis-photon-fips/safe.Dockerfile"

# Pushes the “Aegis Safe” container to the public registry.
safe-push-ist:
	./hack/push.sh "aegis-ist-safe" $(VERSION) "aegishub/aegis-ist-safe"

# Pushes the “Aegis Safe” container to the public registry.
safe-push-ist-fips:
	./hack/push.sh "aegis-ist-fips-safe" \
		$(VERSION) "aegishub/aegis-ist-fips-safe"

# Pushes the “Aegis Safe” (Photon OS) container to the public registry.
safe-push-photon:
	./hack/push.sh "aegis-photon-safe" \
		$(VERSION) "aegishub/aegis-photon-safe"

# Pushes the “Aegis Safe” (Photon OS) container to the public registry.
safe-push-photon-fips:
	./hack/push.sh "aegis-photon-fips-safe" \
		$(VERSION) "aegishub/aegis-photon-fips-safe"

# Pushes the “Aegis Safe” container image to the local registry.
safe-push-ist-local:
	./hack/push.sh "aegis-ist-safe" $(VERSION) "localhost:5000/aegis-ist-safe"

# Pushes the “Aegis Safe” container image to the local registry.
safe-push-ist-fips-local:
	./hack/push.sh "aegis-ist-fips-safe" \
		$(VERSION) "localhost:5000/aegis-ist-fips-safe"

# Pushes the “Aegis Safe” (Photon OS) container image to the local registry.
safe-push-photon-local:
	./hack/push.sh "aegis-photon-safe" \
		$(VERSION) "localhost:5000/aegis-photon-safe"

# Pushes the “Aegis Safe” (Photon OS) container image to the local registry.
safe-push-photon-fips-local:
	./hack/push.sh "aegis-photon-fips-safe" \
		$(VERSION) "localhost:5000/aegis-photon-fips-safe"
