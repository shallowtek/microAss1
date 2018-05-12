//Matt Shallow 14-Mar-18
package main

import (
	
	//"flag"
	
	"log"
	//"os"
	"net/http"
	//"golang.org/x/net/context"
	//"net"
	//"strconv"
	//rs "github.com/shallowtek/microAss1/RedisGateway/proto"
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	//"github.com/gorilla/mux"
//	"google.golang.org/grpc/credentials"
//	"google.golang.org/grpc"
//	"google.golang.org/grpc/testdata"
)



type Result struct{
	
	Key	string   `json:"key,omitempty"` 
	Value string   `json:"value,omitempty"` 	
}

//var results []Result

//GET RESULTS
func GetResult(w http.ResponseWriter, r *http.Request) {
	
	//defer r.Body.Close()
//	conn, _ := redis.Dial("tcp", "redis:6379")
//	defer conn.Close()	
//	
//	val, _ := redis.String(conn.Do("GET", "trump"))

	json.NewEncoder(w).Encode(&Result{Key: "trump", Value: "value"})
	
	
}

//SEND RESULTS

func SetResult(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	key := string(r.FormValue("Key"))
	value := string(r.FormValue("Value"))
	
	conn, _ := redis.Dial("tcp", "redis:6379")
	defer conn.Close()	

	conn.Do("SET", key, value)

	
}

	
func main() {
	
	http.HandleFunc("/getresult", GetResult)
	http.HandleFunc("/setresult", SetResult)

	
 	log.Fatal(http.ListenAndServe(":10006", nil))
 	
}
