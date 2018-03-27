FROM golang:latest

WORKDIR /webApp

COPY . .

RUN go build webService.go
EXPOSE 8080

ENTRYPOINT ["/webApp/webService"]

