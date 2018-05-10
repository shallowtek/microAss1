//Matt Shallow 14-Mar-18
package main

import (
	
	//"flag"
	
	"log"
	//"os"
	"net/http"
	//"golang.org/x/net/context"
	
	//"strconv"
	//rs "github.com/shallowtek/microAss1/RedisGateway/proto"
	"encoding/json"
	//"google.golang.org/grpc"
	"github.com/gomodule/redigo/redis"
	//"github.com/gorilla/mux"
	
)

//var(
//	port       = flag.Int("port", 10010, "The server port")
//	tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
//	certFile   = flag.String("cert_file", "", "The TLS cert file")
//	keyFile    = flag.String("key_file", "", "The TLS key file")
//
//)


//
//func (s *RedisGatewayServer) SetData(ctx context.Context, in *rs.KeyRequest) (*rs.Empty, error) {
//				
//	//set to redis service
//	conn, _ := redis.Dial("tcp", "redis:6379")
//	defer conn.Close()		
//		
//	conn.Do("SET", in.Key, in.Value)
//	
//	return &rs.Empty{}, nil
//	
//}

type Result struct{
	
	Key	string   `json:"key,omitempty"` 
	Value string   `json:"value,omitempty"` 	
}

//var results []Result

//GET RESULTS
func GetResult(w http.ResponseWriter, r *http.Request) {
	
	
	
//	conn, _ := redis.Dial("tcp", "redis:6379")
//	defer conn.Close()	
//	
//	val, _ := redis.String(conn.Do("GET", "trump"))


	
	//json.NewEncoder(w).Encode(&Result{Key: params["key"], Value: val})
	
	//w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(&Result{Key: "trump", Value: "value"})
	
	
}

//SEND RESULTS

func SendResult(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	key := string(r.FormValue("Key"))
	value := string(r.FormValue("Value"))
	
	conn, _ := redis.Dial("tcp", "redis:6379")
	defer conn.Close()	

	conn.Do("SET", key, value)

	
}

//func GetAllResults(w http.ResponseWriter, r *http.Request) {
//	conn, _ := redis.Dial("tcp", "redis:6379")
//	defer conn.Close()	
//	json.NewEncoder(w).Encode(results)
//}



func main() {
	
//	results = append(results, Result{Key: "trump", Value: "0.2"})
//	results = append(results, Result{Key: "hilary", Value: "0.3"})
//	results = append(results, Result{Key: "obama", Value: "0.4"})
	//router := mux.NewRouter()	
	//router.HandleFunc("/getresult", GetResult).Methods("GET")
//    router.HandleFunc("/getallresults", GetAllResults).Methods("GET")
    //router.HandleFunc("/sendresult", SendResult).Methods("POST")
	
	
    //log.Fatal(http.ListenAndServe(":8081", router))
    
    
    http.HandleFunc("/getresult", GetResult)
    log.Fatal(http.ListenAndServe(":8081", nil))
   
 
	
}
