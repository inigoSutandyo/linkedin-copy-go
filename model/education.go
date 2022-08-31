package model

import (
	"time"

	"gorm.io/gorm"
)

type Education struct {
	gorm.Model
	Name      string    `json:"name"`
	StartTime time.Time `json:"start"`
	EndTime   time.Time `json:"end"`
	UserID    uint      `json:"-"`
	User      User      `json:"user"`
}
