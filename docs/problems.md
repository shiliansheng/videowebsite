# Problems

## 编译问题

### falg should be defined gracefully

默认import包为：
```
"github.com/astaxie/beego/orm"
"github.com/astaxie/beego"
```
则会出现：
```
C:\Users\24848\AppData\Local\Temp\go-build4251115413\b001\exe\main.exe flag redefined: graceful
panic: C:\Users\24848\AppData\Local\Temp\go-build4251115413\b001\exe\main.exe flag redefined: graceful
goroutine 1 [running]:
flag.(*FlagSet).Var(0xc0000d6000, {0xec3ee0, 0x1399dd1}, {0xdd7e20, 0x8}, {0xdedcc2, 0x21})
        C:/Program Files/Go/src/flag/flag.go:879 +0x2f4
flag.BoolVar(...)
        C:/Program Files/Go/src/flag/flag.go:638
github.com/beego/beego/v2/server/web/grace.init.0()
        E:/VS/Go/pkg/mod/github.com/beego/beego/v2@v2.0.2/server/web/grace/grace.go:93 +0x52
exit status 2
```

问题解决：  
```
将引入的包
"github.com/astaxie/beego"
替换为
"github.com/astaxie/beego"
```


## 功能问题

### 自动路由找不到函数

使用AutoRouter时，找不到`admin/post`，但是可以找到`admin/login`  
使用的包为：
```go
beego "github.com/beego/beego/v2/server/web"
```
解决办法，将包替换为：
```go
"github.com/astaxie/beego"
```

### 路径寻找

#### 写入文件
想要写入的路径是：`/static/store/demo.txt`  
实际使用的路径应该是：`./static/store/demo.txt`
将该路径上传到数据库中，路径变成了：`static/store/demo.txt`

#### 读取文件
一般是读取图片  
如果在数据库中的路径是：`static/store/image/demo.jpg`  
展示在网页的`img.src`中就应该是`../static/store/image/demo.jpg`

## 第三方库问题

### cropper

主要问题：获取canvas的内容失败了  
可以借鉴的东西是使用是，croppers post数据到服务端是使用`FormData`的




## 其他内容

play.js加载内容问题  
遗留代码：
```js
// 获取侧边视频列表
$.ajax({
    url: "gethotvideo.json"
    , type: "GET"
    , data: { "limit": 10 }
    , success: function (res) {
        var sideList = $("#side-video-list"),
            data = res.data;
        sideList.html('<div class="list-title">热门电影</div>');
        for (var i = 0, len = data.size; i < len; i++) {
            sideList.append(
                '<li class="list-item">' +
                '    <a href="/play?id=' + data.id[i] +'" class="img-box">' +
                '        <img src="' + data.logo[i] +'" alt="">' +
                '    </a>' +
                '    <div class="item-info">' +
                '        <a href="/play?id=' + data.id[i] + '" class="item-title">' + data.name[i] +'</a>' +
                '        <p class="item-type info-piece">' + data.type[i] +'</p>' +
                '        <p class="item-score info-piece">' + data.socre[i] +'</p>' +
                '        <p class="item-introduction info-piece">' + data.introduction[i] +'</p>' +
                '    </div>' +
                '</li>'
            );
        }
    }
});
// 获取评论列表
$.ajax({
    url: "getreview.json"
    , type: "GET"
    , data: {"id": id, "limit": "10"}
    , success: function(res) {
        var $remarkList = $("#remark-list");
        $remarkList.html('<div class="list-title">评论列表</div>');
        for (var i, len = res.Count; i < len; i++) {
            $remarkList.append(
                '<li class="remark-item">' +
                '    <div class="remark-userlogo">' +
                '        <img src="'+ logo +'" alt="">' +
                '    </div>' +
                '    <div class="remark-info">' +
                '        <div class="remark-username">'+ username +'</div>' +
                '        <div class="remark-detail">评论发布于 '+ pubtime +'</div>' +
                '        <div class="remark-content">'+ content +'</div>' +
                '    </div>' +
                '</li>',
            );
        }
    }
});
```