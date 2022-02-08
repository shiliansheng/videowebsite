package controllers

import (
	"fmt"
	"strings"
	"videowebsite/models"
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
		fmt.Println(user)
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
		// c.SetSession("user", user)
		c.History("登录成功", aimUrl)
	}
	c.TplName = "login.html"
}

func (c *AdminController) Index() {
	c.TplName = "admin/index.html"
}
