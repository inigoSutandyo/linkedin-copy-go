package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	models "github.com/inigoSutandyo/linkedin-copy-go/model"
)

func GetUserByIdHandler(id uint) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		user := models.GetUserById(3)
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, gin.H{
			"message": "Some handler for my beautiful api",
			"user":    user,
		})
	}

	return gin.HandlerFunc(fn)
}

func GetAllUsersHandler(c *gin.Context) {
	users := models.GetAllUsers()
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"users":   users,
	})
}

func RegisterUserHandler(c *gin.Context) {
	message := "Succesfull"

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
