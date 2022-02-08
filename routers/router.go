package routers

import (
	"videowebsite/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/index.html", &controllers.IndexController{})
	beego.Router("/", &controllers.IndexController{})
	beego.AutoRouter(&controllers.AdminController{})
}
