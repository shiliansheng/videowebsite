package routers

import (
	"videowebsite/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.AutoRouter(&controllers.AdminController{})
	beego.AutoRouter(&controllers.CommonController{})
	beego.Router("/", &controllers.VideoController{}, "get:Home")
	beego.Router("/:module([\\w]+)/", &controllers.VideoController{}, "*:Library")
	beego.Router("/*/reloadvideo.json", &controllers.VideoController{}, "*:Reloadvideo")
	beego.Router("/reloadvideo", &controllers.VideoController{}, "*:Reloadvideo")
	beego.Router("/searcher", &controllers.VideoController{}, "*:Searcher")
	beego.Router("/play", &controllers.VideoController{}, "*:Play")
	beego.Router("/userzone", &controllers.CommonController{}, "*:Userzone")
	beego.Router("/register", &controllers.CommonController{}, "*:Register")
	// beego.Router("/captcha/*", &controllers.CommonController{}, "*:Captcha")
	beego.AutoRouter(&controllers.VideoController{})

	// beego.Router("/video/:modulename([\\w]+)[.]?*", &controllers.VideoController{})
}
