//Matt Shallow 14-Mar-18
//this is a very straigh forward web server. There are three handler functions to deal with requests from /home, /submit and /score
//the submit handler reads the get request from compute service and stores in a variable to be read by the score handler and display on page.

package main

import (
    "fmt"
    "log"
    "net/http"

	
)
var twitScore = ""
var bbcScore = ""

func handler(w http.ResponseWriter, r *http.Request) {
	
    fmt.Fprintf(w, "This is the Home")
	
}

func handlerScore(w http.ResponseWriter, r *http.Request) {
	
    fmt.Fprintf(w, "This is the Twitter Sentiment score: %s \n", twitScore)

    fmt.Fprintf(w, "This is the Bbc Sentiment score: %s", bbcScore)
}

func handlerSubmitTwit(w http.ResponseWriter, r *http.Request) {
	twitScore = r.URL.Path[12:]
    fmt.Fprintf(w, "You have submiited new twitter score")
}

func handlerSubmitBbc(w http.ResponseWriter, r *http.Request) {
	bbcScore = r.URL.Path[11:]
    fmt.Fprintf(w, "You have submiited new bbc score")
}


func main() {

	
	http.HandleFunc("/home", handler)
	http.HandleFunc("/score", handlerScore)
	http.HandleFunc("/submitTwit/", handlerSubmitTwit)
	http.HandleFunc("/submitBbc/", handlerSubmitBbc)
	
    	log.Fatal(http.ListenAndServe(":8080", nil))



	

	
    	
}
