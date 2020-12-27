package models

import (
	"github.com/astaxie/beego/orm"
)

type Video struct {
	Id                 int    `json:"id" orm:"column(id)"`
	Title              string `json:"title"`
	SubTitle           string `json:"sub_title"`
	Status             int    `json:"status"`
	Img                string `json:"img"`
	Img1               string `json:"img1"`
	ChannelId          int    `json:"channel_id"`
	RegionId           int    `json:"region_id"`
	TypeId             int    `json:"type_id"`
	AddTime            int64  `json:"add_time"`
	EpisodesCount      int    `json:"episodes_count"`
	EpisodesUpdateTime int64  `json:"episodes_update_time"`
	IsEnd              int    `json:"is_end"`
	IsHot              int    `json:"is_hot"`
	IsRecommend        int    `json:"is_recommend"`
	Comment            int    `json:"comment"`
}

type VideoEpisodes struct {
	Id            int    `json:"id" orm:"column(id)"`
	Title         string `json:"title"`
	AddTime       int64  `json:"add_time"`
	Num           int    `json:"add_time"`
	VideoId       int
	PlayUrl       string
	Status        int
	Comment       int
	AliyunVideoId string
}

type VedioData struct {
	Id            int    `json:"id"`
	Title         string `json:"title"`
	SubTitle      string `json:"sub_title"`
	Img           string `json:"img"`
	Img1          string `json:"img1"`
	AddTime       int64  `json:"add_time"`
	EpisodesCount int    `json:"episodes_count"`
	IsEnd         int    `json:"is_end"`
}

func init() {
	orm.RegisterModel(new(Video), new(VideoEpisodes))
}

// 获取顶部广告
func GetChannelHotList(channelId int) (int64, []*Video, error) {
	//defer coverPanic()

	var (
		vedios []*Video
		num    int64
		err    error
	)

	o := orm.NewOrm()
	qs := o.QueryTable("video").Filter("channel_id", channelId)
	qs = qs.Filter("status", 1).Filter("is_hot", 1).OrderBy("-episodes_update_time")
	num, err = qs.All(&vedios, "id", "channel_id", "title", "sub_title", "img", "img1", "add_time", "episodes_count", "is_end")
	return num, vedios, err
}

// 根据频道下地区id获取推荐视频
func GetChannelRegionRecommend(channelId int, regionId int) (int64, []*VedioData, error) {
	defer coverPanic()
	var (
		vedios []*VedioData
	)
	o := orm.NewOrm()
	num, err := o.Raw("SELECT `id`,`title`,`sub_title`,`img`,`img1`,`add_time`,`episodes_count`,`is_end` FROM `video` WHERE `channel_id` = ? AND `region_id` = ? AND `is_recommend` = 1 AND `status` = 1 ORDER BY `episodes_update_time` DESC LIMIT 9", channelId, regionId).QueryRows(&vedios)

	return num, vedios, err
}

// 根据频道下类型id获取推荐视频
func GetChannelTypeRecommend(channelId int, typeId int) (int64, []VedioData, error) {
	defer coverPanic()
	var (
		vedios []VedioData
	)
	o := orm.NewOrm()
	num, err := o.Raw("SELECT `id`,`title`,`sub_title`,`img`,`img1`,`add_time`,`episodes_count`,`is_end` FROM `video` WHERE `channel_id` = ? AND `type_id` = ? AND `is_recommend` = 1 AND `status` = 1 ORDER BY `episodes_update_time` DESC LIMIT 9", channelId, typeId).QueryRows(&vedios)
	return num, vedios, err
}

// 频道下根据不同条件和排序方式获取视频信息
func GetChannelVideoList(channelId int, regionId int, typeId int, end string, sort string, offset int, limit int) (int64, []map[string]interface{}, error) {
	o := orm.NewOrm()
	var vedios []orm.Params

	qs := o.QueryTable("video")
	qs = qs.Filter("channel_id", channelId)
	if regionId > 0 {
		qs = qs.Filter("region_id", regionId)
	}
	if typeId > 0 {
		qs = qs.Filter("type_id", typeId)
	}
	if end == "n" {
		// 没有完结
		qs = qs.Filter("is_end", 0)
	} else if end == "y" {
		// 已经完结
		qs = qs.Filter("is_end", 1)
	}

	// 排序
	if sort == "episodesUpdateTime" {
		qs.OrderBy("-episodes_update_time")
	} else if sort == "comment" {
		qs.OrderBy("-comment")
	} else if sort == "addTime" {
		qs.OrderBy("-add_time")
	} else {
		qs.OrderBy("-add_time")
	}

	qs = qs.Limit(limit, offset)
	nums, err := qs.Values(&vedios, "id", "title", "sub_title", "add_time", "img", "img1", "episodes_count", "is_end")

	// 返回转换
	ooo := make([]map[string]interface{}, nums)
	for i, v := range vedios {
		ooo[i] = map[string]interface{}{
			"id":             v["Id"],
			"title":          v["Title"],
			"sub_title":      v["SubTitle"],
			"add_time":       v["AddTime"],
			"img":            v["Img"],
			"img1":           v["Img1"],
			"episodes_count": v["EpisodesCount"],
			"is_end":         v["IsEnd"],
		}
	}

	return nums, ooo, err
}

// 获取单条视频信息
func GetVideoInfo(id int) (Video, error) {
	o := orm.NewOrm()
	var video Video
	err := o.Raw("SELECT * FROM video WHERE id = ? Limit 1", id).QueryRow(&video)
	return video, err
}

// 获取剧集列表信息
func GetVideoEpisodesList(id int) (int64, []map[string]interface{}, error) {
	o := orm.NewOrm()
	var videoPisodesList []orm.Params

	qs := o.QueryTable("video_episodes")
	qs = qs.Filter("video_id", id)
	qs.OrderBy("num")

	num, err := qs.Values(&videoPisodesList, "id", "title", "add_time", "num", "play_url", "comment")

	ooo := make([]map[string]interface{}, num)
	for i, v := range videoPisodesList {
		ooo[i] = map[string]interface{}{
			"id":       v["Id"],
			"title":    v["Title"],
			"add_time": v["AddTime"],
			"num":      v["Num"],
			"playUrl":  v["PlayUrl"],
			"comment":  v["Comment"],
		}
	}

	return num, ooo, err
}

// 频道排行榜
func GetChannelTop(channelId int, limit int) (int64, []Video, error) {
	o := orm.NewOrm()
	qs := o.QueryTable("video")
	qs = qs.Filter("channel_id", channelId)
	qs = qs.Filter("status", 1)
	qs = qs.OrderBy("-comment")
	qs = qs.Limit(limit)
	var videos []Video
	num, err := qs.Count()
	_, err = qs.All(&videos)
	return num, videos, err
}

// 类型排行榜
func GetTypeTop(typeId int, limit int) (int64, []Video, error) {
	o := orm.NewOrm()
	qs := o.QueryTable("video")
	qs = qs.Filter("type_id", typeId)
	qs = qs.Filter("status", 1)
	qs = qs.OrderBy("-comment")
	qs = qs.Limit(limit)
	var videos []Video
	num, err := qs.Count()
	_, err = qs.All(&videos)
	return num, videos, err
}
