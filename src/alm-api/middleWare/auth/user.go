package auth

import (
	"github.com/gin-gonic/gin"
	"alm-api/session"
)

/*
	
 */

func UserEnum() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
	}
}

func IsLogin(ctx *gin.Context) bool {
	if session.Get(ctx).Id > 0 {
		return true
	}
	return false
}
