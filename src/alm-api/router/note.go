package rout

import (
	"github.com/gin-gonic/gin"
)

/*

 */

func noteRout() {
	rg := RG.Group("/note")
	rg.GET("", func(c *gin.Context) {
	})
	rg.POST("", func(c *gin.Context) {
	})

	rg.Use()
	{
	}
}
