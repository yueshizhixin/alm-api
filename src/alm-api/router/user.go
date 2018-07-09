package rout

import (
	"github.com/gin-gonic/gin"
	"alm-api/midWare/auth"
	"alm-api/service/user"
	"strings"
	"alm-api/model/user"
)

func userRout(r *gin.Engine) {
	var ok bool
	var err error

	rg := r.Group("/user")
	rg.POST("", func(ctx *gin.Context) {
		var acc = strings.TrimSpace(ctx.PostForm("acc"))
		var pwd = strings.TrimSpace(ctx.PostForm("pwd"))
		var pwd2 = strings.TrimSpace(ctx.PostForm("pwd2"))
		if pwd != pwd2 {
			ctx.JSON(200, gin.H{

			})
		}
		var u = user.User{
			Acc: acc,
			Pwd: pwd,
		}
		if ok, err = svUser.Add(&u); ok {
			ctx.JSON(200, gin.H{

			})
		} else {
			ctx.JSON(200, gin.H{
				"code": 123,
				"msg":  err,
				"data": 222,
			})
		}
	})

	rg.Use(auth.AuthUserD())
	{
	}
}
