var $ = layui.jquery,
    layer = layui.layer,
    postTable;

$(function () {
    $("#add-post").on("click", function () {
        var videoid = $("input[name=videoid]:eq(0)").val()
            , postlogo = $("input[name=postlogo]:eq(0)").val()
            , params = {
                time: 5000,
                icon: 2,
            };
        if (videoid == "") {
            layer.msg("视频ID为空，请输入视频ID", params);
            return;
        } else if (!/^[0-9]+$/.test(videoid)) {
            layer.msg("视频ID格式错误，请重新输入视频ID", params);
            $("input[name=videoid]:eq(0)").select();
            return;
        } else if (postlogo == "") {
            layer.msg("未上传海报，请上传海报", params);
            return;
        }
        $.ajax({
            url: "../video/poster.json"
            , type: "post"
            , data: {
                videoid: videoid,
                postlogo: postlogo,
            }, success: function (res) {
                if (res.code == 0) {
                    layer.msg("上传成功");
                    window.location.reload();
                } else {
                    layer.msg(res.msg, params)
                }
            }
        });
    });

});

layui.use(["form", "table", "upload"], function () {
    var $ = layui.jquery,
        form = layui.form,
        table = layui.table,
        upload = layui.upload,
        mapper = new Map();
    postTable = table.render({
        elem: "#postlist-table",
        url: "../video/poster.json",
        defaultToolbar: [
            "filter",
            "exports",
            "print",
        ],
        cols: [
            [
                { field: "id", width: 100, title: "海报ID", sort: true },
                { field: "videoid", width: 100, title: "视频ID", sort: true },
                { field: "postlogo", width: 160, title: "海报", align: "center", templet: function (d) { return '<img src="' + d.postlogo + '">' } },
                { field: "pubtime", width: 160, title: "发布时间" },
                { field: "videoname", minWidth: 240, title: "视频名称" },
            ],
        ],
        limit: 5,
        even: true,
    });

    // video upload
    // upload.render({
    //     elem: "#post-upload-btn",
    //     url: "../common/uploader?type=image-post",
    //     accept: "images",
    //     acceptMine: "image/*",
    //     // size:,
    //     field: "file",
    //     number: 1,
    //     done: function (res) {
    //         if (res.code == 0) {
    //             layer.msg('上传成功');
    //             $("#post-image").attr("src", res.data.src);
    //             $("#post-input").val(res.data.src);
    //         } else {
    //             layer.msg(
    //                 "上传图片出错", {
    //                 icon: 2,
    //                 time: 5000,
    //             });
    //         }
    //     },
    //     error: function () {
    //         layer.msg(
    //             "上传图片出错", {
    //             icon: 2,
    //             time: 5000,
    //         });
    //     }
    // });
});

layui.config({
    base: '/static/lib/cropper/' //layui自定义layui组件目录
}).use(['form', 'croppers'], function () {
    var $ = layui.jquery
        , croppers = layui.croppers;
    croppers.render({
        elem: '#post-upload-btn'
        , saveW: 730     //保存宽度
        , saveH: 320
        , mark: 73 / 32    //选取比例
        , area: '900px'  //弹窗宽度
        , url: "../common/uploader?type=image-post"  //图片上传接口返回和（layui 的upload 模块）返回的JOSN一样
        , done: function (url) { //上传完毕回调
            layer.msg('上传成功');
            $("#post-image").attr("src", url);
            $("#post-input").val(url);
        }, error: function () {
            layer.msg(
                "上传图片出错", {
                icon: 2,
                time: 5000,
            });
        }
    });
});