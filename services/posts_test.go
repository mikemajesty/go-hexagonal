package services

import (
	"golang/projects/go-hexagonal/entities"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func CreateMock() IPostService {
	return &MockRepository{}
}

func (m *MockRepository) Validate(post *entities.Post) error {
	panic("unimplemented")
}

func (m *MockRepository) Create(post *entities.Post) (*entities.Post, error) {
	args := m.Called(post)
	return args.Get(0).(*entities.Post), args.Error(1)
}

func (m *MockRepository) FindAll() ([]entities.Post, error) {
	args := m.Called()
	return args.Get(0).([]entities.Post), args.Error(1)
}

func TestFindAll(t *testing.T) {
	postService := CreateMock()

	post := entities.Post{
		Title: "title",
		Text:  "text",
	}

	posts := []entities.Post{post}

	postService.(*MockRepository).On("FindAll").Return(posts, nil)

	result, _ := postService.FindAll()

	postService.(*MockRepository).AssertExpectations(t)

	assert.Equal(t, posts, result)
}

func TestValidateEmptyPost(t *testing.T) {
	postService := Create()
	err := postService.Validate(nil)

	assert.NotNil(t, err)
	assert.Equal(t, "the post is empty", err.Error())
}

func TestValidateEmptyPostTitle(t *testing.T) {
	postService := Create()
	err := postService.Validate(&entities.Post{
		Title: "",
		Text:  "text",
	})

	assert.NotNil(t, err)
	assert.Equal(t, "the post title is empty", err.Error())
}

func TestValidateEmptyPostText(t *testing.T) {
	postService := Create()
	err := postService.Validate(&entities.Post{
		Title: "title",
		Text:  "",
	})

	assert.NotNil(t, err)
	assert.Equal(t, "the post text is empty", err.Error())
}

func TestValidatePost(t *testing.T) {
	postService := Create()
	err := postService.Validate(&entities.Post{
		Title: "title",
		Text:  "text",
	})

	assert.Nil(t, err)
}
