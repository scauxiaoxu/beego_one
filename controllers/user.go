package controllers

import (
	"apiproject/models"
	"github.com/astaxie/beego"
	"github.com/scauxiaoxu/gotool/commonmark"
	"regexp"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// 注册方法
// @router /register/save [post]
func (this *UserController) SaveRegister() {
	var (
		mobile   string
		password string
	)
	mobile = this.GetString("mobile")
	password = this.GetString("password")

	if mobile == "" {
		this.Data["json"] = ReturnError(4001, "手机号不能为空")
		this.ServeJSON()
		this.StopRun()
	}
	if password == "" {
		this.Data["json"] = ReturnError(4001, "密码不能为空")
		this.ServeJSON()
		this.StopRun()
	}
	isorno, _ := regexp.MatchString(`^1(3|4|5|7|8)[0-9]\d{8}$`, mobile)
	if !isorno {
		this.Data["json"] = ReturnError(4002, "手机号格式不正确")
		this.ServeJSON()
		this.StopRun()
	}
	// 判断手机号码是否被注册
	status := models.IsUserMobile(mobile)
	if status {
		this.Data["json"] = ReturnError(4001, "手机号已经注册")
		this.ServeJSON()
		this.StopRun()
	} else {
		md5Pass := commonmark.Md5(mobile + KEY)
		err := models.UserSave(mobile, md5Pass)
		if err == nil {
			this.Data["json"] = ReturnSuccess(0, "注册成功", nil, 0)
			this.ServeJSON()
		} else {
			this.Data["json"] = ReturnError(5000, err)
			this.ServeJSON()
			this.StopRun()
		}
	}

}

//用户登录
// @router /login/do [*]
func (this *UserController) LoginDo() {
	var (
		mobile   string
		password string
	)
	mobile = this.GetString("mobile")
	password = this.GetString("password")

	if mobile == "" {
		this.Data["json"] = ReturnError(4001, "手机号不能为空")
		this.ServeJSON()
		this.StopRun()
	}
	if password == "" {
		this.Data["json"] = ReturnError(4001, "密码不能为空")
		this.ServeJSON()
		this.StopRun()
	}
	isorno, _ := regexp.MatchString(`^1(3|4|5|7|8)[0-9]\d{8}$`, mobile)
	if !isorno {
		this.Data["json"] = ReturnError(4002, "手机号格式不正确")
		this.ServeJSON()
		this.StopRun()
	}

	id, name := models.IsMobileLogin(mobile, commonmark.Md5(password+KEY))
	if id == 0 {
		this.Data["json"] = ReturnError(5000, "手机号或密码不正确")
		this.ServeJSON()
		this.StopRun()
	}

	this.Data["json"] = ReturnSuccess(0, "ok", map[string]interface{}{"uid": id, "name": name}, 0)
	this.ServeJSON()
}

//----------------------------------------------
