package model

import (
	"fmt"
	"math/big"
	"time"

	utils "github.com/inigoSutandyo/linkedin-copy-go/utils"
	"gorm.io/gorm"
)

type User struct {
	// tableName struct{} `pg:"users"`
	gorm.Model
	Email     string `json:"email" gorm:"unique"`
	Password  []byte `json:"-"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Phone     string `json:"phone" gorm:"unique"`
	Dob       time.Time
	// Posts         []Post
	// LikedPosts    []*Post `gorm:"many2many:user_likedposts"`
	// Comments      []Comment
	// LikedComments []*Comment `gorm:"many2many:user_likedcomments"`
	// Replies       []Reply
	// SharedFrom    []*Share `gorm:"many2many:user_sharedfrom"`
	// SharedTo      []*Share `gorm:"many2many:user_sharedto"`
	Template Template `gorm:"embedded"`
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

func CreateUser(email string, password []byte) (User, error) {
	user := User{
		Email:    email,
		Password: password,
	}
	err := utils.DB.Create(&user).Error
	return user, err
}
