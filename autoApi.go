package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
)

type car struct {
	make  string `json: make`
	model string `json: model`
	color string `json: color`
	plate string `json: plate`
	name  string `json: name`
}

type apartment struct {
	name string `json: name`
	id   int    `json: id`
}

type email struct {
	address string `json: address`
	id      int    `json: id`
}

type permits struct {
	permit_id   int    `json: permit_id`
	car_id      int    `json: car_id`
	active_time string `json: active_time`
	location    string `json: location`
	active      bool   `json: active`
}

var wg sync.WaitGroup

func welcomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to AutoParkingAPI")
}

func genQueueEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to AutoParkingAPI queue")
}

func viewRecordsEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to AutoParkingAPI view")
	//viewPermFunc()
}

func addRecordsEndpoint(w http.ResponseWriter, r *http.Request) {

}
func deleteRecordEndpoint(w http.ResponseWriter, r *http.Request) {

}

func handleRequest() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", welcomePage)
	router.HandleFunc("/queue", genQueueEndpoint)
	router.HandleFunc("/add", addRecordsEndpoint).Methods("UPDATE")
	router.HandleFunc("/view", viewRecordsEndpoint).Methods("GET")
	router.HandleFunc("/delete", deleteRecordEndpoint).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8081", router))
}

func startHttpServer(wg *sync.WaitGroup) *http.Server {
	srv = &http.Server{Addr: ":8080"}
	handleRequest()

	go func() {
		defer wg.Done()
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			fmt.Println(err)
		}
	}()

	return srv
}
