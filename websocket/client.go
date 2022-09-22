package ws

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	ID     string
	Conn   *websocket.Conn
	Pool   *Pool
	ChatID uint
}

type Message struct {
	Type int    `json:"type"`
	Body string `json:"body"`
}

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {

		messageType, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		broadcast := Broadcast{
			ChatID:  c.ChatID,
			Content: Message{Type: messageType, Body: string(p)},
		}
		c.Pool.Broadcast <- broadcast
		fmt.Printf("Message Received: %+v\n", broadcast.Content)
	}
}
