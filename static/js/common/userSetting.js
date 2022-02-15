layui.use(["form", "miniTab", 'laydate'], function () {
	var form = layui.form,
		layer = layui.layer,
		miniTab = layui.miniTab,
		laydate = layui.laydate,
		$ = layui.$;
	form.val("baseinfo", userInfo);
	// laydate实现
	laydate.render({
		elem: '#dateselect',
	});
	form.verify({
		verifyEmail: function (value) {
			var ePattern =
				/^([A-Za-z0-9_\-\.])+\@([A-Za-z0-9_\-\.])+\.([A-Za-z]{2,4})$/;
			if (value != "" && !ePattern.test(value)) {
				return "邮箱格式不正确";
			}
		},
	});
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
});

layui.config({
	base: '../../static/lib/cropper/' //layui自定义layui组件目录
}).use(['form', 'croppers'], function () {
	var $ = layui.jquery
		, form = layui.form
		, croppers = layui.croppers
		, layer = layui.layer;
	croppers.render({
		elem: '#userlogo'
		, saveW: 120	//保存宽度
		, saveH: 120	//保存高度
		, mark: 1 / 1	//选取比例
		, area: '800px'	//弹窗宽度
		, method: 'post'
		, url: "upload_img.json"  //图片上传接口返回和（layui 的upload 模块）返回的JOSN一样
		, done: function (url) { //上传完毕回调
		}
	});
});