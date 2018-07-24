FROM golang:alpine as builder

RUN apk add --no-cache git && \
    go get -d github.com/jkernech/go-http-proxy && \
    CGO_ENABLED=0 GOOS=linux go build -a --ldflags '-extldflags "-static"' github.com/jkernech/go-http-proxy

FROM alpine

RUN apk add --no-cache ca-certificates openssl
COPY --from=builder /go/go-http-proxy /bin/http_proxy

HEALTHCHECK --interval=5s --timeout=3s --retries=3 CMD curl -f http://localhost:8080 || exit 1

CMD ["/bin/http_proxy"]
