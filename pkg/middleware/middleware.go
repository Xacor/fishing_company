package middleware

import (
	"net/http"

	"github.com/Xacor/fishing_company/pkg/globals"
	"github.com/Xacor/fishing_company/pkg/utils"

	"github.com/casbin/casbin/v2"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func AuthRequired(c *gin.Context) {

	session := sessions.Default(c)
	user := session.Get(globals.Userkey)

	if user == nil {
		log.Info("User not logged in", c.Request.UserAgent())
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
			log.Error(err)
			c.Error(err)
			c.Abort()
			return
		}
		if !ok {
			log.Warnf("User %v trys to access privileged resource", session.Get(globals.Userkey))
			utils.FlashMessage(c, "У вас недостаточно прав на это действие")
			session.Save()
			c.Redirect(http.StatusSeeOther, "/")
			c.Abort()
			return
		}

		c.Next()
	}
}
