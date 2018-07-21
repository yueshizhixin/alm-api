package session

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"alm-api/model/user"
)

/**
	优化
		1.反射优化配置
 */

func Get(ctx *gin.Context) user.User {
	s := sessions.Default(ctx)
	return user.User{
		Acc: s.Get("acc").(string),
		Pwd: s.Get("pwd").(string),
	}

}

func Set(ctx *gin.Context, user *user.User) {
	s := sessions.Default(ctx)
	s.Set("acc", &user.Acc)
	s.Set("pwd", &user.Pwd)
	s.Save()
}

func Clear(ctx *gin.Context) {
	s := sessions.Default(ctx)
	s.Clear()
	s.Save()
}
