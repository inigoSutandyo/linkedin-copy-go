package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
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

	fmt.Println("ID = " + id)
	user = models.GetUserById(id)
	message := "success"
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
		"image_type": nil,
		"message":    message,
		"likedposts": postIds,
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
	fmt.Print("USER = ")
	fmt.Println(user.ID)

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
