layui.use(["form", "table"], function () {
    var $ = layui.jquery,
        form = layui.form,
        table = layui.table,
        mapper = new Map();
    table.render({
        elem: "#userlist-table",
        url: "userlist.json",
        toolbar: "#tableToolBar",
        defaultToolbar: [
            "filter",
            "exports",
            "print",
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
                { field: "status", width: 100, title: "身份" },
                { field: "createat", width: 160, title: "添加时间", sort: true },
                { field: "remark", minWidth: 200, title: "备注" },
                { title: "操作", width: 120, toolbar: "#itemToolBar", fixed: "right", align: "center" },
            ],
        ],
        limits: [10, 15, 20, 25, 50, 100],
        limit: 10,
        page: true,
        even: true
    });

    // 监听搜索操作
    form.on("submit(data-search-btn)", function (data) {
        var result = JSON.stringify(data.field);
        //执行搜索重载
        table.reload(
            "userlist-table",
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

    // table toolbar监听事件
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
            var checkStatus = table.checkStatus("userlist-table"),
                data = checkStatus.data;
            var delInfo = "";
            var idlist = [];
            
            for (var idx in data) {
                delInfo += "</br>用户" + (parseInt(idx) + 1) + " : " + data[idx].username
                idlist.push(data[idx].id)
            }
            layer.confirm("确定删除用户？" + delInfo, function (index) {
                $.ajax({
                    url: "userdel",
                    type: "post",
                    data: {idlist: JSON.stringify(idlist)},
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
                        var successlist = res.data;
                        for (var i in successlist) {
                            mapper[successlist[i]].del();
                        }
                        table.reload("userlist-table");
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

    // item toolbar监听事件
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
                content: "useredit.html?id=" + data.id,
            });
            $(window).on("resize", function () {
                layer.full(index);
            });
            return false;
        } else if (obj.event === "delete") {
            layer.confirm("确定删除用户？", function (index) {
                $.ajax({
                    url: "userdel",
                    type: "post",
                    data: { idlist: JSON.stringify([data.id])},
                    success: function (res) {
                        if (res.code == 0) {
                            layer.msg(res.msg);
                            obj.del();
                            table.reload("userlist-table");
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
