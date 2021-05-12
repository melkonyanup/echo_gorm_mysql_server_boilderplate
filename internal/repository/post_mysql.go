package repository

import (
	"gorm.io/gorm"
	"myapp/internal/models"
)

type PostMySQL struct {
	db *gorm.DB
}

func NewPostMySQL(db *gorm.DB) *PostMySQL {
	return &PostMySQL{db: db}
}

func (r *PostMySQL) GetAllPosts() ([]models.Post, error) {
	var postsFromDB []models.Post
	result := r.db.Find(&postsFromDB)
	if result.Error != nil {
		return nil, result.Error
	}

	return postsFromDB, nil
}

func (r *PostMySQL) GetPost(postId int) (*models.Post, error) {
	foundPost := new(models.Post)
	result := r.db.First(foundPost, postId)
	if result.Error != nil {
		return nil, result.Error
	}

	return foundPost, nil
}

// CreatePostRecord creates a post record in the database
func (r *PostMySQL) CreatePost(user *models.Post) error {
	result := r.db.Create(user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// UpdatePost updates a post record in the database
func (r *PostMySQL) UpdatePost(post *models.Post) error {
	result := r.db.Save(post)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// DeletePost deletes post record in the database
func (r *PostMySQL) DeletePost(post *models.Post) error {
	result := r.db.Unscoped().Delete(post)
	if result.Error != nil {
		return result.Error
	}

	return nil
}