#get golang
FROM golang:1.8

#select working directory
WORKDIR /ComputeService

#copy over the files into the new working directory
COPY . .

#get all the imports and install
run go get -d -v ./...

#build the go file
RUN go build computeService.go

#start the service
ENTRYPOINT ["/ComputeService/computeService"]

