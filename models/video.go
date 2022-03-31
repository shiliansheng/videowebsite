package models

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"videowebsite/utils"
)

type Video struct {
	Id            int     `json:"id" orm:"pk"`   // 视频编号
	Videoname     string  `json:"videoname"`     // 视频名称
	Classifiction string  `json:"classifiction"` // 视频分类
	Typename      string  `json:"typename"`      // 视频类型名称
	Introduction  string  `json:"introduction"`  // 视频介绍
	Videologo     string  `json:"videologo"`     // 视频图片地址
	Keywords      string  `json:"keywords"`      // 视频关键字
	Videoresource string  `json:"videoresource"` // 视频资源地址
	Copyright     string  `json:"copyright"`     // 版权所有(原创,转载)
	Username      string  `json:"username"`      // 发布者用户名
	Pubtime       string  `json:"pubtime"`       // 发布时间
	Viewnum       int     `json:"viewnum"`       // 视频观看次数
	Scorenum      int     `json:"scorenum"`      // 视频打分人数
	Remarknum     int     `json:"remarknum"`     // 视频评论次数
	Averscore     float64 `json:"averscore"`     // 用户评分平均分
	Totalscore    int64   `json:"totalscore"`    // 用户评分总分
	Passed        string  `json:"passed"`        // 审核状态(待审核;通过审核)
	Recommand     int     `json:"recommand"`     // 视频推荐(0:不推荐,1:推荐)
}

// ### base function

func (m Video) TableName() string {
	return TableName("video")
}

// ### 获取INFO

// 根据观看次数排序，获取前指定数目(num)的视频
//  @param  num [int] 指定数目
//  @return [[]Video] video数组 
func (m Video) GetHotVideos(num int) []Video {
	videos := []Video{}
	_, err := Orm.QueryTable(m.TableName()).OrderBy("viewnum").Limit(num).All(&videos)
	if err != nil {
		fmt.Println("[ERROR]Get hot video list failed:", err)
	}
	return videos
}


func (m Video) GetVideoCount() int {
	count, _ := Orm.QueryTable(m.TableName()).Count()
	return int(count)
}

func (m Video) GetVideoList(page, limit int, filterMap map[string]interface{}) RespJson {
	resp := RespJson {
		Code: DO_SUCCESS,
		Count: 0,
	}
	list, count, err := m.getVideoList(page, limit, filterMap)
	if err != nil {
		resp.Code = DO_ERROR
		resp.Msg = "获取视频列表失败<br/>" + err.Error()
	} else {
		resp.Msg = "获取视频列表成功"
		resp.Data = list
	}
	resp.Count = count
	return resp
}

func (m Video) getVideoList(page, limit int, mapper map[string]interface{}) ([]Video, int, error) {
	list := []Video{}
	seter := Orm.QueryTable(m.TableName())
	for key, value := range mapper {
		if value == "" {
			continue
		}
		seter = seter.Filter(key+"__icontains", value)
	}
	count, _ := seter.Count()
	_, err := seter.Limit(limit, limit*(page - 1)).All(&list)
	return list, int(count), err
}

// ### CRUD

// 添加参数中的视频
//  @param  v [Video] 待添加的视频
//  @return [RespJson] 
func (m Video) Add(v Video) RespJson {
	resp := RespJson{Code: DO_ERROR}
	if v.Videoname == "" || v.Videoresource == "" || v.Typename == "" || v.Username == "" {
		resp.Msg = "信息不全，添加失败!"
		return resp
	}
	// 设置初始化信息
	v.Pubtime = utils.GetNowTimeString()
	v.Passed = "待审核"

	_, err := Orm.Insert(&v)
	if err != nil {
		resp.Msg = "添加视频失败<br/>" + err.Error()
	} else {
		resp.Code = DO_SUCCESS
		resp.Msg = "添加成功"
	}
	return resp
}

// 更新旧的video，使用旧的进行调用，参数为新的video
//  @param  v [Video] 新的video
//  @return [RespJson]
func (m Video) Update(v Video) RespJson {
	cols := m.GetDifCols(v)
	resp := RespJson{Code: DO_ERROR}
	if len(cols) == 0 {
		resp.Msg = "信息未更改，更新失败"
		return resp
	}
	_, err := Orm.Update(&v, cols...)
	if err != nil {
		resp.Msg =  "更新信息失败<br/>" + err.Error()
	} else {
		resp.Msg = "更新信息成功"
		resp.Code = DO_SUCCESS
	}
	return resp
}

