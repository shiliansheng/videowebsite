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

func (c *CommonController) User_setting() {
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
		user := c.GetSession("user").(models.User)
		bytes, _ := json.Marshal(user)
		c.Data["UserInfoJson"] = string(bytes)
		c.Data["Userlogo"] = user.Userlogo
		c.TplName = "common/userSetting.html"
	}
}

func (c *CommonController) User_password() {
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
				user.Password = new_pwd
				resp = user.Update(user)
			}
		}
		c.Data["json"] = resp
		c.ServeJSON()
	} else {
		c.TplName = "common/userPassword.html"
	}
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

// ### file

// 上传文件，调用方式：common/uploader?type=filetype-belong
func (c *CommonController) Uploader() {
	typeinfo := strings.Split(c.Input().Get("type"), "-")
	storepath := beego.AppConfig.String("storepath")
	resp := models.FileRespJson{Code: models.DO_ERROR}
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
				buffer := make([]byte, 1024*1024) // 创建1K的缓冲区
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
