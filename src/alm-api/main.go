package main

import (
	"alm-api/global"
	"github.com/gin-gonic/gin"
	"alm-api/router"
)

func main() {
	defer glb.DB.Close()
	gin.SetMode(gin.DebugMode)
	var r = gin.Default()
	rout.InitRouter(r)
	r.Run(":80")
}
