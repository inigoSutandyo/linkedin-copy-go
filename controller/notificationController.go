package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/inigoSutandyo/linkedin-copy-go/model"
)

func UserNotifications(c *gin.Context) {
	id := getUserID(c)
	if id == "" {
		abortError(c, http.StatusUnauthorized, "Unauthorized")
		return
	}
	var notifications []model.Notification
	err := model.GetNotifications(id, &notifications)

	if err != nil {
		abortError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":       "success",
		"notifications": notifications,
	})
}

func RemoveNotification(c *gin.Context) {
	id, _ := c.GetQuery("id")

	if id == "" {
		abortError(c, http.StatusBadRequest, "Error")
		return
	}

	err := model.DeleteNotification(id)

	if err != nil {
		abortError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
