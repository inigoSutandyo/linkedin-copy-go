package model

import (
	"fmt"

	utils "github.com/inigoSutandyo/linkedin-copy-go/utils"
)

type User struct {
	// tableName struct{} `pg:"users"`
	Id       uint
	Email    string
	Password string
}

func GetUserById(id uint) User {
	db := utils.GetDatabase()
	var user User
	db.Raw("SELECT id, email, password FROM users WHERE id = ?", id).Scan(&user)
	fmt.Println(user)
	return user
}
