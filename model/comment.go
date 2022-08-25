package model

import (
	"github.com/inigoSutandyo/linkedin-copy-go/utils"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Content string  `json:"content"`
	Likes   int     `json:"likes"`
	Replies []Reply `json:"replies"`
	UserID  uint    `json:"userid"`
	User    User    `json:"user"`
	PostID  uint    `json:"postid"`
	Post    Post    `json:"post"`
}

func CreateComment(user *User, post *Post, comment *Comment) error {
	comment.UserID = user.ID
	comment.User = *user
	err := utils.DB.Model(post).Association("Comments").Append(comment)
	return err
}

func GetCommentByPost(id string, comments *[]Comment) error {
	// err := utils.DB.Where("post_id = ?", id).Find(comments).Error
	err := utils.DB.Joins("User").Joins("Post").Find(comments, "comments.post_id = ?", id).Error
	return err
}

func GetCommentById(id uint) (Comment, error) {
	var comment Comment
	err := utils.DB.First(&comment, id).Error
	return comment, err
}
