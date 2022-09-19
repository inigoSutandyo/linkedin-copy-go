package controller

import (
	"log"

	"github.com/gin-gonic/gin"
	websocket "github.com/inigoSutandyo/linkedin-copy-go/websocket"
	ws "github.com/inigoSutandyo/linkedin-copy-go/websocket"
)

func ServeWebsocket(pool *ws.Pool) gin.HandlerFunc {

	return func(c *gin.Context) {

		conn, err := websocket.Upgrade(c.Writer, c.Request)

		if err != nil {
			log.Println(err)
			return
		}

		client := &ws.Client{
			Conn: conn,
			Pool: pool,
		}

		pool.Register <- client
		client.Read()
	}
}
