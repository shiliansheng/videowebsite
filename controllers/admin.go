package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"videowebsite/models"
	mutils "videowebsite/utils"
)

type AdminController struct {
	BaseController
}

////////////////// 页面函数 //////////////////

func (c *AdminController) Get() {
	if c.GetSession("user") != nil {
		c.Redirect("index.html", 302)
	} else {
		c.Redirect("login.html", 302)
	}
}

////////////////// common //////////////////

func (c *AdminController) Common() {
	pageName := c.Ctx.Input.GetData("0")
	fmt.Println(c.Ctx.Input, pageName)
	if pageName == "userSetting.html" {
		c.TplName = "../common/userSetting.html"
	} else if pageName == "userPassword" {
		c.TplName = "../common/userPassword.html"
	}
}

////////////////// 主页面 //////////////////

func (c *AdminController) Index() {
	user, _ := c.GetSession("user").(models.User)
	c.Data["Nickname"] = user.Nickname
	c.TplName = "admin/index.html"
}

func (c *AdminController) Getmenulist() {
	systemInit := new(models.SystemMenu).GetSystemInit()
	c.Data["json"] = systemInit
	c.ServeJSON()
}

func (c *AdminController) Welcome() {
	c.Data["UserCount"] = new(models.User).GetUserCount()
	c.Data["VideoCount"] = new(models.Video).GetVideoCount()
	c.Data["VideoTypeCount"] = new(models.VideoType).GetVideoTypeCount()
	c.TplName = "admin/welcome.html"
}

func (c *AdminController) Login() {
	if c.Ctx.Request.Method == "POST" {
		username := c.GetString("username")
		password := c.GetString("password")
		user := models.User{Username: username}
		c.Orm.Read(&user, "username")
		if user.Password == "" {
			c.History("账户不存在", "")
		}
		if password != strings.Trim(user.Password, " ") {
			c.History("密码错误", "")
		}
		var aimUrl string
		if strings.HasSuffix(user.Status, "管理员") {
			aimUrl = "index.html"
		} else {
			aimUrl = ""
		}
		c.SetSession("user", user)
		c.History("登录成功", aimUrl)
	}
	c.TplName = "login.html"
}

////////////////// 用户管理界面  //////////////////

func (c *AdminController) Userlist() {
	c.TplName = "admin/userlist.html"
}

func (c *AdminController) Getuserlist() {
	var filterMap = make(map[string]interface{})
	filterString := c.Input().Get("searchParams")
	if filterString != "" {
		json.Unmarshal([]byte(filterString), &filterMap)
	}
	getNil := false
	if c.GetSession("user").(models.User).Status != "超级管理员" {
		if filterMap["status"] == "管理员" {
			getNil = true
		} else {
			filterMap["status"] = "普通用户"
		}
	}
	page, _ := strconv.Atoi(c.Input().Get("page"))
	limit, _ := strconv.Atoi(c.Input().Get("limit"))
	userListJson, err := new(models.User).GetUserListJson(page, limit, filterMap, getNil)
	if err != nil {
		c.Ctx.WriteString("<script>alert('获取用户列表失败');window.history.go(-1);</script>")
		return
	}
	c.Data["json"] = userListJson
	c.ServeJSON()
}

func (c *AdminController) Useradd() {
	ext := c.Ctx.Input.Param(":ext")
	if c.Ctx.Request.Method == "POST" {
		user := models.User{
			CreateAt: mutils.GetNowTimeString(),
			UpdateAt: mutils.GetNowTimeString(),
		}
		user.SetUser(c.Input())
		resp := Responser{}
		resp.Code, resp.Msg = user.Add(user)
		c.Data["json"] = resp
		c.ServeJSON()
	}
	if c.GetSession("user").(models.User).Status == "管理员" {
		c.Data["Disabled"] = "disabled"
	}
	if ext == "html" {
		c.TplName = "admin/useradd.html"
	}
}

func (c *AdminController) Useredit() {
	if c.Ctx.Request.Method == "GET" {
		user := models.User{}
		user.SetUser(c.Input())
		bytes, _ := json.Marshal(user)
		c.Data["UserInfoJson"] = string(bytes)
		if user.Status == "超级管理员" {
			// c.Data["SuperAdminShow"] = "1"
			c.Data["Disabled"] = "disabled"
		}
	} else if c.Ctx.Request.Method == "POST" {
		user := models.User{Id: func() int { ret, _ := strconv.Atoi(c.Input().Get("id")); return ret }()}
		c.Orm.Read(&user)
		var newUser models.User = user
		newUser.SetUser(c.Input())
		resp := Responser{}
		resp.Code, resp.Msg = user.Update(newUser, user.GetDifCols(newUser)...)
		c.Data["json"] = resp
		c.ServeJSON()
	}
	if c.GetSession("user").(models.User).Status == "管理员" {
		c.Data["Disabled"] = "disabled"
	}
	c.TplName = "admin/useredit.html"
}

