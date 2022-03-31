package models

import (
	"net/url"
	"strconv"
	"strings"
	"time"
	"videowebsite/utils"
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

func (c User) GetNickname(name string) string {
	if name == "" {
		name = "stranger"
	}
	return name
}

func (m User) TableName() string {
	return TableName("user")
}

// ### 功能操作

// 用户名、密码进行登录，返回登录结果，登录用户结构体放在RespJson.Data中
//  @param  username [string]
//  @param  password [string]
//  @return [RespJson]
func (m User) Login(username, password string) RespJson {
	resp := RespJson{Code: DO_ERROR}
	user := User{Username: username}
	err := Orm.Read(&user, "username")
	if err != nil {
		resp.Msg = "登录失败，用户名或密码错误</br>" + err.Error()
	} else if user.Password == "" {
		resp.Msg = "登录失败，用户名不存在"
	} else if user.Password != strings.TrimSpace(password) {
		resp.Msg = "登录失败，密码错误"
	} else {
		resp.Code = DO_SUCCESS
		resp.Data = user
		resp.Msg = "登录成功"
	}
	return resp
}

// ### 获取INFO操作

func (m User) GetUserCount() int {
	count, _ := Orm.QueryTable(m.TableName()).Count()
	return int(count)
}

func (m User) GetUserList(page, limit int, filterMap map[string]interface{}) RespJson {
	resp := RespJson{
		Code:  DO_SUCCESS,
		Count: 0,
	}

	userlist, count, err := m.getUserList(page, limit, filterMap)
	if err != nil {
		resp.Code = DO_ERROR
		resp.Msg = "获取用户列表失败: " + err.Error()
		return resp
	}
	resp.Msg = "获取用户列表成功!"
	resp.Count = count
	resp.Data = userlist
	return resp
}

func (m User) getUserList(page, limit int, filterMap map[string]interface{}) ([]User, int, error) {
	var userList []User
	seter := Orm.QueryTable(new(User).TableName()).Exclude("status__exact", "超级管理员")
	for key, value := range filterMap {
		seter = seter.Filter(key+"__icontains", value)
	}
	count, _ := seter.Count()
	_, err := seter.Limit(limit, limit*(page-1)).All(&userList)
	return userList, int(count), err
}

// ### CRUD操作

// 添加内容为参数
//  @param  u [User] 待添加User
//  @return [RespJson]
func (m User) Add(u User) RespJson {
	u.setNickname()
	resp := RespJson{Code: DO_SUCCESS}
	timeStr := utils.GetNowTimeString()
	u.CreateAt, u.UpdateAt = timeStr , timeStr
	_, err := Orm.Insert(&u)
	if err != nil {
		resp.Code = DO_ERROR
		resp.Msg = "添加用户失败</br>" + err.Error()
	} else {
		resp.Msg = "添加用户成功"
	}
	return resp
}

// oldUser调用Update方法，参数为newUser
//  @param  newUser [User] 更改信息的User
//  @return [RespJson]
func (m User) Update(newUser User) RespJson {
	resp := RespJson{Code: DO_ERROR}
	newUser.setNickname()
	cols := m.GetDifCols(newUser)
	if len(cols) == 0 {
		resp.Msg = "信息未更改，更新失败"
		return resp
	}
	_, err := Orm.Update(&newUser, cols...)
	if err != nil {
		resp.Msg = "更新信息失败<br/>" + err.Error()
	} else {
		resp.Code = DO_SUCCESS
		resp.Msg = "更新成功"
	}
	return resp
}

// 操作用户调用Delte方法删除参数User
//  @param  user [User] 待删除User，仅包含id
//  @return [RespJson]
func (m User) Delete(user User) RespJson {
	resp := RespJson{Code: DO_ERROR}
	Orm.Read(&user)
	if user.Id == m.Id {
		resp.Msg = "删除用户 " + user.Username + " 失败<br/>禁止删除自己"
	} else if user.Status == "管理员" && m.Status != "超级管理员" {
		resp.Msg = "删除用户 " + user.Username + " 失败<br/>禁止删除管理员"
	} else {
		_, err := Orm.Delete(&user)
		if err != nil {
			resp.Msg = "删除用户 " + user.Username + " 失败<br/>" + err.Error()
		} else {
			resp.Code = DO_SUCCESS
			resp.Msg = "删除用户 " + user.Username + " 成功"
		}
	}
	return resp
}

// ### 填充INFO操作

// oldUser调用Update方法，与参数newUser进行比对，返回不同的字段名数组
//  @param  new [User] 待更新的新User
//  @return [[]string] 不同的字段名数组
func (m User) GetDifCols(new User) []string {
	dif := []string{}
	if m.Password != new.Password {
		dif = append(dif, "password")
	}
	if m.Nickname != new.Nickname {
		dif = append(dif, "nickname")
	}
	if m.Userlogo != new.Userlogo {
		dif = append(dif, "userlogo")
	}
	if m.Sex != new.Sex {
		dif = append(dif, "sex")
	}
	if m.Email != new.Email {
		dif = append(dif, "email")
	}
	if m.Birthday != new.Birthday {
		dif = append(dif, "birthday")
	}
	if m.Introduction != new.Introduction {
		dif = append(dif, "introduction")
	}
	if m.Status != new.Status {
		dif = append(dif, "status")
	}
	if m.State != new.State {
		dif = append(dif, "state")
	}
	if m.Remark != new.Remark {
		dif = append(dif, "remark")
	}
	return dif
}

func (m *User) SetUser(source url.Values) error {
	if value := source.Get("id"); value != "" {
		m.Id = utils.Atoi(value)
	}
	if value := source.Get("username"); value != "" {
		m.Username = value
	}
	if value := source.Get("password"); value != "" {
		m.Password = value
	}
	m.Nickname = source.Get("nickname")
	m.setNickname()
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

// 设置user的昵称，如果昵称为空，默认为stranger
func (m *User) setNickname() {
	if m.Nickname == "" {
		m.Nickname = "stranger"
	}
}
