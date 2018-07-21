package session

import (
	"github.com/gin-gonic/gin"
)

func clear()  {
	
}

func save()  {
}

func AuthUserD() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if 1 == 1 {
			ctx.Next()
		} else {
			ctx.JSON(200, gin.H{
				"code": 403,
				"msg":  "访问无权限",
			})
			ctx.Abort()
		}
	}
}