package rout

import (
	"github.com/gin-gonic/gin"
	"alm-api/midWare/auth"
	"alm-api/service/user"
	"strings"
	"alm-api/model/user"
)

func userRout() {
	var ok bool
	var err error
	var r = gin.Default()
	rg := r.Group("/user")
	rg.GET("", func(ctx *gin.Context) {
		//var acc = strings.TrimSpace(ctx.PostForm("acc"))
		//var pwd = strings.TrimSpace(ctx.PostForm("pwd"))
		var u = user.User{
			Acc: strings.TrimSpace(ctx.DefaultQuery("acc", "")),
			Pwd: strings.TrimSpace(ctx.DefaultQuery("pwd", "")),
		}
		if ok, err = userSV.Add(&u); ok {
			ctx.JSON(200, gin.H{
				"code": "200",
				"msg":  "操作成功",
				"data": "{}",
			})
			//return  
		} else {
			ctx.JSON(200, gin.H{
				"code": "500",
				"msg":  err.Error(),
				"data": "{}",
			})
		}
	})

	rg.Use(auth.AuthUserD())
	{
	}
}
