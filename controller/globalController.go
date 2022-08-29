package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	models "github.com/inigoSutandyo/linkedin-copy-go/model"
)

func Search(c *gin.Context) {

	param := c.Query("search")

	var users []models.User
	var posts []models.Post
	err := models.SearchUserByName(&users, param)
	err2 := models.SearchPost(&posts, param)
	if err != nil {
		abortError(c, http.StatusInternalServerError, err.Error())
		return
	}
	if err2 != nil {
		abortError(c, http.StatusInternalServerError, err2.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"users":   users,
		"posts":   posts,
	})
}
