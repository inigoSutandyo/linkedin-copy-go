package model

import (
	"github.com/inigoSutandyo/linkedin-copy-go/utils"
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	Content string `json:"content"`
	User    User   `json:"user"`
	UserID  uint   `json:"-"`
	Chat    Chat   `json:"chat"`
	ChatID  uint   `json:"chatid"`
}

type Chat struct {
	gorm.Model
	Users    []*User   `json:"users" gorm:"many2many:user_chats;"`
	Messages []Message `json:"messages"`
}

func GetRooms(id string) []Chat {
	var chats []Chat
	user := GetUserById(id)
	utils.DB.Preload("Users").Model(&user).Association("Chats").Find(&chats)
	return chats
}

func CreateRoom(users []User) error {
	var chat Chat

	err := utils.DB.Create(&chat).Error
	err = utils.DB.Model(&chat).Association("Users").Append(&users)
	for u := range users {
		err = utils.DB.Model(&u).Association("Chats").Append(chat)
		if err != nil {
			return err
		}
	}
	return err
}

func GetRoomById(id uint) Chat {
	var chat Chat
	utils.DB.Find(&chat, "id = ?", id)
	return chat
}

func CreateMessage(chat *Chat, user *User, message *Message) error {
	message.User = *user
	message.UserID = user.ID
	message.Chat = *chat
	return utils.DB.Model(chat).Association("Messages").Append(message)
}

func GetMessage(chat *Chat) []Message {
	var messages []Message
	utils.DB.Preload("User").Model(chat).Association("Messages").Find(&messages)
	return messages
}
