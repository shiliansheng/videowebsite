package models

import (
	"videowebsite/utils"
)

type Review struct {
	Id      int    `json:"id"`
	Userid  int    `json:"userid"`
	Videoid int    `json:"videoid"`
	Content string `json:"content"`
	Pubtime string `json:"pubtime"`
	Status  string `json:"status"`
}

type ReviewInfo struct {
	// Id       int    `json:"id"`
	// Userid   int    `json:"userid"`
	// Videoid  int    `json:"videoid"`
	Nickname string `json:"nickname"`
	Userlogo string `json:"uerlogo"`
	Content  string `json:"content"`
	Pubtime  string `json:"pubtime"`
	// Status   string `json:"status"`
}

// ####### BASE

func (m Review) TableName() string {
	return TableName("review")
}

// ####### INFO 类

// 获取review信息，传递含有id的review
//  @param  r [*Review]
//  @return [*]
func (m Review) GetInfo(r *Review) error {
	err := Orm.Read(&r)
	return err
}

// 给定videoID和page, limit，返回评论信息列表
//  @param  vid [int]
//  @param  page [int]
//  @param  limit [int]
//  @return [[]ReviewInfo]
func (m Review) GetVideoReviewInfoList(vid, page, limit int) []ReviewInfo {
	list := []Review{}
	infolist := []ReviewInfo{}
	Orm.QueryTable(m.TableName()).Filter("videoid__contains", vid).OrderBy("-pubtime").Limit(limit, limit*(page-1)).All(&list)
	for _, pie := range list {
		user := &User{Id: pie.Userid}
		user.GetInfo(user)
		infolist = append(infolist, ReviewInfo{
			Nickname: user.Nickname,
			Userlogo: user.Userlogo,
			Content:  pie.Content,
			Pubtime:  pie.Pubtime,
		})
	}
	return infolist
}

// ###### CRUD

// 添加，参数为Review指针
//  @param  r [*Review]
//  @return [RespJson]
func (m Review) Add(r *Review) RespJson {
	resp := RespJson{Code: DO_ERROR}
	if r.Content == "" {
		resp.Msg = "添加失败，评论内容为空"
	} else {
		new(Video).AddReviewnum(r.Videoid)
		r.Pubtime = utils.GetNowTimeString()
		r.Status = "Y"
		_, err := Orm.Insert(r)
		if err != nil {
			resp.Msg = "添加失败, " + err.Error()
		} else {
			resp.Code = DO_SUCCESS
			resp.Msg = "添加成功"
		}
	}
	return resp
}
