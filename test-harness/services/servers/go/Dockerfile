FROM golang:1.15.8
USER root
WORKDIR /go/src/app
COPY server.go /go/src/app/server.go
RUN go build .
CMD ["./app"]
