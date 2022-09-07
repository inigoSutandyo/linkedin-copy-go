package migrations

import (
	models "github.com/inigoSutandyo/linkedin-copy-go/model"
	utils "github.com/inigoSutandyo/linkedin-copy-go/utils"
)

func Migrate() {
	utils.DB.AutoMigrate(
		&models.User{},
		&models.Post{},
		&models.Comment{},
		&models.PostLike{},
		&models.Invitation{},
		&models.Education{},
		&models.Experience{},
	)
}
