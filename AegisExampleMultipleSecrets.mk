#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                    <aegis.ist>
#     .\_/.
#

# Packages the “multiple secrets” use case binary into a container image.
example-multiple-secrets-bundle:
	./hack/bundle.sh "example-multiple-secrets" \
		$(VERSION) "dockerfiles/example/multiple-secrets.Dockerfile"

# Pushes the “multiple secrets” use case container image to the public registry.
example-multiple-secrets-push:
	./hack/push.sh "example-multiple-secrets" \
		$(VERSION) "aegishub/example-multiple-secrets"

# Pushes the “multiple secrets” use case container image to the local registry.
example-multiple-secrets-push-local:
	./hack/push.sh "example-multiple-secrets" \
		$(VERSION) "localhost:5000/example-multiple-secrets"

# Deploys the “multiple secrets” use case app from the public registry into the cluster.
example-multiple-secrets-deploy:
	./hack/example-multiple-secrets-deploy.sh

# Deploys the “multiple secrets” use case app from the local registry into the cluster.
example-multiple-secrets-deploy-local:
	./hack/example-multiple-secrets-deploy-local.sh
