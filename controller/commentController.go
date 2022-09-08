package controller

import (
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

	id := getUserID(c)
	user := model.GetUserById(id)

	dbErr := model.CreateComment(&user, &post, &comment)

	if dbErr != nil {
		abortError(c, http.StatusInternalServerError, dbErr.Error())
		return
	}
	var notification = model.Notification{
		Message:   " commented on your post",
		HasSource: true,
	}

	model.CreateNotificationForPost(&post.User, &user, &notification, &post)

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"comment": comment,
	})
}

func GetComments(c *gin.Context) {
	post_id, _ := c.GetQuery("id")
	var comments []model.Comment
	err := model.GetCommentByPost(post_id, &comments)

	if err != nil {
		abortError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "success",
		"comments": comments,
	})
}
