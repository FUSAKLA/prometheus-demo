FROM golang:1.10.0 AS build-image
WORKDIR /go/src/github.com/fusakla/prometheus-demo/error-server
COPY . .
RUN go get -d
RUN go build

FROM debian:stretch
WORKDIR /error-server
COPY --from=build-image /go/src/github.com/fusakla/prometheus-demo/error-server .
EXPOSE 8080
ENTRYPOINT ["/error-server/error-server"]
