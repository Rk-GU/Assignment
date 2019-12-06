package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"

)


func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func handleRequests() {
	
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", helloWorld).Methods("GET")
	myRouter.HandleFunc("/candidates", AllCandidates).Methods("GET")
	myRouter.HandleFunc("/candidate/{name}/{phonenumber}/{status}", NewCandidate).Methods("POST")
	myRouter.HandleFunc("/candidate/{phone}", DeleteCandidate).Methods("DELETE")
	myRouter.HandleFunc("/candidate/{id}", UpdateCandidate).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}


func main() {
	fmt.Println("Server Started")

	handleRequests()
}