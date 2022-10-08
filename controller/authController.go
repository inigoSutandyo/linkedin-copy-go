package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/inigoSutandyo/linkedin-copy-go/model"
	models "github.com/inigoSutandyo/linkedin-copy-go/model"
	"github.com/inigoSutandyo/linkedin-copy-go/utils"
	"golang.org/x/crypto/bcrypt"
)

type AuthData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Google struct {
	Email     string `json:"email"`
	FirstName string `json:"given_name"`
	LastName  string `json:"family_name"`
	ImageURL  string `json:"picture"`
}

func GoogleLogin(c *gin.Context) {
	var data Google
	c.BindJSON(&data)
	// fmt.Println(data)
	if data.Email == "" {
		abortError(c, http.StatusBadRequest, "Email is empty!!")
		return
	}
	user := models.GetGoogleUser(data.Email)
	fmt.Println(user.IsGoogle)
	if user.ID == 0 {
		// if no user create 1
		newUser, err := model.CreateGoogleUser(data.Email, data.FirstName, data.LastName, data.ImageURL)
		if err != nil {
			abortError(c, http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
			"user":    newUser,
		})
		createJWT(c, newUser)
		return
	}
	if user.IsGoogle == false {
		abortError(c, http.StatusBadRequest, "Email already registered (not google)")
		return
	} else {
		createJWT(c, user)
		return
	}
}

func Login(c *gin.Context) {
	message := "success"

	var data AuthData
	c.BindJSON(&data)
	email := data.Email
	password := data.Password

	var user models.User
	user = models.GetUserByEmail(email)
	err := bcrypt.CompareHashAndPassword(user.Password, []byte(password))

	if err != nil {
		message = "User not found"
		c.JSON(http.StatusBadRequest, gin.H{
			"message": message,
			"error":   err.Error(),
			"isError": true,
		})
		return
	}
	createJWT(c, user)
}

func createJWT(c *gin.Context, user model.User) {
	var message string
	c.Header("Content-Type", "application/json")
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, tokenErr := claims.SignedString([]byte(utils.GetEnv("SECRET_KEY")))
	updateUser := user
	updateUser.Token = token

	model.UpdateUser(&user, "password, email, id", updateUser)
	if tokenErr != nil {
		message = "Could not sign in (SERVER ERROR)"
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": message,
			"error":   tokenErr.Error(),
			"isError": true,
		})
		return
	} else {
		c.SetCookie("token", token, 3600*12, "/", "http://127.0.0.1", true, true)
		c.SetCookie("auth", "true", 3600*12, "/", "http://127.0.0.1", false, false)

		c.SetCookie("token", token, 3600*12, "/", "http://localhost", true, true)
		c.SetCookie("auth", "true", 3600*12, "/", "http://localhost", false, false)

		c.JSON(http.StatusOK, gin.H{
			"message": message,
			"isError": false,
		})
		return
	}
}

func Register(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	message := "success"

	var data AuthData
	bindErr := c.BindJSON(&data)
	email := data.Email
	password := data.Password

	if email == "" || password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error!",
		})
		return
	}

	if bindErr != nil {
		message = bindErr.Error()
		c.JSON(http.StatusBadRequest, gin.H{
			"message": message,
		})
		return
	}
	user := models.GetUserByEmail(email)
	if user.ID != 0 {
		abortError(c, http.StatusBadRequest, "Email already registerd")
		return
	}

	pw, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	_, err := models.CreateUser(email, pw)

	if err != nil {
		message = err.Error()
	}
	fmt.Println(data)
	fmt.Println(message)

	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}

func Logout(c *gin.Context) {
	message := "success"
	// c.Cookie("token")
	c.SetCookie("token", "deleting", -1, "/", "http://localhost", false, true)
	c.SetCookie("token", "deleting", -1, "/", "http://127.0.0.1", false, true)

	c.SetCookie("auth", "deleting", -1, "/", "http://localhost", false, true)
	c.SetCookie("auth", "deleting", -1, "/", "http://127.0.0.1", false, true)
	id := getUserID(c)
	user := model.GetUserById(id)
	c.JSON(http.StatusOK, gin.H{
		"message":  message,
		"isgoogle": user.IsGoogle,
	})
}

func CheckAuth(c *gin.Context) (bool, *jwt.Token, error) {
	cookie, err := c.Cookie("token")

	if err != nil {
		return false, nil, err
	}

	token, tokenErr := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(utils.GetEnv("SECRET_KEY")), nil
	})

	// fmt.Print("Token = ")
	// fmt.Println(token

	if tokenErr != nil {
		return false, nil, tokenErr
	}

	return true, token, nil
}
func ClientAuth(c *gin.Context) {
	status, _, _ := CheckAuth(c)
	fmt.Println(status)

	if status {
		cookie, _ := c.Cookie("token")
		fmt.Println(cookie)
		user := model.GetUserFromToken(cookie)

		if user.ID == 0 {
			abortError(c, http.StatusUnauthorized, "Unauthorized User Token not found!!")
			return
		}

		if user.Token != cookie {
			abortError(c, http.StatusUnauthorized, "Unauthorized User Token!!")
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  status,
			"message": "Client is authorized",
		})
	} else {
		abortError(c, http.StatusUnauthorized, "Unauthorized User!!")
		return
	}
}
