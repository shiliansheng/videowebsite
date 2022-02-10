package controllers

import (
	"strconv"
	"strings"
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
		if user.Status == models.USER_ADMIN {
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
	user := models.User{}
	user.Username = c.Input().Get("username")
	user.Password = c.Input().Get("password")
	user.Nickname = c.Input().Get("nickname")
	user.Sex = c.Input().Get("sex")
	user.Email = c.Input().Get("email")
	user.Status, _ = strconv.Atoi(c.Input().Get("status"))
	user.Remark = c.Input().Get("remark")

	if user.Nickname == "" {
		user.Nickname = "stranger"
	}
}
