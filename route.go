package main

import (
	"encoding/json"
	"math/rand"
	"net/http"

	"github.com/AdiPP/go-rest-api/entity"
	"github.com/AdiPP/go-rest-api/repository"
)

var (
	postRepository repository.PostRepository = repository.NewPostRepository()
)

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	
	posts, err := postRepository.FindAll();

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)	
		w.Write([]byte(`{"error": "Error getting the posts."}`))
		return
	}	

	w.WriteHeader(http.StatusOK)
	
	json.NewEncoder(w).Encode(posts)
}

func createPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	
	var post entity.Post
	
	err := json.NewDecoder(r.Body).Decode(&post)
	 
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)	
		w.Write([]byte(`{"error": "Error unmarshaling request."}`))
		return
	}

	post.Id = rand.Int63()

	postRepository.Save(&post)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(post)
}