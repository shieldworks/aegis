#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                  <aegis.z2h.dev>
#     .\_/.
#

VERSION=0.5.16
PACKAGE=aegis-safe
REPO=z2hdev/aegis-safe

all: build bundle push deploy

build:
	go build -o ${PACKAGE} ./cmd/main.go

run:
	go run ./cmd/main.go

bundle:
	go mod vendor
	docker build . -t ${PACKAGE}:${VERSION}

push:
	docker build . -t ${PACKAGE}:${VERSION}
	docker tag ${PACKAGE}:${VERSION} ${REPO}:${VERSION}
	docker push ${REPO}:${VERSION}

deploy:
	kubectl apply -f ./k8s/Namespace.yaml
	kubectl apply -f ./k8s/Identity.yaml
	kubectl apply -f ./k8s/Deployment.yaml
	kubectl apply -f ./k8s/Service.yaml

run-in-container:
	docker run ${PACKAGE}:${VERSION}
