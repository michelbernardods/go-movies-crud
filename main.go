package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/michelbernardods/go-crud-movies/handler"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/movies/listAllMovies", handler.GeAllMovies).Methods("GET")
	router.HandleFunc("/movies/listMovie/{id}", handler.GetMovie).Methods("GET")
	router.HandleFunc("/movies/createMovie", handler.CreateMovie).Methods("POST")
	router.HandleFunc("/movies/updateMovie/{id}", handler.UpdateMovie).Methods("PUT")
	router.HandleFunc("/movies/deleteMovie/{id}", handler.DeleteMovie).Methods("DELETE")

	fmt.Println("Start server port: 8080")
	http.ListenAndServe(":8080", router)
}
