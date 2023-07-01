#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                    <aegis.ist>
#     .\_/.
#

# Builds everything and pushes to public registries.
build: \
	example-sidecar-bundle \
	example-sidecar-push \
	example-sdk-bundle \
	example-sdk-push \
	example-multiple-secrets-bundle \
	example-multiple-secrets-push \
	example-init-container-bundle \
	example-init-container-push \
	safe-bundle \
	safe-push \
	safe-bundle-photon \
	safe-push-photon \
	sidecar-bundle \
	sidecar-push \
	sidecar-bundle-photon \
	sidecar-push-photon \
	sentinel-bundle \
	sentinel-push \
	sentinel-bundle-photon \
	sentinel-push-photon \
	init-container-bundle \
	init-container-push \
	init-container-bundle-photon \
	init-container-push-photon

# Builds everything and pushes to the local registry.
build-local: \
	example-sidecar-bundle \
	example-sidecar-push-local \
	example-sdk-bundle \
	example-sdk-push-local \
	example-multiple-secrets-bundle \
	example-multiple-secrets-push-local \
	example-init-container-bundle \
	example-init-container-push-local \
	safe-bundle \
	safe-push-local \
	safe-bundle-photon \
	safe-push-photon-local \
	sidecar-bundle \
	sidecar-push-local \
	sidecar-bundle-photon \
	sidecar-push-photon-local \
	sentinel-bundle \
	sentinel-push-local \
	sentinel-bundle-photon \
	sentinel-push-photon-local \
	init-container-bundle \
	init-container-push-local \
	init-container-bundle-photon \
	init-container-push-photon-local
