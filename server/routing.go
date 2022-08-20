package server

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/inigoSutandyo/linkedin-copy-go/controller"
)

func Routes(router *gin.Engine) {
	api := router.Group("/api")
	{

		api.GET("/auth/isauth", controllers.ClientAuth)
		api.POST("/auth/register", controllers.Register)
		api.POST("/auth/login", controllers.Login)
		api.POST("/auth/logout", controllers.Logout)

		api.GET("/user/profile", controllers.GetUser)
		api.POST("/user/profile/update", controllers.UpdateProfile)
	}
}