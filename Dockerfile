FROM golang:1.17-alpine

ENV PATH="$PATH:/bin/bash" \
    GO111MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64

RUN apk add --update --upgrade curl unzip bash gcc g++

WORKDIR /go/src

COPY entrypoint.sh /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]

RUN ["chmod", "+x", "/entrypoint.sh"]

EXPOSE 8081