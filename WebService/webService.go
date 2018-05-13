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
     _ "github.com/go-sql-driver/mysql"
    "database/sql"
    //"strings"
//    "io/ioutil"
//    "encoding/json"
    //"flag"
    //generate unique ID
	//"github.com/segmentio/ksuid"
	//"github.com/go-redis/redis"
	//"github.com/gomodule/redigo/redis" 
	
	//"github.com/go-redis/cache"
	
	//"github.com/gorilla/mux"
	//"reflect"
	
	//rs "github.com/shallowtek/microAss1/RedisGateway/proto"
//	"golang.org/x/net/context"
//	"google.golang.org/grpc"
//	"google.golang.org/grpc/credentials"
//	"google.golang.org/grpc/testdata"
)


var(


	//client rs.RedisGatewayClient
	//conn redis.Conn
	//rClient *redis.Client
//	term string
//	newKey string
//	newValue string
	//newTerm string
//	choice string
	
	//result *rs.KeyRequest
)

type Result struct{
	
	Name  string   `json:"name,omitempty"` 
	Value string   `json:"value,omitempty"` 	
}


func handler(w http.ResponseWriter, r *http.Request) {
	  
    fmt.Println("method:", r.Method) //get request method
    if r.Method == "GET" {
        t, _ := template.ParseFiles("form.html")
        t.Execute(w, nil)
    } else {
        r.ParseForm()

	term := string(r.FormValue("Search Term"))
    timeP := string(r.FormValue("Search Time"))
	choice := string(r.FormValue("Choice"))
	
	
	resp, _ := http.PostForm("http://compute-service:8080/start", url.Values{"term": {term}, "time": {timeP}, "choice": {choice}})
	defer resp.Body.Close()
		
	//fmt.Println("Submission sent") 
	http.Redirect(w, r, "/home", http.StatusSeeOther)

    }		
}

func GetResult(w http.ResponseWriter, r *http.Request) {
	
	//db, _ := sql.Open("mysql", "root:mysql@tcp(mysql:3306)/")
	
//	db.Exec("DROP DATABASE resultDB")
//	
//	db.Exec("CREATE DATABASE resultDB")
//
//    db.Exec("USE resultDB")
//    
//    db.Exec("CREATE TABLE resultsTable (name VARCHAR(32), value VARCHAR(32))")

     
	
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
	
	
    
//	err := db.Ping()
//	if errTwo != nil {
//    	//panic(err.Error()) // proper error handling instead of panic in your app
//	}
	
	 // perform a db.Query insert 
//    insert, _ := db.Query("INSERT INTO results VALUES ( 'trump', 'value' )")
    
    // if there is an error inserting, handle it
//    if err != nil {
//        panic(err.Error())
//    }
    
//    results, _ := db.Query("SELECT * FROM results")
//	if err != nil {
//		panic(err.Error()) // proper error handling instead of panic in your app
//	}
	
//	fmt.Fprintf(w, "HEADING:  \n")
//	fmt.Fprintf(w, "New Term: %v \n", results)
//	fmt.Fprintf(w, "New Term: %v \n", err)	
		
	
}

func main() {	
	  
    http.HandleFunc("/home", handler)
    http.HandleFunc("/getresult", GetResult)
    log.Fatal(http.ListenAndServe(":8080", nil))
	
	
    	
}
