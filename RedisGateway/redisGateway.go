//Matt Shallow 14-Mar-18
package main

import (
	
		"flag"
	"fmt"
	"log"
	//"os"
	"net"
	"time"
	//"math"
	//"os/signal"
	//"syscall"
	//"strconv"
	"net/http"
	"io/ioutil"
	"encoding/json"
	//"bufio"
	//"context"
	
	//"strconv"
	rs "github.com/shallowtek/microAss1/RedisGateway/proto"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc"
	"google.golang.org/grpc/testdata"
	"github.com/gomodule/redigo/redis"

	
)

var(
	port       = flag.Int("port", 10010, "The server port")
	tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile   = flag.String("cert_file", "", "The TLS cert file")
	keyFile    = flag.String("key_file", "", "The TLS key file")

)

type RedisGatewayServer struct {}

func (s *RedisGatewayServer) getData(in *rs.keyRequest, reply rs.RedisGateway_GetDataServer) error {
				
	//set to redis service
	conn, _ := redis.Dial("tcp", "redis:6379")
	defer conn.Close()
	val, _ := conn.Do("GET", in.Name)
	
	return &rs.reply{in.Name, val}, nil
	

	
}


func (s *RedisGatewayServer) setData(in *rs.keyRequest, reply rs.RedisGateway_GetDataServer) error {
				
	//set to redis service
	conn, _ := redis.Dial("tcp", "redis:6379")
	defer conn.Close()		
	//convertAvg := strconv.FormatFloat(average, 'f', 6, 64)	
	conn.Do("SET", in.Name, in.Value)
	
	return &rs.reply{}, nil
	
}



func main() {
	
	flag.Parse()
	lis, err := net.Listen("tcp", ":10010")
	if err != nil {
	        log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	rs.RegisterRedisGatewayServer(grpcServer, &RedisGatewayServer{})
	... // determine whether to use TLS
	grpcServer.Serve(lis)
	
	
	
}
