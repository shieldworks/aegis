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
COPY cmd /build/cmd
COPY internal /build/internal
COPY vendor /build/vendor
COPY go.mod /build/go.mod
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -mod vendor -a -o aegis-safe ./cmd/main.go

# generate clean, final image for end users
FROM alpine:3.17
COPY --from=builder /build/aegis-safe .

# executable
ENTRYPOINT [ "./aegis-safe" ]
CMD [ "" ]