package main

import (
	"alm-api/global"
	"github.com/gin-gonic/gin"
	"github.com/braintree/manners"
	"alm-api/router"
)

func main() {
	defer glb.DB.Close()
	gin.SetMode(gin.DebugMode)
	var r = gin.Default()
	rout.InitRouter(r)
	manners.ListenAndServe(":80", r)
}
