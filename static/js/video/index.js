var layer = layui.layer
    , navList = ["movie", "cartoon", "episode", "others", "library"]
    , classMap = new Map([
        ["movie", "电影"],
        ["cartoon", "动画"],
        ["episode", "剧集"],
        ["others", "其他"],
        ["library", "片库"]
    ]);

/******* 
 * 根据query获取variable键对应的值
 * @param  query  window.location.search.substring(1)
 * @param  variable 需要获取的值
 * @return 返回获取的variable
 */
var GetQueryVariable = function (query, variable) {
    var vars = query.split("&");
    for (var i = 0; i < vars.length; i++) {
        var pair = vars[i].split("=");
        if (pair[0] == variable) { return pair[1]; }
    }
    return undefined;
}

// 获取当前page信息localhost:8088://movie/
var GetCurrentPageName = function () {
    var href = window.location.href,
        pieces = href.split("/");
    return pieces[3];
}

/******* 
 * 设置跳转页
 *   如果当前跳转页是/的话，则正常按照a标签进行跳转
 *   如果当前页和跳转页相同
 */
var setNavJump = function () {
    for (var i = 0; i < navList.length; i++) {
        var $obj = $("#" + navList[i] + " a").eq(0);
        $obj.on("click", function () {
            //                  stateObj, title, url
            var jumpPageName = $(this).attr("data-a"),
                currPageName = GetCurrentPageName();
            if (currPageName == "" || currPageName == undefined) {
                window.location.href = "/" + jumpPageName + "/";
            } else if (currPageName != jumpPageName) {
                history.replaceState(null, null, "/" + jumpPageName + "/");
                $(".layui-this").removeClass("layui-this");
                $("#" + jumpPageName).addClass("layui-this");
                ResetFilterAndSort();
                reLoadVideo();
            }
        })
    }
}

/******* 
 * 获取用户登录后的盒子内容
 * @param  logo 用户头像
 * @param  name 用户名
 * @return 返回用户登陆后的盒子内容
 */
var LoadUserBox = function (logo, name) {
    return '<a href="javascript:;" class="user-zone"><img src="' + logo + '" class="layui-nav-img"></a>' +
        '<dl class="layui-nav-child">' +
        '    <dd>' + name + '</dd>' +
        '    <dd><a href="javascript:;" class="user-zone"><i class="fa fa-user" aria-hidden="true"></i> 个人空间</a></dd>' +
        '    <dd><a href="javascript:;" id="login-out"><i class="fa fa-sign-out" aria-hidden="true"></i> 退出登录</a></dd>' +
        '</dl>';
};

/******* 
 * 获取用户未登录盒子内容
 * @return 返回用户未登录盒子内容
 */
var LoadLoginBox = function () {
    return '<div class="login-box">' +
        '    <a href="/video/login.html">登录</a>' +
        '    <a href="/register.html">注册</a>' +
        '</div>';
};

/******* 
 * 重新设置library类页面的筛选选项
 */
var ResetFilterAndSort = function () {
    $(".filter-box").show();
    $(".sort-box").show();
    $(".filter-row .on").removeClass("on");
    $(".filter-row dd").eq(0).addClass("on");
    $("input:radio:first").prop("checked", "checked");
}


/******* 
 * 设置用户登陆后的盒子中的链接跳转
 * @param  userid 用户id
 */
var LoadUserFunc = function (userid) {
    $("#login-out").on('click', function () {
        $.ajax({
            url: "../common/logout.json"
            , type: "post"
            , success: function (res) {
                if (res.code == 0) {
                    layer.msg("成功退出登录");
                    window.location = '/video/login.html';
                } else {
                    layer.msg("退出登录失败");
                }
            }
        })
    });

    $(".user-zone").on('click', function () {
        window.location = '/userzone.html?id=' + userid;
    });
};

