layui.use(["form"], function () {
	var form = layui.form,
		layer = layui.layer,
		$ = layui.$;
	// 校验两次密码是否一致
	form.verify({
		verifyUsername: function (value) {
			var uPattern = /^[@a-zA-Z0-9_-]{3,16}$/;
			if (value.length < 3 || value.length > 16) {
				return "用户名长度至少为3";
			} else if (!uPattern.test(value)) {
				return "用户名只能由 字母，数字，下划线，减号，@ 组成";
			}
		},
		verifyPassword: function (value) {
			//密码强度正则，最少6位，只能由 字母，数字，下划线，减号，特殊符号 组成
			pPattern = /^[!@#$%^&*?a-zA-Z0-9_-]{6,}$/;
			if (!pPattern.test(value)) {
				return "密码最少6位最少6位，只能由 字母，数字，下划线，减号，特殊符号 组成";
			}
		},
		confirmPassword: function (value) {
			if ($("input[name=password]").val() !== value)
				return "两次密码输入不一致！";
		},
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
		$.ajax({
			url: "useradd.html",
			type: "post",
			data: data.field,
			success: function (res) {
				if (res.code == 0) {
                    // layer.msg("添加用户成功");
					layer.close(layer.index);
					window.parent.location.reload();
				} else {
					var index = layer.alert(
						res.msg,
						{
							title: "添加用户失败",
						},
						function () {
							// 关闭弹出层
							layer.close(index);

							var iframeIndex = parent.layer.getFrameIndex(
								window.name
							);
							parent.layer.close(iframeIndex);
						}
					);
				}
			},
		});
		return false;
	});
});
