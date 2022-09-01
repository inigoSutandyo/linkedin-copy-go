package model

import (
	"github.com/inigoSutandyo/linkedin-copy-go/utils"
	"gorm.io/gorm"
)

type Invitation struct {
	gorm.Model
	Note          string `json:"note"`
	SourceID      uint   `json:"sourceid"`
	Source        User   `json:"source"`
	DestinationID uint   `json:"destinationid"`
	Destination   User   `json:"destination"`
}

func CreateInvitation(source *User, destination *User, note string) (Invitation, error) {
	invite := Invitation{
		Note:          note,
		SourceID:      source.ID,
		DestinationID: destination.ID,
	}
	err := utils.DB.Create(&invite).Error
	// fmt.Println(invite)
	if err == nil {
		err = utils.DB.Model(&source).Association("SourceInvitations").Append(&invite)
	}

	if err == nil {
		err = utils.DB.Model(&destination).Association("Invitations").Append(&invite)
	}

	return invite, err
}

func DeleteInvitation(sourceId string, destinationId string) error {
	return utils.DB.Where("source_id = ? AND destination_id = ?", sourceId, destinationId).Delete(&Invitation{}).Error
}

func GetAllInvitations() []Invitation {
	var invitations []Invitation
	utils.DB.Find(&invitations)
	return invitations
}
