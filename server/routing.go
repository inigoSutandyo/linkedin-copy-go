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
		api.POST("/user/profile/image", controllers.UploadProfilePicture)

		api.POST("/home/post/add", controllers.AddPost)
		api.GET("/home/post", controllers.GetPosts)
		api.POST("/home/post/like", controllers.AddLikePost)
		api.POST("/home/post/dislike", controllers.RemoveLikePost)

		api.GET("/home/post/comment", controllers.GetComments)
		api.POST("/home/post/comment/add", controllers.AddComment)

		api.POST("/home/post/comment/reply/add", controllers.AddReply)
		api.GET("/home/post/comment/reply", controllers.GetReplies)
	}
}
