package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/inigoSutandyo/linkedin-copy-go/model"
)

func AddReply(c *gin.Context) {
	var reply model.Reply
	c.BindJSON(&reply)

	reply.Content = sanitizeHtml(reply.Content)
	comment, err := model.GetCommentById(reply.CommentID)

	if err != nil {
		abortError(c, http.StatusInternalServerError, err.Error())
	}
	id := getUserID(c)
	user := model.GetUserById(id)
	dbErr := model.CreateReply(&user, &comment, &reply)

	if dbErr != nil {
		abortError(c, http.StatusInternalServerError, dbErr.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"reply":   reply,
	})
}

func GetReplies(c *gin.Context) {
	var replies []model.Reply
	id_str, _ := c.GetQuery("id")
	comment_id, convErr := toUint(id_str)
	if convErr != nil {
		abortError(c, http.StatusInternalServerError, convErr.Error())
	}
	comment, err := model.GetCommentById(comment_id)

	model.GetRepliesForComments(&comment, &replies)
	if err != nil {
		abortError(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"replies": replies,
	})
}
