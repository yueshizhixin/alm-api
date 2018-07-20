package main

import (
	"alm-api/global"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions"
)

func main() {
	defer glb.DB.Close()
	gin.SetMode(gin.DebugMode)
	var r = gin.Default()
	store := cookie.NewStore([]byte("mysecret"))
	r.Use(sessions.Sessions("mysession", store))
	r.Use(sessions.Sessions("aaa", store))
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
	
	//rout.InitRout(r)
	r.Run(":80")
}
