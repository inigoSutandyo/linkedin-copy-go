package controller

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	models "github.com/inigoSutandyo/linkedin-copy-go/model"
)

func getUserID(c *gin.Context) string {
	status, token, err := CheckAuth(c)
	if status == false || err != nil {
		abortError(c, http.StatusUnauthorized, err.Error())
		// return "", err
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
	// likedPost, err := models.GetLikedPostData(&user)

	// var postIds []uint

	// for _, liked := range likedPost {
	// 	postIds = append(postIds, liked.PostID)
	// }

	// if err != nil {
	// 	abortError(c, http.StatusInternalServerError, err.Error())
	// }
	s := http.DetectContentType(user.Image)
	models.SaveImageMime(&user)
	c.JSON(http.StatusOK, gin.H{
		"user":       user,
		"image_type": s,
		"message":    message,
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
	user := models.GetUserById(id)

	image, _, err := c.Request.FormFile("picture")

	if err != nil {
		abortError(c, http.StatusBadRequest, err.Error())
	}

	buf := bytes.NewBuffer(nil)
	_, err2 := io.Copy(buf, image)
	if err2 != nil {
		abortError(c, http.StatusInternalServerError, err2.Error())
	}
	err3 := models.UploadImageUser(&user, buf)
	if err3 != nil {
		abortError(c, http.StatusInternalServerError, err3.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    "success",
		"user":       user,
		"image_type": user.ImageMime,
	})
}
