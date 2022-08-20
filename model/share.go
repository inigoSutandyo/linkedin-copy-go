package model

import "gorm.io/gorm"

type Share struct {
	gorm.Model
	Id       uint
	FromID   uint     `gorm:"many2many:user_sharedto"`
	ToID     uint     `gorm:"many2many:user_sharedfrom"`
	Template Template `gorm:"embedded"`
}
