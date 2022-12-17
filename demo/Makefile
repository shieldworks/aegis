#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                  <aegis.z2h.dev>
#     .\_/.
#

build:
	go build -o aegis-demo-workload

run:
	go run main.go

bundle:
	docker build . -t aegis-demo-workload:0.5.0

run-in-container:
	docker run aegis-demo-workload:0.5.0


