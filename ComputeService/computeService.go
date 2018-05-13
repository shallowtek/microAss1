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
	//"net/url"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/testdata"
	ts "github.com/shallowtek/microAss1/TwitterService/proto"
	bs "github.com/shallowtek/microAss1/BbcService/proto"
	//rs "github.com/shallowtek/microAss1/RedisGateway/proto"
	"github.com/cdipaolo/sentiment"
	//"github.com/go-redis/redis"
	//"github.com/gomodule/redigo/redis"
	//"encoding/json"
    //"github.com/gorilla/mux"
	
	
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

type Result struct{
	
	Name  string   `json:"name,omitempty"` 
	Value string   `json:"value,omitempty"` 	
}


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
				
		analysis := model.SentimentAnalysis(tweet.Text, sentiment.English)
		score += int(analysis.Score)
		count++
		average = (float64)(score)/(float64)(count)		
				
		fmt.Printf("SCORE %d %d %6.1f \n",score, count, average)		
		
		
	}//end for loop	
	
	
	convertAvg := strconv.FormatFloat(average, 'f', 6, 64)
	
	db, _ := sql.Open("mysql", "root:mysql@tcp(mysql:3306)/resultDB") 
	defer db.Close()
	
	db.Query("INSERT INTO resultsTable ( name, value) VALUES ('" + in.Name + "', '" + convertAvg + "') ")
	
	
	
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
	
	convertAvg := strconv.FormatFloat(average, 'f', 6, 64)
	db, _ := sql.Open("mysql", "root:mysql@tcp(mysql:3306)/resultDB") 
	defer db.Close()
	
	db.Query("INSERT INTO resultsTable ( name, value) VALUES ('" + in.Name + "', '" + convertAvg + "') ")
	
	
}

//This function starts the twitter stream using the form data. A connection is made to Twitter service and the print features function 
//is called. A value is returned and sent to the web service to be displayed.
//func startTwitter(term string, timeN string, opts []grpc.DialOption){
//
//	conn, err := grpc.Dial("twitter-service:10000", opts...)
//	if err != nil {
//		log.Fatalf("fail to dial: %v", err)
//	}
//	defer conn.Close()
//
//	client := ts.NewTwitterServiceClient(conn)
//
//		printFeatures(client, &ts.TweetsRequest{
//		Name: term,
//		Minutes: timeN,
//		
//	})
//
//	
//}

//This function starts the bbc stream using the form data. A connection is made to bbc service and the print news function 
//is called. A value is returned and sent to the web service to be displayed.
//func startBbc(term string, timeN string, opts []grpc.DialOption){
//
//	conn, err := grpc.Dial("bbc-service:10005", opts... )
//	if err != nil {
//		log.Fatalf("fail to dial: %v", err)
//	}
//	defer conn.Close()
//
//	clientBbc := bs.NewBbcServiceClient(conn)
//
//		printNews(clientBbc, &bs.NewsRequest{
//		Name: term,
//		Minutes: timeN,
//		
//	})	
//		
//	
//}



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
	
	if choice == "Twitter"{

	
	conn, _ := grpc.Dial("twitter-service:10000", opts...)
	
	defer conn.Close()

	client := ts.NewTwitterServiceClient(conn)

		printFeatures(client, &ts.TweetsRequest{
		Name: term,
		Minutes: timeN,
		
	})
	
	}else if choice == "Bbc"{

	conn, _ := grpc.Dial("bbc-service:10005", opts... )
	
	defer conn.Close()

	clientBbc := bs.NewBbcServiceClient(conn)

		printNews(clientBbc, &bs.NewsRequest{
		Name: term,
		Minutes: timeN,
		
	})	

	}
	
}

func test(w http.ResponseWriter, r *http.Request) {
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
	
	conn, err := grpc.Dial("twitter-service:10000", opts...)
	
	defer conn.Close()

	client := ts.NewTwitterServiceClient(conn)

		printFeatures(client, &ts.TweetsRequest{
		Name: "trump",
		Minutes: "0.1",
		
	})
	
	fmt.Fprintf(w, "TEST PAGE \n")	
	fmt.Fprintf(w, "New Term: %v \n", err)	
	
	
}


func check(w http.ResponseWriter, r *http.Request) {
	
	db, _ := sql.Open("mysql", "root:mysql@tcp(mysql:3306)/resultDB") 
	defer db.Close()
	
	//db.Query("INSERT INTO resultsTable ( name, value) VALUES ('hilary', 'hilaryValue') ")
	
	results, err := db.Query("SELECT name, value FROM resultsTable") 
    
    fmt.Fprintf(w, "Sentiment Analysis Search Results: \n",)
    for results.Next() {
		var result Result
		// for each row, scan the result into our tag composite object
		results.Scan(&result.Name, &result.Value)
		
		fmt.Fprintf(w, "Term: %v	Result: %v\n", result.Name, result.Value)
			 
		}
	
	fmt.Fprintf(w, "New Term: %v \n", err)
	
	
}


func main() {
		
		
    http.HandleFunc("/start", handler)
	http.HandleFunc("/test", test)
	
	http.HandleFunc("/check", check)
	log.Fatal(http.ListenAndServe(":9090", nil))
	
	
	
}
