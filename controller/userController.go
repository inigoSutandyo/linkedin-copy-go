package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/inigoSutandyo/linkedin-copy-go/model"
)

func getUserID(c *gin.Context) string {
	status, token, err := CheckAuth(c)
	if status == false || err != nil {
		abortError(c, http.StatusUnauthorized, err.Error())
		return ""
	}

	claims := token.Claims.(*jwt.StandardClaims)
	c.Next()
	return claims.Issuer

}

func FindUserByEmail(c *gin.Context) {
	email := c.Query("email")
	user := model.GetUserByEmail(email)

	c.JSON(200, gin.H{
		"message": "success",
		"user":    user,
	})
}

func GetUser(c *gin.Context) {
	var user model.User
	var id string
	id = getUserID(c)
	user = model.GetUserById(id)
	model.GetConnection(&user)
	model.GetFollowing(&user)
	model.GetInvitations(&user)
	model.GetEducations(&user)
	model.GetExperiences(&user)

	likedPost, err := model.GetLikedPostData(&user)

	var postIds []uint

	for _, liked := range likedPost {
		postIds = append(postIds, liked.PostID)
	}

	if err != nil {
		abortError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user":       user,
		"message":    "success",
		"likedposts": postIds,
	})

}

func GetOtherUser(c *gin.Context) {
	id, _ := c.GetQuery("id")
	if id == "" {
		abortError(c, http.StatusBadRequest, "Unknown ID")
		return
	}
	user := model.GetUserById(id)

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"user":    user,
	})
}

func UpdateProfile(c *gin.Context) {
	var user model.User

	id := getUserID(c)
	user = model.GetUserById(id)

	var updateUser model.User
	bindErr := c.BindJSON(&updateUser)

	if bindErr != nil {
		c.Error(bindErr)
		abortError(c, http.StatusBadRequest, bindErr.Error())
		return
	}

	model.UpdateUser(&user, "password, email, id", updateUser)

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"user":    user,
	})
}

func UploadProfilePicture(c *gin.Context) {
	id := getUserID(c)
	if id == "" {
		abortError(c, http.StatusBadRequest, "Not Authorized")
		return
	}
	user := model.GetUserById(id)
	url := c.Query("url")
	publicid := c.Query("publicid")

	err := model.UploadImageUser(&user, url, publicid)

	if err != nil {
		abortError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"user":    user,
	})
}

func ConnectUser(c *gin.Context) {
	userId := c.Query("userId")
	connectId := c.Query("connectId")

	if userId == "" || connectId == "" {
		abortError(c, http.StatusBadRequest, "Data not found")
		return
	}

	user := model.GetUserById(userId)
	connect := model.GetUserById(connectId)

	err := model.CreateConnection(&user, &connect)
	if err != nil {
		abortError(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func UserConnections(c *gin.Context) {
	id := c.Query("id")

	if id == "" {
		abortError(c, http.StatusBadRequest, "Data not found")
		return
	}

	user := model.GetUserById(id)
	err := model.GetConnection(&user)

	if err != nil {
		abortError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"user":    user,
	})
}

func InviteUser(c *gin.Context) {
	var invitation model.Invitation

	c.BindJSON(&invitation)

	source := model.GetUserByIdInt(invitation.SourceID)
	destination := model.GetUserByIdInt(invitation.DestinationID)

	err := model.CreateInvitation(&source, &destination, &invitation)

	if err != nil {
		abortError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    "success",
		"invitation": invitation,
	})
}

func IgnoreInvite(c *gin.Context) {
	sourceId := c.Query("source")
	destinationId := c.Query("destination")

	err := model.DeleteInvitation(sourceId, destinationId)
	if err != nil {
		abortError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func AcceptInvite(c *gin.Context) {
	sourceId := c.Query("source")
	destinationId := c.Query("destination")

	err := model.DeleteInvitation(sourceId, destinationId)
	if err != nil {
		abortError(c, http.StatusInternalServerError, err.Error())
		return
	}
	source := model.GetUserById(sourceId)
	destination := model.GetUserById(destinationId)
	err = model.CreateConnection(&source, &destination)

	if err != nil {
		abortError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func RemoveConnection(c *gin.Context) {
	userId := c.Query("user")
	connectId := c.Query("connect")

	user := model.GetUserById(userId)
	connect := model.GetUserById(connectId)

	err := model.DeleteConnection(&user, &connect)

	if err != nil {
		abortError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func AddEducation(c *gin.Context) {
	var education model.Education
	c.BindJSON(&education)
	id := getUserID(c)
	if id == "" {
		abortError(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	user := model.GetUserById(id)
	model.AppendEducation(&user, &education)

	c.JSON(http.StatusOK, gin.H{
		"message":   "success",
		"education": education,
	})
}

func AddExperience(c *gin.Context) {
	var experience model.Experience
	c.BindJSON(&experience)
	id := getUserID(c)
	if id == "" {
		abortError(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	user := model.GetUserById(id)
	model.AppendExperience(&user, &experience)

	c.JSON(http.StatusOK, gin.H{
		"message":   "success",
		"education": experience,
	})
}

func FindRecommendation(c *gin.Context) {
	id := getUserID(c)
	user := model.GetUserById(id)
	err := model.GetConnection(&user)

	if err != nil {
		abortError(c, http.StatusInternalServerError, err.Error())
		return
	}

	var recommendations []model.User
	id_list := make(map[uint]int)
	id_list[user.ID] = 1

	for _, u := range user.Connections {
		id_list[u.ID] = 1
	}

	for _, u := range user.Connections {
		err := model.GetConnection(u)

		if err != nil {
			abortError(c, http.StatusInternalServerError, err.Error())
			return
		}

		for _, c := range u.Connections {
			if id_list[c.ID] == 0 {
				recommendations = append(recommendations, *c)
			}
			id_list[c.ID] = 1
		}
	}

	c.JSON(200, gin.H{
		"message":         "success",
		"recommendations": recommendations,
	})

}

func FollowUser(c *gin.Context) {
	id := getUserID(c)
	user := model.GetUserById(id)

	follow_id := c.Query("follow")
	follow := model.GetUserById(follow_id)

	err := model.CreateFollowing(&user, &follow)

	if err != nil {
		abortError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
	})
}

func UnfollowUser(c *gin.Context) {
	id := getUserID(c)
	user := model.GetUserById(id)

	follow_id := c.Query("follow")
	follow := model.GetUserById(follow_id)

	err := model.DeleteFollowing(&user, &follow)

	if err != nil {
		abortError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
	})
}
