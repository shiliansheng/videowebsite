layui.use(["form", "miniTab"], function () {
    var form = layui.form,
        layer = layui.layer,
        miniTab = layui.miniTab,
        $ = layui.$;
    //监听提交
    form.on("submit(saveBtn)", function (data) {
        layer.confirm("是否修改密码?", function () {
            $.ajax({
                type: "post",
                url: "userpwd.json?action=update",
                data: data.field,
                success: function (res) {
                    if (res.code == 0) {
                        layer.msg("修改密码成功", {icon: 1});
                        miniTab.deleteCurrentByIframe();
                    } else {
                        layer.msg(res.msg, {icon: 2});
                        $("input[name=old_password]").select();
                    }
                },
            });
        });
    });
    form.verify({
        verifyPassword: function (value) {
            //密码强度正则，最少6位，只能由 字母，数字，下划线，减号，特殊符号 组成
            pPattern = /^[!@#$%^&*?a-zA-Z0-9_-]{6,}$/;
            if (!pPattern.test(value)) {
                return "密码最少6位最少6位，只能由 字母，数字，下划线，减号，特殊符号 组成";
            } else if ($("input[name=old_password]").val() == value) {
                return "新的密码不能和旧密码一致";
            }
        },
        confirmPassword: function (value) {
            if ($("input[name=new_password]").val() !== value)
                return "两次密码输入不一致！";
        },
    });
});
