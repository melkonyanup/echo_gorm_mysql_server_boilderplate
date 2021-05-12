package service

import (
	"myapp/internal/models"
	"myapp/internal/repository"
	"myapp/internal/shared/payloads"
)

//go:generate mockery --dir . --name User --output ./mocks
type User interface {
	SignUp(payload *payloads.SignUpPayload) (error)
	SignIn(payload *payloads.SignInPayload) (string, error)
	GetUserProfile(email string) (*models.User, error)
}

//go:generate mockery --dir . --name Post --output ./mocks
type Post interface {
	GetAllPosts() ([]models.Post, error)
	GetPost(postId int) (*models.Post, error)
	CreatePost(post *models.Post) error
	UpdatePost(userIdFromToken, postId int, payload *payloads.UpdatePostPayload) error
	DeletePost(userIdFromToken, postId int) error
}

type Service struct {
	User
	Post
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		User: NewUserService(repos.User),
		Post: NewPostService(repos.Post),
	}
}

