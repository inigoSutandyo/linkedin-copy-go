package model

import (
	"github.com/inigoSutandyo/linkedin-copy-go/utils"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title        string     `json:"-"`
	Content      string     `json:"content" gorm:"text"`
	Attachment   string     `json:"attachment"`
	Likes        int        `json:"likes"`
	FileUrl      string     `json:"fileurl"`
	FilePublicID string     `json:"fileid"`
	UserID       uint       `json:"-"`
	User         User       `json:"user"`
	Comments     []Comment  `json:"-"`
	PostLikes    []PostLike `json:"-"`
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

func UploadFilePost(post *Post, url string, publicid string) {
	post.FileUrl = url
	post.FilePublicID = publicid
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

func SearchPost(posts *[]Post, param string, offset int) error {
	param = "%" + param + "%"
	err := utils.DB.Preload("User").Order("posts.created_at desc").Limit(5).Offset(offset).Find(posts, "posts.content ILIKE ?", param).Error

	for _, post := range *posts {
		getPostLikeCount(&post)
	}

	return err
}

func DeletePost(id string) error {
	var post Post
	// utils.DB.Find(&post, "id = ?", id).Error
	return utils.DB.Delete(&post, "id = ?", id).Error
}
