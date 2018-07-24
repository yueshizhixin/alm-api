package session

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"alm-api/model/user"
	"alm-api/global"
)

/*
	优化
		1.反射优化配置
 */

func Get(ctx *gin.Context) user.User {
	s := sessions.Default(ctx)
	return user.User{
		Id:       glb.If(s.Get("id")==nil,uint32(0),s.Get("id")).(uint32),
		Acc:      glb.If(s.Get("acc")==nil,"",s.Get("acc")).(string),
		Phone:    glb.If(s.Get("phone")==nil,"",s.Get("phone")).(string),
		Email:    glb.If(s.Get("email")==nil,"",s.Get("email")).(string),
		UserName: glb.If(s.Get("userName")==nil, "",s.Get("userName")).(string),
	}
}

func Set(ctx *gin.Context, user *user.User) {
	s := sessions.Default(ctx)
	s.Set("id", user.Id)
	s.Set("acc", user.Acc)
	s.Set("phone", user.Phone)
	s.Set("email", user.Email)
	s.Set("userName", user.UserName)
	s.Save()
}

func Clear(ctx *gin.Context) {
	s := sessions.Default(ctx)
	s.Clear()
	s.Save()
}
