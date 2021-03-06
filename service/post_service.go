package service

import (
	"errors"
	"math/rand"

	"github.com/AdiPP/go-rest-api/entity"
	"github.com/AdiPP/go-rest-api/repository"
)

type PostService interface {
	Validate(post *entity.Post) error
	Create(post *entity.Post) (*entity.Post, error)
	FindAll()  ([]entity.Post, error)
}

type Service struct {}

var (
	postRepository repository.PostRepository
)

func NewPostService(r repository.PostRepository) PostService {
	postRepository = r
	return &Service{}
}

func (s *Service) Validate(post *entity.Post) error {
	if post == nil {
		err := errors.New("The post is empty")
		return err
	}

	if post.Title == "" {
		err := errors.New("The post title is empty")
		return err
	}

	return nil
}

func (s *Service) Create(post *entity.Post) (*entity.Post, error) {
	post.Id = rand.Int63()

	return postRepository.Save(post)
}

func (s *Service) FindAll()  ([]entity.Post, error) {
	return postRepository.FindAll()
}