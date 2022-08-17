package utils

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Utility method to get database
func Connect() {

	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=hotel port=%s sslmode=disable", GetEnv("DB_USER"), GetEnv("DB_PASSWORD"), GetEnv("DB_PORT"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	fmt.Println(err)
	DB = db
}
