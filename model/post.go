package model

import (
	"github.com/inigoSutandyo/linkedin-copy-go/utils"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title      string    `json:"-"`
	Content    string    `json:"content" gorm:"text"`
	Attachment string    `json:"attachment"`
	Likes      int       `json:"likes"`
	UserID     uint      `json:"-"`
	User       User      `json:"user"`
	Comments   []Comment `json:"-"`
	// Template   Template `gorm:"embedded"`s
	// PostLikes  []PostLike
}

func GetPostByID(id uint) (Post, error) {
	var post Post
	err := utils.DB.First(&post, id).Error

	return post, err
}

func CreatePost(user *User, post *Post) error {
	// fmt.Println(user.Email)
	// fmt.Println(post.Content)
	err := utils.DB.Model(user).Association("Posts").Append(post)
	return err
}

func GetAllPost(posts *[]Post, users *[]User) error {
	// err := utils.DB.Find(posts).Error
	err := utils.DB.Preload("User").Find(posts).Error
	return err
}
