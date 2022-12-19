package middleware

import (
	"fishing_company/pkg/globals"
	"fishing_company/pkg/utils"
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
		utils.FlashMessage(c, "Для этого действия необходима аутентификация")
		session.Save()
		c.Redirect(http.StatusSeeOther, "/auth/login")
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
			utils.FlashMessage(c, "У вас недостаточно прав на это действие")
			session.Save()
			c.Redirect(http.StatusSeeOther, "/")
			c.Abort()
			return
		}

		c.Next()
	}
}
