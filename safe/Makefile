#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                  <aegis.z2h.dev>
#     .\_/.
#

build:
	go build -o aegis-safe

run:
	go run main.go

bundle:
	docker build . -t aegis-safe:0.5.0

push:
	docker build . -t aegis-safe:0.5.0
	docker tag aegis-safe:0.5.0 localhost:5000/aegis-safe:0.5.0
	docker push localhost:5000/aegis-safe:0.5.0

deploy:
	minikube kubectl -- apply -f ./k8s/Deployment.yaml

run-in-container:
	docker run aegis-safe:0.5.0


