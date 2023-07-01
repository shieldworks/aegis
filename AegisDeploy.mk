#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                    <aegis.ist>
#     .\_/.
#

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
deploy-photon:
	./hack/deploy-photon.sh
deploy-photon-local:
	./hack/deploy-photon-local.sh

#
# ## Tests ##
#

# Integration tests.
test-remote:
	./hack/test.sh "remote"
test-local:
	./hack/test.sh

#
# ## Versioning ##
#

# tags a release
tag:
	./hack/tag.sh $(VERSION)
