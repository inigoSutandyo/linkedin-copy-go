package controller

import (
	"strconv"

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

func toUint(num string) (uint, error) {
	num64, err := strconv.ParseUint(num, 10, 32)
	if err != nil {
		return 0, err
	}

	res := uint(num64)
	return res, err
}
