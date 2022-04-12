package models

import (
	"fmt"
	"log"
	"net/url"
	"strconv"
	"strings"
	"videowebsite/utils"
)

type Video struct {
	Id             int    `json:"id" orm:"pk"`    // 视频编号
	Videoname      string `json:"videoname"`      // 视频名称
	Classification string `json:"classification"` // 视频分类
	Typename       string `json:"typename"`       // 视频类型名称
	Introduction   string `json:"introduction"`   // 视频介绍
	Videologo      string `json:"videologo"`      // 视频图片地址
	Keywords       string `json:"keywords"`       // 视频关键字
	Videoresource  string `json:"videoresource"`  // 视频资源地址
	Copyright      string `json:"copyright"`      // 版权所有(原创,转载)
	Username       string `json:"username"`       // 发布者用户名
	Pubtime        string `json:"pubtime"`        // 发布时间
	Viewnum        int    `json:"viewnum"`        // 视频观看次数
	Scorenum       int    `json:"scorenum"`       // 视频打分人数
	Reviewnum      int    `json:"reviewnum"`      // 视频评论次数
	Averscore      string `json:"averscore"`      // 用户评分平均分
	Totalscore     int64  `json:"totalscore"`     // 用户评分总分
	Passed         string `json:"passed"`         // 审核状态(待审核;通过审核)
	Recommand      int    `json:"recommand"`      // 视频推荐(0:不推荐,1:推荐)
	State          int    `json:"state"`          // 视频状态(0: 有效,1:无效)
	CreateAt       string `json:"createat"`       // 创建时间
	UpdateAt       string `json:"updateat"`       // 更新时间
}

type VideoShowInfo struct {
	Id             int    `json:"id" orm:"pk"`    // 视频编号
	Videoname      string `json:"videoname"`      // 视频名称
	Classification string `json:"classification"` // 视频分类
	Typename       string `json:"typename"`       // 视频类型名称
	Introduction   string `json:"introduction"`   // 视频介绍
	Videologo      string `json:"videologo"`      // 视频图片地址
	Viewnum        int    `json:"viewnum"`        // 视频观看次数
	Keywords       string `json:"keywords"`       // 视频关键字
	Averscore      string `json:"averscore"`      // 用户评分平均分
}

type VideoPlayInfo struct {
	Id             int    `json:"id" orm:"pk"`    // 视频编号
	Videoname      string `json:"videoname"`      // 视频名称
	Classification string `json:"classification"` // 视频分类
	Typename       string `json:"typename"`       // 视频类型名称
	Introduction   string `json:"introduction"`   // 视频介绍
	Keywords       string `json:"keywords"`       // 视频关键字
	Videoresource  string `json:"videoresource"`  // 视频资源地址
	Averscore      string `json:"averscore"`      // 用户评分平均分
}

// classification Map
var ClassificationMap = map[string]string{
	"movie":   "电影",
	"cartoon": "卡通",
	"episode": "剧集",
	"others":  "其他",
	"library": "片库",
}

// ### base function

func (m Video) TableName() string {
	return TableName("video")
}

// ### 获取INFO

// 获取video info，video必须含有id
//  @param  v [*Video] 含有含有id值
//  @return [error]
func (m Video) GetInfo(v *Video) error {
	v.State = 0;
	err := Orm.Read(v, []string{"id", "state"}...)
	return err
}

// 获取基础类别视频数量
//  @return [[]string] 基础类别数组
//  @return [[]int] 对应的数量
func (m Video) GetClassificationCount() ([]string, []int) {
	classes := []string{}
	counts := []int{}
	_, err := Orm.Raw("SELECT COUNT(*), `classification` FROM `"+m.TableName()+"` WHERE `state`=0 GROUP BY `classification`").QueryRows(&counts, &classes)
	if err != nil {
		log.Println("获取基础类别视频数量失败：", err)
	}
	for i := range classes {
		classes[i] = ClassificationMap[classes[i]]
	}
	return classes, counts
}

