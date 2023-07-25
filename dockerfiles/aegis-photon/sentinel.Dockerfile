#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                    <aegis.ist>
#     .\_/.
#

# builder image
FROM golang:1.20.1-alpine3.17 as builder
RUN mkdir /build
COPY app /build/app
COPY core /build/core
COPY vendor /build/vendor
COPY go.mod /build/go.mod
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -mod vendor -a -o aegis ./app/sentinel/cmd/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -mod vendor -a -o sloth ./app/sentinel/busywait/main.go

# generate clean, final image for end users
FROM photon:5.0

LABEL "maintainers"="Volkan Özçelik <volkan@aegis.ist>"
LABEL "version"="0.18.2"
LABEL "website"="https://aegis.ist/"
LABEL "repo"="https://github.com/shieldworks/aegis-sentinel"
LABEL "documentation"="https://aegis.ist/docs/"
LABEL "contact"="https://aegis.ist/contact/"
LABEL "community"="https://aegis.ist/contact/#community"
LABEL "changelog"="https://aegis.ist/changelog"

# Copy the required binaries
COPY --from=builder /build/aegis /bin/aegis
COPY --from=builder /build/sloth /bin/sloth

ENV HOSTNAME sentinel

# Prevent root access.
ENV USER nobody
USER nobody

# Keep the container alive.
ENTRYPOINT ["/bin/sloth"]
CMD [""]
