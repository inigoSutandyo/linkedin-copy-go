package model

import (
	"fmt"
	"math"
	"time"

	"github.com/inigoSutandyo/linkedin-copy-go/utils"
	"gorm.io/gorm"
)

type Notification struct {
	gorm.Model
	UserID    uint     `json:"userid"`
	User      User     `json:"user"`
	FromID    *uint    `json:"fromid"`
	From      *User    `json:"from"`
	Message   string   `json:"message"`
	HasSource bool     `json:"hassource"`
	PostID    *uint    `json:"postid"`
	Post      *Post    `json:"post"`
	CommentID *uint    `json:"commentid"`
	Comment   *Comment `json:"comment"`
}

func CreateNotificationForPost(user *User, from *User, notification *Notification, post *Post) {
	if user.ID == from.ID {
		return
	}

	notification.UserID = user.ID
	notification.User = *user

	notification.FromID = &from.ID
	notification.From = from

	notification.PostID = &post.ID
	notification.Post = post

	flag := isSameNotif(notification, user.ID, from.ID, "post")
	fmt.Println(flag)
	if flag {
		return
	}
	utils.DB.Create(notification)
}

func CreateNotificationForComment(user *User, from *User, notification *Notification, comment *Comment) {
	if user.ID == from.ID {
		return
	}

	notification.UserID = user.ID
	notification.User = *user

	notification.FromID = &from.ID
	notification.From = from

	notification.CommentID = &comment.ID
	notification.Comment = comment
	flag := isSameNotif(notification, user.ID, from.ID, "comment")

	if flag {
		return
	}
	utils.DB.Create(notification)
}

func CreateNotification(user *User, notification *Notification) {
	notification.UserID = user.ID
	notification.User = *user
	utils.DB.Create(notification)
}

func isSameNotif(notification *Notification, userId uint, fromId uint, notifType string) bool {
	if notifType == "" {
		return false
	}

	var notif Notification
	if notifType == "comment" {
		err := utils.DB.Where("user_id = ? AND from_id = ? AND comment_id = ?", userId, fromId, notification.CommentID).Find(&notif).Error
		if err != nil || notif.ID < 1 {
			return false
		}
		currTime := time.Now()
		timeDiff := math.Abs(notif.CreatedAt.Sub(currTime).Seconds())

		if timeDiff <= 15 {
			return true
		}
	} else if notifType == "post" {
		err := utils.DB.Where("user_id = ? AND from_id = ? AND post_id = ?", userId, fromId, notification.PostID).Find(&notif).Error
		if err != nil || notif.ID < 1 {
			return false
		}
		currTime := time.Now()
		timeDiff := math.Abs(notif.CreatedAt.Sub(currTime).Seconds())

		if timeDiff <= 15 {
			return true
		}
	}
	return false
}
