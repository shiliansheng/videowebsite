layui.use(["form"], function () {
    var form = layui.form,
        layer = layui.layer,
        $ = layui.$;
    // 校验两次密码是否一致
    form.verify({
        verifyUsername: function (value) {

            if (value.length < 3 || value.length > 16) {
                return "用户名长度至少为3";
            } else if (!uPattern.test(value)) {
                return "用户名只能由 字母，数字，下划线，减号，@ 组成";
            }
        },
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
            var ePattern =
                /^([A-Za-z0-9_\-\.])+\@([A-Za-z0-9_\-\.])+\.([A-Za-z]{2,4})$/;
            if (value != "" && !ePattern.test(value)) {
                return "邮箱格式不正确";
            }
        },
    });
    //监听提交
    form.on("submit(saveBtn)", function (data) {
        $.ajax({
            url: "/common/register.json",
            type: "post",
            data: data.field,
            success: function (res) {
                if (res.code == 0) {
                    layer.msg("注册成功", function() {
                        window.location = "/"
                    });
                } else {
                    layer.alert(res.msg, {
                        title: "添加用户失败",
                        time: 5000,
                    },
                    );
                }
            },
        });
        return false;
    });
    $("input[name=username]").on("change", function () {
        var unameStr = $("input[name=username]").eq(0).val(),
            $uniqueLabel = $("#uname-unqiue-label");
        if (uPattern.test(unameStr)) {
            $.ajax({
                url: "../common/unameunique.json",
                type: "post",
                data: {
                    "username": unameStr
                },
                success: function (res) {
                    if (res.code == 0) {
                        $uniqueLabel.removeClass("layui-icon-about");
                        $uniqueLabel.addClass("layui-icon-ok-circle");
                        $uniqueLabel.css("color", "green");
                    } else {
                        layer.tips("用户名不可用", '#uname-unqiue-label');
                        $uniqueLabel.removeClass("layui-icon-about");
                        $uniqueLabel.addClass("layui-icon-about");
                        $uniqueLabel.css("color", "red");
                    }
                },
            });
        } else {
            $uniqueLabel.removeClass("layui-icon-ok-circle");
            $uniqueLabel.removeClass("layui-icon-about");
        }
    });
});


var uPattern = /^[@a-zA-Z0-9_-]{3,16}$/;