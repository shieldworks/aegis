#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                  <aegis.z2h.dev>
#     .\_/.
#

VERSION=0.5.0
PACKAGE=aegis-workload-demo
REPO=z2hdev/aegis-workload-demo

build:
	go build -o ${PACKAGE}

run:
	go run main.go

bundle:
	docker build . -t ${PACKAGE}:${VERSION}

push:
	docker build . -t ${PACKAGE}:${VERSION}
	docker tag ${PACKAGE}:${VERSION} ${REPO}:${VERSION}
	docker push ${REPO}:${VERSION}

deploy:
	minikube kubectl -- apply -f ./k8s/

run-in-container:
	docker run ${PACKAGE}:${VERSION}