func (c *AdminController) Userdel() {
	if c.Ctx.Request.Method == "POST" {
		more := c.Input().Get("more")
		resp := Responser{}
		userlistString, userlist, successlist := "", []models.User{}, []int{}
		endmsg, endcode := "", 0
		for k := range c.Input() {
			if k == "more" {
				continue
			}
			userlistString = k
		}
		if more == "false" {
			userlistString = "[" + userlistString + "]"
		}
		err := json.Unmarshal([]byte(userlistString), &userlist)
		if err != nil {
			endcode = models.DO_JSON_ERR
			endmsg = "解析数据失败，传递数据有误"
		} else {
			for _, user := range userlist {
				ulogin := c.GetSession("user").(models.User)
				msg, code := ulogin.Delete(user)
				if code == 0 {
					successlist = append(successlist, user.Id)
				}
				endcode += code
				endmsg += msg + "<br/>"
			}
		}
		resp.Code, resp.Msg, resp.Data = endcode, endmsg, successlist
		c.Data["json"] = resp
		c.ServeJSON()
	}
}

////////////////// 视频类型管理界面  //////////////////

func (c *AdminController) Videotypelist() {
	ext := c.Ctx.Input.Param(":ext")
	action := c.Input().Get("action")
	if ext == "html" {
		c.Data["NoPicPath"] = c.getImageSrc("")
		c.TplName = "admin/videotypelist.html"
	} else if ext == "json" {
		if action == "getlist" {
			filtermap := make(map[string]interface{})
			filterString := c.Input().Get("searchParams")
			if filterString != "" {
				json.Unmarshal([]byte(filterString), &filtermap)
			}
			page, _ := strconv.Atoi(c.Input().Get("page"))
			limit, _ := strconv.Atoi(c.Input().Get("limit"))
			vtlistJson := new(models.VideoType).GetVideoTypeListJson(page, limit, filtermap)
			c.Data["json"] = vtlistJson
		} else if action == "gettree" {
			vtTreeList := new(models.VideoType).GetVideoTreeList()
			c.Data["json"] = vtTreeList
		}
		c.ServeJSON()
	}
}

func (c *AdminController) Videotypeedit() {
	action := c.Input().Get("action")
	if action == "getinfo" {
		id, _ := strconv.Atoi(c.Input().Get("id"))
		vtype := models.VideoType{Id: id}
		vtype, _ = vtype.GetVideoTypeInfo()
		c.Data["Vtype"] = &vtype
		if vtype.Id == 0 {
			c.Data["Disabled"] = "disabled"
		} else {
			c.Data["Disabled"] = ""
		}
		c.Data["PidName"] = vtype.GetNameById(vtype.Pid)
		c.Data["InitTypeLogoPath"] = c.getImageSrc(vtype.Vtypelogo)
	} else if action == "update" {
		vtype := models.VideoType{Id: mutils.Atoi(c.Input().Get("id"))}
		c.Orm.Read(&vtype)
		var newType models.VideoType = vtype
		newType.SetVideoType(c.Input())
		resp := Responser{}
		resp.Code, resp.Msg = vtype.Update(newType, vtype.GetDifCols(newType)...)
		c.Data["json"] = resp
		c.ServeJSON()
	}
	c.TplName = "admin/videotypeedit.html"
}

func (c *AdminController) Videotypeadd() {
	ext := c.Ctx.Input.Param(":ext")
	if ext == "html" {
		c.Data["InitTypeLogoPath"] = c.getImageSrc("")
		c.TplName = "admin/videotypeadd.html"
	} else if ext == "json" {
		action := c.Input().Get("action")
		resp := Responser{Code: 0, Msg: ""}
		if action == "add" {
			vtype := models.VideoType{}
			vtype.SetVideoType(c.Input())
			userSession := c.GetSession("user")
			if userSession == nil {
				resp.Code, resp.Msg = models.DO_ERROR, "当前未登录，请登录后添加"
			} else {
				vtype.Addid = userSession.(models.User).Id
				resp.Code, resp.Msg = vtype.Add(vtype)
			}
		}
		c.Data["json"] = resp
		c.ServeJSON()
	}
}

