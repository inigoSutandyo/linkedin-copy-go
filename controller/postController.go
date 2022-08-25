package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/inigoSutandyo/linkedin-copy-go/model"
	"github.com/inigoSutandyo/linkedin-copy-go/utils"
	// "github.com/fmpwizard/go-quilljs-delta/delta"
)

func AddPost(c *gin.Context) {
	id := getUserID(c)
	var post model.Post
	c.BindJSON(&post)

	post.Content = sanitizeHtml(post.Content)

	user := model.GetUserById(id)

	dbErr := model.CreatePost(&user, &post)

	if dbErr != nil {
		fmt.Print("ERROR = ")
		fmt.Println(dbErr.Error())
		abortError(c, http.StatusInternalServerError, dbErr.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"post":    post,
	})
}

func GetPosts(c *gin.Context) {
	var posts []model.Post
	var users []model.User
	err := model.GetAllPost(&posts, &users)

	if err != nil {
		abortError(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
		// "users": users,
	})
}

func AddLikePost(c *gin.Context) {
	id := getUserID(c)
	user := model.GetUserById(id)

	id_str, _ := c.GetQuery("id")
	post_id, err := toUint(id_str)
	if err != nil {
		abortError(c, http.StatusInternalServerError, err.Error())
	}

	post, err2 := model.GetPostByID(post_id)

	if err2 != nil {
		abortError(c, http.StatusInternalServerError, err2.Error())
	}

	err3 := model.CreatePostLike(&user, &post)
	if err3 != nil {
		abortError(c, http.StatusInternalServerError, err3.Error())
	}
	post.Likes = post.Likes + 1
	utils.DB.Save(&post)
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
