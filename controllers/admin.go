package controllers

import (
	"strconv"
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
	// fmt.Println(c.Input())
	page, _ := strconv.Atoi(c.Input().Get("page"))
	limit, _ := strconv.Atoi(c.Input().Get("limit"))
	userListJson, err := new(models.User).GetUserListJson(page, limit)
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
			Nickname: func() string { name := c.GetNickname(c.Input().Get("nickname")); return name }(),
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

func (c *AdminController) Useredit() {
	if c.Ctx.Request.Method == "GET" {
		c.Data["ID"] = c.Input().Get("id")
		c.Data["Username"] = c.Input().Get("username")
		c.Data["Password"] = c.Input().Get("password")
		c.Data["Nickname"] = c.Input().Get("nickname")
		checkmap := map[string]int{
			"保密": 0, "男": 1, "女": 2, "普通用户": 0, "管理员": 1,
		}
		checkarr := [3][]string{
			{"checked", " ", " "}, {" ", "checked", " "}, {" ", " ", "checked"},
		}
		index := checkmap[c.Input().Get("sex")]
		c.Data["Sex_Secret"] = checkarr[index][0]
		c.Data["Sex_Male"] = checkarr[index][1]
		c.Data["Sex_Female"] = checkarr[index][2]
		c.Data["Email"] = c.Input().Get("email")

		index = checkmap[c.Input().Get("status")]
		c.Data["Status_Common"] = checkarr[index][0]
		c.Data["Status_Manager"] = checkarr[index][1]
		c.Data["Remark"] = c.Input().Get("remark")
	} else if c.Ctx.Request.Method == "POST" {
		user := models.User{Id: func() int { ret, _ := strconv.Atoi(c.Input().Get("id")); return ret }()}
		c.Orm.Read(&user)
		user.Username = c.Input().Get("username")
		user.Password = c.Input().Get("password")
		user.Nickname = func() string { name := c.GetNickname(c.Input().Get("nickname")); return name }()
		user.Sex = c.Input().Get("sex")
		user.Email = c.Input().Get("email")
		user.Status = c.Input().Get("status")
		user.Remark = c.Input().Get("remark")
		user.UpdateAt = time.Now()
		resp := make(map[string]interface{})
		_, err := c.Orm.Update(&user)
		if err != nil {
			resp["code"] = "202"
			resp["msg"] = "update user failed\n" + err.Error()
		} else {
			resp["code"] = "0"
			resp["msg"] = "update user success"
		}
		c.Data["json"] = resp
		c.ServeJSON()
	}
	c.TplName = "admin/useredit.html"
}

func (c *AdminController) Userdel() {
	if c.Ctx.Request.Method == "POST" {
		id, _ := strconv.Atoi(c.Input().Get("id"))
		resp := make(map[string]interface{})
		if id == c.GetSession("user").(models.User).Id {
			resp["code"] = "203"
			resp["msg"] = "不能删除自己"
		} else {
			user := models.User{Id: id}
			_, err := c.Orm.Delete(&user)
			if err != nil {
				resp["code"] = "204"
				resp["msg"] = "删除用户失败\n" + err.Error()
			} else {
				resp["code"] = "0"
				resp["msg"] = "删除用户成功"
			}
		}
		c.Data["json"] = resp
		c.ServeJSON()
	}
}
