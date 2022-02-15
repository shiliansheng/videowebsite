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


## 第三方库问题

### cropper
```js
layui.config({
	base: '/static/lib/cropper/' //layui自定义layui组件目录
}).use(["form", "miniTab", 'laydate', 'upload', 'croppers'], function () {
	var form = layui.form
		, layer = layui.layer
		, miniTab = layui.miniTab
		, laydate = layui.laydate
		, croppers = layui.croppers
		, upload = layui.upload
		, $ = layui.$;
	// laydate实现
	laydate.render({
		elem: '#dateselect',
		//type: 'datetime',
		//range: true
	});
	//创建一个头像上传组件
	croppers.render({
		elem: '#logoimage'
		, saveW: 150	//保存宽度
		, saveH: 150	//保存高度
		, mark: 1 / 1	//选取比例
		, area: '900px'	//弹窗宽度
		, url: "{:url('admin/upload/img_save',['type'=>'admin'])}"  //图片上传接口返回和（layui 的upload 模块）返回的JOSN一样
		, done: function (data) { //上传完毕回调
			if (data.code == 1) {
				$('#demo1').attr('src', data.url);
				$('#logoimage').attr('data-src', data.url);  //成功返回路径存到数据库
			} else {
				return layer.msg('上传失败');
			}
			/* $("#inputimgurl").val(url);
			 $("#srcimgurl").attr('src',url);*/
		}
	});
	// 图片上传初始化
	/*
	upload.render({
		elem: "#logoimage",
		accpet: "images", // 指定允许上传时校验的文件类型
		acceptMime: 'image/jpg, image/png, image/jpeg', // 规定打开文件选择框时，筛选出的文件类型
		exts: "jpg|png|jpeg", // 允许上传的文件后缀
		number: 1, // 文件上传限制
		auto: false,
		bindAction: "#saveBtn",
		done: function (res, index, upload) {
			//获取当前触发上传的元素，一般用于 elem 绑定 class 的情况，注意：此乃 layui 2.1.0 新增
			var item = this.item;
		},
		choose: function (obj) {
			obj.preview(function (index, file, result) {
				document.getElementById("userlogo").setAttribute("src", result);
				// var logo = document.getElementById("userlogo");
				// logo.style.zoom = (150 / logo.width) * 1.0;
			});
		},
	});
	*/
	//监听提交
	form.on("submit(saveBtn)", function (data) {
		layer.confirm('确认修改?', function () {
			console.log(data.field);
			$.ajax({
				type: "post",
				url: "user_setting.json?action=changeSetting",
				data: data.field,
				success: function (res) {
					if (res.code == 0) {
						layer.msg('修改信息成功');
						miniTab.deleteCurrentByIframe();
					} else {
						layer.msg(res.msg);
					}
				},
			});
		});
	});

	form.val("baseinfo", userInfo);

	form.verify({
		verifyEmail: function (value) {
			var ePattern =
				/^([A-Za-z0-9_\-\.])+\@([A-Za-z0-9_\-\.])+\.([A-Za-z]{2,4})$/;
			if (value != "" && !ePattern.test(value)) {
				return "邮箱格式不正确";
			}
		},
	});
});
```

