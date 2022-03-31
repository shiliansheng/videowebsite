package routers

import (
	"videowebsite/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.AutoRouter(&controllers.AdminController{})
	beego.AutoRouter(&controllers.CommonController{})
	beego.Router("/", &controllers.VideoController{})
	// beego.AutoRouter(&controllers.VideoController{})
}
