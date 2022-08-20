package model

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Id         uint   `json:"id"`
	Title      string `json:"title"`
	Content    string
	Attachment string
	LikeUsers  []User  
	Template   Template `gorm:"embedded"`
}