```js
layui.use(["form", "miniTab", 'laydate', 'upload'], function () {
	var form = layui.form,
		layer = layui.layer,
		miniTab = layui.miniTab,
		laydate = layui.laydate,
		upload = layui.upload,
		$ = layui.$;
	// laydate实现
	laydate.render({
		elem: '#dateselect',
		//type: 'datetime',
		//range: true
	});
	// 图片上传初始化
	/*
	upload.render({
		elem: "#logoimage",
		accpet: "images", // 指定允许上传时校验的文件类型
		acceptMime: 'image/jpg, image/png, image/jpeg', // 规定打开文件选择框时，筛选出的文件类型
		exts: "jpg|png|jpeg", // 允许上传的文件后缀
		number: 1, // 文件上传限制
		auto: false,
		bindAction: "#saveBtn",
		done: function (res, index, upload) {
			//获取当前触发上传的元素，一般用于 elem 绑定 class 的情况，注意：此乃 layui 2.1.0 新增
			var item = this.item;
		},
		choose: function (obj) {
			obj.preview(function (index, file, result) {
				document.getElementById("userlogo").setAttribute("src", result);
				// var logo = document.getElementById("userlogo");
				// logo.style.zoom = (150 / logo.width) * 1.0;
			});
		},
	});
	*/
	//监听提交
	form.on("submit(saveBtn)", function (data) {
		layer.confirm('确认修改?', function () {
			console.log(data.field);
			$.ajax({
				type: "post",
				url: "user_setting.json?action=changeSetting",
				data: data.field,
				success: function (res) {
					if (res.code == 0) {
						layer.msg('修改信息成功');
						miniTab.deleteCurrentByIframe();
					} else {
						layer.msg(res.msg);
					}
				},
			});
		});
	});

	form.val("baseinfo", userInfo);

	form.verify({
		verifyEmail: function (value) {
			var ePattern =
				/^([A-Za-z0-9_\-\.])+\@([A-Za-z0-9_\-\.])+\.([A-Za-z]{2,4})$/;
			if (value != "" && !ePattern.test(value)) {
				return "邮箱格式不正确";
			}
		},
	});
});

layui.config({
	base: '/static/lib/cropper/' //layui自定义layui组件目录
}).use(['form', 'croppers'], function () {
	var $ = layui.jquery
		, form = layui.form
		, croppers = layui.croppers
		, layer = layui.layer;

	//创建一个头像上传组件
	croppers.render({
		elem: '#logoimage'
		, saveW: 150	//保存宽度
		, saveH: 150	//保存高度
		, mark: 1 / 1	//选取比例
		, area: '900px'	//弹窗宽度
		, url: "{:url('admin/upload/img_save',['type'=>'admin'])}"  //图片上传接口返回和（layui 的upload 模块）返回的JOSN一样
		, done: function (data) { //上传完毕回调
			if (data.code == 1) {
				$('#demo1').attr('src', data.url);
				$('#logoimage').attr('data-src', data.url);  //成功返回路径存到数据库
			} else {
				return layer.msg('上传失败');
			}
			/* $("#inputimgurl").val(url);
			 $("#srcimgurl").attr('src',url);*/
		}
	});
});
```
```js
// layui.extend({
//   	croopers: '../static/../lib/cropper/'
// })

layui.use(["form", "miniTab", 'laydate', 'upload'], function () {
	var form = layui.form
		, layer = layui.layer
		, miniTab = layui.miniTab
		, laydate = layui.laydate
		//, upload = layui.upload
		, $ = layui.$;
	// laydate实现
	laydate.render({
		elem: '#dateselect',
	});
	//创建一个头像上传组件
	
	// 图片上传初始化
	/*
	upload.render({
		elem: "#logoimage",
		accpet: "images", // 指定允许上传时校验的文件类型
		acceptMime: 'image/jpg, image/png, image/jpeg', // 规定打开文件选择框时，筛选出的文件类型
		exts: "jpg|png|jpeg", // 允许上传的文件后缀
		number: 1, // 文件上传限制
		auto: false,
		bindAction: "#saveBtn",
		done: function (res, index, upload) {
			//获取当前触发上传的元素，一般用于 elem 绑定 class 的情况，注意：此乃 layui 2.1.0 新增
			var item = this.item;
		},
		choose: function (obj) {
			obj.preview(function (index, file, result) {
				document.getElementById("userlogo").setAttribute("src", result);
				// var logo = document.getElementById("userlogo");
				// logo.style.zoom = (150 / logo.width) * 1.0;
			});
		},
	});
	*/
	//监听提交
	form.on("submit(saveBtn)", function (data) {
		layer.confirm('确认修改?', function () {
			console.log(data.field);
			$.ajax({
				type: "post",
				url: "user_setting.json?action=changeSetting",
				data: data.field,
				success: function (res) {
					if (res.code == 0) {
						layer.msg('修改信息成功');
						miniTab.deleteCurrentByIframe();
					} else {
						layer.msg(res.msg);
					}
				},
			});
		});
	});

	form.val("baseinfo", userInfo);

	form.verify({
		verifyEmail: function (value) {
			var ePattern =
				/^([A-Za-z0-9_\-\.])+\@([A-Za-z0-9_\-\.])+\.([A-Za-z]{2,4})$/;
			if (value != "" && !ePattern.test(value)) {
				return "邮箱格式不正确";
			}
		},
	});
});
//http://localhost:8088/static/js/lay-module/layuimini/miniTab.js?v=1644930291321
//http://localhost:8088/lib/cropper/croppers.js?v=1644930449080
layui.config({
	base: '/static/lib/cropper/' //layui自定义layui组件目录
}).use(['croppers'], function () {
	var croppers = layui.croppers
		, $ = layui.$;
	croppers.render({
		elem: '#logoimage'
		, saveW: 150	//保存宽度
		, saveH: 150	//保存高度
		, mark: 1 / 1	//选取比例
		, area: '900px'	//弹窗宽度
		, url: "/00/"  //图片上传接口返回和（layui 的upload 模块）返回的JOSN一样
		, done: function (data) { //上传完毕回调
			if (data.code == 1) {
				$('#demo1').attr('src', data.url);
				$('#logoimage').attr('data-src', data.url);  //成功返回路径存到数据库
			} else {
				return layer.msg('上传失败');
			}
		}
	});
});
```