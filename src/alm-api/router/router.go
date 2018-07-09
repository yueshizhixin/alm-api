package rout

import "github.com/gin-gonic/gin"

/**
	路由索引
 */

func InitRout(r * gin.Engine)  {
	userRout(r)
	noteRout(r)
}


