//Matt Shallow 14-Mar-18
package main

import (
	"flag"
	"io"
	"log"
	"fmt"
	"time"
	"math"	
	"strconv"
	"net/http"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/testdata"
	ts "github.com/shallowtek/microAss1/TwitterService/proto"
	bs "github.com/shallowtek/microAss1/BbcService/proto"
	"github.com/cdipaolo/sentiment"
	"github.com/go-redis/redis"
	

	
)

var (
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "", "The file containning the CA root cert file")
	serverHostOverride = flag.String("server_host_override", "x.test.youtube.com", "The server name use to verify the hostname returned by TLS handshake")
	rClient *redis.Client
	start time.Time
	end time.Time
	val string
	bbcVal string
)

//This function is called by the startTwitter function. It contains the GRPC communication method to communicate with the
//twitter service. A stream is returned and sentiment calculated and stored on redis.
func printFeatures(client ts.TwitterServiceClient, in *ts.TweetsRequest) string{
	score := 0
	count := 0
	var elapsed float64
	var rounded float64
	average := 0.0

	model, err := sentiment.Restore()

	if err != nil {
    	panic(fmt.Sprintf("Could not restore model!\n\t%v\n", err))
	}

	stream, err := client.GetTweets(context.Background(), in)
	if err != nil {
		log.Fatalf("%v.GetTweets(_) = _, %v", client, err)
	}

	start := time.Now()
	f,_ := strconv.ParseFloat(in.Minutes, 64)
	dur := f * 60

	for rounded < dur {

		end = time.Now()
		elapsed = end.Sub(start).Seconds()
		rounded = math.Floor(elapsed)

		tweet, err := stream.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error with stream")
		}
				
		analysis := model.SentimentAnalysis(tweet.Text, sentiment.English)
		score += int(analysis.Score)
		count++
		average = (float64)(score)/(float64)(count)		
				
		fmt.Printf("SCORE %d %d %6.1f \n",score, count,average)
		
		err = rClient.Set("twitScore", average, 0).Err()

		if err != nil {
			panic(err)
		}		
	}

	fmt.Println("pulling from redis")
		val, err := rClient.Get("twitScore").Result()

		if err != nil {
			panic(err)
		}


	return val 
	
	
}
//This function is called by the startBbc function. It contains the GRPC communication method to communicate with the
//bbc service. A stream is returned and sentiment calculated and stored on redis.
func printNews(client bs.BbcServiceClient, in *bs.NewsRequest) string{
	score := 0
	count := 0
	var elapsed float64
	var rounded float64
	average := 0.0

	model, err := sentiment.Restore()

	if err != nil {
    	panic(fmt.Sprintf("Could not restore model!\n\t%v\n", err))
	}

	stream, err := client.GetNews(context.Background(), in)
	if err != nil {
		log.Fatalf("%v.GetNews(_) = _, %v", client, err)
	}

	start := time.Now()
	f,_ := strconv.ParseFloat(in.Minutes, 64)
	dur := f * 60

	for rounded < dur {

		end := time.Now()
		elapsed = end.Sub(start).Seconds()
		rounded = math.Floor(elapsed)

		news, err := stream.Recv()
		
		if err != nil {
			fmt.Println("stream returned")
			break;
		}
			
		analysis := model.SentimentAnalysis(news.Text, sentiment.English)
		score += int(analysis.Score)
		count++
		average = (float64)(score)/(float64)(count)		
				
		fmt.Printf("SCORE %d %d %6.1f \n",score, count,average)
			
		rClient.Set("bbcScore", average, 0)

				
	}

	fmt.Println("pulling from redis")
		val, _ := rClient.Get("bbcScore").Result()

		
	return val 
	
	
}



//This function starts the twitter stream using the form data. A connection is made to Twitter service and the print features function 
//is called. A value is returned and sent to the web service to be displayed.
func startTwitter(term string, timeN string, opts []grpc.DialOption){

	conn, err := grpc.Dial("twitter-service:10000", opts... )
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := ts.NewTwitterServiceClient(conn)

	val = printFeatures(client, &ts.TweetsRequest{
		Name: term,
		Minutes: timeN,
		
	})

	fmt.Println("This is Val: " + val + "\n")
	
	conn.Close()

	_, errs := http.Get("http://web-service:8080/submitTwit/" + val)
	
	if errs != nil {
		log.Fatalf("fail to submit value twitter: %v", errs)
	}
}

//This function starts the bbc stream using the form data. A connection is made to bbc service and the print news function 
//is called. A value is returned and sent to the web service to be displayed.
func startBbc(term string, timeN string, opts []grpc.DialOption){

	conn, err := grpc.Dial("bbc-service:10005", opts... )
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	clientBbc := bs.NewBbcServiceClient(conn)

	bbcVal = printNews(clientBbc, &bs.NewsRequest{
		Name: term,
		Minutes: timeN,
		
	})	
	
	fmt.Println("This is BBC Val: " + bbcVal + "\n")	
	conn.Close()

	_, errs := http.Get("http://web-service:8080/submitBbc/" + bbcVal)
	
	if errs != nil {
		log.Fatalf("fail to submit value bbc: %v", errs)
	}
}

//This function handles the post request from web-service. It extracts the form data and determines which service to use (twitter or bbc)
//Extracted data is the search term and how long you want to search for
func handler(w http.ResponseWriter, r *http.Request) {
	
	flag.Parse()
	
	var opts []grpc.DialOption
	if *tls {
		if *caFile == "" {
			*caFile = testdata.Path("ca.pem")
		}
		creds, err := credentials.NewClientTLSFromFile(*caFile, *serverHostOverride)
		if err != nil {
			log.Fatalf("Failed to create TLS credentials %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}
    
        r.ParseForm()
        
	term := r.FormValue("term")
	timeN := r.FormValue("time")
	choice := r.FormValue("choice")
	//id := r.FormValue("id")

	if choice == "Twitter"{

	startTwitter(term, timeN, opts)
	
	}else if choice == "Bbc"{

	startBbc(term, timeN, opts)

	}
	
}

func main() {


	rClient = redis.NewClient(&redis.Options{	
		Addr: "redis:6379",		
	})

	rClient.FlushAll()
		
	http.HandleFunc("/start", handler)	
    	log.Fatal(http.ListenAndServe(":9090", nil))
	
	
	
}
