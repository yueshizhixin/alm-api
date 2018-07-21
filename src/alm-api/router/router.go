package rout

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"alm-api/config"
	"github.com/gin-gonic/gin"
)

/**
	路由索引
 */

var (
	R *gin.Engine
)

func init()  {
	R=gin.Default()
}

func InitRout() {
	R.Use(sessions.Sessions("alm_api", cookie.NewStore([]byte(conf.SESSION_SECRET))))
	userRout()
	noteRout()
}
