package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	models "github.com/inigoSutandyo/linkedin-copy-go/model"
)

func SearchUserHandler(c *gin.Context) {

	param := c.Query("search")
	offset := c.Query("offset")
	var users []models.User
	err := models.SearchUserByName(&users, param, offset)
	if err != nil {
		abortError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"users":   users,
	})
}

func SearchPostHandler(c *gin.Context) {
	param := c.Query("search")
	offset, _ := strconv.Atoi(c.Query("offset"))
	fmt.Println(offset)
	var posts []models.Post
	err2 := models.SearchPost(&posts, param, offset)

	if err2 != nil {
		abortError(c, http.StatusInternalServerError, err2.Error())
		return
	}

	var hasMore bool
	if len(posts) < 5 {
		hasMore = false
	} else {
		hasMore = true
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"posts":   posts,
		"hasmore": hasMore,
	})
}
