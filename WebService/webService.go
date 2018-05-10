//Matt Shallow 14-Mar-18
//this is a very straigh forward web server. There are three handler functions to deal with requests from /home, /submit and /score
//the submit handler reads the get request from compute service and stores in a variable to be read by the score handler and display on page.
//Adding some test comments dvsdcdsdcsdc
package main

import (
    "fmt"
    "log"
    "net/http"
    "text/template"
    "net/url"
    "time"
//    "io/ioutil"
//    "encoding/json"
    "flag"
    //generate unique ID
	//"github.com/segmentio/ksuid"
	//"github.com/go-redis/redis"
	//"github.com/gomodule/redigo/redis" 
	
	//"github.com/gorilla/mux"
	//"reflect"
	
	rs "github.com/shallowtek/microAss1/RedisGateway/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/testdata"
)



var(

	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "", "The file containning the CA root cert file")
	serverHostOverride = flag.String("server_host_override", "x.test.youtube.com", "The server name use to verify the hostname returned by TLS handshake")
	rClient rs.RedisGatewayClient
	conn *grpc.ClientConn
	id time.Time
	//client rs.RedisGatewayClient
	//conn redis.Conn
	//rClient *redis.Client
	term string
	choice string
	//result *rs.KeyRequest
)

type Result struct{
	
	Key	string   `json:"key,omitempty"` 
	Value string   `json:"value,omitempty"` 	
}


//func genKsuid() string{
//	id = ksuid.New().String()
//	//fmt.Printf("github.com/segmentio/ksuid:  %s\n", id.String())
//	return id
//}

func GetResult(w http.ResponseWriter, r *http.Request) {
	
//	flag.Parse()
//	var opts []grpc.DialOption
//	if *tls {
//		if *caFile == "" {
//			*caFile = testdata.Path("ca.pem")
//		}
//		creds, err := credentials.NewClientTLSFromFile(*caFile, *serverHostOverride)
//		if err != nil {
//			log.Fatalf("Failed to create TLS credentials %v", err)
//		}
//		opts = append(opts, grpc.WithTransportCredentials(creds))
//	} else {
//		opts = append(opts, grpc.WithInsecure())
//	}
//	
//	
//	conn, err := grpc.Dial("redis-gateway:10006", opts...)
//	if err != nil {
//		fmt.Fprintf(w, "First Error: %v \n", err)
//	}
//	defer conn.Close()
	
//	rClient := rs.NewRedisGatewayClient(conn)
	result , _ := rClient.GetData(context.Background(), &rs.KeyRequest{Key: "trump", Value: " "})
	
	
    
	fmt.Fprintf(w, "This is the Twitter Sentiment score: %s \n", result)	
	fmt.Fprintf(w, "This is the Twitter Sentiment score: %s \n", result.GetKey())
	fmt.Fprintf(w, "This is the Twitter Sentiment score: %s \n", result.GetValue())
	
	//params := mux.Vars(r)
    //key := params["key"]
    
//    rClient = redis.NewClient(&redis.Options{
//		Addr:     "redis-master:6379",
//		Password: "", // no password set
//		DB:       0,  // use default DB
//	})
//	defer rClient.Close()
//	//rClient.Set("trump", "trumpValue", 0).Err()
//	
//	
//	
//	val, _ := rClient.Get("trump").Result()
	
	//val, _ := redis.String(conn.Do("GET", "trump"))
//	var netClient = &http.Client{
//  		Timeout: time.Second * 10,
//	}
//	resp, _ := netClient.Get("http://redis-gateway:8081/getresult")
//	defer resp.Body.Close()	
		
//	decoder := json.NewDecoder(resp.Body)
//    var data Result
//    decoder.Decode(&data)
	//fmt.Println(string(responseData))
	
//	conn, _ := redis.Dial("tcp", "redis:6379") 
//	defer conn.Close()
//	//conn.Do("SET", "trump", "value")
//	val, _ := redis.String(conn.Do("GET", "trump")) 

	
}

