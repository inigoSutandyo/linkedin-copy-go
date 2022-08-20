package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	models "github.com/inigoSutandyo/linkedin-copy-go/model"
	"github.com/inigoSutandyo/linkedin-copy-go/utils"
)

func GetUser(c *gin.Context) {

	var user models.User
	status, token, err := CheckAuth(c)
	if status == false {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status":  status,
			"message": err.Error(),
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	fmt.Println("ID = " + claims.Issuer)
	utils.DB.First(&user, "id = ?", claims.Issuer)
	fmt.Print("USER = ")
	fmt.Println(user)

	message := "success"

	c.JSON(http.StatusOK, gin.H{
		"user":    user,
		"message": message,
	})

}

func UpdateProfile(c *gin.Context) {
	var user models.User

	status, token, _ := CheckAuth(c)

	if status == false {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status":  status,
			"message": "Error Unauthorized",
		})
	}
	claims := token.Claims.(*jwt.StandardClaims)
	utils.DB.First(&user, "id = ?", claims.Issuer)

	var updateUser models.User
	bindErr := c.BindJSON(&updateUser)

	if bindErr != nil {
		c.Error(bindErr)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": bindErr.Error(),
		})
	}

	if updateUser.ID != user.ID {
		// c.Error(bindErr)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "User not found",
		})
	}
	utils.DB.Model(&user).Omit("id, password").Updates(updateUser)
	fmt.Print("USER = ")
	fmt.Println(user.ID)
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"user": user,
	})
}
