layui.use(["form", "table"], function () {
    var $ = layui.jquery,
        form = layui.form,
        table = layui.table,
        mapper = new Map();
    table.render({
        elem: "#currentTableId",
        url: "videotypelist.json?action=getlist",
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
                { field: "pid", width: 80, title: "父ID", sort: true },
                { field: "typename", width: 120, title: "类型名称" },
                { field: "addid", width: 90, title: "添加人ID" },
                { field: "vtypelogo", width: 130, title: "类型logo", align: 'center', templet: function (d) {
                    var path = d.vtypelogo
                    if (path == "") {
                        path = NoPicPath
                    }
                    return '<img src="' + path + '">'
                }},
                { field: "createat", width: 160, title: "添加时间" },
                { field: "sequence", width: 110, title: "显示顺序", sort: true },
                { field: "discription", minWidth: 120, title: "描述", sort: true },
                { title: "操作", width: 120, toolbar: "#currentTableBar", align: 'center', fixed: "right", },
            ],
        ],
        limits: [5, 10, 15, 20],
        limit: 5,
        page: true,
        request: {
            page: "page", //页码的参数名称，默认：page
            limit: "limit", //每页数据量的参数名，默认：limit
        }
    });

    // 监听搜索操作
    form.on("submit(data-search-btn)", function (data) {
        var result = JSON.stringify(data.field);
        console.log(data.field)
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
                title: "添加视频分类",
                type: 2,
                shade: 0.2,
                maxmin: true,
                shadeClose: true,
                area: ["100%", "100%"],
                content: "videotypeadd.html",
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
                delUserInfo += "</br>类型" + (parseInt(idx) + 1) + " : " + data[idx].typename
            }
            var deldata = new Object()
            deldata.Data = JSON.stringify(data) 
            layer.confirm("确定删除类型？" + delUserInfo, function (index) {
                $.ajax({
                    url: "videotypedel?more=true",
                    type: "post",
                    data: deldata,
                    success: function (res) {
                        if (res.code == 0) {
                            layer.msg(res.msg);
                            table.reload("currentTableId");
                        } else {
                            var index = layer.alert(
                                res.msg,
                                {
                                    title: "信息",
                                    icon: 2,
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
                        table.reload("currentTableId");
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
    table.on("tool(currentTableFilter)", function (obj) {
        var data = obj.data;
        if (obj.event === "edit") {
            var index = layer.open({
                title: "编辑视频类型",
                type: 2,
                shade: 0.2,
                maxmin: true,
                shadeClose: true,
                area: ["100%", "100%"],
                content: "videotypeedit.html?action=getinfo&id=" + data.id,
            });
            $(window).on("resize", function () {
                layer.full(index);
            });
            return false;
        } else if (obj.event === "delete") {
            layer.confirm("确定删除该类型？", function (index) {
                var deldata = new Object()
                deldata.Data = JSON.stringify(data) 
                $.ajax({
                    url: "videotypedel?more=false",
                    type: "post",
                    data: deldata,
                    success: function (res) {
                        if (res.code == 0) {
                            layer.open({
                                type: 1,
                                content: res.msg
                            });
                            obj.del();
                            table.reload("currentTableId");
                        } else {
                            var index = layer.alert(
                                res.msg,
                                {
                                    title: "信息",
                                    icon: 2,
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
