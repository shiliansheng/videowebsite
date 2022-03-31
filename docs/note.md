# beego notes

## router

### 自动匹配路由
用户首先需要把需要路由的控制器注册到自动路由中：

`beego.AutoRouter(&controllers.ObjectController{})`
那么 beego 就会通过反射获取该结构体中所有的实现方法，你就可以通过如下的方式访问到对应的方法中：
```
/object/login   调用 ObjectController 中的 Login 方法
/object/logout  调用 ObjectController 中的 Logout 方法
```
除了前缀两个 /:controller/:method 的匹配之外，剩下的 url beego 会帮你自动化解析为参数，保存在 `this.Ctx.Input.Params` 当中：

`/object/blog/2013/09/12`  调用 `ObjectController` 中的 Blog 方法，参数如下：`map[0:2013 1:09 2:12]`
方法名在内部是保存了用户设置的，例如 Login，url 匹配的时候都会转化为小写，所以，`/object/LOGIN` 这样的 url 也一样可以路由到用户定义的 Login 方法中。

现在已经可以通过自动识别出来下面类似的所有 url，都会把请求分发到 controller 的 simple 方法：
```
/controller/simple
/controller/simple.html
/controller/simple.json
/controller/simple.xml
```
可以通过 `this.Ctx.Input.Param(":ext")` 获取后缀名。

## view

### 模板函数
```go
// 获取config
{{config "String" "httpport" "8088"}}

```


# layui note

## 模板操作

### 给表格添加模板

不能使用`<template>`，因为模板种的`{{d.data}}`和go中的模板格式相同，会出问题，而修改beego中的模板比较麻烦，所以使用函数传递模板  
函数传递的方式
```js
templet: function (d) {
    return '<img src="' + d.vtypelogo + '">'
}
```

## js操作

### 弹出框

```javascript
// 设置普通弹出框
layer.alert(info, {
    title: "titleName",
});

// 设置确认框
layer.confirm(info, function{})
```


# sql note

## tables
表VideoType：视频类别表 
```sql
vtid             Int     视频类型编号
typeName         String  视频类型名称
typeDescription  String  视频类型描述
userId           int     添加人编号
vtDate           String  视频类型添加时间
isPassed         String  审核状态(W:待审核;Y:通过审核)
typeLogo         String  类型Logo
sequence         Int     显示顺序

1.1. 首页（固定）： 
1.2. 一级分类（固定）：（来自腾讯视频）
剧情 喜剧 动作 爱情 惊悚 犯罪 悬疑 战争 科幻 动画 恐怖 家庭 传记 冒险 奇幻 武侠 历史 运动 音乐 记录 伦理
1.3. 二级分类： 
后台管理员动态添加，以关键字形式出现，例： 按领域：天文、地理、屋里、化学
```
表VideoTypeImage：视频类别广告表（每个一级分类首页都有滚动播放广告）  
```sql
vtiid            int     视频类别广告ID
vtiTitle         String  广告标题
imageURL         String  图片链接
URL              String  广告链接
sequence         int     显示顺序 
```
表Video：视频表 
```sql
vid              int     视频编号
videoTypeid      int     视频类型编号
videoTypeName    String  视频类型名称
videoName        String  视频名称
introduction     String  视频介绍
videoLogo        String  视频图片地址
keywords         String  视频关键字
videoResources   String  视频资源地址
copyright        int     版权所有(0:原创,1:转载)
uid              int     发布者编号
username         String  发布者用户名
publishedTime    String  发布时间
numOfViewed      int     视频观看次数
numOfDownload    int     视频下载次数
numOfScoring     int     视频打分人数
numOfRemarker    int     视频回复次数
averageScores    float   用户评分平均分
totalScores      float   用户评分总分
isPassed         String  审核状态(W:待审核;Y:通过审核)
default          Int     视频首页显示(0:不再首页显示,1:首页显示)
recommend        Int     视频推荐(0:不推荐,1:推荐)
```
表VideoReview：视频评论表 
```sql
rid            int     评论编号
title          String  评论标题
videoid        int     视频编号
scores         int     评分
reviews        String  评论内容
status         String  评论状态:(N:未被屏蔽Y:已屏蔽)
userid         int     评论人编号
publishedTime  String  发布时间
```
表Users：用户表  
```sql
id                int     用户id
username          String  用户名(邮箱)
password          String  密码
nickname          String  昵称
realname          String  真实姓名
logoimage         String  用户头像
introduction      String  简介
position          String  职位
sex               String  性别(1:男,0:女)
phone             String  电话
qq                int     qq
birthday          String  生日
usertype          String  用户类型(C:普通用户,S:系统用户)
userstatus        String  用户状态(Y:正常,D:冻结)
registrationtime  String  注册时间
lastlogintime     String  最后一次登录时间
numoflogin        long    登录次数
numoflink         long    浏览次数
```
表roletousers：用户角色关联表 
```sql
roletousersid  int  关联id
roleid         int  角色id
userid         int  用户id 
```
表operationtorole：角色操作表 
```sql
operationtoroleid  int  关联id
operationid        int  操作id
roleid             int  角色id
```

每个视频一级分类为一个操作对象，只有分配了该对象的管理员用户才可以添加、修改、删除、审核该分类下的视频。    

# Record

## 正则表达式
```
结构体GetDifCols()
^(.*) (.*)\"(.*)\"(.*)$
if m.$1 != v.$1 { dif = append\(dif, "$3"\)	}

```