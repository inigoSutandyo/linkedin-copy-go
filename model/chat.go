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
	PostID  *uint  `json:"postid"`
	Post    *Post  `json:"post"`
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

func CreateRoom(users []User) (Chat, error) {
	var chat Chat

	err := utils.DB.Create(&chat).Error
	err = utils.DB.Model(&chat).Association("Users").Append(users)
	// for _, user := range users {
	// 	err = utils.DB.Model(&user).Association("Chats").Append(chat)
	// 	if err != nil {
	// 		return chat, err
	// 	}
	// }
	return chat, err
}

func GetRoomById(id uint) Chat {
	var chat Chat
	utils.DB.Find(&chat, "id = ?", id)
	return chat
}

func CreateMessage(chat_id uint, user_id string, message *Message) Message {
	user := GetUserById(user_id)
	chat := GetRoomById(chat_id)

	message.User = user
	message.UserID = user.ID
	message.Chat = chat
	utils.DB.Model(&chat).Association("Messages").Append(message)
	return *message
}

func CreateSendPost(user_id string, dest_id string, post_id string) Message {
	chats := GetRooms(user_id)
	user := GetUserById(dest_id)
	src := GetUserById(user_id)
	var chat Chat
	found := false
	for _, c := range chats {
		if len(c.Users) > 2 {
			continue
		}
		for _, u := range c.Users {
			if u.ID == user.ID {
				chat = c
				found = true
				break
			}
		}
	}

	if found == false {
		users := []User{src, user}
		chat, _ = CreateRoom(users)
	}

	post, _ := GetPostByIDString(post_id)

	var message Message = Message{
		Content: "Sent Post",
		Post:    &post,
		PostID:  &post.ID,
	}

	utils.DB.Model(&post).Update("send_count", post.SendCount+1)
	return CreateMessage(chat.ID, user_id, &message)
}

func GetMessage(chat *Chat) []Message {
	var messages []Message
	utils.DB.Preload("User").Preload("Post").Model(chat).Order("chats.created_at desc").Association("Messages").Find(&messages)
	return messages
}
