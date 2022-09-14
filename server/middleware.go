package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CORS(c *gin.Context) {
	origin := c.Request.Header.Get("Origin")
	// fmt.Println("host = " + origin)
	// if origin == "http://127.0.0.1:5173" || origin == "http://localhost:5173" {
	// }
	c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, x-xsrf-token")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

	if c.Request.Method != "OPTIONS" {

		c.Next()

	} else {
		fmt.Println("OK")
		c.AbortWithStatus(http.StatusOK)
	}
}
