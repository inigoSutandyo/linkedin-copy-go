package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/microcosm-cc/bluemonday"
)

func sanitizeHtml(html string) string {
	p := bluemonday.UGCPolicy()

	sanitized := p.Sanitize(html)
	return sanitized
}

func abortError(c *gin.Context, code int, message string) {
	c.AbortWithStatusJSON(code, gin.H{
		"status":  false,
		"message": message,
	})
}