// 删除参数中的video，只需要提供id属性
//  @param  v [Video] 
//  @return [RespJson] 
func (m Video) Delete(v Video) RespJson {
	resp := RespJson{Code: DO_ERROR}
	_, err := Orm.Delete(&v, "id")
	if err != nil {
		resp.Msg = "删除类型" + v.Videoname + "失败: " + err.Error()
	} else {
		resp.Code = DO_SUCCESS
		resp.Msg = "删除类型" + v.Videoname + "成功"
	}
	return resp
}

// ### 设置INFO

func (m Video) GetDifCols(v Video) []string {
	dif := []string{}
	if m.Classifiction != v.Classifiction {
		dif = append(dif, "classifiction")
	}
	if m.Typename != v.Typename {
		dif = append(dif, "typename")
	}
	if m.Videoname != v.Videoname {
		dif = append(dif, "videoname")
	}
	if m.Introduction != v.Introduction {
		dif = append(dif, "introduction")
	}
	if m.Videologo != v.Videologo {
		dif = append(dif, "videologo")
	}
	if m.Keywords != v.Keywords {
		dif = append(dif, "keywords")
	}
	if m.Videoresource != v.Videoresource {
		dif = append(dif, "videoresource")
	}
	if m.Copyright != v.Copyright {
		dif = append(dif, "copyright")
	}
	if m.Username != v.Username {
		dif = append(dif, "username")
	}
	if m.Pubtime != v.Pubtime {
		dif = append(dif, "pubtime")
	}
	if m.Viewnum != v.Viewnum {
		dif = append(dif, "viewnum")
	}
	if m.Scorenum != v.Scorenum {
		dif = append(dif, "scorenum")
	}
	if m.Remarknum != v.Remarknum {
		dif = append(dif, "remarknum")
	}
	if m.Averscore != v.Averscore {
		dif = append(dif, "averscore")
	}
	if m.Totalscore != v.Totalscore {
		dif = append(dif, "totalscore")
	}
	if m.Passed != v.Passed {
		dif = append(dif, "passed")
	}
	if m.Recommand != v.Recommand {
		dif = append(dif, "recommand")
	}
	return dif
}

func (m *Video) SetVideo(source url.Values) {
	if value := source.Get("classifiction"); value != "" {
		m.Classifiction = value
	}
	for key, value := range source {
		if strings.HasPrefix(key, "typename") {
			m.Typename += value[0] + "/"
		}
	}
	m.Typename = m.Typename[:len(m.Typename)-1]
	if value := source.Get("typename"); value != "" {
		m.Typename = value
	}
	if value := source.Get("videoname"); value != "" {
		m.Videoname = value
	}
	if value := source.Get("introduction"); value != "" {
		m.Introduction = value
	}
	if value := source.Get("videologo"); value != "" {
		m.Videologo = value
	}
	if value := source.Get("keywords"); value != "" {
		m.Keywords = value
	}
	if value := source.Get("videoresource"); value != "" {
		m.Videoresource = value
	}
	if value := source.Get("copyright"); value != "" {
		m.Copyright = value
	}
	if value := source.Get("username"); value != "" {
		m.Username = value
	}
	if value := source.Get("pubtime"); value != "" {
		m.Pubtime = value
	}
	if value := source.Get("viewnum"); value != "" {
		m.Viewnum = utils.Atoi(value)
	}
	if value := source.Get("scorenum"); value != "" {
		m.Scorenum = utils.Atoi(value)
	}
	if value := source.Get("remarknum"); value != "" {
		m.Remarknum = utils.Atoi(value)
	}
	if value := source.Get("averscore"); value != "" {
		aver, _ := strconv.ParseFloat(value, 32)
		m.Averscore = aver
	}
	if value := source.Get("totalscore"); value != "" {
		total, _ := strconv.ParseInt(value, 10, 64)
		m.Totalscore = total
	}
	if value := source.Get("passed"); value != "" {
		m.Passed = value
	}
	if value := source.Get("recommand"); value != "" {
		m.Recommand = utils.Atoi(value)
	}
}
