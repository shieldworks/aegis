#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                    <aegis.ist>
#     .\_/.
#

# Packages the “Aegis Sentinel” binary into a container image.
sentinel-bundle-ist:
	./hack/bundle.sh "aegis-ist-sentinel" \
		$(VERSION) "dockerfiles/aegis-ist/sentinel.Dockerfile"

# Packages the “Aegis Sentinel” binary into a container image for FIPS.
sentinel-bundle-ist-fips:
	./hack/bundle.sh "aegis-ist-fips-sentinel" \
		$(VERSION) "dockerfiles/aegis-ist-fips/sentinel.Dockerfile"

# Packages the “Aegis Sentinel” binary into a container image for Photon OS.
sentinel-bundle-photon:
	./hack/bundle.sh "aegis-photon-sentinel" \
		$(VERSION) "dockerfiles/aegis-photon/sentinel.Dockerfile"

# Packages the “Aegis Sentinel” binary into a container image for Photon OS and FIPS.
sentinel-bundle-photon-fips:
	./hack/bundle.sh "aegis-photon-fips-sentinel" \
		$(VERSION) "dockerfiles/aegis-photon-fips/sentinel.Dockerfile"

# Pushes the “Aegis Sentinel” container image the the public registry.
sentinel-push-ist:
	./hack/push.sh "aegis-ist-sentinel" \
		$(VERSION) "aegishub/aegis-ist-sentinel"

# Pushes the “Aegis Sentinel” (Photon OS) container image to the public registry.
sentinel-push-ist-fips:
	./hack/push.sh "aegis-ist-fips-sentinel" \
		$(VERSION) "aegishub/aegis-ist-fips-sentinel"

# Pushes the “Aegis Sentinel” (Photon OS) container image to the public registry.
sentinel-push-photon:
	./hack/push.sh "aegis-photon-sentinel" \
		$(VERSION) "aegishub/aegis-photon-sentinel"

# Pushes the “Aegis Sentinel” (Photon OS) container image to the public registry.
sentinel-push-photon-fips:
	./hack/push.sh "aegis-photon-fips-sentinel" \
		$(VERSION) "aegishub/aegis-photon-fips-sentinel"

# Pushes the “Aegis Sentinel” container image to the local registry.
sentinel-push-ist-local:
	./hack/push.sh "aegis-ist-sentinel" \
		$(VERSION) "localhost:5000/aegis-ist-sentinel"

# Pushes the “Aegis Sentinel” (Photon OS) container image to the local registry.
sentinel-push-ist-fips-local:
	./hack/push.sh "aegis-ist-fips-sentinel" \
		$(VERSION) "localhost:5000/aegis-ist-fips-sentinel"

# Pushes the “Aegis Sentinel” (Photon OS) container image to the local registry.
sentinel-push-photon-local:
	./hack/push.sh "aegis-photon-sentinel" \
		$(VERSION) "localhost:5000/aegis-photon-sentinel"

# Pushes the “Aegis Sentinel” (Photon OS) container image to the local registry.
sentinel-push-photon-fips-local:
	./hack/push.sh "aegis-photon-fips-sentinel" \
		$(VERSION) "localhost:5000/aegis-photon-fips-sentinel"
