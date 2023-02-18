package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

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


func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index],  movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}


func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var movie Movie

	_ = json.NewDecoder(r.Body).Decode(&movie)

	movie.ID = strconv.Itoa(rand.Intn(100000000))

}

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


	movies = append(movies, Movie{
		ID: "2",
		Isbn: "543",
		Title: "Love",
		Director: &Director{
			FirstName: "Lokesh",
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