package user

import "time"

/**
	@url    `https://github.com/yueshizhixin/go-orm-generator`
	@author `埃罗芒老兄(https://yueshizhixin.github.io/)`
	@scheme `alm`
	@table  `t_user`
	@time   `2018-07-09 17:35:04`
 */
type User struct {
	Id         uint32    `json:"-" gorm:"primary_key;"`
	Uid        string    `json:"uid" gorm:"column:uid;type:char(32);"`                /*用户标识*/
	Acc        string    `json:"acc" gorm:"column:acc;type:varchar(16);"`             /*注册账号*/
	Phone      string    `json:"phone" gorm:"column:phone;type:varchar(32);"`         /*手机号*/
	Email      string    `json:"email" gorm:"column:email;type:varchar(64);"`         /*电子邮箱*/
	SignType   int8      `json:"signType" gorm:"column:signType;type:tinyint(4);"`    /*注册方式*/
	Pwd        string    `json:"pwd" gorm:"column:pwd;type:varchar(16);"`             /*密码*/
	UserName   string    `json:"userName" gorm:"column:userName;type:varchar(64);"`   /*用户姓名*/
	Sex        int8      `json:"sex" gorm:"column:sex;type:tinyint(4);default:'-1';"` /*性别*/
	Brithday   string    `json:"brithday" gorm:"column:brithday;type:varchar(16);"`   /*出生日期*/
	VipLv      int8      `json:"vipLv" gorm:"column:vipLv;type:tinyint(4);"`          /*会员级别*/
	CreateTime time.Time `json:"createTime" gorm:"column:createTime;type:datetime;"`  /*创建时间*/
	LatestTime time.Time `json:"latestTime" gorm:"column:latestTime;type:datetime;"`  /*最后登录时间*/
}

func (User) TableName() string {
	return "t_user"
}
