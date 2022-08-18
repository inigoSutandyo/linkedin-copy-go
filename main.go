package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	controllers "github.com/inigoSutandyo/linkedin-copy-go/controller"
	utils "github.com/inigoSutandyo/linkedin-copy-go/utils"
)

func CORS(c *gin.Context) {

	// c.Header("Access-Control-Allow-Origin", "http://localhost:5173")
	// c.Header("Access-Control-Allow-Methods", "*")
	// c.Header("Access-Control-Allow-Headers", "*")
	// c.Header("Access-Control-Allow-Credentials", "true")
	// c.Header("Content-Type", "application/json")
	origin := c.Request.Header.Get("Origin")
	fmt.Println("host = " + origin)
	if origin == "http://127.0.0.1:5173" || origin == "http://localhost:5173" {
		c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
	}
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, x-xsrf-token")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

	if c.Request.Method != "OPTIONS" {

		c.Next()

	} else {
		fmt.Println("OK")
		c.AbortWithStatus(http.StatusOK)
	}
}

func main() {
	utils.Connect() // connect to DB

	router := gin.Default()

	router.Use(CORS)

	api := router.Group("/api")
	{

		api.GET("/auth/isauth", controllers.CheckAuth)
		api.POST("/auth/register", controllers.Register)
		api.POST("/auth/login", controllers.Login)
		api.POST("/auth/logout", controllers.Logout)

		api.GET("/user/profile", controllers.GetUser)
	}

	router.Run(":8080")
}
