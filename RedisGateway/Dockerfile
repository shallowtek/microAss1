#get golang
FROM golang:1.8

#select working directory
WORKDIR /RedisGateway

#copy over the files into the new working directory
COPY . .

#get all the imports and install
run go get -d -v ./...


#build the go file
RUN go build redisGateway.go

EXPOSE 10006
#start the service
ENTRYPOINT ["/RedisGateway/redisGateway"]


