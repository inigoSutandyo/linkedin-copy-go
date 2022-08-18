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
	message := "success"
	cookie, err := c.Cookie("token")

	token, tokenErr := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(utils.GetEnv("SECRET_KEY")), nil
	})
	if err != nil {
		message = "Not Allowed"
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": message,
			"error":   err.Error(),
			"isError": true,
		})
		return
	} else if tokenErr != nil {
		message = "Could not get user"
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": message,
			"error":   tokenErr.Error(),
			"isError": true,
		})
		return
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User
	fmt.Println(claims.Issuer)
	utils.DB.Where("id = ?", claims.Issuer).First(&user)

	c.JSON(http.StatusOK, gin.H{
		"user":    user,
		"message": message,
	})

}
