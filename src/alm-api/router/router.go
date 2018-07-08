package rout

import "github.com/gin-gonic/gin"

/**
	路由索引
 */

func InitRouter(r * gin.Engine)  {
	user(r)
}


