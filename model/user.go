package model

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	utils "github.com/inigoSutandyo/linkedin-copy-go/utils"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email     string `json:"email" gorm:"unique"`
	Password  []byte `json:"-"`
	Headline  string `json:"headline"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Phone     string `json:"phone" gorm:"unique"`
	Image     []byte `json:"image"`
	ImageMime string `json:"imagemime"`
	Dob       time.Time
	Posts     []Post     `json:"-"`
	Comments  []Comment  `json:"-"`
	Replies   []Reply    `json:"-"`
	PostLikes []PostLike `json:"-"`
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

func UploadImageUser(user *User, buf *bytes.Buffer) error {
	s := http.DetectContentType(buf.Bytes())
	user.Image = buf.Bytes()
	user.ImageMime = s
	err := utils.DB.Save(user).Error
	return err
}

func SaveImageMime(user *User) {
	mime := http.DetectContentType(user.Image)
	user.ImageMime = mime
	utils.DB.Save(user)
}
