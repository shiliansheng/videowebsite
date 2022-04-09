package models

import (
	"log"
	"videowebsite/utils"
)

type Collect struct {
	Id      int    `json:"id"`
	Userid  int    `json:"userid"`
	Videoid int    `json:"videoid"`
	State   int    `json:"state"` // 0:存在，1:不存在
	Pubtime string `json:"pubtime"`
}

type CollectInfo struct {
	ID        int           `json:"id"`
	Pubtime   string        `json:"pubtime"`
	Videoinfo VideoShowInfo `json:"videoinfo"`
}

func (m Collect) TableName() string {
	return TableName("collect")
}

// 获取用户收藏列表，返回收藏列表数组指针
//  @param  id [int] 用户id
//  @return [*[]CollectInfo]
func (m Collect) GetUserCollectList(id int) *[]CollectInfo {
	clist := []Collect{}
	uclist := []CollectInfo{}
	if _, err := Orm.QueryTable(m.TableName()).Filter("userid", id).Filter("state", 0).All(&clist); err != nil {
		log.Println("获取用户收藏列表失败", err)
		return nil
	}
	tmpV := new(Video)
	for i := range clist {
		uclist = append(uclist, CollectInfo{
			ID:        clist[i].Id,
			Pubtime:   clist[i].Pubtime,
			Videoinfo: *tmpV.GetVideoShowInfo(clist[i].Videoid),
		})
	}
	return &uclist
}

// 获取collect内容，需要提供userid和videoid，返回resp.data为id
//  @param  c [*Collect]
//  @return [RespJson]
func (m Collect) Get(c *Collect) RespJson {
	resp := NewRespJson()
	c.State = 0
	if err := Orm.Read(c, "userid", "videoid", "state"); err != nil {
		resp.Msg = "获取内容失败"
		log.Println(resp.Msg, err)
	} else {
		resp.Code = DO_SUCCESS
		resp.Data = c.Id
	}
	return *resp
}

// 添加collect，需要提供userid和videoid
//  @param  c [*Collect]
//  @return [RespJson]
func (m Collect) Add(c *Collect) RespJson {
	resp := NewRespJson()
	if c.Videoid == 0 || c.Userid == 0 {
		resp.Msg = "添加失败，所需内容为空"
		log.Println(resp.Msg, *c)
	} else {
		if err := Orm.Read(c, "userid", "videoid"); err != nil { // 不存在数据
			c.Pubtime = utils.GetNowTimeString()
			c.State = 0
			_, err := Orm.Insert(c)
			if err != nil {
				resp.Msg = "添加失败, " + err.Error()
				log.Println(resp.Msg)
			} else { // 存在数据
				resp.Code = DO_SUCCESS
				resp.Data = c.Id
				resp.Msg = "添加成功"
			}
		} else {
			c.State = 0
			resp = m.Update(c, "state")
		}

	}
	return *resp
}

// 删除collect，需要提供collect的id
//  @param  c [*Collect]
//  @return [RespJson]
func (m Collect) Delete(c *Collect) RespJson {
	resp := NewRespJson()
	if c.Id == 0 {
		resp.Msg = "删除失败，所需内容为空"
		log.Println(resp.Msg, *c)
	} else {
		c.State = 1
		resp = m.Update(c, "state")
	}
	return *resp
}

func (m Collect) Update(c *Collect, cols ...string) *RespJson {
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
