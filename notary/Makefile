#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                  <aegis.z2h.dev>
#     .\_/.
#

build:
	go build -o aegis-notary

run:
	go run main.go

bundle:
	docker build . -t aegis-notary:0.5.0

push:
	docker build . -t aegis-notary:0.5.0
	docker tag aegis-notary:0.5.0 localhost:5000/aegis-notary:0.5.0
	docker push localhost:5000/aegis-notary:0.5.0

deploy:
	minikube kubectl -- apply -f ./k8s/Deployment.yaml

run-in-container:
	docker run aegis-notary:0.5.0


