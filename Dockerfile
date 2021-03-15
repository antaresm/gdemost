FROM golang:latest

WORKDIR /go/src/gdemost
ADD . .

RUN go mod tidy

RUN GO111MODULE=off go get -u -f github.com/qor/bindatafs/...
RUN bindatafs config/bindatafs

RUN go install gdemost

ENTRYPOINT /go/bin/gdemost

