package model

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Content string  `json:"content"`
	Likes   int     `json:"likes"`
	Replies []Reply `json:"replies"`
	UserID  uint    `json:"userid"`
	User    User    `json:"user"`
}
