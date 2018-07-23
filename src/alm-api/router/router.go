package rout

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"alm-api/config"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/cors"
)

/*
	路由索引
 */

var (
	R  *gin.Engine
	RG *gin.RouterGroup
)

func init() {
	R = gin.Default()
	R.Use(sessions.Sessions(conf.SESSION_NAME, cookie.NewStore([]byte(conf.SESSION_SECRET))))
	R.Use(cors.Default())
	RG = R.Group("/api")
}

func InitRout() {
	userRout()
	noteRout()
}
