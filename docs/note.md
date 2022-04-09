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

`/object/blog/2013/09/12` 调用 `ObjectController` 中的 Blog 方法，参数如下：`map[0:2013 1:09 2:12]`
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

不能使用`<template>`，因为模板种的`{{d.data}}`和 go 中的模板格式相同，会出问题，而修改 beego 中的模板比较麻烦，所以使用函数传递模板  
函数传递的方式

```js
templet: function (d) {
    return '<img src="' + d.vtypelogo + '">'
}
```

## js 操作

### 弹出框

```javascript
// 设置普通弹出框
layer.alert(info, {
    title: "titleName",
});

// 设置确认框
layer.confirm(info, function{})
```

# html css js note

## HTML

###

## CSS

### 文本行省略内容

```css
/******** 单行 ********/
white-space: nowrap;
overflow: hidden;
text-overflow: ellipsis;

/******** 多行 ********/
/* 高度必设 */
height: 40px;
display: -webkit-box !important;
text-overflow: ellipsis;
word-break: break-all;
-webkit-box-orient: vertical;
/* 设置行数 */
-webkit-line-clamp: 2;
overflow: hidden;
```

## JS

### 设置`input:radio`

```js
// 设置第一个input:radio为checked
$("input:radio:first").prop("checked", "checked");
```



### 设置元素伪类内容

```js
// html
<div id="diver"></div>
// css
#diver {
    content: attr(diver-content);
}
// js
$("#diver").attr("diver-content", content-text);
```

### ajax 一般设置

```js
（1）type: 请求方式，(默认: "GET") 请求方式 ("POST" 或 "GET")， 默认为 "GET"。
注意：其它 HTTP 请求方法，如 PUT 和 DELETE 也可以使用，但仅部分浏览器支持。
（2）url: 请求的地址；类型：string
（3）async: 默认true，类型：bool
true-请求为异步请求
false-请求为同步请求（同步请求将锁住浏览器，用户其他操作必须等 待请求完成才可以执行）
（4）timeout: 设置请求超时时间（毫秒）；类型：int
（5）cache：默认为true（当dataType为script时，默认为false）； 设置为false将不会从浏览器缓存中加载请求信息；
（6）data: 发送到服务器的数据（例：{a:"a",b:"b"}  $('#formid').serialize()自动转换form表单)；类型：string
（7）dataType: 预期服务器返回的数据类型。如果不指定，JQuery将自动根据http包mime信息返回responseXML或responseText，并作为回调函数参数传递。)；类型：string
可用的类型如下：
xml：返回XML文档，可用JQuery处理。
html：返回纯文本HTML信息；包含的script标签会在插入DOM时执行。
script：返回纯文本JavaScript代码。不会自动缓存结果。除非设置了cache参数。注意在远程请求时（不在同一个域下），所有post请求都将转为get请求。
json：返回JSON数据。
jsonp：JSONP格式。使用SONP形式调用函数时，例如myurl?callback=?，JQuery将自动替换后一个“?”为正确的函数名，以执行回调函数。
text：返回纯文本字符串。
（8）beforeSend：发送请求前可以修改XMLHttpRequest对象的函数，例如添加自定义HTTP头。在beforeSend中如果返回false可以取消本次ajax请求。XMLHttpRequest对象是惟一的参数。
function(XMLHttpRequest){
    this;   //调用本次ajax请求时传递的options参数
}
（9）complete：请求完成后调用的回调函数（请求成功或失败时均调用）。参数：XMLHttpRequest对象和一个描述成功请求类型的字符串。
function(XMLHttpRequest, textStatus){
    this;    //调用本次ajax请求时传递的options参数
}
（10）success：请求成功后调用的回调函数，有两个参数。
1)由服务器返回，并根据dataType参数进行处理后的数据。
2)描述状态的字符串。
function(data, textStatus){
    //data可能是xmlDoc、jsonObj、html、text等等
    this;  //调用本次ajax请求时传递的options参数
}
（12）error：请求失败时被调用的函数。该函数有3个参数，即XMLHttpRequest对象、错误信息、捕获的错误对象(可选)。
ajax事件函数如下：
function(XMLHttpRequest, textStatus, errorThrown){
    //通常情况下textStatus和errorThrown只有其中一个包含信息
    this;   //调用本次ajax请求时传递的options参数
}
（13）contentType：当发送信息至服务器时，内容编码类型默认为"application/x-www-form-urlencoded"。该默认值适合大多数应用场合；类型：string
（14）dataFilter：给Ajax返回的原始数据进行预处理的函数。
提供data和type两个参数。data是Ajax返回的原始数据，type是调用jQuery.ajax时提供的dataType参数。函数返回的值将由jQuery进一步处理。
function(data, type){
    //返回处理后的数据
    return data;
}
（15）global：默认为true。表示是否触发全局ajax事件。设置为false将不会触发全局ajax事件，ajaxStart或ajaxStop可用于控制各种ajax事件；类型：bool
（16）ifModified：默认为false。仅在服务器数据改变时获取新数据。 服务器数据改变判断的依据是Last-Modified头信息。默认值是false，即忽略头信息；类型：bool
（17）jsonp：，在一个jsonp请求中重写回调函数的名字。 该值用来替代在"callback=?"这种GET或POST请求中URL参数里的"callback"部分；类型：string
例如:{jsonp:'onJsonPLoad'}会导致将"onJsonPLoad=?"传给服务器。
（18）username：用于响应HTTP访问认证请求的用户名；类型：string
（19）password：用于响应HTTP访问认证请求的密码；类型：string
（20）processData：默认为true。默认情况下，发送的数据将被转换为对象（从技术角度来讲并非字符串）以配合默认内容类型"application/x-www-form-urlencoded"；类型：bool
如果要发送DOM树信息或者其他不希望转换的信息，请设置为false。
（21）scriptCharset：只有当请求时dataType为"jsonp"或者"script"，并且type是GET时才会用于强制修改字符集(charset)。通常在本地和远程的内容编码不同时使用；类型：string
```

# sql note

## tables

表 VideoType：视频类别表

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

其他种类的分类：
剧情 科幻 动作 喜剧 爱情 冒险 儿童 歌舞 音乐 奇幻 动画 恐怖 惊悚 丧尸 战争 传记 纪录 犯罪 悬疑 西部 灾难 古装 武侠 家庭 短片 校园 文艺 运动 青春 同性 励志 人性 美食 女性 治愈 历史 真人秀 脱口秀

1.3. 二级分类： 
后台管理员动态添加，以关键字形式出现，例： 按领域：天文、地理、屋里、化学
```

表 VideoTypeImage：视频类别广告表（每个一级分类首页都有滚动播放广告）

```sql
vtiid            int     视频类别广告ID
vtiTitle         String  广告标题
imageURL         String  图片链接
URL              String  广告链接
sequence         int     显示顺序 
```

表 Video：视频表

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

表 VideoReview：视频评论表

```sql
id             int     评论编号
userid         int     评论人编号
videoid        int     视频编号
content        String  评论内容
status         String  评论状态:(N:未被屏蔽Y:已屏蔽)
pubtime  	   String  发布时间
socr
```

表 Users：用户表

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

# Record

## 正则表达式

```
结构体GetDifCols()
^(.*) (.*)\"(.*)\"(.*)$
if m.$1 != v.$1 { dif = append\(dif, "$3"\)	}

```
