package main

import (
	"github.com/gin-gonic/gin"
	server "github.com/inigoSutandyo/linkedin-copy-go/server"
	utils "github.com/inigoSutandyo/linkedin-copy-go/utils"
)

func main() {
	utils.Connect() // connect to DB

	router := gin.Default()

	router.Use(server.CORS)
	server.Routes(router)
	router.Run(":8080")
}
