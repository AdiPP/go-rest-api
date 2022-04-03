package service

import (
	"testing"

	"github.com/AdiPP/go-rest-api/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (mock MockRepository) Save(post *entity.Post) (*entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*entity.Post), args.Error(1)
}

func (mock MockRepository) FindAll() ([]entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.Post), args.Error(1)
}

func TestFindAll(t *testing.T) {
	mockRepository := new(MockRepository)

	var identifier int64 = 1

	post := entity.Post{
		Id: identifier,
		Title: "A",
		Text: "B",
	}

	// Setup Expectaction
	mockRepository.On("FindAll").Return([]entity.Post{post}, nil)

	testingService := NewPostService(mockRepository)

	result, _ := testingService.FindAll()
	
	// Mock Assertion: Behavioral
	mockRepository.AssertExpectations(t)

	// Data Assertion
	assert.Equal(t, identifier, result[0].Id)
	assert.Equal(t, "A", result[0].Title)
	assert.Equal(t, "B", result[0].Text)
}

func TestCreate(t *testing.T) {
	mockRepository := new(MockRepository)
	
	var identifier int64 = 1

	post := entity.Post{
		Title: "A",
		Text: "B",
	}

	mockRepository.On("Save").Return(&post, nil)

	testingService := NewPostService(mockRepository)

	result, _ := testingService.Create(&post)

	mockRepository.AssertExpectations(t)

	assert.NotNil(t, identifier, result.Id)
	assert.Equal(t, "A", result.Title)
	assert.Equal(t, "B", result.Text)
}

func TestValidateEmptyPost(t *testing.T) {
	testingService := NewPostService(nil)

	err := testingService.Validate(nil)

	assert.NotNil(t, err)

	assert.Equal(t, err.Error(), "The post is empty")
}

func TestValidateEmptyPostTitle(t *testing.T) {
	post := entity.Post{
		Id: 1,
		Title: "",
		Text: "Test",
	}

	testingService := NewPostService(nil)

	err := testingService.Validate(&post)

	assert.NotNil(t, err)
	
	assert.Equal(t, err.Error(), "The post title is empty")
}

