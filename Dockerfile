FROM golang:latest

WORKDIR /go/src/gdemost
ADD . .

RUN go get -u -f github.com/qor/bindatafs/...
RUN go mod tidy
RUN go install gdemost

ENTRYPOINT /go/bin/gdemost

