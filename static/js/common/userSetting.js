layui.use(["form", "miniTab"], function () {
	var form = layui.form,
		layer = layui.layer,
		miniTab = layui.miniTab;

	//监听提交
	form.on("submit(saveBtn)", function (data) {
		var index = layer.alert(
			JSON.stringify(data.field),
			{
				title: "最终的提交信息",
			},
			function () {
				layer.close(index);
				miniTab.deleteCurrentByIframe();
			}
		);
		return false;
	});

	form.val("baseinfo", userInfo);
});
