package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/inigoSutandyo/linkedin-copy-go/model"
)

func AddReply(c *gin.Context) {
	var reply model.Comment
	commentId := c.Query("id")
	c.BindJSON(&reply)

	reply.Content = sanitizeHtml(reply.Content)
	comment, err := model.GetCommentById(commentId)

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
	var replies []model.Comment
	id_str, _ := c.GetQuery("id")
	comment, err := model.GetCommentById(id_str)

	model.GetRepliesForComments(&comment, &replies)
	if err != nil {
		abortError(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"replies": replies,
	})
}
