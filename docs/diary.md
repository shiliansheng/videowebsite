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