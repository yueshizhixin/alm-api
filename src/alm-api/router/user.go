package rout

import (
	"github.com/gin-gonic/gin"
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

	})
	rg.POST("", func(c *gin.Context) {
		var u user.User
		err = c.Bind(&u)
		if err != nil {
			glb.JSON500(c, err.Error(), nil)
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
