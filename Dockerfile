FROM golang:1.10-alpine

WORKDIR /go/src/github.com/taland/reverse-proxy

ADD . .
RUN go install .

ENTRYPOINT /go/bin/reverse-proxy -addr=:5555

EXPOSE 5555