package model

import (
	"time"

	"gorm.io/gorm"
)

type Education struct {
	gorm.Model
	Institute    string    `json:"institute" gorm:"not null"`
	Degree       string    `json:"degree"`
	FieldOfStudy string    `json:"fieldofstudy"`
	Grade        string    `json:"grade"`
	Description  string    `json:"description"`
	Activities   string    `json:"activities"`
	StartTime    time.Time `json:"start"`
	EndTime      time.Time `json:"end"`
	UserID       uint      `json:"-"`
	User         User      `json:"user"`
}
