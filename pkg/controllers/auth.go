package controllers

import (
	"fishing_company/pkg/db"
	"fishing_company/pkg/globals"
	"fishing_company/pkg/models"
	"fishing_company/pkg/utils"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type Credentials struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type RegisterCred struct {
	Username     string `form:"username"`
	Password     string `form:"password"`
	ConfPassword string `form:"conf-password"`
	Role         int    `form:"role"`
}

func LoginForm(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(globals.Userkey)
	if user != nil {
		c.HTML(http.StatusBadRequest, "login.html",
			gin.H{
				"content": "Please logout first",
				"user":    user,
			})
		return
	}
	log.Printf("%+v", user)
	c.HTML(http.StatusOK, "login.html", gin.H{
		"user": user,
	})
}

func Login(c *gin.Context) {
	session := sessions.Default(c)
	sessionUser := session.Get(globals.Userkey)
	if sessionUser != nil {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{"content": "Please logout first"})
		return
	}

	var creds Credentials
	if err := c.ShouldBind(&creds); err != nil {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{"content": "Form data is not valid"})
		return
	}

	if utils.EmptyUserPass(creds.Username, creds.Password) {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{"content": "Parameters can't be empty"})
		return
	}

	user, err := utils.CheckUserPass(creds.Username, creds.Password)
	if err != nil {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{"content": "Invalid credentials"})
		return
	}
	log.Println(user.Role.Name)
	session.Set(globals.Userkey, user.Name)
	session.Set(globals.Rolekey, user.Role.Name)

	log.Println("User role: ", session.Get(globals.Rolekey))

	if err := session.Save(); err != nil {
		c.HTML(http.StatusInternalServerError, "login.html", gin.H{"content": "Failed to save session"})
		return
	}

	c.Redirect(http.StatusMovedPermanently, "/")
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(globals.Userkey)

	log.Println("logging out user:", user)
	if user == nil {
		c.HTML(http.StatusInternalServerError, "login.html", gin.H{"content": "User not found"})
		return
	}

	session.Clear()
	if err := session.Save(); err != nil {
		c.HTML(http.StatusInternalServerError, "login.html", gin.H{"content": "Failed to clear session"})
		return
	}

	c.Redirect(http.StatusMovedPermanently, "/auth/login")
}

func RegisterForm(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(globals.Userkey)
	c.HTML(http.StatusOK, "register.html", gin.H{"user": user})
}

func Register(c *gin.Context) {
	var creds RegisterCred
	if err := c.ShouldBind(&creds); err != nil {
		log.Println(err)
		c.HTML(http.StatusBadRequest, "register.html", gin.H{"content": "Form data is not valid"})
		return
	}

	log.Printf("%+v", creds)

	if creds.Password != creds.ConfPassword {
		log.Println("Passwords not equal")
		c.HTML(http.StatusBadRequest, "register.html", gin.H{"content": "Passwords are not equal"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "register.html", gin.H{"content": "Cannot hash password"})
		log.Println(err)
		return
	}

	// проверка на существование пользователя с таким именем
	// полученная модель опускается, важен только err
	if err := db.DB.Where(&models.User{Name: creds.Username}).First(&models.User{}).Error; err == nil {
		log.Println(err)
		c.HTML(http.StatusOK, "register.html", gin.H{"content": "User with this name already exists"})
		return
	}

	newUser := models.User{
		Name:     creds.Username,
		Password: string(hash),
		RoleID:   creds.Role,
	}
	if err := db.DB.Create(&newUser).Error; err != nil {
		log.Println(err)
		c.HTML(http.StatusInternalServerError, "register.html", gin.H{"content": "Cannot create user"})
		return
	}

	c.Redirect(http.StatusMovedPermanently, "/auth/login")
}
