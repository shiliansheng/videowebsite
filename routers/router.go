package routers

import (
	"videowebsite/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.AutoRouter(&controllers.AdminController{})
}
//http://localhost:8088/admin/index.html#/page/welcome.html