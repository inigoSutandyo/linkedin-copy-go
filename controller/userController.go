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
	posts := models.GetUserPost(&user)
	likedPost, err := models.GetLikedPostData(&user)
	if err != nil {
		abortError(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, gin.H{
		"user":       user,
		"posts":      posts,
		"message":    message,
		"likedposts": likedPost,
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
