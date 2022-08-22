package model

import (
	"fmt"

	"github.com/inigoSutandyo/linkedin-copy-go/utils"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title      string `json:"-"`
	Content    string `json:"content" gorm:"text"`
	Attachment string `json:"attachment"`
	Likes      int    `json:"like"`
	UserID     uint
	// Template   Template `gorm:"embedded"`s
	// PostLikes  []PostLike
}

func CreatePost(user *User, post *Post) error {
	fmt.Println(user.Email)
	fmt.Println(post.Content)
	err := utils.DB.Model(user).Association("Posts").Append(post)

	return err
}
