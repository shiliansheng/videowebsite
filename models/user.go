package models

import "time"

type User struct {
	Id       int    // 用户ID
	Username string // 用户账户名
	Password string // 用户密码
	Nickname string // 昵称
	Gender   string // 性别
	Email    string // e-mail
	Status   int    // 用户身份
	// LastTime time.Time
	// LastIp   string
	State    int8      // 用户状态
	Remark   string    // 备注
	CreateAt time.Time // 创建时间
	UpdateAt time.Time // 更新时间
	DeleteAt time.Time // 删除时间
}

func (m *User) TableName() string {
	return TableName("user")
}

// user const values
const (

	// status
	USER_NORMAL = 1
	USER_ADMIN  = 2
)
