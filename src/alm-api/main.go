package main

import (
	"alm-api/global"
	"github.com/gin-gonic/gin"
	"alm-api/router"
)

func main() {
	defer glb.DB.Close()
	gin.SetMode(gin.DebugMode)
	rout.InitRout()
	gin.Default().Run(":80")
}
