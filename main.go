package main

import (
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	controllers "github.com/inigoSutandyo/linkedin-copy-go/controller"
	utils "github.com/inigoSutandyo/linkedin-copy-go/utils"
)

func main() {
	utils.Connect() // connect to DB
	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	api := router.Group("/api")
	{
		api.GET("/", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "test",
			})
		})
	}

	// api.GET("/user", controllers.GetUserByIdHandler(3))
	api.POST("/register", controllers.RegisterUserHandler)
	api.GET("/users", controllers.GetAllUsersHandler)
	router.Run(":8080")
}
