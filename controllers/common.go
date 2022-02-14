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
				code, msg := user.UpdateUser(user, "password")
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
			email := c.Input().Get("email")
			nickname := c.Input().Get("nickname")
			sex := c.Input().Get("sex")
			remark := c.Input().Get("remark")
			if email == user.Email && nickname == user.Nickname && sex == user.Sex && remark == user.Remark {
				resp.Code = models.U_DO_REMAIN
				resp.Msg = "信息未改变，修改失败"
			} else {
				user.Email = email
				user.Nickname = nickname
				user.Remark = remark
				user.Sex = sex
				code, msg := user.UpdateUser(user, "email", "nickname", "sex", "remark")
				if code == 0 {
					c.SetSession("user", user)
				}
				resp.Code, resp.Msg = code, msg
			}
		}
		c.Data["json"] = resp
		c.ServeJSON()
	}
	c.TplName = "common/userSetting.html"
}
