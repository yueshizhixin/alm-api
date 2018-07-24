package glb

import (
	"github.com/gin-gonic/gin"
)

/*
	JSON格式化
 */

func formatData(code int, msg interface{}, data interface{}) interface{} {
	if data == nil {
		data = "{}"
	}
	if msg == nil {
		msg = ""
	}
	return gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	}
}

func JSON200(ctx *gin.Context, msg interface{}, data interface{}) {
	ctx.JSON(200, formatData(200, msg, data))
}

func JSON500(ctx *gin.Context, msg interface{}, data interface{}) {
	ctx.JSON(200, formatData(500, msg, data))
}

func JSON404(ctx *gin.Context, msg interface{}, data interface{}) {
	ctx.JSON(200, formatData(404, msg, data))
}