// 获取以当前日期为基准的一整周视频上传数量
//  @return [[]string] 月-日 格式数组
//  @return [[]int] 当天上传的数量数组
func (m Video) GetWeekUploadData() ([]string, []int) {
	names, values := []string{}, []int{}
	_, err := Orm.Raw("SELECT DATE_FORMAT(`create_at`,'%m-%d') AS DATA_TIME, COUNT(*) FROM `"+
		m.TableName()+"` WHERE `state`=0 and `create_at` > ADDDATE(CURDATE(),INTERVAL -6 DAY) GROUP BY DATA_TIME ORDER BY DATA_TIME;").QueryRows(&names, &values)
	if err != nil {
		log.Println("获取用户注册数量失败:", err)
	}
	return names, values
}

// 根据观看次数排序，获取前指定数目(num)的视频
//  @param  num [int] 指定数目
//  @param  classification [...string]
//  @return [[]Video] video数组
func (m Video) GetHotVideos(num int, classification ...string) []Video {
	videos := []Video{}
	seter := Orm.QueryTable(m.TableName()).Filter("state", 0)
	if len(classification) != 0 {
		seter = seter.Filter("classification", classification[0])
	}
	_, err := seter.OrderBy("-viewnum").Limit(num).All(&videos)
	if err != nil {
		fmt.Println("[ERROR]Get hot video list failed:", err)
	}
	for i := range videos {
		videos[i].Classification = ClassificationMap[videos[i].Classification]
	}
	return videos
}

// 获取视频数量，参数为classicition
//  @param  classicition [...string] 为空则为全部
//  @return [int]
func (m Video) GetVideoCount(classicition ...string) int {
	seter := Orm.QueryTable(m.TableName()).Filter("state", 0)
	if len(classicition) != 0 {
		seter = seter.Filter("classification", classicition[0])
	}
	count, _ := seter.Count()
	return int(count)
}