func SendResult(w http.ResponseWriter, r *http.Request){
	
	
	
//	 params := mux.Vars(r)
//     key := params["Key"]
//     value := params["Value"]
//	r.ParseForm()
//	
//	key := string(r.FormValue("Key"))
//	value := string(r.FormValue("Value"))
	
	
//	rClient = redis.NewClient(&redis.Options{
//		Addr:     "redis-master:6379",
//		Password: "", // no password set
//		DB:       0,  // use default DB
//	})
//	defer rClient.Close()
//	
//	rClient.Set("trump", "value", 0).Err()

//	conn, _ := redis.Dial("tcp", "redis-master:6379") 
//	defer conn.Close()
//	
//	conn.Do("SET", "trump", "value")
	
	
	
}

func handler(w http.ResponseWriter, r *http.Request) {
	  
    fmt.Println("method:", r.Method) //get request method
    if r.Method == "GET" {
        t, _ := template.ParseFiles("form.html")
        t.Execute(w, nil)
    } else {
        r.ParseForm()

	term = string(r.FormValue("Search Term"))
    timeP := string(r.FormValue("Search Time"))
	choice = string(r.FormValue("Choice"))
	id := time.Now()
	ts := id.String()
	uniqueKey := choice + ts
	
	resp, _ := http.PostForm("http://compute-service:9090/start", url.Values{"term": {term}, "time": {timeP}, "choice": {choice}, "uniqueKey": {uniqueKey}})
	defer resp.Body.Close()
	
//	conn, _ := redis.Dial("tcp", "redis:6379") 
//	defer conn.Close()
//	
//	conn.Do("SET", "bbcTrump", "bbc")
//	conn.Do("SET", "twitTrump", "twit")
	
	//fmt.Println("Submission sent") 
	http.Redirect(w, r, "/home", http.StatusSeeOther)

	
    }
		
}

//func handlerScore(w http.ResponseWriter, r *http.Request) {
//	
//	//params := mux.Vars(r)
//	//res, _ := http.Get("http://redis-Gateway:8000/getresult/twitbush")
//	
//	client := http.Client{
//		Timeout: time.Second * 2, // Maximum of 2 secs
//	}
//	
//	url := "http://redis-Gateway:8000/getresult/twitBbc.json"
//	req, err := http.NewRequest(http.MethodGet, url, nil)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	res, getErr := client.Do(req)
//	if getErr != nil {
//		log.Fatal(getErr)
//	}
//	
//	//defer res.Body.Close()
//	
//	body, readErr := ioutil.ReadAll(res.Body)
//	if readErr != nil {
//		log.Fatal(readErr)
//	}
//	
//	result := Result{}
//	errJson := json.Unmarshal(body, &result)
//	if errJson != nil {
//		fmt.Println(errJson)
//		return
//	}
//	//fmt.Println(people1.Number)
//	
//  	fmt.Fprintf(w, "This is the Twitter Sentiment score: %s \n", result.Value)
//
//   
//}



func main() {	
	
	
	//conn.Do("CONFIG SET", "DAEMONIZE", "YES")	
	//defer conn.Close()
//	router := mux.NewRouter()
//	
//	router.HandleFunc("/home", handler)
//    router.HandleFunc("/score", handlerScore).Methods("GET")
//    router.HandleFunc("/getresult", GetResult)
////    router.HandleFunc("/getallresults", GetAllResults).Methods("GET")
//    router.HandleFunc("/sendresult", SendResult)
//    log.Fatal(http.ListenAndServe(":8080", router))
  
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
	
	conn, _ = grpc.Dial("redis-gateway:10006", opts...)
	
	
	
	rClient = rs.NewRedisGatewayClient(conn)
	
	http.HandleFunc("/home", handler)
	http.HandleFunc("/getresult", GetResult)
 	log.Fatal(http.ListenAndServe(":8080", nil))
		
    	
}
