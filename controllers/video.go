package controllers

import (
	"strconv"
	"videowebsite/models"

	"github.com/astaxie/beego"
)

type VideoController struct {
	BaseController
}

//////////////////////////////////// 页面函数 ////////////////////////////////////

////////////////// video type //////////////////

func (c *VideoController) Videotypelist() {
	ext := c.Ctx.Input.Param(":ext")
	action := c.Input().Get("action")
	if ext == "html" {
		c.Data["Httpport"] = beego.AppConfig.String("httpport")
		c.TplName = "video/videotypelist.html"
	} else if ext == "json" {
		if action == "getlist" {
			filtermap := make(map[string]interface{})
			page, _ := strconv.Atoi(c.Input().Get("page"))
			limit, _ := strconv.Atoi(c.Input().Get("limit"))
			vtlistJson := new(models.VideoType).GetVideoTypeListJson(page, limit, filtermap)
			c.Data["json"] = vtlistJson
		}
		c.ServeJSON()
	}
}