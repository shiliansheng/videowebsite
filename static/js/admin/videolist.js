layui.use(["form", "table"], function () {
    var $ = layui.jquery,
        form = layui.form,
        table = layui.table,
        mapper = new Map();
    table.render({
        elem: "#videolist-table",
        url: "videolist.json",
        toolbar: "#tableToolBar",
        defaultToolbar: [
            "filter",
            "exports",
            "print",
        ],
        cols: [
            [
                { type: "checkbox",},
                { field: "id", width: 60, title: "ID", sort: true },
                { field: "videoname",  title: "视频名称" },
                { field: "classification", width: 87, title: "分类" },
                { field: "typename", width: 130, title: "类型" },
                // { field: "keywords",  title: "关键字" },
                { field: "copyright", width: 60, title: "版权" },
                { field: "introduction",  title: "视频介绍" },
                { field: "pubtime",  width: 160, title: "上传时间" },
                { field: "username", width: 100, title: "发布者" },
                { field: "viewnum", width: 87, title: "观看次数" },
                { field: "reviewnum", width: 87, title: "评论次数" },
                { field: "averscore", width: 87, title: "用户评分" },
                { field: "passed", width: 87, title: "审核状态" },
                { title: "操作", width: 120, toolbar: "#currentTableBar", align: "center", fixed: "right", },
            ],
        ],
        limits: [5, 10, 15, 20],
        limit: 10,
        page: true,
        even: true,
    });

    // 监听搜索操作
    form.on("submit(data-search-btn)", function (data) {
        var result = JSON.stringify(data.field);
        //执行搜索重载
        table.reload(
            "videolist-table",
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
                title: "添加视频",
                type: 2,
                shade: 0.2,
                maxmin: true,
                shadeClose: true,
                area: ["100%", "100%"],
                content: "videoadd.html",
            });
            $(window).on("resize", function () {
                layer.full(index);
            });
        } else if (obj.event === "delete") {
            // 监听删除操作
            var checkStatus = table.checkStatus("videolist-table"),
                data = checkStatus.data;
            var delInfo = "", idlist = [];
            for (var idx in data) {
                delInfo += "</br>类型" + (parseInt(idx) + 1) + " : " + data[idx].typename
                idlist.push(data[idx].id)
            }
            layer.confirm("确定删除？" + delInfo, function (index) {
                $.ajax({
                    url: "videodel",
                    type: "post",
                    data: {idlist: JSON.stringify(idlist)},
                    success: function (res) {
                        if (res.code == 0) {
                            layer.msg(res.msg);
                        } else {
                            var index = layer.msg(
                                res.msg,
                                {
                                    title: "信息",
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
                        table.reload("videolist-table");
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
                content: "videoedit.html?id=" + data.id,
            });
            $(window).on("resize", function () {
                layer.full(index);
            });
            return false;
        } else if (obj.event === "delete") {
            layer.confirm("确定删除该类型？", function (index) {
                $.ajax({
                    url: "videodel",
                    type: "post",
                    data: {idlist: JSON.stringify([data.id])},
                    success: function (res) {
                        if (res.code == 0) {
                            layer.msg(res.msg);
                            obj.del()
                            table.reload("videolist-table");
                        } else {
                            var index = layer.msg(
                                res.msg,
                                {
                                    title: "删除视频失败",
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
