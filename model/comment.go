package model

import (
	"github.com/inigoSutandyo/linkedin-copy-go/utils"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Content   string `json:"content"`
	Likes     int    `json:"likes"`
	ParentID  *uint
	Replies   []Comment `json:"replies" gorm:"foreignKey:ParentID"`
	UserID    uint      `json:"userid"`
	User      User      `json:"user"`
	PostID    uint      `json:"postid"`
	Post      Post      `json:"post"`
	IsReply   bool      `json:"isreply"`
	MentionID *uint     `json:"mentionid"`
	Mention   *User     `json:"mention"`
}

func CreateComment(user *User, post *Post, comment *Comment) error {
	comment.UserID = user.ID
	comment.User = *user
	err := utils.DB.Preload("Mention").Model(post).Association("Comments").Append(comment)
	return err
}

func GetCommentByPost(id string, comments *[]Comment) error {
	err := utils.DB.Joins("User").Joins("Post").Find(comments, "comments.post_id = ? AND comments.is_reply = false", id).Error
	return err
}

func GetCommentById(id string) (Comment, error) {
	var comment Comment
	err := utils.DB.First(&comment, id).Error
	return comment, err
}

func GetRepliesForComments(comment *Comment, replies *[]Comment) {
	utils.DB.Preload("Mention").Preload("User").Model(comment).Association("Replies").Find(replies, "comments.is_reply = true")
}

func CreateReply(user *User, comment *Comment, reply *Comment, mentionId uint) error {
	reply.UserID = user.ID
	reply.User = *user
	post, err := GetPostByID(comment.PostID)
	err = utils.DB.Model(&post).Association("Comments").Append(reply)

	if mentionId != 0 {
		var mention User
		utils.DB.First(&mention, mentionId)
		utils.DB.Model(&mention).Association("Mentions").Append(reply)
	}

	if err == nil {
		return utils.DB.Model(comment).Association("Replies").Append(reply)
	}
	return err
}
