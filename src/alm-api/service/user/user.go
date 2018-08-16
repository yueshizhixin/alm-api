package userSV

import (
	"alm-api/model/user"
	"errors"
	"alm-api/global"
)

/*
	
 */

func Add(user *user.User) error {
	if user.Acc == "" || user.Pwd == "" {
		return errors.New("账号或密码不能为空")
	}
	if user.SignType == 0 {
		return errors.New("请选择注册类型")
	}
	user.Init()
	if isNew := glb.DB.Create(&user).NewRecord(user); isNew {
		return errors.New(glb.TIP_FAIL)
	}
	return nil
}

func IsExist(user *user.User) bool {
	return true
}
