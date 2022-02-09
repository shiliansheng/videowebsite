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
 |——— utils
 |   |——— functions.go
 |   └─── 
 |——— views
 |   |——— admin
 |   |   |——— index.html
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