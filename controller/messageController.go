package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/inigoSutandyo/linkedin-copy-go/model"
	websocket "github.com/inigoSutandyo/linkedin-copy-go/websocket"
	ws "github.com/inigoSutandyo/linkedin-copy-go/websocket"
)

func ServeWebsocket(pool *ws.Pool) gin.HandlerFunc {

	return func(c *gin.Context) {
		id, _ := toUint(c.Query("id"))
		user_id, _ := toUint(c.Query("user"))
		conn, err := websocket.Upgrade(c.Writer, c.Request)

		if err != nil {
			log.Println(err)
			return
		}

		client := &ws.Client{
			Conn:   conn,
			Pool:   pool,
			ChatID: id,
			UserID: user_id,
		}

		pool.Register <- client
		client.Read()
		// go websocket.Writer(conn)
		// websocket.Reader(conn)
	}
}

func GetChatRooms(c *gin.Context) {
	id := getUserID(c)
	chats := model.GetRooms(id)

	c.JSON(200, gin.H{
		"message": "sucesss",
		"chats":   chats,
	})
}

// still error
func CreateNewChat(c *gin.Context) {
	var users []model.User
	id := getUserID(c)
	if id == "" {
		return
	}
	id2 := c.Query("id")

	if id == "" || id2 == "" {
		abortError(c, http.StatusBadRequest, "Bad Request")
		return
	}

	chats := model.GetRooms(id)
	count := 0

	user := model.GetUserById(id)
	user2 := model.GetUserById(id2)
	var tmp model.Chat
	for _, c := range chats {
		if len(c.Users) > 2 {
			continue
		}
		count = 0
		for _, u := range c.Users {
			if u.ID == user.ID || u.ID == user2.ID {
				count = count + 1
				break
			}
		}

		if count >= 2 {
			tmp = c
			break
		}
	}

	if count >= 2 {
		// abortError(c, http.StatusBadRequest, "Chat is created")
		c.JSON(200, gin.H{
			"message": "chat already created..",
			"tmp":     tmp,
		})
		return
	}

	users = append(users, user)
	users = append(users, model.GetUserById(id2))
	// fmt.Println(users)
	chat, err := model.CreateRoom(users)

	if err != nil {
		abortError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
		"chat":    chat,
	})
}

func AddMessage(c *gin.Context) {
	id := getUserID(c)
	var message model.Message
	c.BindJSON(&message)

	model.CreateMessage(message.ChatID, id, &message)

	// if err != nil {
	// 	abortError(c, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	c.JSON(200, gin.H{
		"message":      "success",
		"chat_message": message,
	})
}

func GetMessageByChat(c *gin.Context) {
	chat_id, _ := toUint(c.Query("id"))
	chat := model.GetRoomById(chat_id)

	messages := model.GetMessage(&chat)

	c.JSON(200, gin.H{
		"message":  "success",
		"messages": messages,
	})
}

func SendPost(c *gin.Context) {
	id := getUserID(c)
	post_id := c.Query("post_id")
	user_id := c.Query("user_id")

	message := model.CreateSendPost(id, user_id, post_id)
	c.JSON(200, gin.H{
		"message": "success",
		"data":    message,
	})
}
func SendProfile(c *gin.Context) {
	id := getUserID(c)
	profile_id, _ := toUint(c.Query("profile_id"))
	user_id := c.Query("user_id")

	message := model.CreateSendProfile(id, user_id, profile_id)

	c.JSON(200, gin.H{
		"message": "success",
		"data":    message,
	})

}

func SendImage(c *gin.Context) {
	chat_id, _ := toUint(c.Query("chat_id"))
	id := getUserID(c)
	var message model.Message
	c.BindJSON(&message)
	model.CreateMessage(chat_id, id, &message)

}
