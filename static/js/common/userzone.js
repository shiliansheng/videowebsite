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

var userid = GetQueryVariable(window.location.search.substring(1), "id"),
    layer = layui.layer;
$(function () {
    userid = GetQueryVariable(window.location.search.substring(1), "id");
    setNavAndContent($(".nav-this").eq(0));
    $(".nav-item").on("click", function () {
        setNavAndContent($(this));
    });
    DoAjax("../video/getucollect.json", "GET", { id: userid }, function (res) {
        if (res.code == 0) {
            setListInfo("collect", res.count);
            appendListItemToDom($(".collect-box").eq(0), "collect", res.data);
        } else {
            layer.alert(res.msg, {
                icon: 2,
                time: 5000,
            });
        }
    });
    DoAjax("../video/getuhistory.json", "GET", { id: userid }, function (res) {
        if (res.code == 0) {
            setListInfo("history", res.count);
            appendListItemToDom($(".history-box").eq(0), "collect", res.data);
        } else {
            layer.alert(res.msg, {
                icon: 2,
                time: 5000,
            });
        }
    });
})


var setNavAndContent = function (obj) {
    if (obj.attr("data-a") == "") {
        window.location.href = "/";
    } else {
        $(".nav-this").removeClass("nav-this");
        obj.addClass("nav-this");
        $(".main-content").css("display", "none");
        var module = obj.attr("data-a"),
            boxStyle = "flex";
        if (module == "collect" || module == "history") {
            boxStyle = "grid"
        }
        if (module == "collect") {
           
        } else if (module == "history") {
           
        }
        $(".main-content").filter("[for=" + module + "]").css("display", boxStyle)
    }
}

layui.use(["form", "miniTab", 'laydate'], function () {
    var form = layui.form,
        miniTab = layui.miniTab,
        laydate = layui.laydate,
        layer = layui.layer,
        $ = layui.$;
    // 初始化form信息
    $.ajax({
        url: "../common/userinfo.json?",
        type: "get",
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
    //监听提交
    form.on("submit(userinfo-saveBtn)", function (data) {
        layer.confirm('确认修改?', function () {
            DoAjax("userzone.json?action=update", "post", data.field, function (res) {
                if (res.code == 0) {
                    layer.msg(res.msg, { icon: 1 });
                    location.reload();
                } else {
                    layer.msg(res.msg, { icon: 2 });
                }
            });
        });
    });
    form.on("submit(userpass-saveBtn)", function (data) {
        layer.confirm("是否修改密码?", function () {
            DoAjax("common/userpwd.json?action=update", "post", data.field, function (res) {
                if (res.code == 0) {
                    layer.msg("修改密码成功", { icon: 1 });
                } else {
                    layer.msg(res.msg, { icon: 2 });
                    $("input[name=old_password]").select();
                }
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


/******* 
 * 给DOM元素添加内容
 * @param  obj DOM对象
 * @param  prefix 标志前缀
 * @param  data 对象数组，含有内容id, pubtime, videoinfo{}
 */
var appendListItemToDom = function (obj, prefix, data) {
    obj.html('');
    for (var i = 0, len = data.length; i < len; i++) {
        obj.append(
            '<div class="list-item" id="' + prefix + '-' + data[i].id + '">' +
            '    <a href="javascript:;" class="delete-btn" data-a="' + prefix + '" data-b="' + data[i].id + '"><i class="fa fa-trash" aria-hidden="true"></i></a>' +
            '    <div class="img-box">' +
            '        <img src="' + data[i].videoinfo.videologo + '" alt="">' +
            '    </div>' +
            '    <div class="item-info">' +
            '        <a href="/play.html?id=' + data[i].videoinfo.videologo + '" class="item-title">' + data[i].videoinfo.videoname + '</a>' +
            '        <div class="item-type"><i class="fa fa-bookmark" aria-hidden="true"></i> ' + data[i].videoinfo.classification + '/' + data[i].videoinfo.typename + '</div>' +
            '        <div class="item-time"><i class="fa fa-clock-o" aria-hidden="true"></i> ' + data[i].pubtime + '</div>' +
            '    </div>' +
            '</div>'
        )
    };
    var confirmText = "是否取消收藏？";
    if (prefix == "history") {
        confirmText = "是否删除该历史记录？";
    }
    $(".delete-btn").on("click", function () {
        var itemtype = $(this).attr("data-a"),
            itemid = $(this).attr("data-b");
        layer.confirm(confirmText, function () {
            DoAjax("../video/delete" + itemtype + ".json", "POST", {
                "id": itemid
            }, function (res) {
                if (res.code == 0) {
                    layer.msg("删除成功");
                    $("#" + itemtype + "-" + itemid).remove();
                    setListInfo(itemtype, $("." + itemtype + "-box .list-item").length)
                } else {
                    layer.alert(res.msg, {
                        icon: 2,
                        time: 7000,
                    })
                }
            });
        })
    });
}

/******* 
 * 进行ajax操作
 * @param  url 请求地址
 * @param  type 请求类型
 * @param  data 参数对象
 * @param  func success函数
 */
var DoAjax = function (url, type, data, func) {
    if (type == "DELETE") {
        url = url + "?collectid=" + data.collectid;
    }
    $.ajax({
        url: url
        , type: type
        , data: data
        , dataType: "json"
        , success: func
    })
}

/******* 
 * 设置listname
 * @param  listname history, collect 
 * @param  num 记录数量
 * @return 
 */
var setListInfo = function (listname, num) {
    var $listbox = $("." + listname + "-box").eq(0),
        nameText,
        infoText;
    if (listname == "history") {
        nameText = "观看"
    } else if (listname == "collect") {
        nameText = "收藏"
    }
    infoText = "您的 " + nameText + " 记录为空，当前记录数量为 0 条";
    if (num != undefined) {
        infoText = "您的 " + nameText + " 记录共 " + num + " 条"
    }
    $listbox.attr(listname + "-info", infoText);
}