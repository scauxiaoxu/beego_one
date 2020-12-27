package models

import (
	"github.com/astaxie/beego/orm"
	gotooltime "github.com/scauxiaoxu/gotool/time"
)

type User struct {
	Id       int
	Name     string
	Password string
	AddTime  int64
	Status   int
	Mobile   string
	Avatar   string
}

func init() {
	// register model
	orm.RegisterModel(new(User))
}

//根据手机号判断用户是否存在
func IsUserMobile(mobile string) bool {
	o := orm.NewOrm()
	user := User{Mobile: mobile}
	err := o.Read(&user, "Mobile")
	if err == orm.ErrNoRows {
		return false
	} else if err == orm.ErrMissPK {
		return false
	}
	return true
}

//保存用户
func UserSave(mobile, password string) error {
	o := orm.NewOrm()
	var user User
	user.Name = ""
	user.Password = password
	user.Mobile = mobile
	user.Status = 1
	user.AddTime = gotooltime.Time()
	_, err := o.Insert(&user)
	return err
}

// 判断用户是否注册
func IsMobileLogin(mobile, password string) (thisid int, thisname string) {
	o := orm.NewOrm()
	var user User

	err := o.QueryTable("user").Filter("mobile", mobile).Filter("password", password).One(&user)
	if err == orm.ErrNoRows {
		return
	} else if err == orm.ErrMissPK {
		return
	}
	return user.Id, user.Name
}

//根据用户ID获取用户信息
func GetUserInfo(uid int) (User, error) {
	o := orm.NewOrm()
	u := User{Id: uid}
	err := o.Read(&u)
	return u, err
}
