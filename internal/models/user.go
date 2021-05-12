package models

// User defines the user in database
type User struct {
	ID        int    `gorm:"primary_key;AUTO_INCREMENT"`
	FirstName string `gorm:"type:VARCHAR(100);NOT NULL"`
	LastName  string `gorm:"type:VARCHAR(100);NOT NULL"`
	Email     string `gorm:"type:VARCHAR(255);NOT NULL;UNIQUE"`
	Password  string `gorm:"type:CHAR(60);NOT NULL"`
	Posts     []Post
}
