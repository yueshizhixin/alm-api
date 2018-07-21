package rout

import (
	"github.com/gin-gonic/gin"
	"alm-api/midWare/auth"
)

func noteRout() {
	rg := R.Group("/note")
	rg.GET("", func(c *gin.Context) {
	})
	rg.POST("", func(c *gin.Context) {
	})

	rg.Use(auth.AuthUserD())
	{
	}
}
