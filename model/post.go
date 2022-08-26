package model

import (
	"github.com/inigoSutandyo/linkedin-copy-go/utils"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title      string     `json:"-"`
	Content    string     `json:"content" gorm:"text"`
	Attachment string     `json:"attachment"`
	Likes      int        `json:"likes"`
	UserID     uint       `json:"-"`
	User       User       `json:"user"`
	Comments   []Comment  `json:"-"`
	PostLikes  []PostLike `json:"-"`
	// Template   Template `gorm:"embedded"`s
	// PostLikes  []PostLike
}

func GetPostByID(id uint) (Post, error) {
	var post Post
	err := utils.DB.Preload("User").First(&post, id).Error
	getPostLikeCount(&post)

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
	err := utils.DB.Preload("User").Order("posts.created_at desc").Find(posts).Error

	for _, post := range *posts {
		getPostLikeCount(&post)
	}

	return err
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
