package model

import (
	"time"

	"gorm.io/gorm"
)

type Template struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Search struct {
	Param string `json:"param"`
}
