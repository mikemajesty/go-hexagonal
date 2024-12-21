package repository

import (
	"golang/projects/go-hexagonal/entities"
	infra_database "golang/projects/go-hexagonal/infra/database"
	"log"
)

type IPostRepository interface {
	SavePost(post *entities.Post) (*entities.Post, error)
	FindAllPosts() ([]entities.Post, error)
}

type Repository struct{}

func Create() IPostRepository {
	return &Repository{}
}

func (*Repository) SavePost(post *entities.Post) (*entities.Post, error) {
	result := infra_database.DB.Create(post)

	if result.Error != nil {
		log.Fatal("Error saving post", result.Error)
		return nil, result.Error
	}

	return post, nil
}

func (*Repository) FindAllPosts() ([]entities.Post, error) {
	posts := []entities.Post{}
	result := infra_database.DB.Find(&posts)

	if result.Error != nil {
		log.Fatal("Error saving post", result.Error)
		return nil, result.Error
	}

	return posts, nil
}
