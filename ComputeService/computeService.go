//Matt Shallow 13-May-18
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
	//I put my proto files on github and they are used here so everything is on cloud not local machine.
	ts "github.com/shallowtek/microAss1/TwitterService/proto"
	bs "github.com/shallowtek/microAss1/BbcService/proto"	
	"github.com/cdipaolo/sentiment"
	"github.com/gomodule/redigo/redis"
    "github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"
    "database/sql"
	
)

var (
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "", "The file containning the CA root cert file")
	serverHostOverride = flag.String("server_host_override", "x.test.youtube.com", "The server name use to verify the hostname returned by TLS handshake")
	//rClient *redis.Client
	start time.Time
	endTwit time.Time
	endBbc time.Time
//	twitName string
//	bbcName string	
	//conn redis.Conn
)

//struct used as part of redis database
type Result struct{
	
	Name  string   `json:"name,omitempty"` 
	Value string   `json:"value,omitempty"` 	
}


//This function contains the GRPC communication method to communicate with the
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

	stream, _ := client.GetTweets(context.Background(), in)


	start := time.Now()
	f,_ := strconv.ParseFloat(in.Minutes, 64)
	dur := f * 60

	for rounded < dur  {

		endTwit = time.Now()
		elapsed = endTwit.Sub(start).Seconds()
		rounded = math.Floor(elapsed)

		tweet, err := stream.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error with stream")
		}
		//Send to sentiment analysis tool and get result		
		analysis := model.SentimentAnalysis(tweet.Text, sentiment.English)
		score += int(analysis.Score)
		count++
		average = (float64)(score)/(float64)(count)		
				
		fmt.Printf("SCORE %d %d %6.1f \n",score, count, average)		
		
		
	}//end for loop	
	
	//this is where the mysql and redis databases are called and the results are stored
	convertAvg := strconv.FormatFloat(average, 'f', 6, 64)
	
	
	db, _ := sql.Open("mysql", "root:mysql@tcp(mysql:3306)/resultDB") 
	defer db.Close()
	
	db.Query("INSERT INTO resultsTable ( name, value) VALUES ('" + in.Name + "-Twitter" + "', '" + convertAvg + "') ")
	
	
	c, _ := redis.Dial("tcp", "redis:6379")
	defer c.Close()
	c.Do("SET", in.Name + "-Twitter", convertAvg )
	c.Do("EXPIRE", in.Name + "-Twitter", 600)
	
	
	
}//end function

//This function contains the GRPC communication method to communicate with the
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

	//for loop used to determine how much time has passed
	for rounded < dur {

		endBbc = time.Now()
		elapsed = endBbc.Sub(start).Seconds()
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
	
	//converts the average into string format
	convertAvg := strconv.FormatFloat(average, 'f', 6, 64)
	//this is where the mysql and redis databases are called and the results are stored
	db, _ := sql.Open("mysql", "root:mysql@tcp(mysql:3306)/resultDB") 
	defer db.Close()
	
	db.Query("INSERT INTO resultsTable ( name, value) VALUES ('" + in.Name + "-BBC" + "', '" + convertAvg + "') ")
	
	c, _ := redis.Dial("tcp", "redis:6379")
    defer c.Close()
    c.Do("SET", in.Name + "-Bbc", convertAvg )
    c.Do("EXPIRE", in.Name + "-Bbc", 600)
}




//This function handles the post request from web-service. It extracts the form data and determines which service to use (twitter or bbc)
//Extracted data is the search term and how long you want to search for
func startHandler(w http.ResponseWriter, r *http.Request) {
	//using mux to extract values from url
	vars := mux.Vars(r)
	term := vars["term"] //term
	timeP := vars["timeP"] //how long to run stream
	choice := vars["choice"] //bbc or twitter
	
	flag.Parse()
	//Grpc needs to use safe dialing options
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
    

	//If choice is twitter then connect to twitter-service
	if choice == "Twitter"{
	
	conn, _ := grpc.Dial("twitter-service:10000", opts...)	
	defer conn.Close()

	client := ts.NewTwitterServiceClient(conn)

		printFeatures(client, &ts.TweetsRequest{
		Name: term,
		Minutes: timeP,
		
	})
	
	//If choice is bbc then connect to bbc-service
	}else if choice == "Bbc"{

	conn, _ := grpc.Dial("bbc-service:10005", opts... )	
	defer conn.Close()

	clientBbc := bs.NewBbcServiceClient(conn)

		printNews(clientBbc, &bs.NewsRequest{
		Name: term,
		Minutes: timeP,
		
	})	

	}//end if statements
	
}//end startHandler


func main() {
				
	//Using mux features to simplify extracting GET request variables
	r := mux.NewRouter()
	r.HandleFunc("/start/{term}/{timeP}/{choice}", startHandler).Methods("GET")

	http.ListenAndServe(":9090", r)
	
}
