package rout

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions"
	"alm-api/config"
)

/**
	路由索引
 */

func InitRout()  {

	var r = gin.Default()
	r.Use(sessions.Sessions("user", cookie.NewStore([]byte(conf.SESSION_SECRET))))
	r.GET("/test", func(c *gin.Context) {
		session := sessions.Default(c)
		var count int
		v := session.Get("count")
		if v == nil {
			count = 0
		} else {
			count = v.(int)
			count++
		}
		session.Set("count", count)
		session.Save()
		c.JSON(200, gin.H{"count": count})
	})
	r.GET("/test2", func(c *gin.Context) {
		session := sessions.Default(c)
		session.Clear()
		session.Save()
		c.JSON(200, gin.H{})
	})


	userRout()
	noteRout()
}


