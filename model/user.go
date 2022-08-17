package model

import (
	"fmt"
	"math/big"
	"time"

	utils "github.com/inigoSutandyo/linkedin-copy-go/utils"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	// tableName struct{} `pg:"users"`
	Id        uint
	Email     string
	Password  []byte
	FirstName string
	LastName  string
	Phone     string
	Dob       time.Time
}

func GetUserById(id big.Int) User {

	var user User
	utils.DB.Raw("SELECT * FROM users WHERE id = ?", id).Scan(&user)
	fmt.Println(user)
	return user
}

func GetUserByEmail(email string) User {

	var user User
	utils.DB.Raw("SELECT id, email, password FROM users WHERE email = ?", email).Scan(&user)
	fmt.Println(user)
	return user
}

func GetAllUsers() []User {

	var users []User
	utils.DB.Find(&users)
	fmt.Println(users)
	return users
}

func GetUserByEmailAndPassword(email string, password string) User {

	var user User
	utils.DB.First(&user, "email = ? AND password = ?", email, password)
	return user
}

func CreateUser(email string, password string) (User, error) {

	pw, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	user := User{
		Email:    email,
		Password: pw,
	}
	err := utils.DB.Create(&user).Error
	return user, err
}
