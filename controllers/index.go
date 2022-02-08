package controllers

import (
	"videowebsite/models"

	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

// 初始化后台框架接口
func (c *IndexController) SystemInit() {
	systemInit := new(models.SystemMenu).GetSystemInit()
	c.Data["json"] = systemInit
	c.ServeJSON()
}

func (c *IndexController) Get() {
	c.TplName = "admin/index.html"
}
