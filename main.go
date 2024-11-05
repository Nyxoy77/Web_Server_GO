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
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func main() {
	movies = append(movies, Movie{ID: "1", Isbn: "2322", Title: "The Fairy Tale", Director: &Director{Firstname: "Shivam", Lastname: "rai"}})
	movies = append(movies, Movie{ID: "2", Isbn: "2333", Title: "Shadow Slave", Director: &Director{Firstname: "Guilty", Lastname: "Three"}})

	r := mux.NewRouter()
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting the server at 8000 port  \n")
	log.Fatal(http.ListenAndServe(":8000", r))

}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(movies); err != nil {
		http.Error(w, "Failed to encode movies", http.StatusInternalServerError)
		return
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
	//Show the movies after they are deleted

	json.NewEncoder(w).Encode(movies)

}

func createMovie(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		http.Error(w, "wrong method", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	var mov Movie
	decoder := json.NewDecoder(r.Body)

	if error := decoder.Decode(&mov); error != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}
	mov.ID = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies, mov)

	json.NewEncoder(w).Encode(movies)

}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	//Find if the id of the movie already exists and if it does delete that movie and add the new movie by taking the data
	//from the body
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var mov Movie
			json.NewDecoder(r.Body).Decode(&mov)
			mov.ID = params["id"]
			movies = append(movies, mov)
			return
		}
	}
	var mov Movie
	json.NewDecoder(r.Body).Decode(&mov)
	movies = append(movies, mov)

}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(movies[index])
			break
		}
	}
}
