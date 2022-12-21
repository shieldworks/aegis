#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                  <aegis.z2h.dev>
#     .\_/.
#

# builder image
FROM golang:1.19.4-alpine3.17 as builder

RUN mkdir /build
ADD *.go /build/
ADD go.mod /build/
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -a -o aegis-workload-demo .

# generate clean, final image for end users
FROM alpine:3.17
COPY --from=builder /build/aegis-workload-demo .

# executable
ENTRYPOINT [ "./aegis-workload-demo" ]
CMD [ "" ]