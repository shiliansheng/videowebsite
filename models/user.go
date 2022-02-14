package models

import (
	"time"
)

type User struct {
	Id           int       `json:"id"`           // 用户ID
	Username     string    `json:"username"`     // 用户账户名
	Password     string    `json:"password"`     // 用户密码
	Nickname     string    `json:"nickname"`     // 昵称
	Logoimage    string    `json:"logoimage"`    // 用户头像
	Sex          string    `json:"sex"`          // 性别
	Email        string    `json:"email"`        // e-mail
	Birthday     string    `json:"birthday"`     // 用户生日
	Introduction string    `json:"introduction"` // 用户简介
	Status       string    `json:"status"`       // 用户身份
	State        int8      `json:"state"`        // 用户状态
	Remark       string    `json:"remark"`       // 备注
	CreateAt     string    `json:"createat"`     // 创建时间
	UpdateAt     string    `json:"updateat"`     // 更新时间
	DeleteAt     time.Time `json:"deleteat"`     // 删除时间
}

type UserJson struct {
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	Count int    `json:"count"`
	Data  []User `json:"data"`
}

func (m User) TableName() string {
	return TableName("user")
}

func (m User) GetUserCount() int {
	userCount, _ := Orm.QueryTable(new(User).TableName()).Count()
	return int(userCount)
}

func (m User) GetUserListJson(page, limit int, mapper map[string]interface{}, getNil bool) (UserJson, error) {
	userJson := UserJson{
		Code:  0,
		Msg:   "",
		Count: 0,
		Data:  nil,
	}
	if getNil {
		return userJson, nil
	}
	userlist, count, err := getUserList(page, limit, mapper)
	if err != nil {
		userJson.Code = MSG_FAIL
		userJson.Msg = "获取用户列表失败"
		return userJson, err
	}
	userJson.Count = count
	userJson.Data = userlist
	return userJson, nil

}

func getUserList(page, limit int, mapper map[string]interface{}) ([]User, int, error) {
	var userList []User
	//_, err := Orm.QueryTable(new(User).TableName()).All(&userList)
	seter := Orm.QueryTable(new(User).TableName())
	for key, value := range mapper {
		seter = seter.Filter(key+"__icontains", value)
	}
	count, _ := seter.Count()
	_, err := seter.Limit(limit, limit*(page-1)).All(&userList)
	return userList, int(count), err
}

func (m User) Update(user User, cols ...string) (int, string) {
	if len(cols) == 0 {
		return DO_REMAIN, "信息未更改，更新失败"
	}
	_, err := Orm.Update(&user, cols...)
	if err != nil {
		return DO_UP_ERROR, "更新信息失败<br/>" + err.Error()
	}
	return DO_SUCCESS, "更新信息成功"
}

// {"id", "username", "password", "nickname", "logoimage", "sex", "email", "birthday", "introduction", "status", "state", "remark"},

func (m User) GetDifCols(base, new User) []string {
	dif := []string{}
	if base.Password != new.Password {
		dif = append(dif, "password")
	}
	if base.Nickname != new.Nickname {
		dif = append(dif, "nickname")
	}
	if base.Logoimage != new.Logoimage {
		dif = append(dif, "logoimage")
	}
	if base.Sex != new.Sex {
		dif = append(dif, "sex")
	}
	if base.Email != new.Email {
		dif = append(dif, "email")
	}
	if base.Birthday != new.Birthday {
		dif = append(dif, "birthday")
	}
	if base.Introduction != new.Introduction {
		dif = append(dif, "introduction")
	}
	if base.Status != new.Status {
		dif = append(dif, "status")
	}
	if base.State != new.State {
		dif = append(dif, "state")
	}
	if base.Remark != new.Remark {
		dif = append(dif, "remark")
	}
	return dif
}
