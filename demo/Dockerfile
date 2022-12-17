# builder image
FROM golang:1.19.4-alpine3.17 as builder

RUN mkdir /build
ADD *.go /build/
ADD go.mod /build/
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -a -o aegis-demo-workload .


# generate clean, final image for end users
FROM alpine:3.17
COPY --from=builder /build/aegis-demo-workload .

# executable
ENTRYPOINT [ "./aegis-demo-workload" ]
CMD [ "" ]