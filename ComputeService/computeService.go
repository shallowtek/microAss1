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

	//This proto location was used for local testing. I then realised that I can just use github so that docker can just pull from there
	//Like it does with the other imports instead of configuring to work from local mahcine	.
	//ts "microAss1/TwitterService/proto"

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
	/*
	score uint64 = 0
	count uint32 = 0
	average float64 = 0.0
	elapsed uint64
	rounded float64
	*/
	start time.Time
	end time.Time
	val string
	bbcVal string
)

/*
This function is responsible for the main computation. A call to the twitter service is made and a stream of tweets are received.
The tweets are passed into the sentiment analysis tool and the score extracted. For each new tweet, the score 
is added on so that an average can be taken. A value is then returned.
*/
func printFeatures(client ts.TwitterServiceClient, in *ts.TweetsRequest) string{
	score := 0
	count := 0
	var elapsed float64
	var rounded float64
	average := 0.0
	//Create a new sentiment model
	model, err := sentiment.Restore()

	if err != nil {
    	panic(fmt.Sprintf("Could not restore model!\n\t%v\n", err))
	}

	//Make call to the twitter service method passing in context and the tweets request struct with variables (grpc)
	stream, err := client.GetTweets(context.Background(), in)
	if err != nil {
		log.Fatalf("%v.GetTweets(_) = _, %v", client, err)
	}

	//I use time to determined how long a stream should last before being stopped.
	start := time.Now()
	//Get the Minutes var's value and convert to a float64 so it can work with time
	f,_ := strconv.ParseFloat(in.Minutes, 64)
	//create var duration to convert to seconds
	dur := f * 60

	//rounded describes how much time has elapsed since function started. So this will loop until duration has reached e.g 1 minute
	for rounded < dur {

		//need end time so this will keep updating and be used to get elapsed time
		end = time.Now()
		elapsed = end.Sub(start).Seconds()
		rounded = math.Floor(elapsed)


		//get the tweets from the stream
		tweet, err := stream.Recv()

		//some error checking from stream. EOF not really needed as it is a stream.
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error with stream")
		}
		
		//Pass each tweet into sentiment analysis tool and get the score. Count each tweet so you can get average.		
		analysis := model.SentimentAnalysis(tweet.Text, sentiment.English)
		score += int(analysis.Score)
		count++
		average = (float64)(score)/(float64)(count)		
				
		//print score so I can it is computing correctly
		fmt.Printf("SCORE %d %d %6.1f \n",score, count,average)
			
		//Here I set the new score and push to Redis service
		err = rClient.Set("Score", average, 0).Err()

		if err != nil {
			panic(err)
		}		
	}

	/*
	Once the loop has finished and an average score has been pushed to Redis service I can now pull from Redis
	the final score.
	*/
	fmt.Println("pulling from redis")
		val, err := rClient.Get("Score").Result()

		if err != nil {
			panic(err)
		}

	
	//I return the value
	return val 
	
	
}
/*
------------------------------------------------------------------------------------------------------
BBC NEWS FUNCTION
-------------------------------------------------------------------------------------------------------
*/
func printNews(client bs.BbcServiceClient, in *bs.NewsRequest) string{
	score := 0
	count := 0
	var elapsed float64
	var rounded float64
	average := 0.0
	//Create a new sentiment model
	model, err := sentiment.Restore()

	if err != nil {
    	panic(fmt.Sprintf("Could not restore model!\n\t%v\n", err))
	}

	//Make call to the bbc service method passing in context and the tweets request struct with variables (grpc)
	stream, err := client.GetNews(context.Background(), in)
	if err != nil {
		log.Fatalf("%v.GetNews(_) = _, %v", client, err)
	}

	//I use time to determined how long a stream should last before being stopped.
	start := time.Now()
	//Get the Minutes var's value and convert to a float64 so it can work with time
	f,_ := strconv.ParseFloat(in.Minutes, 64)
	//create var duration to convert to seconds
	dur := f * 60

	//rounded describes how much time has elapsed since function started. So this will loop until duration has reached e.g 1 minute
	for rounded < dur {

		//need end time so this will keep updating and be used to get elapsed time
		end := time.Now()
		elapsed = end.Sub(start).Seconds()
		rounded = math.Floor(elapsed)


		//get the tweets from the stream
		news, err := stream.Recv()

		//some error checking from stream. EOF not really needed as it is a stream.
		
		if err != nil {
			fmt.Println("stream returned")
			break;
		}
		
		//Pass each tweet into sentiment analysis tool and get the score. Count each 			tweet so you can get average.		
		analysis := model.SentimentAnalysis(news.Text, sentiment.English)
		score += int(analysis.Score)
		count++
		average = (float64)(score)/(float64)(count)		
				
		//print score so I can it is computing correctly
		fmt.Printf("SCORE %d %d %6.1f \n",score, count,average)
			
		//Here I set the new score and push to Redis service
		rClient.Set("Score2", average, 0)

				
	}

	/*
	Once the loop has finished and an average score has been pushed to Redis service I can now pull from Redis
	the final score.
	*/
	fmt.Println("pulling from redis")
		val, _ := rClient.Get("Score2").Result()

		
	
	//I return the value
	return val 
	
	
}




func startTwitter(term, time, opts){

	conn, err := grpc.Dial("twitter-service:10000", opts... )
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := ts.NewTwitterServiceClient(conn)

	val = printFeatures(client, &ts.TweetsRequest{
		Name: term,
		Minutes: time,
		
	})

	fmt.Println("This is Val: " + val + "\n")
	
	conn.Close()

	_, errs := http.Get("http://web-service:8080/submitTwit/" + val)
	
	if errs != nil {
		log.Fatalf("fail to submit value twitter: %v", errs)
	}
	



}

func startBbc(term, time, opts){

	conn, err := grpc.Dial("bbc-service:10005", opts... )
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	clientBbc := bs.NewBbcServiceClient(conn)

	bbcVal = printNews(clientBbc, &bs.NewsRequest{
		Name: term,
		Minutes: time,
		
	})	
	
	fmt.Println("This is BBC Val: " + bbcVal + "\n")	
	conn.Close()

	_, errs = http.Get("http://web-service:8080/submitBbc/" + bbcVal)
	
	if errs != nil {
		log.Fatalf("fail to submit value bbc: %v", errs)
	}
}




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
	time := r.FormValue("time")
	choice := r.FormValue("choice")
	//choice = r.Form["Search Choice"]

	if choice == "Twitter"{

	startTwitter(term, time, opts)
	
	}else if choice == "Bbc"{

	startBbc(term, time, opts)

	}
	
}

func main() {


	

	//select the search term and how long to run search for
	//word := "trump"
	//mins := "1"

	
	//Create a new redis client by connecting to the redis container launched with docker-compose
	rClient = redis.NewClient(&redis.Options{
	
		Addr: "redis:6379",
		
	})

	//If you want to do a second search its good to flush previous data on Redis
	rClient.FlushAll()
	


	
	http.HandleFunc("/start", handler)	
    	log.Fatal(http.ListenAndServe(":9090", nil))
	
	
	
}
