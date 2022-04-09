package controllers

import (
	"videowebsite/models"
	"videowebsite/utils"
)

type VideoController struct {
	BaseController
}

// ################ 网页首页 home页
func (c *VideoController) Home() {
	c.Data["TopVideo"] = new(models.Video).GetHotVideos(6)
	c.Data["RecommandVideo"] = new(models.Video).GetHotVideos(14)
	c.TplName = "video/home.html"
}

// ################ library页

func (c *VideoController) Library() {
	module := c.Ctx.Input.Param(":module")
	if _, ok := models.ClassificationMap[module]; !ok && module != "s" {
		c.TplName = "common/404.html"
	} else {
		ext := c.Ctx.Input.Param(":ext")
		if ext == "json" {

		} else {
			// c.Data["VideoCount"] = new(models.Video).GetVideoCount(modules...)
			// c.Data["Videolist"] = new(models.Video).GetHotVideos(42, modules...)
			c.Data["Typenames"] = new(models.VideoType).GetAllVideoTypeName()
			c.TplName = "video/library.html"
		}
	}
}

// ################ login

// 进行登录, 登录成功设置session
//  @suffix json, html(...)
//  @method GET
//  @param  json: username, password
//  @return
func (c *VideoController) Login() {
	ext := c.Ctx.Input.Param(":ext")
	if ext == "json" {
		resp := new(models.User).Login(c.Input().Get("username"), c.Input().Get("password"))
		if resp.Code == models.DO_SUCCESS {
			// user := resp.Data.(models.User)
			c.SetSession("user", resp.Data)
			resp.Data = "/"
		} else {
			resp.Data = nil
		}
		c.Data["json"] = resp
		c.ServeJSON()
	} else {
		c.TplName = "video/login.html"
	}
}

// ################ video

// 根据筛选和排序内容返回视频列表
//  @method  GET
//  @param   limit
//  @param   page
//  @param   sort
//  @param   filter
//  @param   classification
//  @return  RespJson.Data=[]VideoShowInfo
func (c *VideoController) Reloadvideo() {
	vsort := c.Input().Get("sort")
	limit := utils.Atoi(c.Input().Get("limit"))
	page := utils.Atoi(c.Input().Get("page"))
	classification := c.Input().Get("classification")
	var mapper = make(map[string]interface{})

	mapper["typename"] = c.Input().Get("filter")
	if classification != "library" {
		mapper["classification"] = classification
	}
	mapper["typename"] = c.Input().Get("filter")

	resp := new(models.Video).GetSortVideoShowInfoList(page, limit, mapper, vsort)
	c.Data["json"] = resp
	c.ServeJSON()
}

// 根据筛选和排序内容返回视频列表
//  @method  GET
//  @param   limit
//  @param   sort
//  @param   classification
//  @return  RespJson
func (c *VideoController) Gethotvideo() {
	// limit := utils.Atoi(c.Input().Get("limit"))
	// sort := "view"

}

// searcher

// 进行搜索
//  @method  GET
//  @param   search 搜索内容，主要为
//  @return  RespJson.Data=[]VideoShowInfo
func (c *VideoController) Searcher() {
	search := c.Input().Get("search")
	page := 1
	limit := 999
	var mapper = map[string]interface{} {
		"videoname": search,
	}
	resp := new (models.Video).GetSortVideoShowInfoList(page, limit, mapper)
	c.Data["json"] = resp
	c.ServeJSON()
}

// 创建Play页
//  @method  GET
//  @param   Video, Reviewlist, Hotlist
//  @return
func (c *VideoController) Play() {
	id := utils.Atoi(c.Input().Get("id"))
	tmpv := models.Video{}
	tmpv.GetInfo(&tmpv)
	vPlayInfo := tmpv.GetVideoPlayInfo(id)
	c.Data["Video"] = vPlayInfo
	reviewList := new(models.Review).GetVideoReviewInfoList(id, 1, 10)
	c.Data["Reviewlist"] = reviewList
	hotList := tmpv.GetHotVideos(10)
	c.Data["Hotlist"] = hotList
	c.TplName = "video/play.html"
}

// ################ review

// 提交评论
//  @method POST
//  @param  userid, videoid, content
//  @return RespJson
func (c *VideoController) Submitreview() {
	userid := utils.Atoi(c.Input().Get("userid"))
	videoid := utils.Atoi(c.Input().Get("videoid"))
	content := c.Input().Get("content")
	review := &models.Review{
		Userid:  userid,
		Videoid: videoid,
		Content: content,
	}
	resp := review.Add(review)
	c.Data["json"] = resp
	c.ServeJSON()
}

