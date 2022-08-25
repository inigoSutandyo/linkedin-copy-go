package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/inigoSutandyo/linkedin-copy-go/model"
)

func AddComment(c *gin.Context) {
	var comment model.Comment
	c.BindJSON(&comment)

	comment.Content = sanitizeHtml(comment.Content)

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
		abortError(c, http.StatusInternalServerError, dbErr.Error())
	}

	// AddReplyDebug(&comment, &user)

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"comment": comment,
	})
}

func AddReplyDebug(comment *model.Comment, user *model.User) {
	var reply = model.Reply{
		Content:   "<p>Testing</p>",
		CommentID: comment.ID,
	}
	model.CreateReply(user, comment, &reply)
}

func GetComments(c *gin.Context) {
	post_id, _ := c.GetQuery("id")
	var comments []model.Comment
	err := model.GetCommentByPost(post_id, &comments)

	if err != nil {
		abortError(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "success",
		"comments": comments,
	})
}
