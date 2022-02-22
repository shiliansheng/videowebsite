package models

import (
	"fmt"
	"net/url"
	"strconv"
	"videowebsite/utils"
)

type VideoType struct {
	Id          int    `json:"id" orm:"pk"` // video type id 视频类型编号
	Pid         int    `json:"pid"`         // 视频类型父ID
	Typename    string `json:"typename"`    // 视频类型名称
	Discription string `json:"discription"` // 视频类型描述
	Addid       int    `json:"addid"`       // 添加人员id
	Createat    string `json:"createat"`    // 类型添加时间
	Vtypelogo   string `json:"vtypelogo"`   // 类型logo
	Sequence    int    `json:"sequence"`    // 显示顺序
}

type Video struct {
	Id int `json:"id" orm:"pk"` // 视频编号
	//Typeid        int     `json:"typeid"`        // 视频类型编号
	Videoname     string `json:"videoname"`     // 视频名称
	Typename      string `json:"typename"`      // 视频类型名称
	Introduction  string `json:"introduction"`  // 视频介绍
	Videologo     string `json:"videologo"`     // 视频图片地址
	Keywords      string `json:"keywords"`      // 视频关键字
	Videoresource string `json:"videoresource"` // 视频资源地址
	Copyright     string `json:"copyright"`     // 版权所有(原创,转载)
	//Uerid         int     `json:"userid"`        // 发布者ID
	Username   string  `json:"username"`   // 发布者用户名
	Pubtime    string  `json:"pubtime"`    // 发布时间
	Viewnum    int     `json:"viewnum"`    // 视频观看次数
	Scorenum   int     `json:"scorenum"`   // 视频打分人数
	Remarknum  int     `json:"remarknum"`  // 视频评论次数
	Averscore  float64 `json:"averscore"`  // 用户评分平均分
	Totalscore int64   `json:"totalscore"` // 用户评分总分
	Passed     string  `json:"passed"`     // 审核状态(待审核;通过审核)
	Recommand  int     `json:"recommand"`  // 视频推荐(0:不推荐,1:推荐)
	// DownloadNum   int     `json:"downloadNum"`   // 视频下载次数
	// Default       int     `json:"default"`       // 视频首页显示(0:不再首页显示,1:首页显示)
}

type VideoTypeTree struct {
	Title    string          `json:"title"`
	Id       int             `json:"id"`
	Children []VideoTypeTree `json:"children"`
}

func (m *VideoType) TableName() string {
	return TableName("videotype")
}

func (m *Video) TableName() string {
	return TableName("video")
}

////////////////// 功能函数 //////////////////

// #### video type ####

func (m VideoType) GetVideoTypeInfo() (VideoType, error) {
	vt := VideoType{Id: m.Id}
	err := Orm.Read(&vt, "id")
	return vt, err
}

func (m VideoType) GetVideoTypeCount() int {
	count, _ := Orm.QueryTable(m.TableName()).Count()
	return int(count)
}

func (m VideoType) GetNameById(id int) string {
	if id == 0 {
		return "结点为根结点"
	}
	newVt := VideoType{Id: id}
	Orm.Read(&newVt)
	return newVt.Typename
}

// 获取二级列表，主要获取的内容有name, id, children
func (m VideoType) GetVideoTreeList() []*VideoTypeTree {
	vtypeList := []VideoType{}
	_, err := Orm.QueryTable(m.TableName()).All(&vtypeList)
	if err != nil {
		fmt.Println("获取视频类型列表失败", err)
	}
	vtTreeList := []*VideoTypeTree{}
	var id2node = make(map[int]*VideoTypeTree)
	// 设定只有二级菜单
	for _, vt := range vtypeList {
		if vt.Pid == 0 {
			node := VideoTypeTree{Title: vt.Typename, Id: vt.Id}
			vtTreeList = append(vtTreeList, &node)
			id2node[vt.Id] = &node
		} else {
			pnode, ok := id2node[vt.Pid]
			if !ok {
				fmt.Println("node not found...")
				continue
			}
			node := VideoTypeTree{Title: vt.Typename, Id: vt.Id}
			pnode.Children = append(pnode.Children, node)
		}
	}
	return vtTreeList
}

