package controllers

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"videowebsite/models"

	"github.com/astaxie/beego"
	"github.com/google/uuid"
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
		if c.Data["LogoimgPath"] = "../" + user.Userlogo; user.Userlogo == "" {
			c.Data["LogoimgPath"] = "../" + filepath.Join(beego.AppConfig.String("storepath"), beego.AppConfig.String("nopic_path"))
		}
	} else if ext == "json" {
		action := c.Input().Get("action")
		user := c.GetSession("user").(models.User)
		resp := Responser{}
		if action == "changeSetting" {
			var newUser models.User = user
			newUser.SetUser(c.Input())
			colarr := user.GetDifCols(newUser)
			if len(colarr) == 0 {
				resp.Code = models.DO_REMAIN
				resp.Msg = "信息未改变，修改失败"
			} else {
				resp.Code, resp.Msg = user.Update(newUser, colarr...)
				if resp.Code == 0 {
					c.SetSession("user", newUser)
				}
			}
		}
		c.Data["json"] = resp
		c.ServeJSON()
	}
	c.TplName = "common/userSetting.html"
}

func (c *CommonController) Uploader() {
	info := strings.Split(c.Input().Get("type"), "-")
	storepath := beego.AppConfig.String("storepath")
	resp := Responser{}
	file, _ := os.OpenFile("./out.txt", os.O_CREATE|os.O_RDWR, 0666)
	file.WriteString(fmt.Sprint(c.GetFile("upfile")))
	if info[1] == "image" {
		storepath = filepath.Join(storepath, info[1], info[0], uuid.NewString()+".jpg")
		base64data := c.Input().Get(info[1])
		base64data = base64data[strings.Index(base64data, ",")+1:]
		source, err := base64.StdEncoding.DecodeString(base64data)
		if err != nil {
			resp.Code = models.DO_ERROR
			resp.Msg = "图片解码失败<br/>" + err.Error()
		} else {
			err = ioutil.WriteFile(storepath, source, 0666)
			if err != nil {
				resp.Code = models.DO_ERROR
				resp.Msg = "图片上传失败<br/>" + err.Error()
			} else {
				resp.Code = models.DO_SUCCESS
				resp.Msg = "图片上传成功"
				resp.Data = c.getImageSrc(storepath)
			}
		}
	}
	c.Data["json"] = resp
	c.ServeJSON()
}

// 上传图片，接口格式：filetype-belong
func (c *CommonController) Uploadfile() {
	reqinfo := strings.Split(c.Input().Get("type"), "-")
	storepath := beego.AppConfig.String("storepath")
	resp := FileResponser{}
	if len(reqinfo) == 0 {
		resp.Msg = "请求出错"
		resp.Code = models.DO_REQUST_ERR
	} else {
		ftype, fbelong := reqinfo[0], ""
		if len(reqinfo) > 1 {
			fbelong = reqinfo[1]
		}
		file, handler, err := c.GetFile("file")
		if err != nil {
			resp.Msg = "获取文件出错"
			resp.Code = models.DO_OBTAIN_ERR
		} else {
			fext := filepath.Ext(handler.Filename)
			storepath = filepath.Join(storepath, ftype, fbelong, uuid.NewString()+fext)
			outfile, err := os.OpenFile(storepath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
			if err != nil {
				resp.Code = models.DO_ERROR
				resp.Msg = "创建文件失败，请重试!<br/>" + err.Error()
			} else {
				reader := bufio.NewReader(file)
				writer := bufio.NewWriter(outfile)
				buffer := make([]byte, 1024*1024) // 创建1K的缓冲区
				for {
					_, err := reader.Read(buffer)
					if err != nil {
						//读到文件末尾时退出读写循环
						if err == io.EOF {
							break
						}
						fmt.Println(err)
					} else {
						//将桶中的数据写出到目标文件
						writer.Write(buffer)
					}
				}
				writer.Flush()
				resp.Data.Src = "..\\" + storepath
				resp.Msg = "上传图片成功！"
			}
			defer outfile.Close()
		}
	}
	c.Data["json"] = resp
	c.ServeJSON()
}
