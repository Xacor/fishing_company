package auth

import (
	"fishing_company/pkg/db"
	"fishing_company/pkg/models"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

const tokenkey = "session_token"

type Credentials struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func LoginForm(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

// Session data stored in memory, but session token stored in clients cookies
func Login(c *gin.Context) {
	var creds Credentials

	session := sessions.Default(c)
	err := c.ShouldBind(&creds)
	if err != nil {
		// render some html
		c.Error(err)
	}

	var user models.User
	result := db.DB.Where("name = ?", creds.Username).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{"error": "invalid credentials"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": "invalid credentials, password dont match"})
		return
	}

	sessionToken := uuid.NewString()
	session.Set(sessionToken, creds.Username)

	err = session.Save()
	if err != nil {
		c.Error(err)
	}

	c.SetCookie(tokenkey, sessionToken, 3600, "/", "localhost", false, false)

	//replace with render some html
	c.JSON(http.StatusOK, gin.H{"message": "Successfully authenticated user"})
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	if err := session.Save(); err != nil {
		//change it
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		token, err := c.Cookie(tokenkey)
		if err != nil {
			c.Error(err)
		}

		user := session.Get(token)
		if user == nil {
			// render html or flash message???
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authorized"})
			c.Abort()
		}
	}
}

func Profile(c *gin.Context) {
	session := sessions.Default(c)
	token, _ := c.Cookie(tokenkey)
	user := session.Get(token)
	c.JSON(http.StatusOK, gin.H{"user": user})
}

// Обновляет время действия токена при каждмом действии пользователя.
// Возможно надо переделать по таймауту или как то еще
func TokenTimeoutRefresh() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie(tokenkey)
		if err != nil {
			c.Error(err)

		}
		c.SetCookie(tokenkey, token, 3600, "/", "localhost", false, false)
	}
}

type RegisterCred struct {
	Username string `form:"username"`
	Pwd      string `form:"password"`
	ConfPwd  string `form:"conf-password"`
}

func RegisterForm(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{})
}

func Register(c *gin.Context) {
	var creds RegisterCred
	err := c.ShouldBind(&creds)
	if err != nil {
		c.Error(err)
	}

	if creds.Pwd == creds.ConfPwd {
		hash, e := bcrypt.GenerateFromPassword([]byte(creds.Pwd), bcrypt.DefaultCost)
		if e != nil {
			c.Error(e)
		}

		newUser := models.User{
			Name:     creds.Username,
			Password: string(hash),
		}

		//проверка на существование пользователя с таким именем
		//полученная модель опускается, важен только reult
		result := db.DB.Where(&models.User{Name: newUser.Name}).First(&models.User{})
		if result.Error == nil {
			c.JSON(http.StatusOK, gin.H{"error": "user with this name already exist"})
			return
		}

		if result = db.DB.Create(&newUser); result.Error != nil {
			c.AbortWithError(http.StatusNotFound, result.Error)
			return
		}
		c.Redirect(http.StatusMovedPermanently, "/login")

	} else {
		c.JSON(http.StatusOK, gin.H{"error": "passwords dont match"})
	}

}
