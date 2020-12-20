package controllers

import (
	"apiproject/models"
	"github.com/astaxie/beego"
)

// Operations about Users
type Basecontroller struct {
	beego.Controller
}

// 获取频道下地区
// @router /channel/region [*]
func (this *Basecontroller) ChannelRegion() {
	channelId, _ := this.GetInt("channelId")
	if channelId == 0 {
		this.Data["json"] = ReturnError(4001, "必须指定频道")
		this.ServeJSON()
		this.StopRun()
	}
	num, regions, err := models.GetChannelRegion(channelId)

	if err != nil {
		this.Data["json"] = ReturnError(4004, "查询异常")
		this.ServeJSON()
	} else if num == 0 {
		this.Data["json"] = ReturnEmpty(0, []string{}, num)
		this.ServeJSON()
	} else {
		this.Data["json"] = ReturnSuccess(0, "success", regions, num)
		this.ServeJSON()
	}
}

// 获取频道下类型
// @router /channel/type [*]
func (this *Basecontroller) ChannelType() {
	channelId, _ := this.GetInt("channelId")
	if channelId == 0 {
		this.Data["json"] = ReturnError(4001, "必须指定频道")
		this.ServeJSON()
		this.StopRun()
	}
	num, types, err := models.GetChannelType(channelId)

	if err != nil {
		this.Data["json"] = ReturnError(4004, "查询异常")
		this.ServeJSON()
	} else if num == 0 {
		this.Data["json"] = ReturnEmpty(0, []string{}, num)
		this.ServeJSON()
	} else {
		this.Data["json"] = ReturnSuccess(0, "success", types, num)
		this.ServeJSON()
	}
}
