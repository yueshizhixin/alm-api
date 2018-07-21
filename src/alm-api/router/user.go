package rout

import (
	"github.com/gin-gonic/gin"
	"strings"
	"alm-api/service/user"
	"alm-api/model/user"
	"alm-api/session"
	"alm-api/global"
)

/*

 */

func userRout() {
	var ok bool
	var err error

	rg := RG.Group("/user")
	rg.GET("/enum", func(c *gin.Context) {
		glb.JSON200(c, nil, nil)
		return
	})
	rg.GET("", func(c *gin.Context) {
		var u = user.User{
			Acc: strings.TrimSpace(c.DefaultQuery("acc", "1")),
			Pwd: strings.TrimSpace(c.DefaultQuery("pwd", "2")),
		}
		if ok, err = userSV.Add(&u); ok {
			session.Set(c, &u)
			glb.JSON200(c, "操作成功", nil)
			return
		} else {
			glb.JSON500(c, err.Error(), nil)
			return
		}
	})
	rg.POST("", func(c *gin.Context) {
		
	})

	rg.Use()
	{
	}
}
