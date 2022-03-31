layui.use(["form", "miniTab", 'laydate'], function () {
    var form = layui.form,
        layer = layui.layer,
        miniTab = layui.miniTab,
        laydate = layui.laydate,
        $ = layui.$;
    // 初始化form信息
    $.ajax({
        url: "../common/userinfo.json?",
        type: "post",
        success: function (res) {
            if (res.code == 0) {
                form.val("userinfo", JSON.parse(res.data));
            }
        }
    })
    // laydate实现
    laydate.render({ elem: '#dateselect' });
    // 验证信息
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
            $.ajax({
                type: "post",
                url: "user_setting.json?action=update",
                data: data.field,
                success: function (res) {
                    if (res.code == 0) {
                        layer.msg(res.msg, { icon: 1 });
                    } else {
                        layer.msg(res.msg, { icon: 2 });
                    }
                },
            });
        });
    });
});

layui.config({
    base: '/static/lib/cropper/' //layui自定义layui组件目录
}).use(['form', 'croppers'], function () {
    var $ = layui.jquery
        , croppers = layui.croppers;
    croppers.render({
        elem: '#uploadBtn'
        , saveW: 150     //保存宽度
        , saveH: 150
        , mark: 1 / 1    //选取比例
        , area: '900px'  //弹窗宽度
        , url: "../common/uploader?type=image-user"  //图片上传接口返回和（layui 的upload 模块）返回的JOSN一样
        , done: function (url) { //上传完毕回调
            $("#userlogoInput").val(url);
            $("#uploadImg").attr('src', url);
        }
    });
});