FROM golang:latest

WORKDIR /go/src/gdemost
COPY . .

RUN go mod tidy
RUN go install gdemost

ENTRYPOINT /go/bin/gdemost

