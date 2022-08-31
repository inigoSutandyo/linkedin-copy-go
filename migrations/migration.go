package migrations

import (
	models "github.com/inigoSutandyo/linkedin-copy-go/model"
	utils "github.com/inigoSutandyo/linkedin-copy-go/utils"
)

func Migrate() {
	// utils.DB.Migrator().CreateConstraint(&models.User{}, "Posts")
	// utils.DB.Migrator().DropTable( &models.PostLike{}, &models.Comment{}, &models.Reply{}, &models.Post{}, &models.User{})
	// utils.DB.Migrator().DropTable(&models.Comment{})
	// utils.DB.Migrator().DropTable(&models.PostLike{})
	// utils.DB.Migrator().DropColumn(&models.Post{}, "FileMime")
	// utils.DB.Migrator().DropColumn(&models.Post{}, "File")
	// utils.DB.Migrator().DropTable(&models.Invitation{})
	utils.DB.AutoMigrate(
		&models.User{},
		&models.Post{},
		&models.Comment{},
		&models.Reply{},
		&models.PostLike{},
		&models.Invitation{})
}
