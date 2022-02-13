package routers

import (
	"videowebsite/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.AutoRouter(&controllers.AdminController{})
	// beego.Router("common/*", &controllers.CommonController{})
	beego.AutoRouter(&controllers.CommonController{})
}
//http://localhost:8088/admin/index.html#/page/welcome.html