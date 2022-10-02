package auth

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const tokenkey = "session_token"

var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

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
	var session = sessions.Default(c)
	err := c.ShouldBind(&creds)
	if err != nil {
		// render some html
		c.Error(err)
	}

	expectedPassword, ok := users[creds.Username]
	if !ok || expectedPassword != creds.Password {
		//render some html
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
