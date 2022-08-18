package model

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Id         uint   `json:"id"`
	Title      string `json:"title"`
	Content    string
	DateCreate time.Time
	Attachment string
	UserID     uint
	LikeUsers  []*User `gorm:"many2many:user_likeposts"`
}
