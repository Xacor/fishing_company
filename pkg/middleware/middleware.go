package middleware

import (
	"net/http"
	"time"

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
			log.Warnf("User %v tries to access privileged resource", session.Get(globals.Userkey))
			utils.FlashMessage(c, "У вас недостаточно прав на это действие")
			session.Save()
			c.Redirect(http.StatusSeeOther, "/")
			c.Abort()
			return
		}

		c.Next()
	}
}

func Logger(c *gin.Context) {
	logClient := log.New()

	start := time.Now()

	// process request
	c.Next()

	// End Time
	end := time.Now()

	//execution time
	latency := end.Sub(start)

	logClient.WithFields(log.Fields{
		"ip":      c.ClientIP(),
		"method":  c.Request.Method,
		"path":    c.Request.URL.Path,
		"code":    c.Writer.Status(),
		"latency": latency,
		"agent":   c.Request.UserAgent(),
	}).Info()
}
