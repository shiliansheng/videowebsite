package models

type VideoType struct {
	Typeid      int    `json:"typeid"`      // video type id 视频类型编号
	Typename    string `json:"typename"`    // 视频类型名称
	Discription string `json:"discription"` // 视频类型描述
	Addid       int    `json:"addid"`       // 添加人员id
	Adddatetime string `json:"adddatetime"` // 类型添加时间
	Logo        string `json:"logo"`        // 类型logo
	Sequence    int    `json:"sequence"`    // 显示顺序
}

type Video struct {
	Videoid       int     `json:"videoid"`       // 视频编号
	Typeid        int     `json:"typeid"`        // 视频类型编号
	Typename      string  `json:"typename"`      // 视频类型名称
	Videoname     string  `json:"videoname"`     // 视频名称
	Introduction  string  `json:"introduction"`  // 视频介绍
	Videologo     string  `json:"videologo"`     // 视频图片地址
	Keywords      string  `json:"keywords"`      // 视频关键字
	VideoResource string  `json:"videoResource"` // 视频资源地址
	Copyright     int     `json:"copyright"`     // 版权所有(0:原创,1:转载)
	Uerid         int     `json:"userid"`        // 发布者编号
	Username      string  `json:"username"`      // 发布者用户名
	Pubtime       string  `json:"pubtime"`       // 发布时间
	ViewNum       int     `json:"viewNum"`       // 视频观看次数
	ScoreNum      int     `json:"scoreNum"`      // 视频打分人数
	RemarkNum     int     `json:"remarkNum"`     // 视频回复次数
	AverScore     float32 `json:"averScore"`     // 用户评分平均分
	TotalScore    float32 `json:"totalScore"`    // 用户评分总分
	Passed        string  `json:"passed"`        // 审核状态(W:待审核;Y:通过审核)
	Recommend     int     `json:"recommend"`     // 视频推荐(0:不推荐,1:推荐)
	// DownloadNum   int     `json:"downloadNum"`   // 视频下载次数
	// Default       int     `json:"default"`       // 视频首页显示(0:不再首页显示,1:首页显示)
}

func (m *VideoType) TableName() string {
	return TableName("video_type")
}

func (m *Video) TableName() string {
	return TableName("video")
}

func (m *VideoType) Update(vtype VideoType, cols ...string) (int, string) {
	if len(cols) == 0 {
		return DO_REMAIN, "信息未更改，更新失败"
	}
	_, err := Orm.Update(&vtype, cols...)
	if err != nil {
		return DO_UP_ERROR, "更新信息失败<br/>" + err.Error()
	}
	return DO_SUCCESS, "更新信息成功"
}

func (m *Video) Update(v Video, cols ...string) (int, string) {
	if len(cols) == 0 {
		return DO_REMAIN, "信息未更改，更新失败"
	}
	_, err := Orm.Update(&v, cols...)
	if err != nil {
		return DO_UP_ERROR, "更新信息失败<br/>" + err.Error()
	}
	return DO_SUCCESS, "更新信息成功"
}

func (m VideoType) GetDifCols(u VideoType) []string {
	dif := []string{}
	if m.Typename != u.Typename {
		dif = append(dif, "typename")
	}
	if m.Discription != u.Discription {
		dif = append(dif, "discription")
	}
	if m.Addid != u.Addid {
		dif = append(dif, "addid")
	}
	if m.Adddatetime != u.Adddatetime {
		dif = append(dif, "adddatetime")
	}
	if m.Logo != u.Logo {
		dif = append(dif, "logo")
	}
	if m.Sequence != u.Sequence {
		dif = append(dif, "sequence")
	}
	return dif
}

func (m Video) GetAverScore() float32 {
	return float32(m.TotalScore * 1.0 / float32(m.ScoreNum))
}

func (m Video) GetDifCols(v Video) []string {
	dif := []string{}
	if m.Typeid != v.Typeid {
		dif = append(dif, "typeid")
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
	if m.VideoResource != v.VideoResource {
		dif = append(dif, "videoResource")
	}
	if m.Copyright != v.Copyright {
		dif = append(dif, "copyright")
	}
	if m.Uerid != v.Uerid {
		dif = append(dif, "userid")
	}
	if m.Username != v.Username {
		dif = append(dif, "username")
	}
	if m.Pubtime != v.Pubtime {
		dif = append(dif, "pubtime")
	}
	if m.ViewNum != v.ViewNum {
		dif = append(dif, "viewNum")
	}
	if m.ScoreNum != v.ScoreNum {
		dif = append(dif, "scoreNum")
	}
	if m.RemarkNum != v.RemarkNum {
		dif = append(dif, "remarkNum")
	}
	if m.AverScore != v.AverScore {
		dif = append(dif, "averScore")
	}
	if m.TotalScore != v.TotalScore {
		dif = append(dif, "totalScore")
	}
	if m.Passed != v.Passed {
		dif = append(dif, "passed")
	}
	if m.Recommend != v.Recommend {
		dif = append(dif, "recommend")
	}
	return dif
}
