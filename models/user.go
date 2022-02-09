package models

import (
	"time"
)

// user const values
const ()

type User struct {
	Id       int    `json:"id"`       // 用户ID
	Username string `json:"username"` // 用户账户名
	Password string `json:"password"` // 用户密码
	Nickname string `json:"nickname"` // 昵称
	Sex      string `json:"sex"`      // 性别
	Email    string `json:"email"`    // e-mail
	Status   int    `json:"status"`   // 用户身份
	// LastTime time.Time
	// LastIp   string
	State    int8      `json:"state"`    // 用户状态
	Remark   string    `json:"remark"`   // 备注
	CreateAt time.Time `json:"createat"` // 创建时间
	UpdateAt time.Time // 更新时间
	DeleteAt time.Time // 删除时间
}

type UserJson struct {
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	Count int    `json:"count"`
	Data  []User `json:"data"`
}

func (m *User) TableName() string {
	return TableName("user")
}

func (m *User) GetUserCount() int {
	userCount, _ := Orm.QueryTable(new(User).TableName()).Count()
	return int(userCount)
}

func (m *User) GetUserListJson() (UserJson, error) {
	userJson := UserJson{
		Code: 0,
		Msg:  "",
	}
	userlist, err := getUserList()
	if err != nil {
		userJson.Code = MSG_FAIL
		userJson.Msg = "获取用户列表失败"
		return userJson, err
	}
	userJson.Count = len(userlist)
	userJson.Data = userlist
	return userJson, nil

}

func getUserList() ([]User, error) {
	var userList []User
	_, err := Orm.QueryTable(new(User).TableName()).All(&userList)
	return userList, err
}
