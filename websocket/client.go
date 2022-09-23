package ws

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gorilla/websocket"
	"github.com/inigoSutandyo/linkedin-copy-go/model"
)

type Client struct {
	ID     string
	Conn   *websocket.Conn
	Pool   *Pool
	ChatID uint
	UserID uint
}

// type Message struct {
// 	Type int    `json:"type"`
// 	Body string `json:"body"`
// }

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {

		_, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		user_id := strconv.FormatUint(uint64(c.UserID), 10)

		broadcast := Broadcast{
			ChatID:  c.ChatID,
			Message: model.CreateMessage(c.ChatID, user_id, &model.Message{Content: string(p)}),
		}
		c.Pool.Broadcast <- broadcast
		fmt.Printf("Message Received: %+v\n", broadcast.Message.Content)
	}
}
