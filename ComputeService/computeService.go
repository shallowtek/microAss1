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
	"net/url"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/testdata"
	ts "github.com/shallowtek/microAss1/TwitterService/proto"
	bs "github.com/shallowtek/microAss1/BbcService/proto"
	rs "github.com/shallowtek/microAss1/RedisGateway/proto"
	"github.com/cdipaolo/sentiment"
	//"github.com/go-redis/redis"
	//"github.com/gomodule/redigo/redis"
	//"encoding/json"
    //"github.com/gorilla/mux"
	
	//force delete docker images docker rmi -f $(docker images | grep '^<none>' | awk '{print $3}')
	//$ docker rm $(docker ps -aq)
	//$ docker rmi $(docker images -q)
	//docker system prune
	
)

var (
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "", "The file containning the CA root cert file")
	serverHostOverride = flag.String("server_host_override", "x.test.youtube.com", "The server name use to verify the hostname returned by TLS handshake")
	//rClient *redis.Client
	start time.Time
	end time.Time
	val string
	bbcVal string
	id string
//	twitName string
//	bbcName string
	
	
	//conn redis.Conn

)

//type Result struct {
//    Value string `json:"value"`   
//}


// Display a single data
//func GetResult(w http.ResponseWriter, r *http.Request) {
//    params := mux.Vars(r)
//    key := params["key"] 
//    
//    conn, _ := redis.Dial("tcp", "redis:6379")
//	  defer conn.Close()
//	
//    val, _:= conn.Do("GET", key)
//    result := Result{Value: val.(string)}
//    	
//    json.NewEncoder(w).Encode(result)
//    return
//}

//This function is called by the startTwitter function. It contains the GRPC communication method to communicate with the
//twitter service. A stream is returned and sentiment calculated and stored on redis.
func printFeatures(client ts.TwitterServiceClient, in *ts.TweetsRequest){
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
				
		fmt.Printf("SCORE %d %d %6.1f \n",score, count, average)		
		
		
	}//end for loop	
	
	
	convertAvg := strconv.FormatFloat(average, 'f', 6, 64)
	resp, _ := http.PostForm("http://web-service:8080/sendresult/", url.Values{"key": {in.Name}, "value": {convertAvg}})
	defer resp.Body.Close()
	
	
	
	
}//end function
//This function is called by the startBbc function. It contains the GRPC communication method to communicate with the
//bbc service. A stream is returned and sentiment calculated and stored on redis.
func printNews(client bs.BbcServiceClient, in *bs.NewsRequest){
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
		
		fmt.Printf("SCORE %d %d %6.1f \n",score, count, average)

		
				
	}//end for loop

	//convertAvg := strconv.FormatFloat(average, 'f', 6, 64)
	//http.Post("http://redis-gateway:8080/sendresult/" + in.Name + "/" + convertAvg)
//	resp, _ := http.PostForm("http://redis-gateway:8000/sendresult", url.Values{"Key": {"trump"}, "Value": {"value"}})
//	defer resp.Body.Close()
	
//	rClient := redis.NewClient(&redis.Options{
//		Addr:     "web-service:6379",
//		Password: "", // no password set
//		DB:       0,  // use default DB
//	})
//	
//	rClient.Set("trump", "value", 0).Err()
	
//	conn, _ := redis.Dial("tcp", "redis-instance:6379") 
//	defer conn.Close()
//	
//	conn.Do("SET", "trump", "value")
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
	
	
	conn, err := grpc.Dial("redis-gateway:10006", opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	
	rClient := rs.NewRedisGatewayClient(conn)
	rClient.SetData(context.Background(), &rs.KeyRequest{Key: "trump", Value: "trumpValue"})
		
	
}

//This function starts the twitter stream using the form data. A connection is made to Twitter service and the print features function 
//is called. A value is returned and sent to the web service to be displayed.
func startTwitter(term string, timeN string, opts []grpc.DialOption){

	conn, err := grpc.Dial("twitter-service:10000")
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := ts.NewTwitterServiceClient(conn)

		printFeatures(client, &ts.TweetsRequest{
		Name: term,
		Minutes: timeN,
		
	})


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

		printNews(clientBbc, &bs.NewsRequest{
		Name: term,
		Minutes: timeN,
		
	})	
	
}

//This function handles the post request from web-service. It extracts the form data and determines which service to use (twitter or bbc)
//Extracted data is the search term and how long you want to search for
func handler(w http.ResponseWriter, r *http.Request) {
	
//	rClient := redis.NewClient(&redis.Options{
//		Addr:     "redis-master:6379",
//		Password: "", // no password set
//		DB:       0,  // use default DB
//	})
//	
//	defer rClient.Close()
//	
//	rClient.Set("trump", "value", 0).Err()
	
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
	uniqueKey := r.FormValue("uniqueKey")

	id = uniqueKey
	
	if choice == "Twitter"{

	startTwitter(term, timeN, opts)
	
	}else if choice == "Bbc"{

	startBbc(term, timeN, opts)

	}
	
}


func main() {
		
	//defer conn.Close()	
    http.HandleFunc("/start", handler)
	
	log.Fatal(http.ListenAndServe(":9090", nil))
	
	
	
}
