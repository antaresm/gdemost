FROM golang:1.16.2

WORKDIR /go/src/gdemost
COPY . .

RUN go mod download
RUN go build -tags bindatafs main.go
RUN go install -tags bindatafs main.go

ENTRYPOINT ["/go/bin/main", "-tags", "bindatafs"]