// 获取视频列表
//  需要提供如下参数，其mapper中如果含有typename键，则将会以+进行分割其值
//  @param  page [int] 页
//  @param  limit [int] 每页数量
//  @param  mapper [map[string]interface{}] 筛选键值对
func (m Video) GetVideoList(page, limit int, filterMap map[string]interface{}) RespJson {
	resp := RespJson{
		Code:  DO_SUCCESS,
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

// 获取视频信息列表
//  需要提供如下参数，其mapper中如果含有typename键，则将会以+进行分割其值
//  @param  page [int] 页
//  @param  limit [int] 每页数量
//  @param  mapper [map[string]interface{}] 筛选键值对
//  @param  sorts [...string] 排序，其类型为枚举类型：add, rate, view, 默认为update
//  @param  RespJson.Data=[]VideoShowInfo
func (m Video) GetSortVideoShowInfoList(page, limit int, filterMap map[string]interface{}, sorts ...string) RespJson {
	resp := NewRespJson()
	vshowInfos := []VideoShowInfo{}
	list, count, err := m.getVideoList(page, limit, filterMap, sorts...)
	if err != nil {
		resp.Code = DO_ERROR
		resp.Msg = "获取视频列表失败<br/>" + err.Error()
	} else {
		resp.Code = DO_SUCCESS
		resp.Msg = "获取视频列表成功"
		resp.Count = len(list)
		for _, pie := range list {
			vshowInfos = append(vshowInfos, *pie.setVideoShowInfo())
		}
	}
	resp.Count = count
	resp.Data = vshowInfos
	return *resp
}

// 获取视频列表
//  需要提供如下参数，其mapper中如果含有typename键，则将会以+进行分割其值
//  @param  page [int] 页
//  @param  limit [int] 每页数量
//  @param  mapper [map[string]interface{}] 筛选键值对
//  @param  sorts [...string] 排序，其类型为枚举类型：add, rate, view, 默认为update
func (m Video) GetSortVideoList(page, limit int, filterMap map[string]interface{}, sorts ...string) VideoRespJson {
	resp := VideoRespJson{
		Code:  DO_SUCCESS,
		Count: 0,
	}
	list, count, err := m.getVideoList(page, limit, filterMap, sorts...)
	if err != nil {
		resp.Code = DO_ERROR
		resp.Msg = "获取视频列表失败<br/>" + err.Error()
	} else {
		resp.Msg = "获取视频列表成功"
		resp.Data.Size = len(list)
		resp.Page = page
		for _, pie := range list {
			resp.Data.Id = append(resp.Data.Id, pie.Id)
			resp.Data.Name = append(resp.Data.Name, pie.Videoname)
			resp.Data.Type = append(resp.Data.Type, pie.Typename)
			resp.Data.Logo = append(resp.Data.Logo, pie.Videologo)
			resp.Data.Score = append(resp.Data.Score, pie.Averscore)
		}
	}
	resp.Count = count
	return resp
}

// 获取视频列表
//  需要提供如下参数，其mapper中如果含有typename键，则将会以+进行分割其值
//  @param  page [int] 页
//  @param  limit [int] 每页数量
//  @param  mapper [map[string]interface{}] 筛选键值对
//  @param  sorts [...string] 排序，其类型为枚举类型：add, rate, view, 默认为update
//  @return [[]Video]
//  @return [int] 排除page和limit影响，返回查询到的数量
//  @return [error]
func (m Video) getVideoList(page, limit int, mapper map[string]interface{}, sorts ...string) ([]Video, int, error) {
	list := []Video{}
	seter := Orm.QueryTable(m.TableName()).Filter("state", 0)
	for key, value := range mapper {
		if value == "" {
			continue
		}
		if key == "typename" {
			pies := strings.Split(value.(string), "+")
			for _, pie := range pies {
				if pie == "" {
					continue
				}
				seter = seter.Filter("typename__icontains", pie)
			}
		}
		seter = seter.Filter(key+"__icontains", value)
	}
	count, _ := seter.Count()
	if len(sorts) > 0 {
		if sorts[0] == "add" {
			seter = seter.OrderBy("-create_at")
		} else if sorts[0] == "update" {
			seter = seter.OrderBy("-update_at")
		} else if sorts[0] == "rate" {
			seter = seter.OrderBy("-averscore")
		} else if sorts[0] == "view" {
			seter = seter.OrderBy("-viewnum")
		} else {
			seter = seter.OrderBy("-update_at")

		}
	}
	_, err := seter.Limit(limit, limit*(page-1)).All(&list)
	return list, int(count), err
}

// ############### CRUD

// 获取视频展示信息，需要提供视频id
//  @param  id [int]
//  @return [*VideoShowInfo]
func (m Video) GetVideoShowInfo(id int) *VideoShowInfo {
	video := &Video{Id: id}
	Orm.Read(video)
	vinfo := video.setVideoShowInfo()
	return vinfo
}

// 获取视频播放信息，需要提供视频id
//  @param  id [int]
//  @return [*VideoPlayInfo]
func (m Video) GetVideoPlayInfo(id int) *VideoPlayInfo {
	video := &Video{Id: id}
	Orm.Read(video)
	vinfo := video.setVideoPlayInfo()
	return vinfo
}

// video调用，返回VideoShowInfo指针
//  @return [*VideoShowInfo]
func (m Video) setVideoShowInfo() *VideoShowInfo {
	return &VideoShowInfo{
		Id:             m.Id,
		Videoname:      m.Videoname,
		Classification: ClassificationMap[m.Classification],
		Typename:       m.Typename,
		Introduction:   m.Introduction,
		Viewnum:        m.Viewnum,
		Videologo:      m.Videologo,
		Keywords:       m.Keywords,
		Averscore:      m.Averscore,
	}
}

// video调用，返回VideoPlayInfo指针
//  @return [*VideoPlayInfo]
func (m Video) setVideoPlayInfo() *VideoPlayInfo {
	return &VideoPlayInfo{
		Id:             m.Id,
		Videoname:      m.Videoname,
		Classification: ClassificationMap[m.Classification],
		Typename:       m.Typename,
		Introduction:   m.Introduction,
		Videoresource:  m.Videoresource,
		// Viewnum:       m.Viewnum,
		// Videologo:     m.Videologo,
		Keywords:  m.Keywords,
		Averscore: m.Averscore,
	}
}

// 增加视频播放量
//  @param  id [int]
//  @return [RespJson]
func (m Video) AddViewnum(id int) RespJson {
	resp := *NewRespJson()
	video := &Video{Id: id}
	if err := Orm.Read(video); err != nil {
		resp.Msg = "读取视频数据失败"
		log.Println(resp.Msg, err)
	} else {
		video.Viewnum++
		if _, err := Orm.Update(video); err != nil {
			resp.Msg = "更新视频数据失败"
			log.Println(resp.Msg, err)
		} else {
			resp.Code = DO_SUCCESS
		}
	}
	return resp
}

// 增加视频评论数
//  @param  id [int]
//  @return [RespJson]
func (m Video) AddReviewnum(id int) RespJson {
	resp := *NewRespJson()
	video := &Video{Id: id}
	if err := Orm.Read(video); err != nil {
		resp.Msg = "读取视频数据失败"
		log.Println(resp.Msg, err)
	} else {
		video.Reviewnum++
		if _, err := Orm.Update(video); err != nil {
			resp.Msg = "更新视频播放数失败"
			log.Println(resp.Msg, err)
		} else {
			resp.Code = DO_SUCCESS
		}
	}
	return resp
}

// 更新视频评分
//  @param  id [int]
//  @param  score [int]
//  @return [RespJson]
func (m Video) UpdateScore(id, score int) RespJson {
	resp := *NewRespJson()
	video := &Video{Id: id}
	if err := Orm.Read(video); err != nil {
		resp.Msg = "读取视频数据失败"
		log.Println(resp.Msg, err)
	} else {
		video.Scorenum++
		video.Totalscore = video.Totalscore + int64(score)
		video.Averscore = fmt.Sprintf("%.01f", float64(video.Totalscore)/float64(video.Scorenum))
		if _, err := Orm.Update(video); err != nil {
			resp.Msg = "更新视频评分失败"
			log.Println(resp.Msg, err)
		} else {
			resp.Code = DO_SUCCESS
		}
	}
	return resp
}

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
	v.Passed = "审核通过"
	v.State = 0
	timeStr := utils.GetNowTimeString()
	v.CreateAt, v.UpdateAt = timeStr, timeStr
	_, err := Orm.Insert(&v)
	if err != nil {
		resp.Msg = "添加视频失败<br/>" + err.Error()
		log.Println(resp.Msg, err)
	} else {
		resp.Code = DO_SUCCESS
		resp.Msg = "添加成功"
	}
	return resp
}

// 更新旧的video，使用旧的进行调用，参数为新的video
//  @param  v [*Video] 新的video
//  @return [RespJson]
func (m Video) Update(v *Video, cols ...string) RespJson {
	if len(cols) == 0 {
		cols = m.GetDifCols(*v)
	}
	resp := RespJson{Code: DO_ERROR}
	if len(cols) == 0 {
		resp.Msg = "信息未更改，更新失败"
		return resp
	}
	_, err := Orm.Update(v, cols...)
	if err != nil {
		resp.Msg = "更新信息失败"
		log.Println(resp.Msg, err)
	} else {
		resp.Msg = "更新信息成功"
		resp.Code = DO_SUCCESS
	}
	return resp
}

// 删除参数中的video，只需要提供id属性，更改为设置state属性为1
//  @param  v [*Video]
//  @return [RespJson]
func (m Video) Delete(v *Video) RespJson {
	v.State = 0
	resp := v.Update(v, "state")
	return resp
}

// ### 设置INFO

func (m Video) GetDifCols(v Video) []string {
	dif := []string{}
	if m.Classification != v.Classification {
		dif = append(dif, "classification")
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
	if m.Reviewnum != v.Reviewnum {
		dif = append(dif, "reviewnum")
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
	if value := source.Get("classification"); value != "" {
		m.Classification = value
	}
	m.Typename = ""
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
	if value := source.Get("reviewnum"); value != "" {
		m.Reviewnum = utils.Atoi(value)
	}
	if value := source.Get("averscore"); value != "" {
		m.Averscore = value
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
