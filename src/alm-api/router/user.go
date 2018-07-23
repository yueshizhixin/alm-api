package rout

import (
	"github.com/gin-gonic/gin"
	"alm-api/service/user"
	"alm-api/model/user"
	"alm-api/session"
	"alm-api/global"
	"fmt"
	"errors"
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
		glb.JSON200(c, "访问成功", nil)
		return
	})
	rg.POST("", func(c *gin.Context) {
		var u user.User
		fmt.Println(glb.ERR_FAIL)
		if c.Bind(&u) != nil {
			glb.JSON500(c, errors.New(glb.TIP_FAIL), nil)
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

	rg.Use()
	{
	}
}
