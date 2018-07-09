package svUser

import (
	"alm-api/model/user"
	"errors"
	"time"
	"alm-api/global"
)

func Add(user *user.User) (bool, error) {
	if user.Acc == "" || user.Pwd == "" {
		return false, errors.New("账号或密码不能为空")
	}
	user.Uid = glb.UID()
	user.CreateTime = time.Now()
	user.LatestTime = time.Now()
	if isNew := glb.DB.Create(&user).NewRecord(user); isNew {
		return false, errors.New("操作失败")
	}
	return true, nil
}
