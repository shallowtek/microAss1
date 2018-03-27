//Matt Shallow 14-Mar-18
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"net"
	"time"
	"math"
	"os/signal"
	"syscall"
	"strconv"
	"github.com/dghubble/go-twitter/twitter"	
	"github.com/dghubble/oauth1"
	ts "github.com/shallowtek/microAss1/TwitterService/proto"
	//ts "microAss1/TwitterService/proto"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc"
	"google.golang.org/grpc/testdata"	
)

var (
	port       = flag.Int("port", 10000, "The server port")
	tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile   = flag.String("cert_file", "", "The TLS cert file")
	keyFile    = flag.String("key_file", "", "The TLS key file")

	start time.Time
	end time.Time
	elapsed uint64
	rounded float64
	
	//These are my personal keys from my twitter api account used to get the stream		
	TWITTER_CONSUMER_KEY="OIary3uSfXs21SE1aOQxMFloQ"
	TWITTER_CONSUMER_SECRET="SIskrpjCS2kg3X3VEwtAH5LubpfBKagax7mbyGFZ5QEgBr1V9h"
	TWITTER_ACCESS_TOKEN="874680197806780416-fh9fyv2adGM4o7zsZiCIHOxL9oBkARi"
	TWITTER_ACCESS_SECRET="tVS3UUIpGgtpsJHki0nmYouqkxXmzCu3DDfEXzYp7V1gC"	

	
)

type TwitterServiceServer struct {}

/*
This function is called by the compute service and uses grpc to communicate. The tweets request struct with variables (trump, 1) is passed in
and a return stream variable is created.
*/
	
func (s *TwitterServiceServer) GetTweets(in *ts.TweetsRequest, stream ts.TwitterService_GetTweetsServer) error {
	
	//Create a new configuration using my keys and tokens	
	config := oauth1.NewConfig(TWITTER_CONSUMER_KEY, TWITTER_CONSUMER_SECRET)
	token := oauth1.NewToken(TWITTER_ACCESS_TOKEN, TWITTER_ACCESS_SECRET)

	// create an authorized httpClient. 
	httpClient := config.Client(oauth1.NoContext, token)

	// Then created the actual twitter client using the authorised http client
	client := twitter.NewClient(httpClient)
	
	fmt.Println("Start Stream...")

	/*
	Based on examples in the stream twitter package they use demux in order to keep a nice steady stream and using
	filtering for particular words. The tweets are added to the tweet reply struct and sent using the return stream variable and send.
	*/
	demux := twitter.NewSwitchDemux()
	demux.Tweet = func(tweet *twitter.Tweet) {

		stream.Send(&ts.TweetsReply{Text: tweet.Text})
		
		fmt.Println(&ts.TweetsReply{Text: tweet.Text})		
	}
	
	//Here is where I filter based on the passed in word (in.Name) which is trump for this test
	filterParams := &twitter.StreamFilterParams{
		Track:         []string{in.Name},
		StallWarnings: twitter.Bool(true),
	}

	//Get a the stream with new word filtered
	streamB, err := client.Streams.Filter(filterParams)
	if err != nil {
		log.Fatal(err)
	}
	
	/*
	This is a go routine I use to stop the server side streaming. I already explained in compute-service comments how grpc has no clean way to cancel a server stream from the client so
	to circumvent this I pass in the length of time I want the stream to run. Since I do not want to interrupt the stream 
	I run the function below in the background that gets the elapsed time and when a minute has reached it will stop the stream.
	*/

	go func(){
		f,_ := strconv.ParseFloat(in.Minutes, 64)
		dur := f * 60
		start := time.Now()
		rounded = 0

	 for rounded <= dur{
		end := time.Now()

		elapsed := end.Sub(start).Seconds()
		rounded = math.Floor(elapsed)
	}

	fmt.Println("Stopping Stream...")
	streamB.Stop()
	
	}()

	// This is another go routine used as part of the demus to start the stream again once the filtering has been done.	
	go demux.HandleChan(streamB.Messages)
	
	// I can also cancel the stream using CTRL+C which is handy for testing
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-ch)

	fmt.Println("Stopping Stream...")
	
	return nil
}



func main() {

	flag.Parse()
	
	//Have server listen on port 10000
	lis, err := net.Listen("tcp", ":10000")
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
	newServer := &TwitterServiceServer{}
	ts.RegisterTwitterServiceServer(grpcServer, newServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	
}
