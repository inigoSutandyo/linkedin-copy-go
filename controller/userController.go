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

	status, token, err := CheckAuth(c)

	if status == false {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status":  status,
			"message": err.Error(),
		})
		return
	}

	message := "success"
	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User
	fmt.Println(claims.Issuer)
	utils.DB.Where("id = ?", claims.Issuer).First(&user)

	c.JSON(http.StatusOK, gin.H{
		"user":    user,
		"message": message,
	})

}
