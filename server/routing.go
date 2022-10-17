package server

import (
	"github.com/gin-gonic/gin"
	"github.com/inigoSutandyo/linkedin-copy-go/controller"
	ws "github.com/inigoSutandyo/linkedin-copy-go/websocket"
)

func Routes(router *gin.Engine) {

	pool := ws.NewPool()

	go pool.Start()

	api := router.Group("/api")
	{
		api.GET("/user/all", controller.FindUsers)

		api.GET("/auth/isauth", controller.ClientAuth)
		api.POST("/auth/register", controller.Register)
		api.POST("/auth/login", controller.Login)
		api.POST("/auth/google/login", controller.GoogleLogin)
		api.POST("/auth/logout", controller.Logout)
		api.POST("/auth/forget", controller.ForgetRequest)
		api.POST("/auth/reset", controller.ResetPassword)
		api.POST("/auth/verify", controller.VerifyUser)

		api.GET("/user/experiences", controller.FetchUserExperience)
		api.GET("/user/profile", controller.GetUser)
		api.GET("/user/otherprofile", controller.GetOtherUser)
		api.POST("/user/profile/update", controller.UpdateProfile)
		api.POST("/user/profile/image", controller.UploadProfilePicture)

		api.GET("/user/email", controller.FindUserByEmail)
		api.GET("/user/invitations", controller.GetAllInvitations)
		api.POST("/user/invite", controller.InviteUser)
		api.POST("/user/invite/accept", controller.AcceptInvite)
		api.POST("/user/invite/ignore", controller.IgnoreInvite)

		api.GET("/user/connection", controller.UserConnections)
		api.DELETE("/user/connection/remove", controller.RemoveConnection)

		api.POST("/user/educations/add", controller.AddEducation)
		api.POST("/user/experiences/add", controller.AddExperience)

		api.GET("/home/post", controller.GetPosts)
		api.POST("/home/post/add", controller.AddPost)
		api.POST("/home/post/file", controller.UploadFilePost)
		api.POST("/home/post/like", controller.AddLikePost)
		api.POST("/home/post/dislike", controller.RemoveLikePost)
		api.DELETE("/home/post/remove", controller.RemovePost)

		api.GET("/home/post/comment", controller.GetComments)
		api.GET("/home/post/comment/count", controller.GetCommentCount)

		api.POST("/home/post/comment/add", controller.AddComment)

		api.POST("/home/post/comment/reply/add", controller.AddReply)
		api.GET("/home/post/comment/reply", controller.GetReplies)

		api.GET("/search/post", controller.SearchPostHandler)
		api.GET("/search/user", controller.SearchUserHandler)

		api.GET("/notifications", controller.UserNotifications)
		api.DELETE("/notifications/remove", controller.RemoveNotification)

		api.GET("/jobs", controller.GetAllJobs)
		api.POST("/jobs/add", controller.AddJob)

		api.GET("/websocket", controller.ServeWebsocket(pool))
		api.GET("/chats", controller.GetChatRooms)
		api.POST("/chats/create", controller.CreateNewChat)
		api.POST("/message/add", controller.AddMessage)
		api.GET("/message", controller.GetMessageByChat)

		api.GET("/recommendation", controller.FindRecommendation)

		api.POST("/user/follow", controller.FollowUser)
		api.DELETE("/user/unfollow", controller.UnfollowUser)

		api.POST("/send/post", controller.SendPost)
		api.POST("/send/profile", controller.SendProfile)
		api.GET("/get/post", controller.GetSinglePost)

		api.POST("/chat/image", controller.SendImage)
	}
}
