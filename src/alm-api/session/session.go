package session

import "github.com/gin-gonic/gin"

func clear() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		gin.Default()
	}
}

func save()  {
	gin.Default()
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