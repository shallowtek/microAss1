//Matt Shallow 13-May-18

package main

import (
    "fmt"
    "log"
    "net/http"
    "text/template"
     _ "github.com/go-sql-driver/mysql"
    "database/sql"
	"github.com/gomodule/redigo/redis" 
)

//This struct is used for mapping data in mysql database so that values can be extracted and displayed
type Result struct{
	
	Name  string   `json:"name,omitempty"` 
	Value string   `json:"value,omitempty"` 	
}

//This function responds to the form submission and sends results to compute-service
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
		
		
		//send request with values to compute-service so that it can start the sentiment anlysis
		//I put it in a go routine so that the user can enter a number of terms one after another and let
		//them be searched in the background. 
		go func(){
			
			resp, _ := http.Get("http://compute-service:9090/start/" + term + "/" + timeP + "/" + choice)			
			resp.Body.Close()
	
		}()

		//redirect to home
		http.Redirect(w, r, "/home", http.StatusSeeOther)

    }		
}

//This function will display the results from the mysql database which is persistent and data stored using mysql volume claim
func GetMysqlResult(w http.ResponseWriter, r *http.Request) {
	
	//connect to mysql database
	db, _ := sql.Open("mysql", "root:mysql@tcp(mysql:3306)/resultDB") 
	defer db.Close()
	
	
	results, _ := db.Query("SELECT name, value FROM resultsTable") 
    
    fmt.Fprintf(w, "MYSQL SENTIMENT RESULTS: \n",)
    
    
    for results.Next() {
		var result Result
		// for each row, scan the result into our result struct
		results.Scan(&result.Name, &result.Value)
		
		fmt.Fprintf(w, "Term: %v \t Result: %v\n", result.Name, result.Value)
				 
		}		
}//end function


//This function will display the results from the redis database
func GetRedisResult(w http.ResponseWriter, r *http.Request) {
	
	c, _ := redis.Dial("tcp", "redis:6379")
    defer c.Close()
       
    keys, _ := redis.Strings(c.Do("KEYS", "*"))
    
	fmt.Fprintf(w, "REDIS SENTIMENT RESULTS:  \n")
	for _, key := range keys {
		
	   value, _ := redis.String(c.Do("GET", key))
	   
	   fmt.Fprintf(w, "Term: %v 	Value: %v\n", key, value)
	   
	}	
	
}

//I use this function for general testing and clearing table results
func DropTable(w http.ResponseWriter, r *http.Request) {
	
	c, _ := redis.Dial("tcp", "redis:6379")
    defer c.Close()
       
	c.Do("FLUSHDB")
	
	//db, _ := sql.Open("mysql", "root:mysql@tcp(mysql:3306)/")
	db, _ := sql.Open("mysql", "root:mysql@tcp(mysql:3306)/resultDB") 
	defer db.Close()
	
	db.Exec("DROP TABLE resultsTable")
	
	db.Exec("CREATE TABLE resultsTable (name VARCHAR(32), value VARCHAR(32))")
	
	http.Redirect(w, r, "/home", http.StatusSeeOther)
}


func main() {	
	
	//brings you to the form page
    http.HandleFunc("/home", handler) 
    //brings you to the mysql results page
    http.HandleFunc("/mysqlresult", GetMysqlResult)
    //brings you to the redis results page
    http.HandleFunc("/redisresult", GetRedisResult)
    //used to drop all tables and other tests
    http.HandleFunc("/droptable", DropTable)
    log.Fatal(http.ListenAndServe(":8080", nil))
	
	
    	
}
