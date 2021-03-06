package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// 菜单
type SystemMenu struct {
	Id       int       `json:"id"`
	Pid      int       `json:"pid"`
	Title    string    `json:"title"`
	Icon     string    `json:"icon"`
	Href     string    `json:"href"`
	Sort     string    `json:"sort"`
	Target   string    `json:"target"`
	Remark   string    `json:"remark"`
	State    int       `json:"state"`
	CreateAt time.Time `json:"create_at"`
}

func (m *SystemMenu) TableName() string {
	return TableName("system_menu")
}

// 初始化结构体
type SystemInit struct {
	HomeInfo struct {
		Title string `json:"title"`
		Href  string `json:"href"`
	} `json:"homeInfo"`
	LogoInfo struct {
		Title string `json:"title"`
		Image string `json:"image"`
	} `json:"logoInfo"`
	MenuInfo []*MenuTreeList `json:"menuInfo"`
}

// 菜单结构体
type MenuTreeList struct {
	Id     int             `json:"id"`
	Pid    int             `json:"pid"`
	Title  string          `json:"title"`
	Icon   string          `json:"icon"`
	Href   string          `json:"href"`
	Target string          `json:"target"`
	Remark string          `json:"remark"`
	Child  []*MenuTreeList `json:"child"`
}

// 获取初始化数据
func (m *SystemMenu) GetSystemInit() SystemInit {
	var systemInit SystemInit

	// 首页
	systemInit.HomeInfo.Title = "首页"
	systemInit.HomeInfo.Href = "homepage.html"

	// logo
	systemInit.LogoInfo.Title = "VIDEO WEB"
	systemInit.LogoInfo.Image = "../static/img/logo.png"
	// 菜单
	systemInit.MenuInfo = m.GetMenuList()

	return systemInit
}

// 获取菜单列表
func (m *SystemMenu) GetMenuList() []*MenuTreeList {
	o := orm.NewOrm()
	var menuList []SystemMenu
	_, _ = o.QueryTable(m.TableName()).Filter("state", 1).OrderBy("-sort").All(&menuList)
	return m.buildMenuChild(0, menuList)
}

//递归获取子菜单
func (m *SystemMenu) buildMenuChild(pid int, menuList []SystemMenu) []*MenuTreeList {
	var treeList []*MenuTreeList
	for _, v := range menuList {
		if pid == v.Pid {
			node := &MenuTreeList{
				Title:  v.Title,
				Icon:   v.Icon,
				Href:   v.Href,
				Target: v.Target,
			}
			child := v.buildMenuChild(v.Id, menuList)
			if len(child) != 0 {
				node.Child = child
			}
			// todo 后续此处加上用户的权限判断
			treeList = append(treeList, node)
		}
	}
	return treeList
}
