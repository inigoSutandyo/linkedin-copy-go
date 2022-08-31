package model

import (
	"github.com/inigoSutandyo/linkedin-copy-go/utils"
	"gorm.io/gorm"
)

type Invitation struct {
	gorm.Model
	Note        string
	Source      []User `gorm:"many2many:invitation_source" json:"source"`
	Destination []User `gorm:"many2many:invitation_destination" json:"destination"`
}

func CreateInvitation(source *User, destination *User, note string) (Invitation, error) {
	invite := Invitation{
		Note: note,
	}
	err := utils.DB.Create(&invite).Error

	if err == nil {
		err = utils.DB.Model(&invite).Association("Source").Append(source)
	}

	if err == nil {
		err = utils.DB.Model(&invite).Association("Destination").Append(destination)
	}

	return invite, err
}
