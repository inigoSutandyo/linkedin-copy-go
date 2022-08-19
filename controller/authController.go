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

type AuthData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	// c.Header("Access-Control-Allow-Origin", "*")
	// c.Header("Access-Control-Allow-Headers", "Content-Type")
	// c.Header("Content-Type", "application/json")
	// c.Header("Access-Control-Allow-Credentials", "true")

	message := "success"

	var data AuthData
	c.BindJSON(&data)
	email := data.Email
	password := data.Password

	var user models.User
	user = models.GetUserByEmail(email)
	err := bcrypt.CompareHashAndPassword(user.Password, []byte(password))

	if err != nil {

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
		c.SetCookie("token", token, 3600*12, "/", "http://127.0.0.1", true, true)
		c.SetCookie("auth", "true", 3600*12, "/", "http://127.0.0.1", false, false)

		c.SetCookie("token", token, 3600*12, "/", "http://localhost", true, true)
		c.SetCookie("auth", "true", 3600*12, "/", "http://localhost", false, false)

		c.JSON(http.StatusOK, gin.H{
			"message": message,
			"isError": false,
		})
	}

}

func Register(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	message := "success"

	var data AuthData
	bindErr := c.BindJSON(&data)
	email := data.Email
	password := data.Password

	if email == "" || password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error!",
		})
		return
	}

	if bindErr != nil {
		message = bindErr.Error()
		c.JSON(http.StatusBadRequest, gin.H{
			"message": message,
		})
		return
	}

	pw, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	_, err := models.CreateUser(email, pw)

	if err != nil {
		message = err.Error()
	}
	fmt.Println(data)
	fmt.Println(message)

	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}

func Logout(c *gin.Context) {
	message := "success"
	// c.Cookie("token")
	c.SetCookie("token", "deleting", -1, "/", "http://localhost", false, true)
	c.SetCookie("token", "deleting", -1, "/", "http://127.0.0.1", false, true)

	c.SetCookie("auth", "deleting", -1, "/", "http://localhost", false, true)
	c.SetCookie("auth", "deleting", -1, "/", "http://127.0.0.1", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}

func CheckAuth(c *gin.Context) (bool, *jwt.Token, error) {
	cookie, err := c.Cookie("token")

	if err != nil {
		return false, nil, err
	}

	token, tokenErr := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(utils.GetEnv("SECRET_KEY")), nil
	})

	if tokenErr != nil {
		return false, nil, tokenErr
	}
	return true, token, nil
}

func ClientAuth(c *gin.Context) {
	status, _, err := CheckAuth(c)

	if status {
		c.JSON(http.StatusOK, gin.H{
			"status":  status,
			"message": "Client is authorized",
		})
	} else {
		c.Error(err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status":  status,
			"message": "Client is unaouthorized",
		})
	}
}
