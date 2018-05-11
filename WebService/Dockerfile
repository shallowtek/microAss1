FROM golang:latest

WORKDIR /webApp

COPY . .
run go get -d -v ./...
RUN go build webService.go
EXPOSE 8080


ENTRYPOINT ["/webApp/webService"]

