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


## layui中js操作

### 弹出框

#### 设置普通弹出框
```javascript
layer.alert(info, {
    title: "titleName",
});
```