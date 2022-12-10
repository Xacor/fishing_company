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
		utils.FlashMessage(c, "Сначала необходимо выйти")
		c.HTML(http.StatusBadRequest, "login.html",
			gin.H{
				"alerts": utils.Flashes(c),
				"user":   user,
			})
		return
	}

	c.HTML(http.StatusOK, "login.html", gin.H{
		"user":   user,
		"alerts": utils.Flashes(c),
	})
}

func Login(c *gin.Context) {
	session := sessions.Default(c)
	sessionUser := session.Get(globals.Userkey)
	if sessionUser != nil {
		utils.FlashMessage(c, "Сначала необходимо выйти")
		c.HTML(http.StatusBadRequest, "login.html", gin.H{"alerts": utils.Flashes(c)})
		return
	}

	var creds Credentials
	if err := c.ShouldBind(&creds); err != nil {
		utils.FlashMessage(c, "Даные формы некорректны")
		c.HTML(http.StatusBadRequest, "login.html", gin.H{"alerts": utils.Flashes(c)})
		return
	}

	if utils.EmptyUserPass(creds.Username, creds.Password) {
		utils.FlashMessage(c, "Поля не могут быть пустыми")
		c.HTML(http.StatusBadRequest, "login.html", gin.H{"alerts": utils.Flashes(c)})
		return
	}

	user, err := utils.CheckUserPass(creds.Username, creds.Password)
	if err != nil {
		utils.FlashMessage(c, "Введены неверные учтеные данные")
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{"alerts": utils.Flashes(c)})
		return
	}

	session.Set(globals.Userkey, user.Name)
	session.Set(globals.Rolekey, user.Role.Name)

	if err := session.Save(); err != nil {
		utils.FlashMessage(c, "Ошибка во время сохранения сессии")
		c.HTML(http.StatusInternalServerError, "login.html", gin.H{"alerts": utils.Flashes(c)})
		return
	}

	c.Redirect(http.StatusSeeOther, "/")
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(globals.Userkey)

	log.Println("logging out user:", user)
	if user == nil {
		utils.FlashMessage(c, "Пользователь не найден")
		c.HTML(http.StatusInternalServerError, "login.html", gin.H{"alerts": utils.Flashes(c)})
		return
	}

	session.Clear()
	if err := session.Save(); err != nil {
		utils.FlashMessage(c, "Ошибка во время удаления сессии")
		c.HTML(http.StatusInternalServerError, "login.html", gin.H{"alerts": utils.Flashes(c)})
		return
	}

	c.Redirect(http.StatusSeeOther, "/auth/login")
}

func RegisterForm(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(globals.Userkey)
	c.HTML(http.StatusOK, "register.html", gin.H{"user": user, "alerts": utils.Flashes(c)})
}

func Register(c *gin.Context) {
	var creds RegisterCred

	if err := c.ShouldBind(&creds); err != nil {
		log.Println(err)
		utils.FlashMessage(c, "Неверные данные формы")
		c.HTML(http.StatusBadRequest, "register.html", gin.H{"alerts": utils.Flashes(c)})
		return
	}

	if creds.Password != creds.ConfPassword {
		utils.FlashMessage(c, "Введеные пароли не совпадают")
		c.HTML(http.StatusBadRequest, "register.html", gin.H{"alerts": utils.Flashes(c)})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.FlashMessage(c, "Ошибка сервера")
		c.HTML(http.StatusInternalServerError, "register.html", gin.H{"alerts": utils.Flashes(c)})
		log.Println(err)
		return
	}

	var userExists models.User
	if err := db.DB.Find(&userExists, "name = ?", creds.Username).Error; err != nil {
		log.Println(err)
		utils.FlashMessage(c, "Ошибка при запросе к безе данных")
		c.HTML(http.StatusOK, "register.html", gin.H{"alerts": utils.Flashes(c)})
		return
	}
	log.Printf("%+v", userExists)

	if userExists.ID != 0 {
		utils.FlashMessage(c, "Пользователь с таким именем уже существует")
		c.HTML(http.StatusOK, "register.html", gin.H{"alerts": utils.Flashes(c)})
		return
	}

	newUser := models.User{
		Name:     creds.Username,
		Password: string(hash),
		RoleID:   creds.Role,
	}
	if err := db.DB.Create(&newUser).Error; err != nil {
		log.Println(err)
		utils.FlashMessage(c, "Ошибка при запросе к безе данных")
		c.HTML(http.StatusInternalServerError, "register.html", gin.H{"alerts": utils.Flashes(c)})
		return
	}

	c.Redirect(http.StatusSeeOther, "/auth/login")
}
