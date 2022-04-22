package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"tittle"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func main() {
	r := mux.NewRouter()

  movies = append(movies, Movie{
    ID: "1",
    Isbn: "438322",
    Title: "Movie one",
    Director: &Director:{Firstname:"Miro", Lastname:"Tamica"}})

  movies = append(movies, Movie{
    ID: "2",
    Isbn: "426590",
    Title: "Movie two", 
    Director: &Director: {Firstname: "Limo", Lastname: "Svetlic"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
  log.fatal(http.ListenAndServe(":8000", r))
}
