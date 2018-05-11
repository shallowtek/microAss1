#get golang gcloud components install docker-credential-gcr
FROM golang:latest

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


