#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                    <aegis.ist>
#     .\_/.
#

# Packages the “Aegis Sentinel” binary into a container image.
sentinel-bundle:
	./hack/bundle.sh "aegis-ist-sentinel" \
		$(VERSION) "dockerfiles/aegis-ist/sentinel.Dockerfile"

# Packages the “Aegis Sentinel” binary into a container image for Photon OS.
sentinel-bundle-photon:
	./hack/bundle.sh "aegis-photon-sentinel" \
		$(VERSION) "dockerfiles/aegis-photon/sentinel.Dockerfile"

# Pushes the “Aegis Sentinel” container image the the public registry.
sentinel-push:
	./hack/push.sh "aegis-ist-sentinel" \
		$(VERSION) "aegishub/aegis-ist-sentinel"

# Pushes the “Aegis Sentinel” (Photon OS) container image to the public registry.
sentinel-push-photon:
	./hack/push.sh "aegis-photon-sentinel" \
		$(VERSION) "aegishub/aegis-photon-sentinel"

# Pushes the “Aegis Sentinel” container image to the local registry.
sentinel-push-local:
	./hack/push.sh "aegis-ist-sentinel" \
		$(VERSION) "localhost:5000/aegis-ist-sentinel"

# Pushes the “Aegis Sentinel” (Photon OS) container image to the local registry.
sentinel-push-photon-local:
	./hack/push.sh "aegis-photon-sentinel" \
		$(VERSION) "localhost:5000/aegis-photon-sentinel"

# Deploys “Aegis Sentinel” from the public registry into the cluster.
sentinel-deploy:
	./hack/sentinel-deploy.sh

# Deploys “Aegis Sentinel” (Photon OS) from the public registry into the cluster.
sentinel-deploy-photon:
	./hack/sentinel-deploy-photon.sh

# Deploys “Aegis Sentinel” from the local registry into the cluster.
sentinel-deploy-local:
	./hack/sentinel-deploy-local.sh

# Deploys “Aegis Sentinel” (Photon OS) from the local registry into the cluster.
sentinel-deploy-photon-local:
	./hack/sentinel-deploy-photon.sh
