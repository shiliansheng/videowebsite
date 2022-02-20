# DIARY

## 2022.2.6

- [x] 初始化项目

## 2022.2.8

- [x] 完善登录
- [x] 上传到github上
- [ ] 完善session

## 2020.2.9

- [x] 完善session
- [x] 完善接收`admin/index`下的ifrmae
- [x] 将菜单目录转接数据库
- [x] 规范数据库格式

- [x] 增加用户列表展示

## 2022.2.10
- [x] [添加用户表单正则验证](https://www.cnblogs.com/raphael1982/p/8012634.html)
- [x] 添加用户页面前后端交互(ajax)
- [x] 修改用户模板部署

## 2022.2.11
- [x] 添加修改用户功能
- [x] 添加删除用户功能（单个）
- [x] 添加删除用户功能（多个）
- [x] 完善删除用户功能（前端+后端）

## 2022.2.13
- [x] 添加筛选用户功能
[layui select使用问题](https://www.cnblogs.com/kcat/p/10650227.html)
- [x] 添加个人修改密码
- [x] 修改用户信息展示方式，使用`form.val(lay-filter, object)`方式
注意：需要引用：`/static/js/lay-config.js?v=1.0.4`文件
- [x] 添加个人信息修改页面，后端还没写

## 2022.2.14
- [x] 完成个人信息修改后端（并增加内容）
- [x] 增加user表字段
- [x] 创建video的model
- [x] 完善更新内容

## 2022.2.15
- [ ] 添加头像裁剪功能
使用的是coppers，结果是失败的

## 2022.2.16
- [x] 添加头像裁剪功能
使用集成cropper的avatar
- [x] 显示路径为存储在数据库中的路径的图片

## 2022.2.17
> 离职后第一天，今天太懒了，唉...

- [x] 添加video type表

## 2022.2.18

- [x] 添加video type菜单
- [x] 添加video type修改及搜索
- [x] 完善图片上传机制
- [x] 完善根据`url.values`获取结构体内容

## 2022.2.20

- [x] 完善video type菜单删除等操作


# app tree
```
root
 |——— conf
 |   └─── app.conf
 |——— controllers
 |   |─── admin.go
 |   |─── base.go
 |   └─── 
 |——— db
 |   |——— init.sql
 |   |——— menu.sql
 |   └─── user.sql
 |——— docs
 |   |——— diary.md
 |   |——— note.md
 |   └─── problems.md
 |——— models
 |   |——— init.go
 |   |——— 
 |   |——— SystemMenu.go
 |   └─── user.go
 |——— routers
 |   └─── router.go
 |——— static
 |   |——— api
 |   |——— css
 |   |——— img
 |   |——— lib
 |   └─── js
 |       └─── admin
 |           |——— useradd.js
 |           |——— useredit.js
 |           └─── js
 |——— utils
 |   |——— functions.go
 |   └─── 
 |——— views
 |   |——— admin
 |   |   |——— index.html
 |   |   |——— userlist.html
 |   |   └─── welcome.html
 |   |——— video
 |   |   |——— index.html
 |   |   └─── 
 |   └─── login.html
 └─── main.go
 |   |——— 
 |   |——— 
 |   └─── 
```

## 用户

- 管理员
- 普通用户

### 管理员
管理员通过`login.html`登录，登录验证后，进入后台`admin/index.html`页面中