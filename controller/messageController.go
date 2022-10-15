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
	id2 := c.Query("id")

	if id == "" || id2 == "" {
		abortError(c, http.StatusBadRequest, "Bad Request")
		return
	}

	users = append(users, model.GetUserById(id))
	users = append(users, model.GetUserById(id2))
	// fmt.Println(users)
	_, err := model.CreateRoom(users)

	if err != nil {
		abortError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
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

	model.CreateSendPost(id, user_id, post_id)
	
}
