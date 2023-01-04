#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                  <aegis.z2h.dev>
#     .\_/.
#

VERSION=0.5.20
PACKAGE=aegis-sidecar
REPO=z2hdev/aegis-sidecar

all: build bundle push

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

run-in-container:
	docker run ${PACKAGE}:${VERSION}

