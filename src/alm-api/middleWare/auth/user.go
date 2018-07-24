package auth

import (
	"github.com/gin-gonic/gin"
	"alm-api/session"
	"alm-api/model/user"
)

/*
	
 */

 
func UserEnum() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
	}
}

func IsLogin(ctx *gin.Context) (bool, user.User) {
	if u := session.Get(ctx); u.Id > 0 {
		return true, u
	}
	return false, user.User{}
}
