package controllers

import (
	"strings"
	"time"
	"videowebsite/models"

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
		if user.Status == "管理员" {
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
	userListJson, err := new(models.User).GetUserListJson()
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
			Username: c.Input().Get("username"),
			Password: c.Input().Get("password"),
			Nickname: func() string {
				name := c.Input().Get("nickname")
				if name == "" {
					name = "stranger"
				}
				return name
			}(),
			Sex:      c.Input().Get("sex"),
			Email:    c.Input().Get("email"),
			Status:   c.Input().Get("status"),
			Remark:   c.Input().Get("remark"),
			CreateAt: time.Now(),
			UpdateAt: time.Now(),
		}
		resp := make(map[string]interface{})
		_, err := c.Orm.Insert(&user)
		if err != nil {
			resp["code"] = "201"
			resp["msg"] = "add user failed\n" + err.Error()
		} else {
			resp["code"] = "0"
			resp["msg"] = "add user success"
		}
		c.Data["json"] = resp
		c.ServeJSON()
	}
	if ext == "html" {
		c.TplName = "admin/useradd.html"
	}
}
