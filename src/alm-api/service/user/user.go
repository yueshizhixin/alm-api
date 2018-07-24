package userSV

import (
	"alm-api/model/user"
	"errors"
	"alm-api/global"
)

/*
	
 */

func Add(user *user.User) (bool, error) {
	if user.Acc == "" || user.Pwd == "" {
		return false, errors.New("账号或密码不能为空")
	}
	if user.SignType == 0 {
		return false, errors.New("请选择注册类型")
	}
	user.Init()
	if isNew := glb.DB.Create(&user).NewRecord(user); isNew {
		return false, glb.ERR_FAIL
	}
	return true, nil
}

func IsExist(user *user.User) bool {
	return true
}
