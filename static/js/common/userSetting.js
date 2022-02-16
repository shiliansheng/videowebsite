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
	base: '/static/lib/cropper/' //layui自定义layui组件目录
}).use(['form', 'avatar'], function () {
	var $ = layui.jquery
		, form = layui.form
		, avatar = layui.avatar
		, layer = layui.layer;
	avatar.render({
		elem: "#userlogo"
		, success: function (base64, size) {
			var data = new Object()
			data.image = base64
			$.ajax({
				type: "post"
				, url: "uploader?type=user-image"
				, data: data
				, success: function(res) {
					if (res.code == 0) {
						layer.msg('上传成功')
						document.getElementById("userlogoInput").value = res.data
						document.getElementById('logoimg').setAttribute("src", base64)
					} else {
						layer.msg(res.msg)
					}
				}
			});
		}
	});
});