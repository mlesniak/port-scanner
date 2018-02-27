FROM golang:alpine
RUN apk add --no-cache git
RUN go get -u github.com/mlesniak/port-scanner

FROM alpine:latest  
COPY --from=0 /go/bin/port-scanner /usr/local/bin
ENTRYPOINT ["/usr/local/bin/port-scanner"]
