package models

import (
	"github.com/astaxie/beego/orm"
)

type Advert struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	SubTitle  string
	ChannelId int64
	Img       string
	Sort      string
	AddTime   int64
	Url       string
	Status    int
}

func init() {
	orm.RegisterModel(new(Advert))
}

// 获取顶部广告
func GetChannelAdvert(channelId int) (int64, error, []Advert) {
	defer coverPanic()
	var (
		adverts []Advert
		num     int64
		err     error
	)
	o := orm.NewOrm()
	qs := o.QueryTable("advert").Filter("channel_id", channelId)
	qs = qs.Filter("status", 1).OrderBy("-sort")
	num, err = qs.All(&adverts, "id", "title", "sub_title", "img", "add_time", "url")
	return num, err, adverts
}
