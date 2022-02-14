package controllers

import (
	"encoding/json"
	"videowebsite/models"
)

type CommonController struct {
	BaseController
}

func (c *CommonController) User_password() {
	if c.Ctx.Input.Param(":ext") == "json" {
		action := c.Input().Get("action")
		user := c.GetSession("user").(models.User)
		resp := Responser{}
		if action == "changePassword" {
			oldPass := c.Input().Get("old_password")
			newPass := c.Input().Get("new_password")
			if oldPass != user.Password {
				resp.Code = models.U_PASS_WRONG
				resp.Msg = "旧密码输入不正确"
			} else {
				user.Password = newPass
				code, msg := user.Update(user, "password")
				resp.Code, resp.Msg = code, msg
			}
		}
		c.Data["json"] = resp
		c.ServeJSON()
	}
	c.TplName = "common/userPassword.html"
}

func (c *CommonController) User_setting() {
	ext := c.Ctx.Input.Param(":ext")
	if ext == "html" {
		user := c.GetSession("user").(models.User)
		bytes, _ := json.Marshal(user)
		c.Data["UserInfoJson"] = string(bytes)
	} else if ext == "json" {
		action := c.Input().Get("action")
		user := c.GetSession("user").(models.User)
		resp := Responser{}
		if action == "changeSetting" {
			var newUser models.User = user
			newUser.Email = c.Input().Get("email")
			newUser.Nickname = c.Input().Get("nickname")
			newUser.Sex = c.Input().Get("sex")
			newUser.Remark = c.Input().Get("remark")
			newUser.Birthday = c.Input().Get("birthday")
			newUser.Introduction = c.Input().Get("introduction")
			colarr := user.GetDifCols(user, newUser)
			if len(colarr) == 0 {
				resp.Code = models.DO_REMAIN
				resp.Msg = "信息未改变，修改失败"
			} else {
				code, msg := user.Update(newUser, colarr...)
				if code == 0 {
					c.SetSession("user", newUser)
				}
				resp.Code, resp.Msg = code, msg
			}
		}
		c.Data["json"] = resp
		c.ServeJSON()
	}
	c.TplName = "common/userSetting.html"
}
