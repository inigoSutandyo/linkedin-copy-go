package model

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Id        uint     `json:"id"`
	Content   string   `json:"content"`
	Likes     []User   `json:"likes"`
	Replies   []Reply  `json:"replies"`
	UserID    uint     `json:"userid"`
	LikeUsers []*User  `json:"likeusers" gorm:"many2many:user_likedcomments"`
	Template  Template `gorm:"embedded"`
}
