package model

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title      string `json:"title"`
	Content    []byte
	Attachment string
	Likes      int
	Template   Template `gorm:"embedded"`
	UserID     uint
	// PostLikes  []PostLike
}
