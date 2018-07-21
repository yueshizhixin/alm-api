package rout

import (
	"github.com/gin-gonic/gin"
	"strings"
	"alm-api/service/user"
	"alm-api/model/user"
	"alm-api/session"
)

/*

 */

func userRout() {
	var ok bool
	var err error

	rg := RG.Group("/user")
	rg.GET("/enum",)
	rg.GET("", func(ctx *gin.Context) {
		//var acc = strings.TrimSpace(ctx.PostForm("acc"))
		//var pwd = strings.TrimSpace(ctx.PostForm("pwd"))
		var u = user.User{
			Acc: strings.TrimSpace(ctx.DefaultQuery("acc", "1")),
			Pwd: strings.TrimSpace(ctx.DefaultQuery("pwd", "2")),
		}
		if ok, err = userSV.Add(&u); ok {
			session.Set(ctx, &u)
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
	
	rg.Use()
	{
	}
}
