package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
)

func welcomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to AutoParkingAPI")
}

func genQueueEndpoint(w http.ResponseWriter, r *http.Request) {
	genQueueFunc()
}

func viewPermitEndpoint(w http.ResponseWriter, r *http.Request) {
	viewPermFunc()
}

func viewCarsEndpoint(w http.ResponseWriter, r *http.Request) {
	viewCarFunc()
}

func handleRequest(){
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", welcomePage)
	router.HandleFunc("/queue", genQueueEndpoint).Methods("POST")
	router.HandleFunc("/view", viewPermitEndpoint).Methods("GET")
	log.Fatal(http.ListenAndServe(":18492", router))

}
