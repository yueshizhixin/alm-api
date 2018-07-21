package userSV

import (
	"alm-api/model/user"
	"errors"
	"time"
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
	user.Uid = glb.UUID()
	user.CreateTime = glb.If(user.CreateTime.IsZero(), time.Now(), user.CreateTime).(time.Time)
	user.LatestTime = glb.If(user.LatestTime.IsZero(), time.Now(), user.LatestTime).(time.Time)
	if isNew := glb.DB.Create(&user).NewRecord(user); isNew {
		return false, errors.New("操作失败")
	}
	return true, nil
}
