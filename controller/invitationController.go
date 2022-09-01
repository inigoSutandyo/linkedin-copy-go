package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/inigoSutandyo/linkedin-copy-go/model"
)

func GetAllInvitations(c *gin.Context) {
	invitations := model.GetAllInvitations()

	c.JSON(http.StatusOK, gin.H{
		"message":     "success",
		"invitations": invitations,
	})
}
