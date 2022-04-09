package models

import (
	"log"
	"videowebsite/utils"
)

type History struct {
	Id      int    `json:"id"`
	Userid  int    `json:"userid"`
	Videoid int    `json:"videoid"`
	State   int    `json:"state"` // 0:存在，1:不存在
	Pubtime string `json:"pubtime"`
}

type HistoryInfo struct {
	Id        int           `json:"id"`
	Pubtime   string        `json:"pubtime"`
	Videoinfo VideoShowInfo `json:"videoinfo"`
}

func (m History) TableName() string {
	return TableName("history")
}

// 获取用户观看历史列表，返回观看历史信息列表数组指针
//  @param  id [int] 用户id
//  @return [*[]HistoryInfo]
func (m History) GetUserHistoryList(id int) *[]HistoryInfo {
	hlist := []History{}
	uhlist := []HistoryInfo{}
	if _, err := Orm.QueryTable(m.TableName()).Filter("userid", id).Filter("state", 0).All(&hlist); err != nil {
		log.Println("获取用户收藏列表失败", err)
		return nil
	}
	tmpV := new(Video)
	for i := range hlist {
		uhlist = append(uhlist, HistoryInfo{
			Id:        hlist[i].Id,
			Pubtime:   hlist[i].Pubtime,
			Videoinfo: *tmpV.GetVideoShowInfo(hlist[i].Videoid),
		})
	}
	return &uhlist
}

// 获取History内容，需要提供userid和videoid，返回resp.data为id
//  @param  h [*History]
//  @return [RespJson]
func (m History) Get(h *History) RespJson {
	resp := NewRespJson()
	h.State = 0
	if err := Orm.Read(h, "userid", "videoid", "state"); err != nil {
		resp.Msg = "获取内容失败"
		log.Println(resp.Msg, err)
	} else {
		resp.Code = DO_SUCCESS
		resp.Data = h.Id
	}
	return *resp
}

// 添加History，需要提供userid和videoid
//  @param  h [*History]
//  @return [RespJson]
func (m History) Add(h *History) RespJson {
	resp := NewRespJson()
	if h.Videoid == 0 || h.Userid == 0 {
		resp.Msg = "添加失败，所需内容为空"
		log.Println(resp.Msg, *h)
	} else {
		if err := Orm.Read(h, "userid", "videoid"); err != nil { // 不存在数据
			h.Pubtime = utils.GetNowTimeString()
			h.State = 0
			_, err := Orm.Insert(h)
			if err != nil {
				resp.Msg = "添加失败, " + err.Error()
				log.Println(resp.Msg)
			} else { // 存在数据
				resp.Code = DO_SUCCESS
				resp.Data = h.Id
				resp.Msg = "添加成功"
			}
		} else {
			h.Pubtime = utils.GetNowTimeString()
			h.State = 0
			resp = m.Update(h, "state", "pubtime")
		}

	}
	return *resp
}

// 删除History，需要提供History的id
//  @param  c [*History]
//  @return [RespJson]
func (m History) Delete(h *History) RespJson {
	resp := NewRespJson()
	if h.Id == 0 {
		resp.Msg = "删除失败，所需内容为空"
		log.Println(resp.Msg, *h)
	} else {
		h.State = 1
		resp = m.Update(h, "state")
	}
	return *resp
}

func (m History) Update(c *History, cols ...string) *RespJson {
	resp := NewRespJson()
	_, err := Orm.Update(c, cols...)
	if err != nil {
		resp.Msg = "更新失败, " + err.Error()
		log.Println(resp.Msg)
	} else {
		resp.Code = DO_SUCCESS
		resp.Data = c.Id
		resp.Msg = "更新成功"
	}
	return resp
}
