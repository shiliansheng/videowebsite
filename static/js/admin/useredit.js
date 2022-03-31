layui.use(["form"], function () {
    var form = layui.form,
        layer = layui.layer,
        $ = layui.$;

    // 初始化form信息
    $.ajax({
        url: "../common/userinfo.json?id=" + $("input[name=id]").val(),
        type: "post",
        success: function (res) {
            if (res.code == 0) {
                form.val("userinfo", JSON.parse(res.data));
            }
        }
    })
    // 校验两次密码是否一致
    form.verify({
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
            var ePattern = /^([A-Za-z0-9_\-\.])+\@([A-Za-z0-9_\-\.])+\.([A-Za-z]{2,4})$/;
            if (value != "" && !ePattern.test(value)) {
                return "邮箱格式不正确";
            }
        },
    });
    //监听提交
    form.on("submit(saveBtn)", function (data) {
        $.ajax({
            url: "useredit.json",
            type: "post",
            data: data.field,
            success: function (res) {
                if (res.code == 0) {
                    layer.msg(res.msg)
                    layer.close(layer.index);
                    window.parent.location.reload();
                } else {
                    console.log(res);
                    layer.alert(
                        res.msg,
                        {
                            title: "修改用户失败",
                        },
                        //  function () {
                        //      // 关闭弹出层
                        //      layer.close(index);
                        //      var iframeIndex = parent.layer.getFrameIndex(
                        //          window.name
                        //      );
                        //      parent.layer.close(iframeIndex);
                        //  }
                    );
                }
            },
        });
        return false;
    });
});
