package user

import (
	"time"
	"alm-api/global"
	"fmt"
)

/*
	@url    : https://github.com/yueshizhixin/go-orm-generator 
	@author : yueshizhixin(https://yueshizhixin.github.io/) 
	@scheme : alm 
	@table  : t_user 
	@time   : 2018-07-23 21:19:00 
 */
type User struct {
	Id         uint32    `json:"id" gorm:"primary_key;"`                                   /*用户标识*/
	Acc        string    `json:"acc" gorm:"column:acc;type:varchar(16);"`                  /*注册账号*/
	Phone      string    `json:"phone" gorm:"column:phone;type:varchar(32);"`              /*手机号*/
	Email      string    `json:"email" gorm:"column:email;type:varchar(64);"`              /*电子邮箱*/
	SignType   int8      `json:"signType" gorm:"column:signType;type:tinyint(4);"`         /*注册方式*/
	Pwd        string    `json:"-" gorm:"column:pwd;type:varchar(16);"`                  /*密码*/
	UserName   string    `json:"userName" gorm:"column:userName;type:varchar(64);"`        /*用户姓名*/
	Sex        int8      `json:"sex" gorm:"column:sex;type:tinyint(4);"`                   /*性别*/
	Brithday   string    `json:"brithday" gorm:"column:brithday;type:varchar(16);"`        /*出生日期*/
	VipLv      int8      `json:"vipLv" gorm:"column:vipLv;type:tinyint(4);"`               /*会员级别*/
	CreateTime time.Time `json:"createTime" gorm:"column:createTime;type:datetime;"`       /*创建时间*/
	LatestTime time.Time `json:"latestTime" gorm:"column:latestTime;type:datetime;"`       /*最后登录时间*/
	IsSign     int8      `json:"isSign" gorm:"column:isSign;type:tinyint(4);default:'1';"` /*是否允许登录*/
}

func (User) TableName() string {
	return "t_user"
}

func (u *User) Init() {
	fmt.Println("user init")
	u.CreateTime = glb.If(u.CreateTime.IsZero(), time.Now(), u.CreateTime).(time.Time)
	u.LatestTime = glb.If(u.LatestTime.IsZero(), time.Now(), u.LatestTime).(time.Time)
}
