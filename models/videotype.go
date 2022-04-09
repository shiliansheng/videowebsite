package models

import (
	"fmt"
	"log"
)

type VideoType struct {
	Id          int    `json:"id" orm:"pk"` // video type id 视频类型编号
	Pid         int    `json:"pid"`         // 视频类型父ID
	Typename    string `json:"typename"`    // 视频类型名称
	Discription string `json:"discription"` // 视频类型描述
	Addid       int    `json:"addid"`       // 添加人员id
	Createat    string `json:"createat"`    // 类型添加时间
	Vtypelogo   string `json:"vtypelogo"`   // 类型logo
	Sequence    int    `json:"sequence"`    // 显示顺序
}

// ### base function

func (m VideoType) TableName() string {
	return TableName("videotype")
}

// ### 获取INFO

// 获取视频类型总数量
//  @return [int] 
//
func ( m VideoType) GetVideoTypeCount() int {
	count := 0
	if err := Orm.Raw("SELECT COUNT(*) FROM `" + m.TableName() + "`;").QueryRow(&count); err != nil {
		log.Println("获取视频类型失败:", err)
	}
	return count
}

// 获取全部视频类型名称
//  @return [[]string] 视频类型名称数组
func (m *VideoType) GetAllVideoTypeName() []string {
	var vtlist []VideoType
	var vtnames []string
	_, err := Orm.QueryTable(m.TableName()).All(&vtlist)
	if err != nil {
		fmt.Println("[ERROR] 获取视频全部类型失败:", err.Error())
	} else {
		for i := range vtlist {
			vtnames = append(vtnames, vtlist[i].Typename)
		}
	}
	return vtnames
}

// ### CRUD

// ### 设置INFO
