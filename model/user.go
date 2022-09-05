package model

import (
	"time"

	utils "github.com/inigoSutandyo/linkedin-copy-go/utils"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email              string       `json:"email" gorm:"unique"`
	Password           []byte       `json:"-"`
	Headline           string       `json:"headline"`
	FirstName          string       `json:"firstname"`
	LastName           string       `json:"lastname"`
	Phone              string       `json:"phone"`
	ImageURL           string       `json:"imageurl"`
	ImagePublicID      string       `json:"imageid"`
	BackgroundURL      string       `json:"backgroundurl"`
	BackgroundPublicID string       `json:"backgroundid"`
	Dob                time.Time    `json:"dob"`
	Posts              []Post       `json:"-"`
	PostLikes          []PostLike   `json:"-"`
	Connections        []*User      `gorm:"many2many:user_connections" json:"connections"`
	Educations         []Education  `json:"-"`
	Invitations        []Invitation `gorm:"foreignKey:DestinationID" json:"invitations"`
	SourceInvitations  []Invitation `gorm:"foreignKey:SourceID" json:"-"`
	Mentions           []Comment    `gorm:"foreignKey:MentionID"`
}

func GetUserById(id string) User {
	var user User
	utils.DB.Raw("SELECT * FROM users WHERE id = ?", id).Scan(&user)
	return user
}

func GetUserByEmail(email string) User {

	var user User
	utils.DB.Raw("SELECT id, email, password FROM users WHERE email = ?", email).Scan(&user)
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

func UploadImageUser(user *User, url string, publicid string) error {
	user.ImageURL = url
	user.ImagePublicID = publicid
	err := utils.DB.Save(user).Error
	return err
}

func SearchUserByName(users *[]User, param string) error {
	param = "%" + param + "%"
	return utils.DB.Raw("SELECT * FROM users WHERE users.first_name ILIKE ? OR users.last_name ILIKE ?", param, param).Scan(users).Error
}

func CreateConnection(user *User, connect *User) error {
	err := utils.DB.Model(user).Association("Connections").Append(connect)
	if err == nil {
		err = utils.DB.Model(connect).Association("Connections").Append(user)
	}
	return err
}

func DeleteConnection(user *User, connect *User) error {
	err := utils.DB.Model(user).Association("Connections").Delete(connect)
	if err == nil {
		err = utils.DB.Model(connect).Association("Connections").Delete(user)
	}
	return err
}

func GetConnection(user *User) error {
	return utils.DB.Preload("Connections").Find(user).Error
}

func GetInvitations(user *User) {
	var invitations []Invitation
	utils.DB.Preload("Source").Find(&invitations, "destination_id = ?", user.ID)
	user.Invitations = invitations
}
