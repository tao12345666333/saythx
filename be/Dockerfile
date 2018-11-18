FROM golang:1.11.1 as builder

WORKDIR /go/src/be
COPY . /go/src/be
RUN go get -u github.com/golang/dep/cmd/dep \
        && dep ensure \
        && go build

FROM debian:stretch-slim
COPY --from=builder /go/src/be/be /usr/bin/be
ENTRYPOINT ["/usr/bin/be"]
EXPOSE 8080
