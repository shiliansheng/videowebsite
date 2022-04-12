package models

import (
	"log"
	"videowebsite/utils"
)

type Post struct {
	Id       int    `json:"id"`
	Videoid  int    `json:"videoid"`
	Postlogo string `json:"postlogo"`
	State    int    `json:"state"`
	Pubtime  string `json:"pubtime"`
}

type PostInfo struct {
	Id        int    `json:"id"`
	Videoid   int    `json:"videoid"`
	Videoname string `json:"videoname"`
	Postlogo  string `json:"postlogo"`
	Pubtime   string `json:"pubtime"`
}

// 有效的post最大数量
const maxValidPostNum int = 5

func (m Post) TableName() string {
	return TableName("post")
}

func (m Post) GetPostCount() int {
	var count int = 0
	if err := Orm.Raw("SELECT COUNT(*) FROM `" + m.TableName() + "` WHERE `state`=0").QueryRow(&count); err != nil {
		log.Println("获取海报有效数失败:", err)
	}
	return count
}

func (m Post) GetOldestPostId() int {
	var id int = 0
	if err := Orm.Raw("SELECT id FROM `" + m.TableName() + "` WHERE `state`=0 ORDER BY `pubtime` LIMIT 1").QueryRow(&id); err != nil {
		log.Println("获取海报有效数失败:", err)
	}
	return id
}

// 获取海报列表，直接获取海报的所有信息
//  @return [RespJson] RespJson.Data = []PostInfo
func (m Post) GetPostList() RespJson {
	resp := NewRespJson()
	list := []Post{}
	infoList := []PostInfo{}
	if _, err := Orm.QueryTable(m.TableName()).Filter("state", 0).OrderBy("-pubtime").All(&list); err == nil {
		resp.Code = DO_SUCCESS
		resp.Msg = "获取海报列表成功"
		resp.Count = len(list)
		tmpv := Video{}
		for _, pie := range list {
			tmpv.Id = pie.Videoid
			tmpv.GetInfo(&tmpv)
			infoList = append(infoList, PostInfo{
				Id:        pie.Id,
				Videoid:   pie.Videoid,
				Videoname: tmpv.Videoname,
				Postlogo:  pie.Postlogo,
				Pubtime:   pie.Pubtime,
			})
		}
		resp.Data = infoList
	} else {
		resp.Msg = "获取海报列表失败"
		log.Println(resp.Msg, err)
	}
	return *resp
}

// 传递Post指针，必须含有Videoid和Postlogo两项内容
//  @param  p [*Post]
//  @return [RespJson]
func (m Post) Add(p *Post) RespJson {
	resp := NewRespJson()
	if p.Videoid == 0 || p.Postlogo == "" {
		resp.Msg = "缺少关键信息"
		log.Println(resp.Msg, *p)
		return *resp
	}
	if err := new (Video).GetInfo(&Video{Id: p.Videoid}); err != nil {
		resp.Msg = "当前视频ID出错，不存在该视频"
		log.Println(resp.Msg, err)
		return *resp
	}
	// 进行要求设定查询
	//   最多存在 maxValidPostNum 个Post
	//   如果超过 maxValidPostNum ，则将当前存在时间最长的Post删除
	if m.GetPostCount() >= maxValidPostNum {
		m.Delete(&Post{Id: m.GetOldestPostId()})
	}

	// 进行添加
	p.Pubtime = utils.GetNowTimeString()
	p.State = 0
	tmp := &Post{Id: p.Id}
	if err := Orm.Read(tmp); err == nil {
		p.Id = tmp.Id
		*resp = p.Update(p, []string{"stauts", "pubtime"}...)
	} else {
		if _, err := Orm.Insert(p); err != nil {
			resp.Msg = "添加海报失败"
			log.Println(resp.Msg, err)
		} else {
			resp.Code = DO_SUCCESS
			resp.Msg = "添加海报成功"
		}
	}
	return *resp
}

// 提供含有Id的Post指针，且更新信息的行信息必须在cols中
//  @param  p [*Post]
//  @param  cols [...string]
//  @return [RespJson]
func (m Post) Update(p *Post, cols ...string) RespJson {
	resp := NewRespJson()
	if len(cols) == 0 {
		resp.Msg = "当前更新信息为空，更新失败，请添加更新信息"
		return *resp
	}
	if _, err := Orm.Update(p, cols...); err != nil {
		resp.Msg = "更新海报失败"
		log.Println(resp.Msg, err)
	} else {
		resp.Msg = "更新海报成功"
		resp.Code = DO_SUCCESS
	}
	return *resp
}

// 传递含有Id项的Post指针，不进行删除，仅进行状态更新
//  Staus: 0 -> 1
//  @param  p [*Post]
//  @return [RespJson]
func (m Post) Delete(p *Post) RespJson {
	resp := NewRespJson()
	if err := Orm.Read(p); err != nil {
		resp.Msg = "删除海报失败, 不存在该海报"
	} else {
		p.State = 1
		*resp = p.Update(p, "state")
	}
	return *resp
}
