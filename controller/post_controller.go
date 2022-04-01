package controller

import (
	"encoding/json"
	"net/http"

	"github.com/AdiPP/go-rest-api/entity"
	"github.com/AdiPP/go-rest-api/errors"
	"github.com/AdiPP/go-rest-api/service"
)

var (
	postService service.PostService
)

type PostController interface {
	GetPosts(w http.ResponseWriter, r *http.Request)
	CreatePost(w http.ResponseWriter, r *http.Request)
}

type Controller struct {}

func NewPostController(s service.PostService) PostController {
	postService = s
	return &Controller{}
}

func (c *Controller) GetPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	
	posts, err := postService.FindAll();

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)	
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "Error getting the posts."})
		return
	}	

	w.WriteHeader(http.StatusOK)
	
	json.NewEncoder(w).Encode(posts)
}

func (c *Controller) CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	
	var post entity.Post
	
	err := json.NewDecoder(r.Body).Decode(&post)
	 
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)	
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "Error unmarshaling request."})
		return
	}

	err2 := postService.Validate(&post)

	if err2 != nil {
		w.WriteHeader(http.StatusInternalServerError)	
		json.NewEncoder(w).Encode(errors.ServiceError{Message: err2.Error()})
		return
	}

	result, err3 := postService.Create(&post)

	if err3 != nil {
		w.WriteHeader(http.StatusInternalServerError)	
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "Error saving the post."})
	}
	
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}