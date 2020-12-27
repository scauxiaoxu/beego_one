package controllers

import (
	"apiproject/models"
	"github.com/astaxie/beego"
)

type TopController struct {
	beego.Controller
}

type Topjson struct {
	Id            int    `json:"id"`
	Title         string `json:"title"`
	SubTitle      string `json:"sub_title"`
	Img           string `json:"img"`
	Img1          string `json:"img1"`
	AddTime       int64  `json:"add_time"`
	EpisodesCount int    `json:"episodes_count"`
	IsEnd         int    `json:"is_end"`
}

// 频道排行榜
// @router /channel/top [get]
func (this *TopController) ChannelTop() {
	channelId, _ := this.GetInt("channelId")
	if channelId == 0 {
		this.Data["json"] = ReturnError(4001, "必须指定频道id")
		this.ServeJSON()
		this.StopRun()
	}
	// 获取10条记录
	var limit = 10
	_, vedios, err := models.GetChannelTop(channelId, limit)
	if err != nil {
		this.Data["json"] = ReturnError(4001, "查询异常")
		this.ServeJSON()
		this.StopRun()
	}

	var data []Topjson
	var item Topjson
	for _, v := range vedios {
		item.Id = v.Id
		item.Title = v.Title
		item.SubTitle = v.SubTitle
		item.Img = v.Img
		item.Img1 = v.Img1
		item.AddTime = v.AddTime
		item.EpisodesCount = v.EpisodesCount
		item.IsEnd = v.IsEnd
		data = append(data, item)
	}
	num := int64(len(data))
	if num == 0 {
		this.Data["json"] = ReturnSuccess(0, "success", []int{}, num)
	} else {
		this.Data["json"] = ReturnSuccess(0, "success", data, num)
	}
	this.ServeJSON()
}

// 根据类型获取排行榜
// @router /type/top [get]
func (this *TopController) TypeTop() {
	typeId, _ := this.GetInt("typeId")
	if typeId == 0 {
		this.Data["json"] = ReturnError(4001, "必须指定类型")
		this.ServeJSON()
		this.StopRun()
	}
	// 获取10条记录
	var limit = 10
	_, vedios, err := models.GetTypeTop(typeId, limit)
	if err != nil {
		this.Data["json"] = ReturnError(4001, "查询异常")
		this.ServeJSON()
		this.StopRun()
	}

	var data []Topjson
	var item Topjson
	for _, v := range vedios {
		item.Id = v.Id
		item.Title = v.Title
		item.SubTitle = v.SubTitle
		item.Img = v.Img
		item.Img1 = v.Img1
		item.AddTime = v.AddTime
		item.EpisodesCount = v.EpisodesCount
		item.IsEnd = v.IsEnd
		data = append(data, item)
	}
	num := int64(len(data))
	if num == 0 {
		this.Data["json"] = ReturnSuccess(0, "success", []int{}, num)
	} else {
		this.Data["json"] = ReturnSuccess(0, "success", data, num)
	}
	this.ServeJSON()
}
