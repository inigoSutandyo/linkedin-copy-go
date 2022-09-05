package model

import (
	"time"

	"gorm.io/gorm"
)

type Education struct {
	gorm.Model
	Institute string    `json:"institute"`
	Degree    string    `json:"degree"`
	StartTime time.Time `json:"start"`
	EndTime   time.Time `json:"end"`
	UserID    uint      `json:"-"`
	User      User      `json:"user"`
}
