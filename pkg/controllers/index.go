package controllers

import (
	"net/http"

	"github.com/Xacor/fishing_company/pkg/globals"
	"github.com/Xacor/fishing_company/pkg/utils"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(globals.Userkey)
	c.HTML(http.StatusOK, "base.html", gin.H{
		"user":   user,
		"alerts": utils.Flashes(c),
	})
}
