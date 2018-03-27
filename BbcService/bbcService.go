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
	ts "github.com/shallowtek/microAss1/BbcService/proto"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc"
	"google.golang.org/grpc/testdata"
		
)

var (
	port       = flag.Int("port", 10005, "The server port")
	tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile   = flag.String("cert_file", "", "The TLS cert file")
	keyFile    = flag.String("key_file", "", "The TLS key file")

	start time.Time
	end time.Time
	elapsed uint64
	rounded float64
	
	//These is my personal key from my bbc api account used to get the stream		
	NEWS_API_KEY="054871042b774af78e9b5c778ab71dd4"
	//url string = "https://newsapi.org/v2/top-headlines?country=us&apiKey=054871042b774af78e9b5c778ab71dd4"	

	
)

type news struct {
	Status   string     `json:"status"`
	Source   string     `json:"source"`
	SortBy   string     `json:"sortBy"`
	Articles []newsItem `json:"articles"`
}

type newsItem struct {
	Author      string `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	UrlToImage  string `json:"urlToImage"`
	PublishedAt string `json:"publishedAt"`
}
type BbcServiceServer struct {}

/*
This function is called by the compute service and uses grpc to communicate. The tweets request struct with variables (trump, 1) is passed in
and a return stream variable is created.
*/
	
func (s *BbcServiceServer) GetNews(in *ts.NewsRequest, stream ts.BbcService_GetNewsServer) error {
		
	var newsResults news
	url := "https://newsapi.org/v2/everything?q="+in.Name+"&apiKey="+NEWS_API_KEY
	resp, _ := http.Get(url)
	bytes, _ := ioutil.ReadAll(resp.Body)

	err := json.Unmarshal(bytes, &newsResults)

	if err != nil {
		fmt.Println("We have an error")
		return err
	}

	resp.Body.Close()
	for _, s := range newsResults.Articles{

		if err := stream.Send(&ts.NewsReply{Text: s.Description});

		 err != nil{

			return err
		}

		fmt.Printf("%s\n", s.Description)

		
	}

	fmt.Println("Stopping Stream...")
	return nil	
		
	
}



func main() {

	flag.Parse()
	
	//Have server listen on port 10000
	lis, err := net.Listen("tcp", ":10005")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	//More of the grpc options similar to client
	var opts []grpc.ServerOption
	if *tls {
		if *certFile == "" {
			*certFile = testdata.Path("server1.pem")
		}
		if *keyFile == "" {
			*keyFile = testdata.Path("server1.key")
		}
		creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
		if err != nil {
			log.Fatalf("Failed to generate credentials %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}

	//create a new grpc server and twitter service server then register them.
	grpcServer := grpc.NewServer(opts...)
	newServer := &BbcServiceServer{}
	ts.RegisterBbcServiceServer(grpcServer, newServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	
}
