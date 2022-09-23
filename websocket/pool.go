package ws

import (
	"fmt"

	"github.com/inigoSutandyo/linkedin-copy-go/model"
)

type Broadcast struct {
	ChatID  uint
	Message model.Message
}

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan Broadcast
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Broadcast),
	}
}

func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
			fmt.Println("Size of Connection Pool: ", len(pool.Clients))
			for client := range pool.Clients {
				fmt.Println(client)
				// client.Conn.WriteJSON(Message{Type: 1, Body: "New User Joined..."})
			}
			break
		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			fmt.Println("Size of Connection Pool: ", len(pool.Clients))
			for client := range pool.Clients {
				fmt.Print("Disconnected : ")
				fmt.Println(client)
				// client.Conn.WriteJSON(Message{Type: 1, Body: "User Disconnected..."})
			}
			break

		case broadcast := <-pool.Broadcast:
			fmt.Println("Sending message to all clients in Pool")
			send := make(map[uint]int)
			for client := range pool.Clients {
				if client.ChatID != broadcast.ChatID {
					continue
				}
				flag := send[client.UserID]

				if flag == 1 {
					continue
				}
				send[client.UserID] = 1
				fmt.Println(client.UserID)
				if err := client.Conn.WriteJSON(broadcast.Message); err != nil {
					fmt.Println(err)
					return
				}
			}
		}
	}
}
