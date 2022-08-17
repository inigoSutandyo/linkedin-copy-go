package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	models "github.com/inigoSutandyo/linkedin-copy-go/model"
	"github.com/inigoSutandyo/linkedin-copy-go/utils"
	"golang.org/x/crypto/bcrypt"
)

func LoginUserHandler(c *gin.Context) {
	message := "Sucessful"

	email := c.PostForm("email")
	password := c.PostForm("password")
	var user models.User
	user = models.GetUserByEmail(email)
	err := bcrypt.CompareHashAndPassword(user.Password, []byte(password))

	if err != nil {
		c.Status(404)
		message = "User not found"
		c.JSON(http.StatusBadRequest, gin.H{
			"message": message,
			"error":   err.Error(),
			"isError": true,
		})
		return
	}

	c.Header("Content-Type", "application/json")
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, tokenErr := claims.SignedString([]byte(utils.GetEnv("SECRET_KEY")))

	if tokenErr != nil {
		message = "Could not sign in (SERVER ERROR)"
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": message,
			"error":   tokenErr.Error(),
			"isError": true,
		})
	} else {
		c.SetCookie("token", token, 3600*6, "/", "127.0.0.1:8080", false, true)
		c.JSON(http.StatusOK, gin.H{
			"message": message,
			"user":    user,
			"isError": false,
		})
	}

}

func RegisterUserHandler(c *gin.Context) {
	message := "Sucesful"

	email := c.PostForm("email")
	password := c.PostForm("password")

	fmt.Println(email)
	pw, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	user, err := models.CreateUser(email, pw)
	c.Header("Content-Type", "application/json")
	if err != nil {
		message = err.Error()
	}

	c.JSON(http.StatusOK, gin.H{
		"message": message,
		"user":    user,
	})
}

func GetAuth(c *gin.Context) {
	message := "Sucessful"
	cookie, err := c.Cookie("token")

	token, tokenErr := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(utils.GetEnv("SECRET_KEY")), nil
	})
	if err != nil {
		message = "Unauthorized"
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

func LogoutHandler(c *gin.Context) {
	message := "Sucessful"
	c.SetCookie("token", "deleting", -1, "/", "127.0.0.1:8080", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}
