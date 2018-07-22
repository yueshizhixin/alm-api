package session

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"alm-api/model/user"
)

/*
	优化
		1.反射优化配置
 */

func Get(ctx *gin.Context) user.User {
	s := sessions.Default(ctx)
	return user.User{
		Id:       s.Get("id").(string),
		Acc:      s.Get("acc").(string),
		Phone:    s.Get("phone").(string),
		Email:    s.Get("email").(string),
		UserName: s.Get("userName").(string),
	}
}

func Set(ctx *gin.Context, user *user.User) {
	s := sessions.Default(ctx)
	s.Set("id", &user.Id)
	s.Set("acc", &user.Acc)
	s.Set("phone", &user.Phone)
	s.Set("email", &user.Email)
	s.Set("userName", &user.UserName)
	s.Save()
}

func Clear(ctx *gin.Context) {
	s := sessions.Default(ctx)
	s.Clear()
	s.Save()
}
