package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/michelbernardods/go-crud-movies/handler"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/movies/listAllMovies", handler.GeAllMovies).Methods("GET")
	router.HandleFunc("/movies/listMovie/{id}", handler.GetMovie).Methods("GET")
	router.HandleFunc("/movies/createMovie", handler.CreateMovie).Methods("POST")
	router.HandleFunc("/movies/updateMovie/{id}", handler.UpdateMovie).Methods("PUT")
	router.HandleFunc("/movies/deleteMovie/{id}", handler.DeleteMovie).Methods("DELETE")

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	fmt.Printf("Start server port%s", os.Getenv("PORT_ENV"))
	http.ListenAndServe(os.Getenv("PORT_ENV"), router)
}
