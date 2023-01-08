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
COPY vendor /build/vendor
COPY internal /build/internal
COPY go.mod /build/go.mod
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -mod vendor -a -o aegis ./cmd/main.go

FROM alpine:3.17.0

# Copy the aegis binary
COPY --from=builder /build/aegis /bin/aegis

ENV HOSTNAME sentinel

ARG USERNAME=sentinel
ARG USER_UID=1000
ARG USER_GID=$USER_UID

# Create a non-root user.
RUN groupadd --gid $USER_GID $USERNAME \
    && useradd --uid $USER_UID --gid $USER_GID -m $USERNAME

# Set the default user.
USER $USERNAME

# Keep the container alive.
ENTRYPOINT ["/bin/sh","-c","sleep infinity"]
# Default command.
CMD ["/bin/sentinel"]