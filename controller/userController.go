package controller

import (
	"fmt"
	"math/big"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	models "github.com/inigoSutandyo/linkedin-copy-go/model"
	"github.com/inigoSutandyo/linkedin-copy-go/utils"
	"golang.org/x/crypto/bcrypt"
)

func GetUserByIdHandler(c *gin.Context) {
	id := new(big.Int)
	_, err := fmt.Sscan(c.Param("id"), id)
	if err != nil {
		fmt.Println("error scanning value:", err)
	} else {
		fmt.Println(id)
	}

	user := models.GetUserById(*id)
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"message": "Some handler for my beautiful api",
		"user":    user,
	})
}

func GetAllUsersHandler(c *gin.Context) {
	users := models.GetAllUsers()
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"users":   users,
	})
}

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
		c.JSON(http.StatusOK, gin.H{
			"message": message,
			"error":   err.Error(),
			"isError": true,
		})
	} else {
		c.Header("Content-Type", "application/json")
		claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
			Issuer:    strconv.Itoa(int(user.Id)),
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		})

		token, tokenErr := claims.SignedString([]byte(utils.GetEnv("SECRET_KEY")))

		if tokenErr != nil {
			message = "Could not sign in (SERVER ERROR)"
			c.JSON(http.StatusOK, gin.H{
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
}

func RegisterUserHandler(c *gin.Context) {
	message := "Sucesful"

	email := c.PostForm("email")
	password := c.PostForm("password")

	fmt.Println(email)

	user, err := models.CreateUser(email, password)
	c.Header("Content-Type", "application/json")
	if err != nil {
		message = err.Error()
	}

	c.JSON(http.StatusOK, gin.H{
		"message": message,
		"user":    user,
	})
}
