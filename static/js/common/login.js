// 粒子线条背景
$(document).ready(function () {
    $('.container').particleground({
        dotColor: '#7ec7fd',
        lineColor: '#7ec7fd'
    });
    $("#captcha").on("click", function() {
        console.log("click the captcha")
    });
});

layui.use(["form", "layer"], function() {
    var form = layui.form,
        layer = layui.layer,
        $ = layui.jquery;
    form.on("submit(login)", function (data) {
        $.ajax({
            url: "login.json"
            , type: "post"
            , data: data.field,
            success: function (res) {
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
                    if (res.code == 1001) {
                        setSrcQuery(document.getElementsByClassName('captcha-img')[0], "reload=" + (new Date()).getTime());
                    }
                }
            }
        })
        return false; // 禁止跳转，未指定跳转时，将会跳转到当前页面
    })
});

function setSrcQuery(e, q) {
    var src = e.src;
    var p = src.indexOf('?');
    if (p >= 0) {
        src = src.substr(0, p);
    }
    e.src = src + "?" + q
}