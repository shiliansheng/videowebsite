package models

import (
	"net/url"
	"strconv"
	"time"
)

type User struct {
	Id           int       `json:"id"`                     // 用户ID
	Username     string    `json:"username"`               // 用户账户名
	Password     string    `json:"password"`               // 用户密码
	Nickname     string    `json:"nickname"`               // 昵称
	Userlogo     string    `json:"userlogo,omitempty"`     // 用户头像
	Sex          string    `json:"sex"`                    // 性别
	Email        string    `json:"email,omitempty"`        // e-mail
	Birthday     string    `json:"birthday,omitempty"`     // 用户生日
	Introduction string    `json:"introduction,omitempty"` // 用户简介
	Status       string    `json:"status"`                 // 用户身份
	State        int8      `json:"state"`                  // 用户状态
	Remark       string    `json:"remark,omitempty"`       // 备注
	CreateAt     string    `json:"createat"`               // 创建时间
	UpdateAt     string    `json:"updateat"`               // 更新时间
	DeleteAt     time.Time `json:"deleteat,omitempty"`     // 删除时间
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

func (m User) Add(u User) (int, string) {
	code, msg := 0, ""
	_, err := Orm.Insert(&u)
	if err != nil {
		code = DO_ERROR
		msg = "添加用户失败</br>" + err.Error()
	} else {
		code = DO_SUCCESS
		msg = "添加用户成功"
	}
	return code, msg
}

func (m User) GetDifCols(base, new User) []string {
	dif := []string{}
	if base.Password != new.Password {
		dif = append(dif, "password")
	}
	if base.Nickname != new.Nickname {
		dif = append(dif, "nickname")
	}
	if base.Userlogo != new.Userlogo {
		dif = append(dif, "userlogo")
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

func (m *User) GetUserInfo(source url.Values) error {
	if value := source.Get("id"); value != "" {
		m.Id = func() int { res, _ := strconv.Atoi(value); return res }()
	}
	if value := source.Get("username"); value != "" {
		m.Username = value
	}
	if value := source.Get("password"); value != "" {
		m.Password = value
	}
	m.Nickname = source.Get("nickname")
	if m.Nickname == "" {
		m.Nickname = "stranger"
	}
	if value := source.Get("userlogo"); value != "" {
		m.Userlogo = value
	}
	if value := source.Get("sex"); value != "" {
		m.Sex = value
	}
	if value := source.Get("email"); value != "" {
		m.Email = value
	}
	if value := source.Get("birthday"); value != "" {
		m.Birthday = value
	}
	if value := source.Get("Birthday"); value != "" {
		m.Birthday = value
	}
	if value := source.Get("introduction"); value != "" {
		m.Introduction = value
	}
	if value := source.Get("status"); value != "" {
		m.Status = value
	}
	if value := source.Get("state"); value != "" {
		m.State = func() int8 { res, _ := strconv.Atoi(value); return int8(res) }()
	}
	if value := source.Get("remark"); value != "" {
		m.Remark = value
	}
	return nil
}
