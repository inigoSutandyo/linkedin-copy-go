package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/inigoSutandyo/linkedin-copy-go/model"
)

func AddReply(c *gin.Context) {
	var comment model.Comment
	c.BindJSON(&comment)

	comment.Content = sanitizeHtml(comment.Content)
}
