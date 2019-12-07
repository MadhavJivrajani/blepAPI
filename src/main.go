package main

import (
	"fmt"
	"net/http"
	"log"
	"github.com/gorilla/mux"
)

func api_test(w http.ResponseWriter, r *http.Request) { 
	//http.ResponseWriter assembles the HTTP server's response
	//http.Request represents the client't request 
	fmt.Fprintf(w, "API endpoint test")
}

func main() {
	r := mux.NewRouter().StrictSlash(true) //handles trailing slash behaviour in route
	r.HandleFunc("/", api_test)
	log.Fatal(http.ListenAndServe(":8000", r))
}
