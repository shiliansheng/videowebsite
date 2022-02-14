layui.use(["form", "miniTab", 'laydate'], function () {
	var form = layui.form,
		layer = layui.layer,
		miniTab = layui.miniTab,
		laydate = layui.laydate,
		$ = layui.$;
	// laydate实现
	laydate.render ({
		elem: '#dateselect',
		//type: 'datetime',
		//range: true
	});
	//监听提交
	form.on("submit(saveBtn)", function (data) {
		layer.confirm('确认修改?', function(){
			console.log(data.field);
			$.ajax({
				type : "post",
				url : "user_setting.json?action=changeSetting",
				data : data.field,
				success : function(res){
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
