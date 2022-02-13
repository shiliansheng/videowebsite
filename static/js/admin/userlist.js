layui.use(["form", "table"], function () {
	var $ = layui.jquery,
		form = layui.form,
		table = layui.table,
        mapper = new Map();
	table.render({
		elem: "#currentTableId",
		url: "http://localhost:" + httpPort + "/admin/getuserlist.json",
		toolbar: "#toolbarDemo",
		defaultToolbar: [
			"filter",
			"exports",
			"print",
			{
				title: "提示",
				layEvent: "LAYTABLE_TIPS",
				icon: "layui-icon-tips",
			},
		],
		cols: [
			[
				{ type: "checkbox", width: 50 },
				{ field: "id", width: 60, title: "ID", sort: true },
				{ field: "username", width: 120, title: "用户名", sort: true },
				{ field: "password", width: 120, title: "密码" },
				{ field: "nickname", width: 120, title: "昵称" },
				{ field: "sex", width: 75, title: "性别", sort: true },
				{ field: "email", width: 200, title: "电子邮箱" },
				{ field: "status", width: 90, title: "身份" },
				{ field: "createat", width: 160, title: "添加时间", sort: true},
				{ field: "remark", minWidth: 200, title: "备注" },
				{ title: "操作", width: 120, toolbar: "#currentTableBar", fixed: "right", },
			],
		],
		limits: [10, 15, 20, 25, 50, 100],
		limit: 10,
		page: true,
		request: {
			page: "page", //页码的参数名称，默认：page
			limit: "limit", //每页数据量的参数名，默认：limit
		},
	});

	// 监听搜索操作
	form.on("submit(data-search-btn)", function (data) {
		var result = JSON.stringify(data.field);
		//执行搜索重载
		table.reload(
			"currentTableId",
			{
				page: {
					curr: 1,
				},
				where: {
					searchParams: result,
				},
			},
			"data"
		);

		return false;
	});

	/**
	 * toolbar监听事件
	 */
	table.on("toolbar(currentTableFilter)", function (obj) {
		if (obj.event === "add") {
			// 监听添加操作
			var index = layer.open({
				title: "添加用户",
				type: 2,
				shade: 0.2,
				maxmin: true,
				shadeClose: true,
				area: ["100%", "100%"],
				content: "useradd.html",
			});
			$(window).on("resize", function () {
				layer.full(index);
			});
		} else if (obj.event === "delete") {
			// 监听删除操作
			var checkStatus = table.checkStatus("currentTableId"),
				data = checkStatus.data;
			// layer.alert(JSON.stringify(data));
            var delUserInfo = ""
            for (var idx in data) {
                delUserInfo += "</br>用户" + (parseInt(idx) + 1) + " : " + data[idx].username
            }
            layer.confirm("确定删除用户？" + delUserInfo, function (index) {
				$.ajax({
					url: "userdel?more=true",
					type: "post",
					data: JSON.stringify(data),
					success: function (res) {
						if (res.code == 0) {
							layer.msg(res.msg);
						} else {
							var index = layer.alert(
								res.msg,
								{
									title: "删除用户失败",
								},
								function () {
									// 关闭弹出层
									layer.close(index);
								}
							);
						}
                        var successlist = res.successlist;
                        for (var i in successlist) {
                            mapper[successlist[i]].del();
                        }
					},
				});
				layer.close(index);
			});
		}
	});

	//监听表格复选框选择
	table.on("checkbox(currentTableFilter)", function (obj) {
		mapper[obj.data.id] = obj
	});
	/**
	 * param 将要转为URL参数字符串的对象
	 * key URL参数字符串的前缀
	 * encode true/false 是否进行URL编码,默认为true
	 * return URL参数字符串
	 */
	var urlEncode = function (param, key, encode) {
		if (param == null) return "";
		var paramStr = "";
		var t = typeof param;
		if (t == "string" || t == "number" || t == "boolean") {
			paramStr +=
				"&" +
				key +
				"=" +
				(encode == null || encode ? encodeURIComponent(param) : param);
		} else {
			for (var i in param) {
				var k =
					key == null
						? i
						: key +
						  (param instanceof Array ? "[" + i + "]" : "." + i);
				paramStr += urlEncode(param[i], k, encode);
			}
		}
		return paramStr;
	};
	table.on("tool(currentTableFilter)", function (obj) {
		var data = obj.data;
		if (obj.event === "edit") {
			var index = layer.open({
				title: "编辑用户",
				type: 2,
				shade: 0.2,
				maxmin: true,
				shadeClose: true,
				area: ["100%", "100%"],
				content: "useredit.html?" + urlEncode(data),
			});
			$(window).on("resize", function () {
				layer.full(index);
			});
			return false;
		} else if (obj.event === "delete") {
			layer.confirm("确定删除用户？", function (index) {
				$.ajax({
					url: "userdel?more=false",
					type: "post",
					data: JSON.stringify(data),
					success: function (res) {
						if (res.code == 0) {
							layer.msg("删除用户成功");
							obj.del();
						} else {
							var index = layer.alert(
								res.msg,
								{
									title: "删除用户失败",
								},
								function () {
									// 关闭弹出层
									layer.close(index);
								}
							);
						}
					},
				});
				layer.close(index);
			});
		}
	});
});
