package model

import (
	"github.com/inigoSutandyo/linkedin-copy-go/utils"
	"gorm.io/gorm"
)

type Job struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Company     string `json:"company"`
	Location    string `json:"location"`
	UserID      uint   `json:"userid"`
	User        User   `json:"user"`
}

func CreateJob(job *Job, user *User) error {
	job.UserID = user.ID
	job.User = *user
	return utils.DB.Create(job).Error
}

func GetJobs() ([]Job, error) {
	var jobs []Job
	err := utils.DB.Preload("User").Find(&jobs).Error
	return jobs, err
}
