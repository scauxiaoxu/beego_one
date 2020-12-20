package controllers

import (
	"apiproject/models"
	"fmt"
	"github.com/astaxie/beego"
)

type VedioController struct {
	beego.Controller
}

// 频道页 顶部广告
// @router /channel/advert [get]
func (this *VedioController) ChannelAdvert() {
	channelId, _ := this.GetInt("channelId")
	if channelId == 0 {
		this.Data["json"] = ReturnError(4001, "必须指定频道")
		this.ServeJSON()
		this.StopRun()
	}
	num, err, re := models.GetChannelAdvert(channelId)
	if err != nil {
		fmt.Printf("%T --> %v\n", err, err)
		this.Data["json"] = ReturnError(4001, err)
		this.ServeJSON()
		this.StopRun()
	}
	this.Data["json"] = ReturnSuccess(0, "success", re, num)
	this.ServeJSON()
}

// 频道页 - 获取正在热播
// @router /channel/hot [get]
func (this *VedioController) ChannelHotList() {
	channelId, _ := this.GetInt("channelId")
	if channelId == 0 {
		this.Data["json"] = ReturnError(4001, "必须指定频道")
		this.ServeJSON()
		this.StopRun()
	}

	num, videos, err := models.GetChannelHotList(channelId)

	if err != nil {
		fmt.Printf("%T --> %v\n", err, err)
		this.Data["json"] = ReturnError(4001, err.Error())
		this.ServeJSON()
		this.StopRun()
	} else if num == 0 {
		this.Data["json"] = ReturnError(4001, "没有相关内容")
		this.ServeJSON()
		this.StopRun()
	}

	this.Data["json"] = ReturnSuccess(0, "ok", videos, num)
	this.ServeJSON()
}

// 根据频道地区获取 推荐视频
// @router /channel/recommend/region [get]
func (this *VedioController) ChannelRegionRecommendList() {
	channelId, _ := this.GetInt("channelId")
	if channelId == 0 {
		this.Data["json"] = ReturnError(4001, "必须指定频道")
		this.ServeJSON()
		this.StopRun()
	}
	regionId, _ := this.GetInt("regionId")
	if regionId == 0 {
		this.Data["json"] = ReturnError(4001, "必须指定频道地区")
		this.ServeJSON()
		this.StopRun()
	}

	num, videos, err := models.GetChannelRegionRecommend(channelId, regionId)

	if err != nil {
		this.Data["json"] = ReturnError(4004, "没有相关内容")
		this.ServeJSON()
	} else if num == 0 {
		fmt.Printf("000 %T --> %v\n", videos, videos)
		this.Data["json"] = ReturnSuccess(0, "success", []string{}, num)
		this.ServeJSON()
	} else {
		this.Data["json"] = ReturnSuccess(0, "success", videos, num)
		this.ServeJSON()
	}
}

// 根据类型获取推荐视频
// @router /channel/recommend/type [get]
func (this *VedioController) ChannelTypeRecommendList() {
	channelId, _ := this.GetInt("channelId")
	if channelId == 0 {
		this.Data["json"] = ReturnError(4001, "必须指定频道")
		this.ServeJSON()
		this.StopRun()
	}

	typeId, _ := this.GetInt("typeId")
	if typeId == 0 {
		this.Data["json"] = ReturnError(4001, "必须指定频道类型")
		this.ServeJSON()
		this.StopRun()
	}

	num, videos, err := models.GetChannelTypeRecommend(channelId, typeId)

	if err != nil {
		this.Data["json"] = ReturnError(4004, "没有相关内容")
		this.ServeJSON()
	} else if num == 0 {
		fmt.Printf("000 %T --> %v\n", videos, videos)
		this.Data["json"] = ReturnSuccess(0, "success", []string{}, num)
		this.ServeJSON()
	} else {
		this.Data["json"] = ReturnSuccess(0, "success", videos, num)
		this.ServeJSON()
	}
}

// 根据传入的参数获取视频列表
// @router /channel/video [*]
func (this *VedioController) ChannelVideo() {
	channelId, _ := this.GetInt("channelId")
	if channelId == 0 {
		this.Data["json"] = ReturnError(4001, "必须指定频道")
		this.ServeJSON()
		this.StopRun()
	}
	regionId, _ := this.GetInt("regionId")
	typeId, _ := this.GetInt("typeId")

	end := this.GetString("end")
	sort := this.GetString("sort")

	offset, _ := this.GetInt("offset")
	limit, _ := this.GetInt("limit")

	if limit == 0 {
		limit = 12 //默认显示12条
	}

	num, list, err := models.GetChannelVideoList(channelId, regionId, typeId, end, sort, offset, limit)

	fmt.Println(list)

	if err != nil {
		this.Data["json"] = ReturnError(4004, "查询异常")
		this.ServeJSON()
	} else if num == 0 {
		this.Data["json"] = ReturnEmpty(0, []string{}, num)
		this.ServeJSON()
	} else {
		this.Data["json"] = ReturnSuccess(0, "success", list, num)
		this.ServeJSON()
	}
}

// 获取视频info
// @router /video/info [*]
func (this *VedioController) VideoInfo() {
	videoId, err := this.GetInt("videoId")
	if err != nil {
		this.Data["json"] = ReturnError(4001, "参数异常")
		this.ServeJSON()
		this.StopRun()
	} else if videoId == 0 {
		this.Data["json"] = ReturnError(4001, "必须指定视频id")
		this.ServeJSON()
		this.StopRun()
	}

	video, err := models.GetVideoInfo(videoId)
	if err != nil {
		fmt.Println("获取视频信息异常:", err)
		this.Data["json"] = ReturnError(4004, "查询异常")
		this.ServeJSON()
	}
	this.Data["json"] = ReturnSuccess(0, "success", video, 1)
	this.ServeJSON()
}

//获取视频剧集列表
// @router /video/episodes/list [*]
func (this *VedioController) VideoEpisodesList() {
	videoId, _ := this.GetInt("videoId")
	if videoId == 0 {
		this.Data["json"] = ReturnError(4001, "必须指定视频id")
		this.ServeJSON()
		this.StopRun()
	}
	num, episodes, err := models.GetVideoEpisodesList(videoId)
	if err != nil {
		fmt.Println("获取视频列表信息异常:", err)
		this.Data["json"] = ReturnError(4004, "查询异常")
		this.ServeJSON()
	}

	this.Data["json"] = ReturnSuccess(0, "success", episodes, num)
	this.ServeJSON()
}
