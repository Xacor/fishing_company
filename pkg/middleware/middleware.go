package middleware

import (
	"fishing_company/pkg/globals"
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

		ok, err := e.Enforce(role, c.Request.URL.Path, c.Request.Method)
		if err != nil {
			log.Println(err)
			c.Error(err)
			c.Abort()
			return
		}
		if !ok {
			log.Println("No privs, redirect")
			c.Redirect(http.StatusMovedPermanently, "/")
			//add flash messages
			c.Abort()
			return
		}

		c.Next()
	}
}
