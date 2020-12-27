package models

import (
	"errors"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/scauxiaoxu/gotool/time"
)

type Comment struct {
	Id          int64
	Content     string
	AddTime     int64
	UserId      int
	Status      int
	Stamp       int
	PraiseCount int //点赞书
	EpisodesId  int //评论视频
	VideoId     int //所属视频
}

func init() {
	orm.RegisterModel(new(Comment))
}

// 获取剧集评论信息
func GetCommentList(episodesId int, offset int, limit int) (int64, []Comment, error) {
	o := orm.NewOrm()
	qs := o.QueryTable("comment")
	qs = qs.Filter("episodes_id", episodesId)
	qs = qs.Filter("status", 1)
	qs = qs.OrderBy("-add_time")
	qs = qs.Limit(limit, offset)
	var commentList []Comment
	num, err := qs.Count()
	_, err = qs.All(&commentList)
	return num, commentList, err
}

// 保存评论
func CommentSave(content string, uid int, episodesId int, videoId int) (int64, error) {
	o := orm.NewOrm()
	// 开启事物
	err := o.Begin()
	if err != nil {
		logs.Error("start the transaction failed")
		return 0, errors.New("事物开启异常")
	}

	var comment Comment
	comment.Content = content
	comment.AddTime = time.Time()
	comment.UserId = uid
	comment.Status = 1 //默认审核
	comment.EpisodesId = episodesId
	comment.VideoId = videoId
	id, err := o.Insert(&comment)
	if err != nil {
		logs.Error("execute transaction's sql fail, rollback.", err)
		err = o.Rollback()
		if err != nil {
			logs.Error("roll back transaction failed", err)
		}
		return 0, errors.New("插入评论异常")
	}
	//修改视频的总评论数
	res, err := o.Raw("UPDATE video SET comment=comment+1 WHERE id=?", videoId).Exec()
	num, _ := res.RowsAffected()
	if err != nil || num == 0 {
		logs.Error("更新剧集评论数异常.", err)
		logs.Error("更新剧集评论数异常.", num)
		err = o.Rollback()
		if err != nil {
			logs.Error("roll back transaction failed", err)
		}
		return 0, errors.New("更新异常1")
	}
	//修改视频剧集的评论数
	res, err = o.Raw("UPDATE video_episodes SET comment=comment+1 WHERE id=?", episodesId).Exec()
	num, _ = res.RowsAffected()
	if err != nil || num == 0 {
		logs.Error("更新分集评论数异常.", err)
		logs.Error("更新分集评论数异常.", num)
		err = o.Rollback()
		if err != nil {
			logs.Error("roll back transaction failed", err)
		}
		return 0, errors.New("更新异常2")
	}

	err = o.Commit()
	if err != nil {
		logs.Error("commit transaction failed.", err)
	}
	return id, err
}
