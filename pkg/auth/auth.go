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

	cookie, err := c.Cookie(tokenkey)
	if err != nil {
		cookie = "NotSet"
		c.SetCookie(tokenkey, sessionToken, 3600, "/", "localhost", false, false)
	}

	//replace with render some html
	c.JSON(http.StatusOK, gin.H{"message": "Successfully authenticated user", "cookie": cookie})
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
			return
		}

		c.Next()
	}
}

func Profile(c *gin.Context) {
	session := sessions.Default(c)
	token, _ := c.Cookie(tokenkey)
	user := session.Get(token)
	c.JSON(http.StatusOK, gin.H{"user": user})
}
