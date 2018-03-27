FROM golang:1.8

WORKDIR /TwitterService

COPY . .

run go get -d -v ./...


RUN go build twitterService.go

EXPOSE 10000

ENTRYPOINT ["/TwitterService/twitterService"]
