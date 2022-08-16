package model

import (
	"fmt"

	utils "github.com/inigoSutandyo/linkedin-copy-go/utils"
)

type User struct {
	// tableName struct{} `pg:"users"`
	Id    uint
	Email string
	Name  string
	Phone string
	Dob   string
}

func GetUserById(id uint) User {
	db := utils.GetDatabase()
	var user User
	db.Raw("SELECT * FROM users WHERE id = ?", id).Scan(&user)
	fmt.Println(user)
	return user
}

func GetUserByEmail(email string) User {
	db := utils.GetDatabase()
	var user User
	db.Raw("SELECT id, email, password FROM users WHERE email = ?", email).Scan(&user)
	fmt.Println(user)
	return user
}

func GetAllUsers() []User {
	db := utils.GetDatabase()
	var users []User
	db.Find(&users)
	fmt.Println(users)
	return users
}
