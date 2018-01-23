FROM golang:alpine

WORKDIR /go/src/github.com/SeedJobs/devops-go-tests/bit.ly
COPY . .

RUN apk add --no-cache git && \
    go-wrapper download && \
    go-wrapper install && \
    go test -v ./... && \
    rm -rf $GOPATH/src

WORKDIR /

EXPOSE 8000

CMD ["/go/bin/bit.ly"]
