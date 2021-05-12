package repository

import (
	"gorm.io/gorm"
	"myapp/internal/models"
)

type UserMySQL struct {
	db *gorm.DB
}

func NewUserMySQL(db *gorm.DB) *UserMySQL {
	return &UserMySQL{db: db}
}

// CreateUserRecord creates a controllers_post record in the database
func (r *UserMySQL) CreateUserRecord(user *models.User) error {
	result := r.db.Create(user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// FindUserByEmail searches for the user in the database by email
func (r *UserMySQL) FindUserByEmail(email *string) (*models.User, error) {
	user := &models.User{}
	result := r.db.Where("email = ?", email).First(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
