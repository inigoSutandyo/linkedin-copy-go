package model

import (
	"gorm.io/gorm"
)

type Reply struct {
	gorm.Model
	Content   string
	UserID    uint
	User      User
	CommentID uint
	Comment   Comment
}
