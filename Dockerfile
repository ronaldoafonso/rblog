
FROM golang:alpine3.10

RUN apk add --update && \
    apk add git

RUN addgroup rblog && \
    adduser -h /home/rblog -s /bin/ash -G rblog -D rblog

RUN mkdir -p /go/src/github.com/ronaldoafonso/rblog && \
    mkdir -p /go/src/github.com/ronaldoafonso/rblog/rarticle/html && \
    chown -R rblog:rblog /go

USER rblog:rblog

ENV CGO_ENABLED 0

RUN go get -v github.com/gorilla/mux

RUN go get -v github.com/gorilla/handlers

WORKDIR /go/src/github.com/ronaldoafonso/rblog

COPY --chown=rblog:rblog *.go /go/src/github.com/ronaldoafonso/rblog/

COPY --chown=rblog:rblog rarticle/*.go /go/src/github.com/ronaldoafonso/rblog/rarticle/

COPY --chown=rblog:rblog rarticle/html/*.html /go/src/github.com/ronaldoafonso/rblog/rarticle/html/

RUN go install -v github.com/ronaldoafonso/rblog

CMD ["rblog"]
