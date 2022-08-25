package model

import (
	"github.com/inigoSutandyo/linkedin-copy-go/utils"
	"gorm.io/gorm"
)

type Reply struct {
	gorm.Model
	Content   string  `json:"content"`
	UserID    uint    `json:"userid"`
	User      User    `json:"-"`
	CommentID uint    `json:"commentid"`
	Comment   Comment `json:"-"`
}

func CreateReply(user *User, comment *Comment, reply *Reply) error {
	reply.UserID = user.ID
	reply.User = *user
	err := utils.DB.Model(comment).Association("Replies").Append(reply)
	err2 := utils.DB.Model(user).Association("Replies").Append(reply)

	if err != nil {
		return err
	}
	if err2 != nil {
		return err2
	}

	return nil
}

func GetRepliesForComments(comment *Comment, replies *[]Reply) {

	utils.DB.Model(comment).Association("Replies").Find(replies)
}
