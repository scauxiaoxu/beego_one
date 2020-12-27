package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/scauxiaoxu/gotool/time"
)

type Message struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	AddTime int64  `json:"add_time"`
}

type MessageUser struct {
	Id        int   `json:"id"`
	MessageId int64 `json:"message_id"`
	UserId    int   `json:"user_id"`
	AddTime   int64 `json:"add_time"`
	Status    int   `json:"status"`
}

func init() {
	orm.RegisterModel(new(Message), new(MessageUser))
}

// 保存通知信息
func SendMessageDo(content string) (int64, error) {
	o := orm.NewOrm()
	var message Message
	message.Content = content
	message.AddTime = time.Time()
	insertId, err := o.Insert(&message)
	return insertId, err
}

// 保存消息接收人
func SendMessageUser(messageID int64, uid int) (int64, error) {
	o := orm.NewOrm()
	var messageUser MessageUser
	messageUser.MessageId = messageID
	messageUser.UserId = uid
	messageUser.Status = 1
	messageUser.AddTime = time.Time()
	insertId, err := o.Insert(&messageUser)
	return insertId, err
}
