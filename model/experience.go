package model

import (
	"time"

	"gorm.io/gorm"
)

type Experience struct {
	gorm.Model
	Title          string    `json:"title" gorm:"not null"`
	EmploymentType string    `json:"employmenttype"`
	CompanyName    string    `json:"company" gorm:"not null"`
	Location       string    `json:"location"`
	Industry       string    `json:"industry"`
	Description    string    `json:"description"`
	IsWorking      bool      `json:"isworking"`
	StartTime      time.Time `json:"start"`
	EndTime        time.Time `json:"end"`
	UserID         uint      `json:"-"`
	User           User      `json:"user"`
}
