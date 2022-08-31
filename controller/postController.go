package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/inigoSutandyo/linkedin-copy-go/model"
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

func UploadFilePost(c *gin.Context) {
	postId := c.Request.FormValue("id")
	publicId := c.Request.FormValue("publicid")
	fileUrl := c.Request.FormValue("fileurl")
	id, _ := toUint(postId)
	post, err := model.GetPostByID(id)

	if err != nil {
		abortError(c, http.StatusBadRequest, err.Error())
		return
	}

	model.UploadFilePost(&post, fileUrl, publicId)

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"post":    post,
	})
}

func GetPosts(c *gin.Context) {
	var posts []model.Post
	var users []model.User
	offset := c.Query("offset")
	limit := c.Query("limit")
	offset_int, _ := strconv.ParseInt(offset, 10, 32)
	limit_int, _ := strconv.ParseInt(limit, 10, 32)
	err := model.GetPostsInRange(&posts, &users, int(offset_int), int(limit_int))
	if err != nil {
		abortError(c, http.StatusInternalServerError, err.Error())
	}
	// fmt.Println(posts[0])
	var hasMore bool
	if len(posts) < int(limit_int) {
		hasMore = false
	} else {
		hasMore = true
	}
	c.JSON(http.StatusOK, gin.H{
		"posts":   posts,
		"hasmore": hasMore,
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
		return
	}

	post, err2 := model.GetPostByID(post_id)

	if err2 != nil {
		abortError(c, http.StatusInternalServerError, err2.Error())
		return
	}
	postLike, _ := model.GetPostLike(user.ID, post_id)
	if postLike.ID > 0 {
		abortError(c, http.StatusBadRequest, "Already Liked")
	}

	err3 := model.CreatePostLike(&user, &post)
	if err3 != nil {
		abortError(c, http.StatusInternalServerError, err3.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "success",
		"likepost": post_id,
		"post":     post,
	})
}

func RemoveLikePost(c *gin.Context) {
	id := getUserID(c)
	postId, _ := c.GetQuery("id")
	model.DeleteLikedPostData(id, postId)
	uint_id, _ := toUint(postId)
	post, err := model.GetPostByID(uint_id)

	if err != nil {
		abortError(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, gin.H{
		"message":  "success",
		"likepost": postId,
		"post":     post,
	})
}
