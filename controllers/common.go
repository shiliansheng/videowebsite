package controllers

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"videowebsite/models"
	"videowebsite/utils"

	"github.com/astaxie/beego"
	"github.com/dchest/captcha"
)

type CommonController struct {
	BaseController
}

// ### user

// 获取userinfo

// 通过参数Id进行获取user信息，如果不给出id参数，则返回当前登录信息的id
func (c *CommonController) Userinfo() {
	resp := models.RespJson{Code: models.DO_ERROR}
	id := c.Input().Get("id")
	var user models.User
	if id == "" {
		user = c.GetSession("user").(models.User)
	} else {
		user = models.User{Id: utils.Atoi(id)}
		c.Orm.Read(&user)
	}
	bytes, err := json.Marshal(&user)
	if err != nil {
		resp.Msg = "获取用户信息失败"
	} else {
		resp.Code = models.DO_SUCCESS
		resp.Msg = "获取信息成功"
		resp.Data = string(bytes)
	}

	c.Data["json"] = resp
	c.ServeJSON()
}

// 获取登录信息
//  @method  GET
//  @return  RespJson{Id, Name, Logo}
func (c *CommonController) Getlogininfo() {
	resp := models.RespJson{Code: models.DO_ERROR}
	userSession := c.GetSession("user")
	if userSession != nil {
		user := userSession.(models.User)
		resp.Code = models.DO_SUCCESS
		resp.Data = struct {
			Id   int    `json:"id"`
			Name string `json:"name"`
			Logo string `json:"logo"`
		}{
			Id:   user.Id,
			Name: user.Nickname,
			Logo: user.Userlogo,
		}
	}
	c.Data["json"] = resp
	c.ServeJSON()
}

func (c *CommonController) Userzone() {
	ext := c.Ctx.Input.Param(":ext")
	if ext == "json" {
		resp := models.RespJson{}
		action := c.Input().Get("action")
		if action == "update" {
			user := c.GetSession("user").(models.User)
			var newUser models.User = user
			newUser.SetUser(c.Input())
			resp = user.Update(newUser)
			if resp.Code == models.DO_SUCCESS {
				c.SetSession("user", newUser)
			}
		}
		c.Data["json"] = resp
		c.ServeJSON()
	} else {
		userSession := c.GetSession("user")
		if userSession == nil {
			c.History("", "/video/login.html")
		} else {
			user := userSession.(models.User)
			c.Data["User"] = &user
			collectInfo := new(models.Collect).GetUserCollectList(user.Id)
			c.Data["CollectInfo"] = collectInfo
			c.Data["CollectCount"] = len(*collectInfo)
			c.TplName = "common/userzone.html"
		}
	}
}

func (c *CommonController) Userset() {
	ext := c.Ctx.Input.Param(":ext")
	if ext == "json" {
		resp := models.RespJson{}
		action := c.Input().Get("action")
		if action == "update" {
			user := c.GetSession("user").(models.User)
			var newUser models.User = user
			newUser.SetUser(c.Input())
			resp = user.Update(newUser)
			if resp.Code == models.DO_SUCCESS {
				c.SetSession("user", newUser)
			}
		}
		c.Data["json"] = resp
		c.ServeJSON()
	} else {
		userSession := c.GetSession("user")
		if userSession == nil {
			c.History("", "/video/login.html")
		} else {
			user := userSession.(models.User)
			c.Data["Userlogo"] = user.Userlogo
			c.TplName = "common/userSetting.html"
		}
	}
}

