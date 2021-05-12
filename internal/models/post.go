package models

// User defines the post in database
type Post struct {
	ID     int    `gorm:"primary_key;AUTO_INCREMENT" json:"id" xml:"id"`
	UserID int    `gorm:"type:BIGINT;NOT NULL" json:"user_id" xml:"user_id"`
	Title  string `gorm:"type:VARCHAR(300);NOT NULL" json:"title" xml:"title" validate:"required"`
	Body   string `gorm:"type:BLOB;NOT NULL" json:"body" xml:"body" validate:"required"`
}
