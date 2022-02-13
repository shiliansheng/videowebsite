package controllers

import (
	"encoding/json"
	"fmt"
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
	if c.Ctx.Input.Param(":ext") == "html" {
		user := c.GetSession("user").(models.User)
		bytes, _ := json.Marshal(user)
		fmt.Println(string(bytes))
		c.Data["UserInfoJson"] = string(bytes)
	}
	c.TplName = "common/userSetting.html"
}