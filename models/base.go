package models

import "github.com/astaxie/beego/orm"

type Region struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Type struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// 获取频道下地区
func GetChannelRegion(channelId int) (int64, []Region, error) {
	o := orm.NewOrm()
	var regions []Region
	num, err := o.Raw("SELECT id, name FROM channel_region WHERE channel_id = ? AND status = 1 ORDER BY sort Desc", channelId).QueryRows(&regions)
	return num, regions, err
}

// 获取频道下类型
func GetChannelType(channelId int) (int64, []Type, error) {
	o := orm.NewOrm()
	var types []Type
	num, err := o.Raw("SELECT id, name FROM channel_type WHERE channel_id = ? AND status = 1 ORDER BY sort Desc", channelId).QueryRows(&types)
	return num, types, err
}
