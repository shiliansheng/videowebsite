package controllers

import (
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type BaseController struct {
	beego.Controller
	Orm            orm.Ormer
	ControllerName string
	ActionName     string
}

type Responser struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (c *BaseController) Prepare() {
	controllerName, actionName := c.GetControllerAndAction()
	c.ControllerName, c.ActionName = controllerName, actionName
	c.Orm = orm.NewOrm()
	if strings.ToLower(c.ControllerName) == "admincontroller" && strings.ToLower(c.ActionName) != "login" {
		if c.GetSession("user") == nil {
			c.History("未登录", "login.html")
		}
	}
}

func (c *BaseController) History(msg string, url string) {
	if url == "" {
		c.Ctx.WriteString("<script>alert('" + msg + "');window.history.go(-1);</script>")
		c.StopRun()
	} else {
		c.Redirect(url, 302)
	}
}

func (c *BaseController) GetNickname(name string) string {
	if name == "" {
		name = "stranger"
	}
	return name
}
