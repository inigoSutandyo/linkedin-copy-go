package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/inigoSutandyo/linkedin-copy-go/model"
	models "github.com/inigoSutandyo/linkedin-copy-go/model"
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

func GetUser(c *gin.Context) {

	var user models.User
	var id string
	id = getUserID(c)

	user = models.GetUserById(id)
	models.GetConnection(&user)
	model.GetInvitations(&user)

	likedPost, err := models.GetLikedPostData(&user)

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
		// "invites":    invites,
	})

}

func GetOtherUser(c *gin.Context) {
	id, _ := c.GetQuery("id")
	if id == "" {
		abortError(c, http.StatusBadRequest, "Unknown ID")
		return
	}
	user := models.GetUserById(id)

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"user":    user,
	})
}

func UpdateProfile(c *gin.Context) {
	var user models.User

	id := getUserID(c)
	user = models.GetUserById(id)

	var updateUser models.User
	bindErr := c.BindJSON(&updateUser)

	if bindErr != nil {
		c.Error(bindErr)
		abortError(c, http.StatusBadRequest, bindErr.Error())
		return
	}

	models.UpdateUser(&user, "password, email, id", updateUser)

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
	user := models.GetUserById(id)
	url := c.Query("url")
	publicid := c.Query("publicid")

	err := models.UploadImageUser(&user, url, publicid)

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

	user := models.GetUserById(userId)
	connect := models.GetUserById(connectId)

	err := models.CreateConnection(&user, &connect)
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

	user := models.GetUserById(id)
	err := models.GetConnection(&user)

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
	sourceId := c.Query("source")
	destinationId := c.Query("destination")
	note := c.Query("note")

	source := models.GetUserById(sourceId)
	destination := models.GetUserById(destinationId)

	invite, err := models.CreateInvitation(&source, &destination, note)

	if err != nil {
		abortError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    "success",
		"invitation": invite,
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
