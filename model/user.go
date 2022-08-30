package model

import (
	"fmt"
	"time"

	utils "github.com/inigoSutandyo/linkedin-copy-go/utils"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email       string `json:"email" gorm:"unique"`
	Password    []byte `json:"-"`
	Headline    string `json:"headline"`
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	Phone       string `json:"phone" gorm:"unique"`
	ImageURL    string `json:"imageurl"`
	Dob         time.Time
	Posts       []Post     `json:"-"`
	Comments    []Comment  `json:"-"`
	Replies     []Reply    `json:"-"`
	PostLikes   []PostLike `json:"-"`
	Connections []*User    `gorm:"many2many:user_connections"`
}

func GetUserById(id string) User {
	var user User
	utils.DB.Raw("SELECT * FROM users WHERE id = ?", id).Scan(&user)
	// fmt.Println(user)
	return user
}

func GetUserByEmail(email string) User {

	var user User
	utils.DB.Raw("SELECT id, email, password FROM users WHERE email = ?", email).Scan(&user)
	fmt.Println(user)
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

func UpdateUser(user *User, omit string, update User) {
	utils.DB.Model(&user).Omit(omit).Updates(update)
}

func GetUserPost(user *User) []Post {
	var post []Post
	utils.DB.Model(user).Association("Posts").Find(&post)
	return post
}

func UploadImageUser(user *User, url string) error {
	user.ImageURL = "https://res.cloudinary.com/dy8gj7hrx/image/upload/v1661868616/users/profiles/zpc2m2sy2taacfe836nj.jpg"
	err := utils.DB.Save(user).Error
	return err
}

func SearchUserByName(users *[]User, param string) error {
	param = "%" + param + "%"
	return utils.DB.Raw("SELECT * FROM users WHERE users.first_name ILIKE ? OR users.last_name ILIKE ?", param, param).Scan(users).Error
}
