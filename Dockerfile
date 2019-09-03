
FROM golang:alpine3.10

RUN apk add --update

RUN addgroup rblog && \
    adduser -h /home/rblog -s /bin/ash -G rblog -D rblog

RUN mkdir -p /go/src/github.com/ronaldoafonso/rblog && \
    chown -R rblog:rblog /go

USER rblog:rblog

ENV CGO_ENABLED 0

RUN go get -d -v ./...

WORKDIR /go/src/github.com/ronaldoafonso/rblog

COPY --chown=rblog:rblog main.go /go/src/github.com/ronaldoafonso/rblog

RUN go install -v github.com/ronaldoafonso/rblog

CMD ["rblog"]
