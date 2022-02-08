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