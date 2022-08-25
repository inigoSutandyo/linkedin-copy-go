package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/inigoSutandyo/linkedin-copy-go/model"
	// "github.com/fmpwizard/go-quilljs-delta/delta"
)

func AddPost(c *gin.Context) {
	id := getUserID(c)
	var post model.Post
	c.BindJSON(&post)

	post.Content = sanitizeHtml(post.Content)

	user := model.GetUserById(id)

	dbErr := model.CreatePost(&user, &post)

	if dbErr != nil {
		fmt.Print("ERROR = ")
		fmt.Println(dbErr.Error())
		abortError(c, http.StatusInternalServerError, dbErr.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"post":    post,
	})
}

func GetPosts(c *gin.Context) {
	var posts []model.Post
	var users []model.User
	err := model.GetAllPost(&posts, &users)

	if err != nil {
		abortError(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
		// "users": users,
	})
}
