package routers

import (
	"videowebsite/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.AutoRouter(&controllers.AdminController{})
	// beego.Router("common/*", &controllers.CommonController{})
	beego.AutoRouter(&controllers.CommonController{})
	beego.Router("/", &controllers.VideoController{})
	beego.AutoRouter(&controllers.VideoController{})
}
//http://localhost:8088/admin/index.html#/page/welcome.html