func (m VideoType) GetVideoTypeListJson(page, limit int, mapper map[string]interface{}) ResposeJson {
	vtJson := ResposeJson{
		Code:  DO_SUCCESS,
		Msg:   "",
		Count: 0,
		Data:  nil,
	}
	vtList, count, err := m.getVideoTypeList(page, limit, mapper)
	if err != nil {
		vtJson.Code = DO_ERROR
		vtJson.Msg = "获取视频类型列表失败<br/>" + err.Error()
		return vtJson
	}
	vtJson.Count, vtJson.Data = count, vtList
	return vtJson
}

func (m VideoType) getVideoTypeList(page, limit int, mapper map[string]interface{}) ([]VideoType, int, error) {
	var vtList []VideoType
	seter := Orm.QueryTable(m.TableName())
	for key, value := range mapper {
		seter = seter.Filter(key+"__icontains", value)
	}
	count, _ := seter.Count()
	_, err := seter.Limit(limit, limit*(page-1)).All(&vtList)
	return vtList, int(count), err
}

// #### video ####

func (m Video) GetVideoCount() int {
	cnt, _ := Orm.QueryTable(m.TableName()).Count()
	return int(cnt)
}

func (m Video) GetVideoListJson(page, limit int, mapper map[string]interface{}) ResposeJson {
	vJson := ResposeJson{
		Code: DO_SUCCESS,
	}
	var vList []Video
	seter := Orm.QueryTable(m.TableName())
	for key, value := range mapper {
		if value == "" {
			continue
		}
		seter = seter.Filter(key+"__icontains", value)
	}
	count, _ := seter.Count()
	// 返回值是什么，是得到的行数还是什么？
	_, err := seter.Limit(limit, limit*(page-1)).All(&vList)
	if err != nil {
		vJson.Code, vJson.Msg = DO_ERROR, "获取视频列表失败<br/>"+err.Error()
	} else {
		vJson.Msg = "成功获取视频列表"
	}
	vJson.Count, vJson.Data = int(count), vList
	return vJson
}

////////////////// 基础CRUD函数 //////////////////

// #### video type ####

func (m VideoType) Add(vtype VideoType) (int, string) {
	msg, code := "", DO_ERROR
	if vtype.Pid <= 0 {
		msg = "添加类型失败<br/>禁止添加根类型"
	}
	vtype.Createat = utils.GetNowTimeString()
	_, err := Orm.Insert(&vtype)
	if err != nil {
		msg = "添加视频类型失败<br/>" + err.Error()
	} else {
		code = DO_SUCCESS
		msg = "添加成功"
	}
	return code, msg
}

func (m VideoType) Delete(vtype VideoType) (string, int) {
	msg, code := "", DO_ERROR
	if vtype.Pid == 0 {
		msg = "删除" + vtype.Typename + "失败: 禁止删除根类型"
	} else {
		_, err := Orm.Delete(&vtype, "id")
		if err != nil {
			msg = "删除类型" + vtype.Typename + "失败: " + err.Error()
		} else {
			code = DO_SUCCESS
			msg = "删除类型" + vtype.Typename + "成功"
		}
	}
	return msg, code

}

// 更新videotype，使用旧的videotype进行调用，参数为新的videotype
//  @param  vtype [VideoType] 新的视频类型
//  @param  cols [...string] 更新的内容
//  @return [int] 返回resp.Code
//  @return [string] 返回resp.Msg
func (m VideoType) Update(vtype VideoType, cols ...string) (int, string) {
	if len(cols) == 0 {
		return DO_REMAIN, "信息未更改，更新失败"
	}
	_, err := Orm.Update(&vtype, cols...)
	if err != nil {
		return DO_UP_ERROR, "更新信息失败<br/>" + err.Error()
	}
	return DO_SUCCESS, "更新信息成功"
}

