package rout

import (
	"github.com/gin-gonic/gin"
	"alm-api/service/user"
	"alm-api/model/user"
	"alm-api/session"
	"alm-api/global"
	"alm-api/middleWare/auth"
	"fmt"
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
		var u user.User
		//断电
		fmt.Println(u)
		if c.Bind(&u) != nil {
			glb.JSON500(c, glb.TIP_DATA_BIND_FAIL, nil)
			return
		}
		fmt.Println(u)
		if u.Id > 0 {
			glb.JSON200(c, nil, gin.H{
				"user": glb.DB.Take(&u),
			})
		} else {
			if auth.IsLogin(c) {
				glb.JSON200(c, nil, gin.H{
					"user": session.Get(c),
				})
				return
			} else {
				glb.JSON404(c, glb.TIP_NOT_EXIST, nil)
			}
		}
	})
	//注册
	rg.POST("", func(c *gin.Context) {
		var u user.User
		if c.Bind(&u) != nil {
			glb.JSON500(c, glb.TIP_DATA_BIND_FAIL, nil)
			return
		}
		if ok, err = userSV.Add(&u); ok {
			session.Set(c, &u)
			glb.JSON200(c, glb.TIP_SUCCESS, nil)
			return
		} else {
			glb.JSON500(c, err.Error(), nil)
			return
		}
	})
	//登录
	rg.POST("signIn", func(c *gin.Context) {
		var u user.User
		if c.Bind(&u) != nil {
			glb.JSON500(c, glb.TIP_DATA_BIND_FAIL, nil)
			return
		}
		if ok = userSV.IsExist(&u); ok {
			session.Set(c, &u)
			glb.JSON200(c, glb.TIP_SUCCESS, nil)
			return
		} else {
			glb.JSON404(c, glb.TIP_NOT_EXIST, nil)
			return
		}
	})

	rg.Use()
	{
	}
}
