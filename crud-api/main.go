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
	Id       string    `json:"id"`
	Year     int       `json:"year"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

var movies []Movie


func getAllMovies (res http.ResponseWriter, req *http.Request ) {
	
	if req.URL.Path != "/movies"{
		http.Error(res, "404 not found", 404)
		return
	}
	if req.Method != "GET"{
		http.Error(res, "Unavaible  for this method", 405)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(movies)
}



func deleteMovie(res http.ResponseWriter, req *http.Request){

	res.Header().Set("Content-Type","application/json")
	params := mux.Vars(req)



	for  index , item := range movies {
		if item.Id == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break;
		}
		fmt.Fprintf(res, "There's no such a movie with id: %v", params["id"])

	}

	json.NewEncoder(res).Encode(movies)


}

func getSingleMovie(res http.ResponseWriter, req *http.Request){

	res.Header().Set("Content-Type","application/json")
	params := mux.Vars(req)



	for  _ , item := range movies {
		if item.Id == params["id"] {
			
		json.NewEncoder(res).Encode(item)
		break;	
		} 
		fmt.Fprintf(res, "There's no such a movie with id: %v", params["id"])
	}



}


	func createMovie(res http.ResponseWriter, req *http.Request){
		res.Header().Set("Content-Type", "application/json")
		var movie Movie 
		_ = json.NewDecoder(req.Body).Decode(&movie)
		movie.Id = strconv.Itoa(rand.Intn(1000000) + 3333) 
		movies = append(movies, movie)
		json.NewEncoder(res).Encode(movie)
	}


	func updateMovie(res http.ResponseWriter, req *http.Request){

		res.Header().Set("Content-Type","application/json")
		params := mux.Vars(req)
	
	
	
		for  index , item := range movies {
			if item.Id == params["id"] {
				movies = append(movies[:index], movies[index+1:]...)
				_ = json.NewDecoder(req.Body).Decode(&movies[index])
				movies[index].Id = strconv.Itoa(rand.Intn(1000000) + 3333) 
				movies = append(movies, movies[index])
				json.NewEncoder(res).Encode(movies[index])

			} else{
				fmt.Fprintf(res, "There's no such a movie with id: %v", params["id"])
			}
		}
	
		json.NewEncoder(res).Encode(movies)
	
	
	}






func main() {

	r := mux.NewRouter()

	movies = append(movies, Movie{
		Id: "412Fqsf342",
		Year: 2019,
		Title: "The Irishman",
		Director: &Director{"421SDQ41", "Martin  Scorsese"},
	}, Movie{
		Id: "2az12e23RGgd2",
		Year: 2023,
		Title: "Oppenheimer",
		Director: &Director{"421SDQ41", "Christopher Nolan"},
	},
	Movie{
		Id: "9SAadD8Adazd3",
		Year: 2024,
		Title: "Dune 2",
		Director: &Director{"421SDQ41", "Danis Chi"},
	},
)

	

	r.HandleFunc("/movies", getAllMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getSingleMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/update/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/delete/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Server running on port 8081")
	log.Fatal(http.ListenAndServe(":8081", r))



}
