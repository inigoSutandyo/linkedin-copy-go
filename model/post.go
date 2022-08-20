package model

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title      string `json:"title"`
	Content    string
	Attachment string
	// PostDetails []PostDetail
	Template Template `gorm:"embedded"`
	UserID   uint
}
