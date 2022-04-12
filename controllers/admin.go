package controllers

import (
	"encoding/json"
	"strconv"
	"strings"
	"videowebsite/models"
	"videowebsite/utils"
)

type AdminController struct {
	BaseController
}

// ### 登录

func (c *AdminController) Login() {
	ext := c.Ctx.Input.Param(":ext")
	if ext == "json" {
		resp := new(models.User).Login(c.Input().Get("username"), c.Input().Get("password"))
		if resp.Code == models.DO_SUCCESS {
			c.SetSession("user", resp.Data)
			resp.Data = "index.html"
		} else {
			resp.Data = nil
		}
		c.Data["json"] = resp
		c.ServeJSON()
	} else { // 其他后缀均为获取页面
		c.TplName = "admin/login.html"
	}
}

// ### 后台主页面

func (c *AdminController) Index() {
	ext := c.Ctx.Input.Param(":ext")
	if ext == "json" {
		c.ServeJSON()
	} else {
		// video := new (models.Video)

		user := c.GetSession("user").(models.User)
		if !strings.HasSuffix(user.Status, "管理员") {
			c.History("您的身份不是管理员", "admin/login.html")
		} else {
			c.Data["Nickname"] = user.Nickname
			c.TplName = "admin/index.html"
		}
	}
}

// 获取菜单json数据
func (c *AdminController) Menulist_api() {
	systemInit := new(models.SystemMenu).GetSystemInit()
	c.Data["json"] = systemInit
	c.ServeJSON()
}

// 获取clear.json
func (c *AdminController) Clear_api() {
	clearApi := models.RespJson{Code: 1, Msg: "服务端清理缓存成功"}
	c.Data["json"] = clearApi
	c.ServeJSON()
}

// homepage

func (c *AdminController) Homepage() {
	tmpVideo := new(models.Video)
	tmpUser := new(models.User)
	className, classValue := tmpVideo.GetClassificationCount()
	pieData := []models.PieStruct{}
	for i := range className {
		pieData = append(pieData, models.PieStruct{Name: className[i], Value: classValue[i]})
	}
	weekName, classWeekValue := tmpVideo.GetWeekUploadData()
	_, userWeekValue := tmpUser.GetWeekRegistData()
	
	c.Data["UserCount"] = tmpUser.GetUserCount()
	c.Data["VideoCount"] = tmpVideo.GetVideoCount()
	c.Data["VideoTypeCount"] = new (models.VideoType).GetVideoTypeCount()
	c.Data["classData"] = pieData
	c.Data["weekName"] = weekName
	c.Data["classWeekValue"] = classWeekValue
	c.Data["userWeekValue"] = userWeekValue
	c.TplName = "admin/homepage.html"
}

// user

func (c *AdminController) Userlist() {
	ext := c.Ctx.Input.Param(":ext")
	if ext == "json" {
		page, _ := strconv.Atoi(c.Input().Get("page"))
		limit, _ := strconv.Atoi(c.Input().Get("limit"))
		filterStr := c.Input().Get("searchParams")
		var filterMap = make(map[string]interface{})
		if filterStr != "" {
			json.Unmarshal([]byte(filterStr), &filterMap)
		}
		c.Data["json"] = new(models.User).GetUserList(page, limit, filterMap)
		c.ServeJSON()
	} else {
		c.TplName = "admin/userlist.html"
	}
}

func (c *AdminController) Useradd() {
	ext := c.Ctx.Input.Param(":ext")
	if ext == "json" {
		user := models.User{}
		user.SetUser(c.Input())
		resp := user.Add(&user)
		c.Data["json"] = resp
		c.ServeJSON()
	} else {
		c.TplName = "admin/useradd.html"
	}
}

// 删除用户，仅接收id数组，通过idlist[]获取
func (c *AdminController) Userdel() {
	// 此处不检查后缀
	resp := models.RespJson{Code: models.DO_SUCCESS}
	idlist := c.Input().Get("idlist")
	idArr, successList := []int{}, []int{}
	if err := json.Unmarshal([]byte(idlist), &idArr); err != nil {
		resp.Msg = "解析数据失败，传递数据有误<br/>" + err.Error()
		resp.Code = models.DO_ERROR
	} else {
		ulogin := c.GetSession("user").(models.User)
		for _, id := range idArr {
			tmpResp := ulogin.Delete(models.User{Id: id})
			if tmpResp.Code != models.DO_ERROR {
				successList = append(successList, id)
			}
			resp.Code |= tmpResp.Code
			resp.Msg += tmpResp.Msg + "<br/>"
		}
	}
	resp.Data = successList
	c.Data["json"] = resp
	c.ServeJSON()
}

