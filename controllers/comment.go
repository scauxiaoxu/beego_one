package controllers

import (
	"apiproject/models"
	"github.com/astaxie/beego"
)

type CommentController struct {
	beego.Controller
}

// 返回评论接口提
type ListStruct struct {
	Id          int64  `json:"id"`
	Content     string `json:"content"`
	AddTime     int64  `json:"add_time"`
	UserId      int    `json:"user_id"`
	Status      int    `json:"status"`
	Stamp       int    `json:"stamp"`
	PraiseCount int    `json:"praise_count"`
	EpisodesId  int    `json:"episodes_id"`
	VideoId     int    `json:"video_id"` //所属视频
	UserInfo    struct {
		Id      int    `json:"id"`
		Name    string `json:"name"`
		AddTime int64  `json:"add_time"`
		Avatar  string `json:"avatar"`
	}
	AddTimeTitle   string `json:"addTimeTitle"`
	PraiseCountNew int    `json:"praiseCount"`
}

// 获取评论列表
// @router /comment/list [*]
func (this *CommentController) List() {
	episodesId, _ := this.GetInt("episodesId")
	limit, _ := this.GetInt("limit")
	offset, _ := this.GetInt("offset") //默认为初始值0

	if episodesId == 0 {
		this.Data["json"] = ReturnError(4001, "必须指定视频剧集")
		this.ServeJSON()
		this.StopRun()
	}

	if limit == 0 {
		limit = 12
	}

	num, comments, err := models.GetCommentList(episodesId, offset, limit)

	if err != nil {
		this.Data["json"] = ReturnError(4001, "查询异常")
		this.ServeJSON()
		this.StopRun()
	}
	var data []ListStruct
	var commentInfo ListStruct
	// 遍历
	for _, v := range comments {
		commentInfo.Id = v.Id
		commentInfo.Content = v.Content
		commentInfo.AddTime = v.AddTime
		commentInfo.UserId = v.UserId
		commentInfo.Status = v.Status
		commentInfo.Stamp = v.Stamp
		commentInfo.PraiseCount = v.PraiseCount
		commentInfo.EpisodesId = v.EpisodesId
		commentInfo.VideoId = v.VideoId
		// 查找用户信息start
		userinfo, _ := models.GetUserInfo(v.UserId)
		commentInfo.UserInfo.Id = userinfo.Id
		commentInfo.UserInfo.Name = userinfo.Name
		commentInfo.UserInfo.AddTime = userinfo.AddTime
		commentInfo.UserInfo.Avatar = userinfo.Avatar
		// 查找用户信息end
		commentInfo.AddTimeTitle = DateFormat(v.AddTime) //format time
		commentInfo.PraiseCountNew = v.PraiseCount
		data = append(data, commentInfo)
	}
	this.Data["json"] = ReturnSuccess(0, "success", data, num)
	this.ServeJSON()
}

// 保存评论
// @router /comment/save [*]
func (this *CommentController) Save() {
	episodesId, _ := this.GetInt("episodesId")
	videoId, _ := this.GetInt("videoId")
	uid, _ := this.GetInt("uid")
	content := this.GetString("content")

	if content == "" {
		this.Data["json"] = ReturnError(4001, "内容不能为空")
		this.ServeJSON()
		this.StopRun()
	}
	if uid == 0 {
		this.Data["json"] = ReturnError(4001, "请先登录")
		this.ServeJSON()
		this.StopRun()
	}
	if episodesId == 0 {
		this.Data["json"] = ReturnError(4001, "必须指定评论视频集数")
		this.ServeJSON()
		this.StopRun()
	}
	if videoId == 0 {
		this.Data["json"] = ReturnError(4001, "必须指定评论视频")
		this.ServeJSON()
		this.StopRun()
	}

	id, err := models.CommentSave(content, uid, episodesId, videoId)

	if err != nil {
		this.Data["json"] = ReturnError(4001, "插入数据异常")
		this.ServeJSON()
		this.StopRun()
	}
	this.Data["json"] = ReturnSuccess(0, "success", nil, id)
	this.ServeJSON()
}
