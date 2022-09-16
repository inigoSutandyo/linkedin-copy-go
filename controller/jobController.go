package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/inigoSutandyo/linkedin-copy-go/model"
)

func AddJob(c *gin.Context) {
	id := getUserID(c)

	if id == "" {
		abortError(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	user := model.GetUserById(id)

	var job model.Job
	c.BindJSON(&job)
	err := model.CreateJob(&job, &user)

	if err != nil {
		abortError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
		"job":     job,
	})
}

func GetAllJobs(c *gin.Context) {
	jobs, err := model.GetJobs()

	if err != nil {
		abortError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
		"jobs":    jobs,
	})
}
