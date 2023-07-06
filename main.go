package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var Movie []Movies

func main() {

	router := mux.NewRouter()
	Movie = append(Movie, Movies{1, "avengers"})
	Movie = append(Movie, Movies{1, "park"})
	router.HandleFunc("/", GetMovies).Methods("GET")
	fmt.Println("started")
	log.Fatal(http.ListenAndServe(":8080", router))

}

type Movies struct {
	ID   int
	Name string
}

func GetMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(Movie)
}
