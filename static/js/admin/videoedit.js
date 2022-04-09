layui.use(["form", "layer"], function () {
    var form = layui.form,
        layer = layui.layer,
        $ = layui.jquery;
    // 初始化videoinfo
    $.ajax ({
        type: "post",
        url: "../common/videoinfo.json?id=" + $("input[name=id]").val(),
        success: function(res) {
            if (res.code == 0) {
                var obj = JSON.parse(res.data),
                    typenames = obj.typename.split("/");
                for (var i in typenames) {
                    obj["typename[" + typenames[i] + "]"] = typenames[i];
                }
                form.val("videoinfo", obj);
                document.getElementById("img-image").setAttribute("src", obj.videologo);
            }

        },
    });
    //监听提交
    form.on("submit(saveBtn)", function (data) {
        $.ajax({
            url: "videoedit.json?action=update",
            type: "post",
            data: data.field,
            success: function (res) {
                if (res.code == 0) {
                    layer.close(layer.index);
                    window.parent.location.reload();
                } else {
                    layer.alert(
                        res.msg,
                        {
                            title: "信息",
                            icon: 2,
                            time: 5000,
                        }
                    );
                }
            },
        });
        return false;
    });
});

layui.config({
    base: '/static/lib/cropper/' //layui自定义layui组件目录
}).use(['form', 'croppers'], function () {
    var $ = layui.jquery
        , croppers = layui.croppers;
    croppers.render({
        elem: '#img-button'
        , saveW: 193     //保存宽度
        , saveH: 289
        , mark: 193 / 289    //选取比例
        , area: '900px'  //弹窗宽度
        , url: "../common/uploader?type=image-video"  //图片上传接口返回和（layui 的upload 模块）返回的JOSN一样
        , done: function (url) { //上传完毕回调
            $("#img-input").val(url);
            $("#img-image").attr('src', url);
        }
    });
});