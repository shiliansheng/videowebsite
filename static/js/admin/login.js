layui.use(['form'], function () {
    var form = layui.form,
        layer = layui.layer;

    // 登录过期的时候，跳出ifram框架
    if (top.location != self.location) top.location = self.location;

    form.on("submit(login)", function (data) {
        $.ajax({
            url: "login.json"
            ,type: "post"
            ,data: data.field,
            success: function(res) {
                if (res.code == 0) {
                    layer.msg(res.msg, {
                        icon: 1,
                        time: 1500,
                    });
                    window.location.href = res.data
                } else {
                    layer.msg(res.msg, {
                        icon: 2,
                        time: 3000,
                    })
                }
            }
        })
        return false; // 禁止跳转，未指定跳转时，将会跳转到当前页面
    })

    // 粒子线条背景
    $(document).ready(function () {
        $('.layui-container').particleground({
            dotColor: '#7ec7fd',
            lineColor: '#7ec7fd'
        });
    });
});