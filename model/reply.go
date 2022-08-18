package model

import "gorm.io/gorm"

type Reply struct {
	gorm.Model
	Id        uint
	Content   string
	UserID    uint
	CommentID uint
}
