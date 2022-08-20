package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/inigoSutandyo/linkedin-copy-go/model"
	// "github.com/fmpwizard/go-quilljs-delta/delta"
)

func AddPost(c *gin.Context) {
	id := getUserID(c)
	var post model.Post
	c.BindJSON(&post)
	fmt.Print("Add Post = ")
	fmt.Println(post.Content)
	fmt.Println(id)
}