// ################ score

// 对score进行操作
//  @method POST(add)、GET(getinfo)
//  @param userid, videoid
//  @param POST: value 评分的值
//  @return RespJson.data=value
func (c *VideoController) Scorer() {
	userId := utils.Atoi(c.Input().Get("userid"))
	videoId := utils.Atoi(c.Input().Get("videoid"))
	method := c.Ctx.Request.Method
	var resp models.RespJson
	if method == "POST" {
		value := utils.Atoi(c.Input().Get("value"))
		score := &models.Score{
			Userid:  userId,
			Videoid: videoId,
			Value:   value,
		}
		resp = score.Add(score)
	} else if method == "GET" {
		score := &models.Score{
			Userid:  userId,
			Videoid: videoId,
		}
		resp = score.GetValue(score)
	}
	c.Data["json"] = resp
	c.ServeJSON()
}

// ################ collect

// 获取用户收藏列表，需要提供用户id
//  @method  GET
//  @param   id(userid)
//  @return  RespJson 其中Data为CollectInfo结构体数组
func (c *VideoController) Getucollect() {
	uid := utils.Atoi(c.Input().Get("id"))
	collectList := new(models.Collect).GetUserCollectList(uid)
	resp := models.NewRespJson()
	if collectList == nil {
		resp.Msg = "获取用户收藏列表失败"
	} else {
		resp.Code = models.DO_SUCCESS
		resp.Msg = "获取用户收藏列表成功"
		resp.Count = len(*collectList)
		resp.Data = *collectList
	}
	c.Data["json"] = resp
	c.ServeJSON()
}

// json类型进行获取，增删collect
//  @mathod  POST、GET、DELETE
//  @param   DELETE: collectid, GET、POST(add): userid 和 videoid
//  @return  RespJson, 其中resp.data=collectid
func (c *VideoController) Collecter() {
	method := c.Ctx.Request.Method
	var resp models.RespJson
	switch method {
	case "POST":
		collect := &models.Collect{
			Userid:  utils.Atoi(c.Input().Get("userid")),
			Videoid: utils.Atoi(c.Input().Get("videoid")),
		}
		resp = collect.Add(collect)
	case "GET":
		collect := &models.Collect{
			Userid:  utils.Atoi(c.Input().Get("userid")),
			Videoid: utils.Atoi(c.Input().Get("videoid")),
		}
		resp = collect.Get(collect)
	case "DELETE":
		collect := &models.Collect{
			Id: utils.Atoi(c.Input().Get("collectid")),
		}
		resp = collect.Delete(collect)
	}
	c.Data["json"] = resp
	c.ServeJSON()
}

// 删除collect
//  @method  POST
//  @param   id(collectid)
//  @return
func (c *VideoController) Deletecollect() {
	cid := utils.Atoi(c.Input().Get("id"))
	collect := &models.Collect{Id: cid}
	resp := collect.Delete(collect)
	c.Data["json"] = resp
	c.ServeJSON()
}

// ################ history

// 获取用户观看历史列表，需要提供用户id
//  @method GET
//  @param  id(userid)
//  @return RespJson 其中Data为HistoryInfo结构体数组
func (c *VideoController) Getuhistory() {
	uid := utils.Atoi(c.Input().Get("id"))
	historyList := new(models.History).GetUserHistoryList(uid)
	resp := models.NewRespJson()
	if historyList == nil {
		resp.Msg = "获取用户观看历史列表失败"
	} else {
		resp.Code = models.DO_SUCCESS
		resp.Msg = "获取用户观看历史列表成功"
		resp.Count = len(*historyList)
		resp.Data = *historyList
	}
	c.Data["json"] = resp
	c.ServeJSON()
}

// 增加video.viewnum，如果登录了，增加用户历史记录
//  @method  POST
//  @param   id(videoid)
//  @return
func (c *VideoController) Addviewnum() {
	uSesson := c.GetSession("user")
	vid := utils.Atoi(c.Input().Get("id"))
	if uSesson != nil {
		history := &models.History{
			Userid:  uSesson.(models.User).Id,
			Videoid: vid,
		}
		history.Add(history)
	}
	c.Data["json"] = new(models.Video).AddViewnum(vid)
	c.ServeJSON()
}

// 删除history
//  @method  POST
//  @param   id(historyid)
//  @return
func (c *VideoController) Deletehistory() {
	cid := utils.Atoi(c.Input().Get("id"))
	history := &models.History{Id: cid}
	resp := history.Delete(history)
	c.Data["json"] = resp
	c.ServeJSON()
}

//
//  @method
//  @param
//  @return
