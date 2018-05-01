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
    //generate unique ID
	//"github.com/segmentio/ksuid"
	//"github.com/go-redis/redis"
	"github.com/gomodule/redigo/redis" 
	rs "github.com/shallowtek/microAss1/RedisGateway/proto"
	//"github.com/gorilla/mux"
	
)

//type Result struct {
//	Value string `json:"value"`
//}

var(


	id time.Time
	//rClient *redis.Client
	//conn redis.Conn
	term string
	choice string
)


//func genKsuid() string{
//	id = ksuid.New().String()
//	//fmt.Printf("github.com/segmentio/ksuid:  %s\n", id.String())
//	return id
//}


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
	http.Redirect(w, r, "/score", http.StatusSeeOther)
    }
		
}

func handlerScore(w http.ResponseWriter, r *http.Request) {
	
	
	conn, err := grpc.Dial("redis-gateway:10010", opts... )
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
    
	client := rs.NewRedisGatewayClient(conn)
	newKey := choice+term
	val, _ := client.GetData(context.Background(), &rs.KeyRequest{newKey, ""})
		
//    combinedTerm := choice + term
//  	
//  	res, _ := http.Get("http://compute-service:9090/result/" + combinedTerm)
  	
//  	body, readErr := ioutil.ReadAll(res.Body)
//	if readErr != nil {
//		log.Fatal(readErr)
//	}
	
//	result := Result{}
//	jsonErr := json.Unmarshal(body, &result)
//	if jsonErr != nil {
//		log.Fatal(jsonErr)
//	}
  	
  	fmt.Fprintf(w, "This is the " + choice + " Sentiment score: %s \n", val.Value)  
  	
   
    
}

func handlerSubmitTwit(w http.ResponseWriter, r *http.Request) {
	
	//twitScore = r.URL.Path[12:]
	
    fmt.Fprintf(w, "You have submiited new twitter score")
}

func handlerSubmitBbc(w http.ResponseWriter, r *http.Request) {
	
	//bbcScore = r.URL.Path[11:]
	
    fmt.Fprintf(w, "You have submiited new bbc score")
}


func main() {
	
//	var pool = newPool()
//	conn = pool.Get()
//	rClient = redis.NewClient(&redis.Options{	
//		Addr: "redis:6379",		
//	})
//	

//	router := mux.NewRouter()
//	router.HandleFunc("/home", handler).Methods("GET")
//	router.HandleFunc("/score", handlerScore).Methods("GET")
//	
//    //router.HandleFunc("/results/{key}", GetResults).Methods("GET")
//		
//    log.Fatal(http.ListenAndServe(":8080", router))
    
    
	http.HandleFunc("/home", handler)
	http.HandleFunc("/score", handlerScore)
//	http.HandleFunc("/submitTwit/", handlerSubmitTwit)
//	http.HandleFunc("/submitBbc/", handlerSubmitBbc)
	
	log.Fatal(http.ListenAndServe(":8080", nil))	
    	
}
