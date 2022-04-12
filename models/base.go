package models

import (
	"fmt"
	"log"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type RespJson struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Count int         `json:"count,omitempty"`
	Data  interface{} `json:"data,omitempty"`
}

type FileRespJson struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Src string `json:"src"`
	} `json:"data"`
}

type VideoRespJson struct {
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	Count int    `json:"count"` // 总数
	Page  int    `json:"page"`
	Data  struct {
		Size    int      `json:"size"` // 当前限制数目
		Id      []int    `json:"id"`   // 视频id
		Name    []string `json:"name"` // 视频名
		Logo    []string `json:"logo"` // 视频图
		Type    []string `json:"type"` // 视频类型
		Score   []string `json:"score"` // 视频评分
	} `json:"data"`
}

type PieStruct struct {
	Name string `json:"name"`
	Value int `json:"value"`
}

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
	orm.RegisterModel(new(SystemMenu), new(User), new(Video), new(VideoType), new(Review), new(Score), new(Collect), new(History), new(Post))

	Orm = orm.NewOrm()
	log.SetFlags(log.LstdFlags | log.Llongfile)
}

// 返回带前缀的表名
//  @param  str [string]
//  @return [string] 前缀-str
func TableName(str string) string {
	prefix := beego.AppConfig.String("dbprefix")
	return prefix + "_" + str
}

// 创建一个Code为错误的指针
func NewRespJson() *RespJson {
	return &RespJson{Code: DO_ERROR}
}
