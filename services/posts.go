package services

import (
	"errors"
	"golang/projects/go-hexagonal/entities"
	repository "golang/projects/go-hexagonal/repositories"
)

var PostRepository repository.IPostRepository = repository.Create()

type IPostService interface {
	Validate(post *entities.Post) error
	Create(post *entities.Post) (*entities.Post, error)
	FindAll() ([]entities.Post, error)
}

type Service struct{}

func Create() IPostService {
	return &Service{}
}

func (*Service) Validate(post *entities.Post) error {
	if post == nil {
		return errors.New("the post is empty")
	}

	if post.Title == "" {
		return errors.New("the post title is empty")
	}

	if post.Text == "" {
		return errors.New("the post text is empty")
	}

	return nil
}

func (*Service) Create(post *entities.Post) (*entities.Post, error) {
	posts, _ := PostRepository.FindAllPosts()
	var id uint = uint(len(posts)) + 1

	post.ID = id
	return PostRepository.SavePost(post)
}

func (*Service) FindAll() ([]entities.Post, error) {
	return PostRepository.FindAllPosts()
}
