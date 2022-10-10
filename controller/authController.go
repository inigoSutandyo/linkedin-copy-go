package controller

import (
	"crypto/tls"
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
	gomail "gopkg.in/gomail.v2"
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
	user, err := models.CreateUser(email, pw)

	if err != nil {
		abortError(c, http.StatusInternalServerError, err.Error())
	}
	token := utils.GenerateVerifToken(user.ID)
	url := utils.GetEnv("CLIENT_URL") + "/verification/" + token

	updateUser := user
	updateUser.VerifToken = token
	updateUser.VerifExpire = time.Now().Add(time.Minute * 30)
	model.UpdateUser(&user, "password, email. id", updateUser)

	emailMessage := "Hello, " + email + " you have registered to LinkHEdIn. Click the link below to verify your email.\n" + url + "\n"
	err = createEmail(email, "LinkHEdIn Verification", emailMessage)
	// fmt.Println(data)
	// fmt.Println(message)

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func VerifyUser(c *gin.Context) {
	token := c.Query("token")

	user := model.GetUserFromVerifToken(token)

	if user.ID == 0 {
		abortError(c, http.StatusBadRequest, "Token is invalid")
		return
	}

	if time.Now().After(user.VerifExpire) {
		abortError(c, http.StatusBadRequest, "Token has expired")
		return
	}

	updateUser := user
	updateUser.IsVerified = true
	model.UpdateUser(&user, "email, password, id", updateUser)

	c.JSON(200, gin.H{
		"message": "success",
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
		user := model.GetUserFromAuthToken(cookie)

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

func ForgetRequest(c *gin.Context) {
	email := c.Query("email")

	user := model.GetUserByEmail(email)
	if user.ID == 0 {
		abortError(c, http.StatusBadRequest, "Email not registered")
		return
	}

	expire := time.Now().Add(time.Minute * 30)
	token := utils.GenerateForgetToken(user.ID)
	url := utils.GetEnv("CLIENT_URL") + "/reset/" + token
	emailMessage := "Hello, " + email + " your request to reset password has been accepted. Click the link below to change your password.\n" + url + "\n"
	updateUser := user
	updateUser.ForgetToken = token
	updateUser.ForgetExpire = expire

	model.UpdateUser(&user, "password, email, id", updateUser)

	if err := createEmail(email, "Forget Password Request", emailMessage); err != nil {
		abortError(c, http.StatusInternalServerError, err.Error())
		fmt.Println(err)
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
	})
}

func createEmail(email string, title string, message string) error {
	from := utils.GetEnv("GMAIL_EMAIL")
	password := utils.GetEnv("GMAIL_PASSWORD")

	m := gomail.NewMessage()

	m.SetHeader("From", from)
	m.SetHeader("To", email)

	m.SetHeader("Subject", title)
	m.SetBody("text/plain", message)

	d := gomail.NewDialer("smtp.gmail.com", 587, from, password)

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	return d.DialAndSend(m)
}

type Reset struct {
	Forget   string `json:"token"`
	Password string `json:"password"`
}

func ResetPassword(c *gin.Context) {
	var reset Reset
	c.BindJSON(&reset)

	user := model.GetUserForgetPassword(reset.Forget)

	if user.ID == 0 {
		abortError(c, http.StatusInternalServerError, "User not found...")
		return
	}

	if time.Now().After(user.ForgetExpire) {
		abortError(c, http.StatusInternalServerError, "Token expired")
		return
	}

	updateUser := user
	pw, _ := bcrypt.GenerateFromPassword([]byte(reset.Password), 14)
	updateUser.Password = pw
	model.UpdateUser(&user, "email, id", updateUser)

	c.JSON(200, gin.H{
		"message": "success",
	})
}
