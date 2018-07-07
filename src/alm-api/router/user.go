package rout

import (
	"github.com/gin-gonic/gin"
	"alm-api/middleWare/auth"
)

func user(r *gin.Engine) {
	rg := r.Group("/user")
	rg.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code":200,
			"msg":"",
			"user":"{a:2,b:3}",
		})
	})

	rg.Use(wm.AuthUserD())
	{
	}
}
