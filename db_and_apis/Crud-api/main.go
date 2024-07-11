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
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	Lastname  string `'json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	jsonBytes, err := json.Marshal(movies)
	if err != nil {
		log.Printf("Failed to marshal movies: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			w.Write([]byte("Movie deleted successfully!"))
			return
		}
	}
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			jsonBytes, err := json.Marshal(item)
			if err != nil {
				log.Printf("Failed to marshal movie: %v", err)
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write(jsonBytes)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	err := json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000000))
	if err != nil {
		log.Printf("Failed to decode movie: %v", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	movies = append(movies, movie)
	jsonBytes, err := json.Marshal(movie)
	if err != nil {
		log.Printf("Failed to marshal movie: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonBytes)
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
			// jsonBytes, _ := json.Marshal(movies)
			w.Write([]byte("Movie updated successfully!"))
			return
		}
	}
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "448743", Title: "Movie One", Director: &Director{FirstName: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "448744", Title: "Movie Two", Director: &Director{FirstName: "Steve", Lastname: "Smith"}})
	movies = append(movies, Movie{ID: "3", Isbn: "448745", Title: "Movie Three", Director: &Director{FirstName: "Jane", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "4", Isbn: "448746", Title: "Movie Four", Director: &Director{FirstName: "Mike", Lastname: "Smith"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")

	fmt.Println("Starting server on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Server failed to start %v", err)
	}
}
