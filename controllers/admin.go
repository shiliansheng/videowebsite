package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"videowebsite/models"
	mutils "videowebsite/utils"

	"github.com/astaxie/beego"
)

type AdminController struct {
	BaseController
}

func (c *AdminController) Get() {
	if c.GetSession("user") != nil {
		c.Redirect("index.html", 302)
	} else {
		c.Redirect("login.html", 302)
	}
}

func (c *AdminController) Common() {
	pageName := c.Ctx.Input.GetData("0")
	fmt.Println(c.Ctx.Input, pageName)
	if pageName == "userSetting.html" {
		c.TplName = "../common/userSetting.html"
	} else if pageName == "userPassword" {
		c.TplName = "../common/userPassword.html"
	}
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

func (c *AdminController) Index() {
	user, _ := c.GetSession("user").(models.User)
	c.Data["Nickname"] = user.Nickname
	c.TplName = "admin/index.html"
}

func (c *AdminController) Welcome() {
	c.Data["UserCount"] = new(models.User).GetUserCount()
	c.Data["VideoCount"] = 123
	c.Data["ViewCount"] = 456
	c.TplName = "admin/welcome.html"
}

func (c *AdminController) Userlist() {
	c.Data["Httpport"] = beego.AppConfig.String("httpport")
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
		user.GetUserInfo(c.Input())
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
		user.GetUserInfo(c.Input())
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
		newUser.GetUserInfo(c.Input())
		resp := Responser{}
		resp.Code, resp.Msg = user.Update(newUser, user.GetDifCols(user, newUser)...)
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
				msg, code := c.deleteUser(user)
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

func (c *AdminController) deleteUser(user models.User) (string, int) {
	msg, code, loginUser := "", models.DO_SUCCESS, c.GetSession("user").(models.User)
	if user.Id == loginUser.Id {
		code = models.U_DEL_SELF
		msg = "删除用户 " + user.Username + " 失败<br/>禁止删除自己"
	} else if user.Status == "管理员" && loginUser.Status != "超级管理员" {
		code = models.U_DEL_MANAGER
		msg = "删除用户 " + user.Username + " 失败<br/>禁止删除管理员"
	} else {
		_, err := c.Orm.Delete(&user)
		if err != nil {
			code = models.U_DEL_ERROR
			msg = "删除用户 " + user.Username + " 失败<br/>" + err.Error()
		} else {
			msg = "删除用户 " + user.Username + " 成功"
		}
	}
	return msg, code
}
