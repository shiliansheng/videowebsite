package models

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

var Orm orm.Ormer

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

	Orm = orm.NewOrm()

	systemInit := new(SystemMenu).GetSystemInit()
	menuJson, _ := json.Marshal(systemInit)
	path := beego.AppConfig.String("menuinitpath")
	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("open file failed:", err)
	}
	file.Write(menuJson)
}

//返回带前缀的表名
func TableName(str string) string {
	prefix := beego.AppConfig.String("dbprefix")
	return prefix + "_" + str
}
