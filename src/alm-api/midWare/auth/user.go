package auth

import "github.com/gin-gonic/gin"

/*
	
 */

 
func UserEnum() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
	}
}