func (c *CommonController) Userpwd() {
	ext := c.Ctx.Input.Param(":ext")
	if ext == "json" {
		action := c.Input().Get("action")
		user := c.GetSession("user").(models.User)
		resp := models.RespJson{Code: models.DO_ERROR}
		if action == "update" {
			old_pwd := c.Input().Get("old_password")
			new_pwd := c.Input().Get("new_password")
			if old_pwd != user.Password {
				resp.Msg = "旧密码输入不正确"
			} else {
				var newUser models.User = user
				newUser.Password = new_pwd
				resp = user.Update(newUser)
				if resp.Code == models.DO_SUCCESS {
					c.SetSession("user", newUser)
				}
			}
		}
		c.Data["json"] = resp
		c.ServeJSON()
	} else {
		c.TplName = "common/userPassword.html"
	}
}

func (c *CommonController) Logout() {
	c.SetSession("user", nil)
	c.Data["json"] = models.RespJson{Code: models.DO_SUCCESS}
	c.ServeJSON()
}

func (c *CommonController) Register() {
	ext := c.Ctx.Input.Param(":ext")
	if ext == "json" {
		user := models.User{}
		user.SetUser(c.Input())
		resp := user.Add(&user)
		c.Data["json"] = resp
		if resp.Code == models.DO_SUCCESS {
			c.SetSession("user", user)
		}
		c.ServeJSON()
	} else {
		c.TplName = "common/register.html"
	}
}

func (c *CommonController) Unameunique() {
	username := c.Input().Get("username")
	c.Data["json"] = new(models.User).UnameUnique(username)
	c.ServeJSON()
}

// ### video

// 获取videoinfo
//  通过参数Id进行获取video信息
func (c *CommonController) Videoinfo() {
	resp := models.RespJson{Code: models.DO_ERROR}
	id := c.Input().Get("id")
	var video models.Video
	if id == "" {
		resp.Msg = "获取信息失败，参数有误"
	} else {
		video = models.Video{Id: utils.Atoi(id)}
		c.Orm.Read(&video)
		bytes, err := json.Marshal(&video)
		if err != nil {
			resp.Msg = "获取视频信息失败"
		} else {
			resp.Code = models.DO_SUCCESS
			resp.Msg = "获取信息成功"
			resp.Data = string(bytes)
		}
	}
	c.Data["json"] = resp
	c.ServeJSON()
}

// ### post

func (c *CommonController) Postlist() {
	resp := new(models.Post).GetPostList()
	c.Data["json"] = resp
	c.ServeJSON()
}

// ###### captcha

func (c *CommonController) Captcha() {
	captcha.Server(140, 38)
}

// ### file

// 上传文件，调用方式：common/uploader?type=filetype-belong
func (c *CommonController) Uploader() {
	typeinfo := strings.Split(c.Input().Get("type"), "-")
	storepath := beego.AppConfig.String("storepath")
	resp := models.FileRespJson{Code: models.STD_ERROR}
	if len(typeinfo) == 0 {
		resp.Msg = "请求出错"
	} else {
		ftype, fbelong := typeinfo[0], ""
		if len(typeinfo) > 1 {
			fbelong = typeinfo[1]
		}
		file, handeler, err := c.GetFile("file")
		if err != nil {
			resp.Msg = "获取文件出错"
		} else {
			fext := filepath.Ext(handeler.Filename)
			storepath = filepath.Join(storepath, ftype, fbelong, utils.UniqueId()+fext)
			outfile, err := os.OpenFile(storepath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
			if err != nil {
				resp.Msg = "创建文件失败，请重试!</br>" + err.Error()
			} else {
				// 使用bufio
				reader := bufio.NewReader(file)
				writer := bufio.NewWriter(outfile)
				buffer := make([]byte, 1024) // 创建1K的缓冲区
				for {
					_, err := reader.Read(buffer)
					if err != nil {
						if err == io.EOF {
							break
						}
						fmt.Println(err)
					} else {
						writer.Write(buffer)
					}
				}
				writer.Flush()
				resp.Code = models.DO_SUCCESS
				resp.Data.Src = "..\\" + storepath
				resp.Msg = "上传成功"
				outfile.Close()
			}
		}
	}
	c.Data["json"] = resp
	c.ServeJSON()
}
