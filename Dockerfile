FROM golang

ADD . /go/src/gdemost

RUN go install gdemost@latest

ENTRYPOINT /go/bin/gdemost

