package models

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	dbhost := beego.AppConfig.String("dbhost")
	dbport := beego.AppConfig.String("dbport")
	dbuser := beego.AppConfig.String("dbuser")
	dbpassword := beego.AppConfig.String("dbpassword")
	dbname := beego.AppConfig.String("dbname")
	if dbport == "" {
		dbport = "3306"
	}
	dbConnStr := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8&loc=Asia%2FShanghai"
	err := orm.RegisterDataBase("default", "mysql", dbConnStr)
	if err != nil {
		fmt.Println("register database(", dbConnStr, ") failed:", err)
	}
	orm.RegisterModel(new(SystemMenu), new(User))
	// orm.RunSyncdb("default", false, true)
}

//返回带前缀的表名
func TableName(str string) string {
	prefix := beego.AppConfig.String("dbprefix")
	return prefix + "_" + str
}
