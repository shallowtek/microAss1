#get golang
FROM golang:1.8

#select working directory
WORKDIR /RedisGateway

#copy over the files into the new working directory
COPY . .

#get all the imports and install
run go get -d -v ./...

EXPOSE 8081

#build the go file
RUN go build redisGateway.go

#start the service
ENTRYPOINT ["/RedisGateway/redisGateway"]


