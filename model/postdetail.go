package model

import "gorm.io/gorm"

type PostDetail struct {
	gorm.Model
	UserID uint 
	PostID uint
}