// #### video ####

func (m Video) Add(v Video) (string, int) {
	msg, code := "", DO_ERROR
	if v.Videoname == "" || v.Videoresource == "" || v.Typename == "" || v.Username != "" {
		return "信息不全，添加失败!", DO_ERROR
	}
	v.Pubtime = utils.GetNowTimeString()
	v.Passed = "待审核"
	_, err := Orm.Insert(&v)
	if err != nil {
		msg = "添加视频失败<br/>" + err.Error()
	} else {
		code = DO_SUCCESS
		msg = "添加成功"
	}
	return msg, code
}

// 更新旧的video，使用旧的进行调用，参数为新的video
//  @param  v [Video] 新的video
//  @param  cols [...string] 更新列
//  @return [string] resp.Msg
//  @return [int] resp.Code
func (m Video) Update(v Video, cols ...string) (string, int) {
	if len(cols) == 0 {
		return "信息未更改，更新失败", DO_REMAIN
	}
	_, err := Orm.Update(&v, cols...)
	if err != nil {
		return "更新信息失败<br/>" + err.Error(), DO_UP_ERROR
	}
	return "更新信息成功", DO_SUCCESS
}

func (m Video) Delete(v Video) (string, int) {
	msg, code := "", DO_ERROR
	_, err := Orm.Delete(&v, "id")
	if err != nil {
		msg = "删除类型" + v.Videoname + "失败: " + err.Error()
	} else {
		code = DO_SUCCESS
		msg = "删除类型" + v.Videoname + "成功"
	}
	return msg, code
}

////////////////// 辅助函数 //////////////////

// #### video type ####

func (m VideoType) GetDifCols(vt VideoType) []string {
	dif := []string{}
	if m.Typename != vt.Typename {
		dif = append(dif, "typename")
	}
	if m.Discription != vt.Discription {
		dif = append(dif, "discription")
	}
	if m.Addid != vt.Addid {
		dif = append(dif, "addid")
	}
	if m.Createat != vt.Createat {
		dif = append(dif, "createat")
	}
	if m.Vtypelogo != vt.Vtypelogo {
		dif = append(dif, "vtypelogo")
	}
	if m.Sequence != vt.Sequence {
		dif = append(dif, "sequence")
	}
	if m.Discription != vt.Discription {
		dif = append(dif, "discription")
	}
	return dif
}

// 根据给定的map内容，设置videotype，需要使用一个初始化的VideoType进行调用
//  @param  source [url.Values] 给定的map内容
func (m *VideoType) SetVideoType(source url.Values) {
	if value := source.Get("id"); value != "" {
		m.Id = utils.Atoi(value)
	}
	if value := source.Get("pid"); value != "" {
		m.Pid = utils.Atoi(value)
	}
	if value := source.Get("typename"); value != "" {
		m.Typename = value
	}
	if value := source.Get("discription"); value != "" {
		m.Discription = value
	}
	if value := source.Get("addid"); value != "" {
		m.Addid = utils.Atoi(value)
	}
	if value := source.Get("createat"); value != "" {
		m.Createat = value
	}
	if value := source.Get("vtypelogo"); value != "" {
		m.Vtypelogo = value
	}
	if value := source.Get("sequence"); value != "" {
		m.Sequence = utils.Atoi(value)
	}
}

// #### video ####

func (m Video) GetAverScore() float64 {
	return float64(m.Totalscore) * 1.0 / float64(m.Scorenum)
}

func (m Video) GetDifCols(v Video) []string {
	dif := []string{}
	// if m.Typeid != v.Typeid {
	// 	dif = append(dif, "typeid")
	// }
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
	// if m.Uerid != v.Uerid {
	// 	dif = append(dif, "userid")
	// }
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
	// if value := source.Get("typeid"); value != "" {
	// 	m.Typeid = utils.Atoi(value)
	// }
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
	// if value := source.Get("userid"); value != "" {
	// 	m.Uerid = utils.Atoi(value)
	// }
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