func (c *AdminController) Videotypedel() {
	more := c.Input().Get("more")
	resp := Responser{}
	reqString, vtypelist, successlist := c.Input().Get("Data"), []models.VideoType{}, []int{}
	finalmsg, finalcode := "", 0
	if more == "false" {
		reqString = "[" + reqString + "]"
	}
	err := json.Unmarshal([]byte(reqString), &vtypelist)
	if err != nil {
		finalcode = models.DO_JSON_ERR
		finalmsg = "解析数据失败，传递数据有误"
	} else {
		var tmp = new(models.VideoType)
		for _, vtype := range vtypelist {
			msg, code := tmp.Delete(vtype)
			if code == 0 {
				successlist = append(successlist, vtype.Id)
			}
			finalcode += code
			finalmsg += msg + "<br/>"
		}
	}
	resp.Code, resp.Msg, resp.Data = finalcode, finalmsg, successlist
	c.Data["json"] = resp
	c.ServeJSON()
}

////////////////// 视频管理界面 //////////////////

func (c *AdminController) Videolist() {
	ext := c.Ctx.Input.Param(":ext")
	if ext == "html" {
		c.Data["NoPicPath"] = c.getImageSrc("")
		c.TplName = "admin/videolist.html"
	} else if ext == "json" {
		action := c.Input().Get("action")
		if action == "getlist" {
			filtermap := make(map[string]interface{})
			filterString := c.Input().Get("searchParams")
			if filterString != "" {
				json.Unmarshal([]byte(filterString), &filtermap)
			}
			page, _ := strconv.Atoi(c.Input().Get("page"))
			limit, _ := strconv.Atoi(c.Input().Get("limit"))
			vlistJson := new(models.Video).GetVideoListJson(page, limit, filtermap)
			c.Data["json"] = vlistJson
		}
		c.ServeJSON()
	}
}

func (c *AdminController) Videoadd() {
	ext := c.Ctx.Input.Param(":ext")
	if ext == "html" {
		c.TplName = "admin/videoadd.html"
	} else if ext == "json" {
		action := c.Input().Get("action")
		resp := Responser{}
		if action == "add" {
			userSession := c.GetSession("user")
			if userSession == nil {
				resp.Code, resp.Msg = models.DO_ERROR, "当前未登录，请登录后添加"
			} else {
				video := models.Video{Username: userSession.(models.User).Username}
				video.SetVideo(c.Input())
				resp.Msg, resp.Code = video.Add(video)
			}
		}
		c.Data["json"] = resp
		c.ServeJSON()
	}
}

func (c *AdminController) Videoedit() {
	ext := c.Ctx.Input.Param(":ext")
	if ext == "html" {
		vid := mutils.Atoi(c.Input().Get("id"))
		video := models.Video{Id: vid}
		c.Orm.Read(&video)
		c.Data["Video"] = video
		c.Data["NoPicPath"] = c.getImageSrc("")
		c.TplName = "admin/videoedit.html"
	} else if ext == "json" {
		action := c.Input().Get("action")
		resp := Responser{}
		if action == "update" {
			video := models.Video{Id: mutils.Atoi(c.Input().Get("id"))}
			c.Orm.Read(&video)
			var newV models.Video = video
			newV.SetVideo(c.Input())
			resp.Msg, resp.Code = video.Update(newV, video.GetDifCols(newV)...)
			c.Data["json"] = resp
		}
		c.ServeJSON()
	}
}

func (c *AdminController) Videodel() {
	more := c.Input().Get("more")
	resp := Responser{}
	reqString, vlist, successlist := c.Input().Get("Data"), []models.Video{}, []int{}
	finalmsg, finalcode := "", 0
	if more == "false" {
		reqString = "[" + reqString + "]"
	}
	err := json.Unmarshal([]byte(reqString), &vlist)
	if err != nil {
		finalcode = models.DO_JSON_ERR
		finalmsg = "解析数据失败，传递数据有误"
	} else {
		var tmp = new(models.Video)
		for _, video := range vlist {
			msg, code := tmp.Delete(video)
			if code == 0 {
				successlist = append(successlist, video.Id)
			}
			finalcode += code
			finalmsg += msg + "<br/>"
		}
	}
	resp.Code, resp.Msg, resp.Data = finalcode, finalmsg, successlist
	c.Data["json"] = resp
	c.ServeJSON()
}