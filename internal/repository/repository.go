package repository

import (
	"gorm.io/gorm"
	"myapp/internal/models"
)

type User interface {
	CreateUserRecord(user *models.User) error
	FindUserByEmail(email *string) (*models.User, error)
}

type Post interface {
	GetAllPosts() ([]models.Post, error)
	GetPost(postId int) (*models.Post, error)
	CreatePost(post *models.Post) error
	UpdatePost(post *models.Post) error
	DeletePost(post *models.Post) error
}

type Repository struct {
	User
	Post
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		User: NewUserMySQL(db),
		Post: NewPostMySQL(db),
	}
}
