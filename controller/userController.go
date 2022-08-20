package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	models "github.com/inigoSutandyo/linkedin-copy-go/model"
	"github.com/inigoSutandyo/linkedin-copy-go/utils"
)

func getUserID(c *gin.Context) string {
	status, token, err := CheckAuth(c)
	if status == false {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status":  status,
			"message": err.Error(),
		})
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
	utils.DB.First(&user, "id = ?", id)
	message := "success"

	c.JSON(http.StatusOK, gin.H{
		"user":    user,
		"message": message,
	})

}

func UpdateProfile(c *gin.Context) {
	var user models.User

	id := getUserID(c)
	utils.DB.First(&user, "id = ?", id)

	var updateUser models.User
	bindErr := c.BindJSON(&updateUser)

	if bindErr != nil {
		c.Error(bindErr)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": bindErr.Error(),
		})
	}

	utils.DB.Model(&user).Omit("id, password").Updates(updateUser)
	fmt.Print("USER = ")
	fmt.Println(user.ID)
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"user":    user,
	})
}
