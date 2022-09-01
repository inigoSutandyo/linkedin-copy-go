package server

import (
	"github.com/gin-gonic/gin"
	"github.com/inigoSutandyo/linkedin-copy-go/controller"
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
		api.GET("/user/otherprofile", controllers.GetOtherUser)
		api.POST("/user/profile/update", controllers.UpdateProfile)
		api.POST("/user/profile/image", controllers.UploadProfilePicture)
		api.POST("/user/invite", controller.InviteUser)
		api.POST("/user/invite/accept", controller.AcceptInvite)
		api.POST("/user/invite/ignore", controller.IgnoreInvite)
		api.GET("/user/connection", controller.UserConnections)

		api.GET("/user/invitations", controller.GetAllInvitations)

		api.GET("/home/post", controllers.GetPosts)
		api.POST("/home/post/add", controllers.AddPost)
		api.POST("/home/post/file", controllers.UploadFilePost)
		api.POST("/home/post/like", controllers.AddLikePost)
		api.POST("/home/post/dislike", controllers.RemoveLikePost)

		api.GET("/home/post/comment", controllers.GetComments)
		api.POST("/home/post/comment/add", controllers.AddComment)

		api.POST("/home/post/comment/reply/add", controllers.AddReply)
		api.GET("/home/post/comment/reply", controllers.GetReplies)

		api.GET("/search", controllers.Search)

	}
}
