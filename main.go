package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	controllers "github.com/inigoSutandyo/linkedin-copy-go/controller"
	utils "github.com/inigoSutandyo/linkedin-copy-go/utils"
)

func CORS(c *gin.Context) {

	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Content-Type", "application/json")

	if c.Request.Method != "OPTIONS" {

		c.Next()

	} else {
		c.AbortWithStatus(http.StatusOK)
	}
}

func main() {
	utils.Connect() // connect to DB

	router := gin.Default()

	router.Use(CORS)

	api := router.Group("/api")
	{

		api.POST("/register", controllers.RegisterUserHandler)
		api.POST("/login", controllers.LoginUserHandler)
		api.POST("/logout", controllers.LogoutHandler)
		api.GET("/auth", controllers.GetAuth)
		api.GET("/users", controllers.GetAllUsersHandler)
	}

	router.Run(":8080")
}
