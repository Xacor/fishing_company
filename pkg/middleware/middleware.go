package middleware

import (
	"fishing_company/pkg/globals"
	"fmt"
	"log"
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(globals.Userkey)
	if user == nil {
		log.Println("User not logged in")
		c.Redirect(http.StatusMovedPermanently, "/auth/login")
		c.Abort()
		return
	}
	c.Next()
}

func Authorization(e *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		role := session.Get(globals.Rolekey)
		user := session.Get(globals.Userkey)
		log.Println(user, role)

		log.Println("Authorization middleware")
		ok, err := e.Enforce(role, c.Request.URL.Path, c.Request.Method)
		if err != nil {
			log.Println(err)
			c.Error(err)
			c.Abort()
			return
		}
		if !ok {
			url := session.Get(globals.LastUrlkey)
			log.Println("No privs, redirect")
			c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("%v", url))
			c.Abort()
			return
		}

		c.Next()
	}
}

func SetReferer(c *gin.Context) {
	session := sessions.Default(c)
	session.Set(globals.LastUrlkey, c.Request.URL.Path)
	if err := session.Save(); err != nil {
		c.Error(err)
	}

	c.Next()

}
