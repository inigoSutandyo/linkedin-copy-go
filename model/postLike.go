package model

import (
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
	postLike.UserID = user.ID
	postLike.PostID = post.ID
	err := utils.DB.Model(post).Association("PostLikes").Append(&postLike)
	err2 := utils.DB.Model(user).Association("PostLikes").Append(&postLike)
	getPostLikeCount(post)

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
	// fmt.Println(postLike)
	return postLike, err
}

func DeleteLikedPostData(userId string, postId string) {
	utils.DB.Where("user_id = ? AND post_id = ?", userId, postId).Delete(&PostLike{})
}
