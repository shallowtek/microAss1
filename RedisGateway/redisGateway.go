//Matt Shallow 14-Mar-18
package main

import (
	
		"flag"
	
	"log"
	//"os"
	"net"
	
	//"math"
	//"os/signal"
	//"syscall"
	//"strconv"
	

	//"bufio"
	"golang.org/x/net/context"
	
	//"strconv"
	rs "github.com/shallowtek/microAss1/RedisGateway/proto"

	"google.golang.org/grpc"

	"github.com/gomodule/redigo/redis"

	
)

var(
	port       = flag.Int("port", 10010, "The server port")
	tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile   = flag.String("cert_file", "", "The TLS cert file")
	keyFile    = flag.String("key_file", "", "The TLS key file")

)

type RedisGatewayServer struct {}

//, reply rs.RedisGateway_GetDataServer
func (s *RedisGatewayServer) getData(ctx context.Context, in *rs.KeyRequest) (*rs.DataReply, error) {
				
	//set to redis service
	conn, _ := redis.Dial("tcp", "redis:6379")
	defer conn.Close()
	val, _ := conn.Do("GET", in.Key)
	stringVal := val.(string)
	return &rs.DataReply{Key: in.Key, Value: stringVal}, nil
	

	
}


func (s *RedisGatewayServer) setData(ctx context.Context, in *rs.KeyRequest) (*rs.Empty, error) {
				
	//set to redis service
	conn, _ := redis.Dial("tcp", "redis:6379")
	defer conn.Close()		
	//convertAvg := strconv.FormatFloat(average, 'f', 6, 64)	
	conn.Do("SET", in.Key, in.Value)
	
	return &rs.Empty{}, nil
	
}



func main() {
	
	flag.Parse()
	lis, err := net.Listen("tcp", ":10010")
	if err != nil {
	        log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	//newServer := &RedisGatewayServer{}
	rs.RegisterRedisGatewayServer(grpcServer, &RedisGatewayServer)
	grpcServer.Serve(lis)
	
	
	
}
