package models

import (
	"log"
	"videowebsite/utils"
)

type Score struct {
	Id      int    `json:"id"`
	Userid  int    `json:"userid"`
	Videoid int    `json:"videoid"`
	Value   int    `json:"value"`
	Pubtime string `json:"pubtime"`
}

// ####### BASE

func (m Score) TableName() string {
	return TableName("score")
}

// ####### INFO 类

// 获取Score信息，传递含有id的Score
//  @param  s [*Score]
//  @return [*]
func (m Score) GetInfo(s *Score) error {
	err := Orm.Read(&s)
	return err
}

// ###### CRUD

// 添加Score，包含videoid, userid, value
//  @param  s [*Score]
//  @return [RespJson]
func (m Score) Add(s *Score) RespJson {
	resp := RespJson{Code: DO_ERROR}
	if s.Videoid == 0 || s.Userid == 0 || s.Value == 0 {
		resp.Msg = "添加失败，所需内容为空"
		log.Println(resp.Msg, *s)
	} else {
		s.Pubtime = utils.GetNowTimeString()
		_, err := Orm.Insert(s)
		if err != nil {
			resp.Msg = "添加失败, " + err.Error()
			log.Println(resp.Msg)
		} else {
			resp.Code = DO_SUCCESS
			resp.Msg = "添加成功"
		}
		new(Video).UpdateScore(s.Videoid, s.Value)
	}
	return resp
}

// 获取评分值
//  @param  s [*Score]
//  @return [RespJson] resp.data=value
func (m Score) GetValue(s *Score) RespJson {
	resp := NewRespJson()
	err := Orm.Read(s, "userid", "videoid")
	if err != nil {
		resp.Msg = "获取评分失败"
		log.Println(resp.Msg, err)
	} else {
		resp.Data = s.Value
		resp.Code = DO_SUCCESS
	}
	return *resp
}
