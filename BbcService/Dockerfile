FROM golang:1.8

WORKDIR /BbcService

COPY . .

run go get -d -v ./...


RUN go build bbcService.go

EXPOSE 10005

ENTRYPOINT ["/BbcService/bbcService"]
