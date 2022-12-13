package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Movie struct {
	ID       string    `json:"id"`
	Title    string    `json:"Title"`
	Isbn     string    `json:"Isbn"`
	Year     string    `json:"Year"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"Firstname"`
	Lastname  string `json:"Lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
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
	movie.ID = strconv.Itoa(len(movies) + 1)
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
}

func main() {
	router := mux.NewRouter()

	movies = append(movies, Movie{ID: "100", Title: "Movie One", Isbn: "12345", Year: "2019", Director: &Director{Firstname: "John", Lastname: "Doe"}})

	router.HandleFunc("/api/movies", getMovies).Methods("GET")
	router.HandleFunc("/api/movies/{id}", getMovie).Methods("GET")
	router.HandleFunc("/api/movies", createMovie).Methods("POST")
	router.HandleFunc("/api/movies/{id}", updateMovie).Methods("PUT")
	router.HandleFunc("/api/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Starting the application on port 3000...")
	log.Fatal(http.ListenAndServe(":3000", router))
}
