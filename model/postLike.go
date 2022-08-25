package model

import (
	"fmt"

	"github.com/inigoSutandyo/linkedin-copy-go/utils"
	"gorm.io/gorm"
)

type PostLike struct {
	gorm.Model
	UserID uint
	PostID uint
	Post   Post
	User   User
}

func CreatePostLike(user *User, post *Post) error {
	var postLike PostLike
	err := utils.DB.Model(post).Association("Replies").Append(&postLike)
	err2 := utils.DB.Model(user).Association("Replies").Append(&postLike)

	if err != nil {
		return err
	}
	if err2 != nil {
		return err2
	}
	return nil
}

func GetLikedPostData(user *User) ([]PostLike, error) {
	var postLike []PostLike
	err := utils.DB.Model(&user).Association("PostLikes").Find(&postLike)
	fmt.Println(postLike)
	return postLike, err
}
