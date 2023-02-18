package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)




type Movie struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}


var movies []Movie





func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{
		ID: "1",
		Isbn: "542",
		Title: "Indian",
		Director: &Director{
			FirstName: "Rakesh",
			LastName: "Sharma",
		},
	})

	r.HandleFunc("/movies",getMovies).Methods("GET")
	r.HandleFunc("/movie/{id}",getMovie).Methods("GET")
	r.HandleFunc("/movies",createMovie).Methods("POST")
	r.HandleFunc("/movie/{id}",updateMovie).Methods("PUT")
	r.HandleFunc("/movie/{id}",deleteMovie).Methods("DELETE")

	fmt.Printf("starting the server at port :8000")

	err := http.ListenAndServe(":8000",r)

	if err !=nil {
		log.Fatal(err)
	}
}