func (c *AdminController) Useredit() {
	ext := c.Ctx.Input.Param(":ext")
	if ext == "json" {
		resp := models.RespJson{Code: models.DO_ERROR}
		user := models.User{Id: utils.Atoi(c.Input().Get("id"))}
		c.Orm.Read(&user)
		newUser := models.User(user)
		newUser.SetUser(c.Input())
		resp = user.Update(newUser)
		c.Data["json"] = resp
		c.ServeJSON()
	} else {
		id, _ := strconv.Atoi(c.Input().Get("id"))
		c.Data["Userid"] = id
		c.TplName = "admin/useredit.html"
	}
}

////// ####### video

func (c *AdminController) Videolist() {
	ext := c.Ctx.Input.Param(":ext")
	if ext == "json" {
		resp := models.RespJson{Code: models.DO_ERROR}
		limit, _ := strconv.Atoi(c.Input().Get("limit"))
		page, _ := strconv.Atoi(c.Input().Get("page"))
		filterStr := c.Input().Get("searchParams")
		var filterMap = make(map[string]interface{})
		if filterStr != "" && json.Unmarshal([]byte(filterStr), &filterMap) != nil {
			resp.Msg = "解析参数出错<br/>"
		} else {
			resp = new(models.Video).GetVideoList(page, limit, filterMap)
		}
		c.Data["json"] = resp
		c.ServeJSON()
	} else {
		c.TplName = "admin/videolist.html"
	}
}

// CRUD

func (c *AdminController) VideoAdd() {
	ext := c.Ctx.Input.Param(":ext")
	if ext == "json" {
		resp := models.RespJson{Code: models.DO_ERROR}
		userLogin := c.GetSession("user")
		if userLogin == nil {
			resp.Msg = "当前未登录，请登陆后添加"
		} else {
			video := models.Video{Username: userLogin.(models.User).Username}
			video.SetVideo(c.Input())
			resp = video.Add(video)
		}
		c.Data["json"] = resp
		c.ServeJSON()
	} else {
		typenames := new(models.VideoType).GetAllVideoTypeName()
		c.Data["Typenames"] = typenames
		c.TplName = "admin/videoadd.html"
	}
}

func (c *AdminController) Videoedit() {
	ext := c.Ctx.Input.Param(":ext")
	if ext == "json" {
		action := c.Input().Get("action")
		if action == "update" {
			video := models.Video{Id: utils.Atoi(c.Input().Get("id"))}
			video.GetInfo(&video)
			var newV models.Video = video
			newV.SetVideo(c.Input())
			resp := video.Update(&newV)
			c.Data["json"] = resp
		}
		c.ServeJSON()
	} else {
		id := c.Input().Get("id")
		c.Data["Videoid"] = id
		typenames := new(models.VideoType).GetAllVideoTypeName()
		c.Data["Typenames"] = typenames
		c.TplName = "admin/videoedit.html"
	}
}

// 删除视频，仅接收id数组，通过idlist[]获取
func (c *AdminController) Videodel() {
	// 此处不检查后缀
	resp := models.RespJson{Code: models.DO_SUCCESS}
	idlist := c.Input().Get("idlist")
	idArr, successList := []int{}, []int{}
	if err := json.Unmarshal([]byte(idlist), &idArr); err != nil {
		resp.Msg = "解析数据失败，传递数据有误<br/>" + err.Error()
		resp.Code = models.DO_ERROR
	} else {
		tmpVideo := models.Video{}
		for _, id := range idArr {
			tmpResp := tmpVideo.Delete(&models.Video{Id: id})
			if tmpResp.Code != models.DO_ERROR {
				successList = append(successList, id)
			}
			resp.Code |= tmpResp.Code
			resp.Msg += tmpResp.Msg + "<br/>"
		}
	}
	resp.Data = successList
	c.Data["json"] = resp
	c.ServeJSON()
}


// ### post

func (c *AdminController) Postlist() {
	c.TplName = "admin/postlist.html"
}