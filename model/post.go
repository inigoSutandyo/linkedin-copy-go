package model

import (
	"net/http"

	"github.com/inigoSutandyo/linkedin-copy-go/utils"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title      string     `json:"-"`
	Content    string     `json:"content" gorm:"text"`
	Attachment string     `json:"attachment"`
	Likes      int        `json:"likes"`
	File       []byte     `json:"file"`
	FileMime   string     `json:"mime"`
	UserID     uint       `json:"-"`
	User       User       `json:"user"`
	Comments   []Comment  `json:"-"`
	PostLikes  []PostLike `json:"-"`
}

func GetPostByID(id uint) (Post, error) {
	var post Post
	err := utils.DB.Preload("User").First(&post, id).Error
	getPostLikeCount(&post)

	return post, err
}

func CreatePost(user *User, post *Post) error {
	err := utils.DB.Model(user).Association("Posts").Append(post)
	err = utils.DB.Preload("User").First(post).Error
	return err
}

func GetAllPost(posts *[]Post, users *[]User) error {
	// err := utils.DB.Find(posts).Error
	err := utils.DB.Preload("User").Order("posts.created_at desc").Find(posts).Error

	for _, post := range *posts {
		getPostLikeCount(&post)
	}

	return err
}

func GetPostsInRange(posts *[]Post, users *[]User, offset int, limit int) error {
	err := utils.DB.Preload("User").Order("posts.created_at desc").Limit(limit).Offset(offset).Find(posts).Error

	for _, post := range *posts {
		getPostLikeCount(&post)
	}
	return err
}

func UploadFilePost(post *Post, data []byte) {
	post.File = data
	post.FileMime = http.DetectContentType(data)
	utils.DB.Save(post)
	utils.DB.Preload("User").First(post)
}

func getPostLikeCount(post *Post) {
	count := utils.DB.Model(post).Association("PostLikes").Count()

	post.Likes = int(count)
	utils.DB.Save(post)
}

func getPostLikeCountById(id string) {
	var post Post
	utils.DB.Find(&post, id)
	count := utils.DB.Model(&post).Association("PostLikes").Count()

	post.Likes = int(count)
	utils.DB.Save(&post)
}
