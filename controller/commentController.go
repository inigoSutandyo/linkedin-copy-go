package controller

import (
	"net/http"
	"strconv"

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

func GetCommentCount(c *gin.Context) {
	post_id, _ := c.GetQuery("id")

	count := model.GetCommentCount(post_id)
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"count":   count,
	})

}

func GetComments(c *gin.Context) {
	post_id, _ := c.GetQuery("id")
	offset, _ := strconv.ParseInt(c.Query("offset"), 10, 32)
	limit, _ := strconv.ParseInt(c.Query("limit"), 10, 32)

	var comments []model.Comment
	var err error

	err = model.GetCommentByPostWithRange(post_id, &comments, int(offset), int(limit))

	if err != nil {
		abortError(c, http.StatusInternalServerError, err.Error())
		return
	}
	hasmore := true
	if len(comments) < int(limit) {
		hasmore = false
	}
	c.JSON(http.StatusOK, gin.H{
		"message":  "success",
		"comments": comments,
		"hasmore":  hasmore,
	})

}
