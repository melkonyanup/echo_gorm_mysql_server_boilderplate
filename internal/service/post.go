package service

import (
	"errors"
	"myapp/internal/models"
	"myapp/internal/repository"
	"myapp/internal/shared/payloads"
)

type PostService struct {
	repo repository.Post
}

func NewPostService(repo repository.Post) *PostService {
	return &PostService{repo: repo}
}

func (s *PostService) GetAllPosts() ([]models.Post, error) {
	return s.repo.GetAllPosts()
}

func (s *PostService) CreatePost(user *models.Post) error {
	return s.repo.CreatePost(user)
}

func (s *PostService) GetPost(postId int) (*models.Post, error) {
	return s.repo.GetPost(postId)
}

func (s *PostService) UpdatePost(userIdFromToken, postId int, payload *payloads.UpdatePostPayload) error {
	// check if user is not updating somebody elses post
	postFromDB, err := s.repo.GetPost(postId)
	if err != nil {
		return err
	}
	if userIdFromToken != postFromDB.UserID {
		return errors.New("don't try to update post that you didn't create")
	}

	postToUpdate := &models.Post{ID: postId, UserID: userIdFromToken, Title: payload.Title, Body: payload.Body}
	return s.repo.UpdatePost(postToUpdate)
}

func (s *PostService) DeletePost(userIdFromToken, postId int) error {
	// check if user is not deleting somebody elses post
	postFromDB, err := s.repo.GetPost(postId)
	if err != nil {
		return err
	}
	if userIdFromToken != postFromDB.UserID {
		return errors.New("don't try to delete post that you didn't create")
	}

	return s.repo.DeletePost(postFromDB)
}