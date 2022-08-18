package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Id        uint
	Content   string
	Likes     []User
	Replies   []Reply
	UserID    uint
	LikeUsers []*User `gorm:"many2many:user_likedcomments"`
}
