package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/inigoSutandyo/linkedin-copy-go/model"
	"github.com/microcosm-cc/bluemonday"
)

func AddComment(c *gin.Context) {
	var comment model.Comment
	c.BindJSON(&comment)

	p := bluemonday.UGCPolicy()

	html := p.Sanitize(comment.Content)
	comment.Content = html

	post, err := model.GetPostByID(comment.PostID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": err.Error(),
		})
	}
	fmt.Println(post.ID)
	id := getUserID(c)
	user := model.GetUserById(id)

	dbErr := model.CreateComment(&user, &post, &comment)

	if dbErr != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": dbErr.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"comment": comment,
	})
}

func GetComments(c *gin.Context) {
	str_id, _ := c.GetQuery("id")
	id64, _ := strconv.ParseUint(str_id, 10, 32)
	id := uint(id64)

	post, err := model.GetPostByID(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": err.Error(),
		})
	}

	var comments []model.Comment
	dbErr := model.GetCommentByPost(&post, &comments)

	if dbErr != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": dbErr.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "success",
		"comments": comments,
	})
}
