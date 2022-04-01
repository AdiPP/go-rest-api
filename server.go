package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main () {
	router := mux.NewRouter()
	const port string = ":8000"

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Up and running...")
	})

	router.HandleFunc("/posts", getPosts).Methods(http.MethodGet)

	router.HandleFunc("/posts", createPost).Methods(http.MethodPost)

	log.Println("Server listening on port", port)
	log.Fatalln(http.ListenAndServe(port, router))
}