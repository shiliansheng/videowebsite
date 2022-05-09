var starEmpty = 'fa-star-half-full'
    , starFull = 'fa-star'
    , actionActiveClass = 'action-active'
    , videoid = GetQueryVariable(window.location.search.substring(1), "id")
    , layer = layui.layer
    , userid;

// 初始化页面
$(function () {
    videoid = GetQueryVariable(window.location.search.substring(1), "id");
    // 配置标签
    var $videoLabelBox = $("#video-label"),
        typenames = ($videoLabelBox.attr("data-i") + "").split("/");
    for (var i = 0, len = typenames.length; i < len; i++) {
        $videoLabelBox.append('<div class="label-piece">' + typenames[i] + '</div>')
    }

    // 配置登录信息
    $.ajax({
        url: "../common/getlogininfo.json",
        type: "GET",
        success: function (res) {
            if (res.code == 0) {
                userid = res.data.id;
                LoadUserFunc();
                // 设置收藏按钮
                DoAjax("../video/scorer.json", "GET", {
                    "userid": userid,
                    "videoid": videoid,
                }, function (res) {
                    if (res.code == 0) {
                        setVideoRate(true, res.data / 2);
                    } else {
                        setVideoRate();
                    }
                });

                // 设置评分
                DoAjax("../video/collecter.json", "GET", {
                    "userid": userid,
                    "videoid": videoid,
                }, function (res) {
                    if (res.code == 0) {
                        modifyCollectBtn("GET", res.data);
                    }
                })

                // 设置用户框
                $(".user-box").eq(0).html(LoadUserBox(res.data.logo, res.data.name, userid))
                $("#review-send-content").attr("placeholder", "发表你的评论吧！");
                $("#review-send-content").removeAttr("disabled");
            } else {
                $(".user-box").eq(0).html(LoadLoginBox());
                $("#review-send-content").attr({
                    "placeholder": "请登录后发表评论.",
                    "disabled": "disabled"
                });
                setVideoRate(true, 0)
            }
        },
    });

    // 设置收藏按钮
    $("#collect-btn").on('click', function () {
        var collectid = $("#collect-btn").attr("data-i");
        if (userid == undefined) {
            layer.msg('收藏失败 请登陆后收藏');
        } else {
            var type = "DELETE"
                , url = "../video/collecter.json"
                , msg = "取消收藏"
                , data = {
                    "collectid": collectid,
                };
            if (collectid == "" || collectid == undefined) {
                msg = "收藏";
                type = "POST";
                data = {
                    "userid": userid,
                    "videoid": videoid,
                };
            }
            layer.confirm("是否" + msg + "本视频？", function () {
                DoAjax(url, type, data, function (res) {
                    if (res.code == 0) {
                        layer.msg(msg + "成功")
                        modifyCollectBtn(type, res.data);
                    } else {
                        layer.msg(msg + "失败")
                    }
                });
            });
        }
    });

    // 配置video
    var player = videojs('play-video', {
        controls: true,
        loop: false,
        width: 915,
        notSupportedMessage: '此视频暂无法播放!', // 无法播放时显示的信息
        controlBar: {
            volumePanel: {
                inline: false
            },
            currentTimeDisplay: true,       // 当前播放位置
            timeDivider: true,              // 时间分割线
            durationDisplay: true,          // 总时间
            progressControl: true,          // 进度条
            remainingTimeDisplay: true,     // 剩余时间
            fullscreenToggle: true,         // 全屏按钮
        }, plugins: {

        },
    }, function onPlayerReady() {
        var firstPlay = true;
        videojs.log('播放器已经准备好了!');
        this.on('ended', function () {
            videojs.log('播放结束了!');
        });
        this.on('play', function () {
            if (firstPlay) {
                firstPlay = false;
                DoAjax("../video/addviewnum.json", "POST", {
                    "id": videoid,
                })
            }
        });
    });

    // 设置评论发布按钮事件
    $("#submit-review-btn").click(function () {
        var reviewContent = $("#review-send-content").val();
        if (userid == undefined) {
            layer.msg('发表失败 请登陆后发送');
        } else if (reviewContent == "") {
            layer.msg('发表失败 评论内容为空');
        } else {
            $.ajax({
                url: "../video/submitreview.json"
                , type: "POST"
                , data: {
                    "userid": userid,
                    "videoid": videoid,
                    "content": reviewContent,
                }, success: function (res) {
                    if (res.code == 0) {
                        $("#review-send-content").val('');
                        layer.msg("发布成功");
                        $("#review-list-content").prepend(
                            '<li class="review-item">' +
                            '   <div class="review-userlogo">' +
                            '       <img src="' + $(".user-box img").eq(0).attr("src") + '" alt="">' +
                            '   </div>' +
                            '   <div class="review-info">' +
                            '       <div class="review-username">' + $(".user-box dd").eq(0).text() + '</div>' +
                            '       <div class="review-detail">评论发布于 ' + new Date().toLocaleString().replace("/", "-").replace("/", "-") + '</div>' +
                            '       <div class="review-content">' + reviewContent + '</div>' +
                            '   </div>' +
                            '</li>'
                        )
                    } else {
                        layer.alert(res.msg, {
                            icon: 2,
                            time: 5000,
                            closeBtn: 0,
                        })
                    }
                }
            });
        }
    });
});

/******* 
 * 设置视频评分组件
 *  如果评分成功，则进行ajax传递信息，返回信息成功则设置该组件只读
 * @param  readonly 是否进行
 * @param  value 组件的初始值
 */
var setVideoRate = function (readonly, value) {
    if (readonly == undefined) {
        readonly = false;
    }
    if (value == undefined) {
        value = 0;
    }
    layui.use(['layer', 'rate'], function () {
        var rate = layui.rate;
        layer = layui.layer
        //渲染
        rate1 = rate.render({
            elem: '#video-rate'  //绑定元素
            , length: 5
            , text: true
            , value: value
            , readonly: readonly
            , half: true
            , setText: function (value) {
                this.span.text(" " + value * 2 + " 分");
            }, choose: function (value) {
                $.ajax({
                    url: "../video/scorer.json"
                    , type: "POST"
                    , data: {
                        "userid": userid,
                        "videoid": videoid,
                        "value": value * 2,
                    }, success: function (res) {
                        if (res.code == 0) {
                            layer.msg("评分成功");
                            rate.render({
                                elem: '#video-rate'  //绑定元素
                                , length: 5
                                , text: true
                                , value: value
                                , readonly: true
                                , half: true
                                , setText: function (value) {
                                    this.span.text(" " + value * 2 + " 分");
                                }
                            });
                        } else {
                            layer.alert(res.msg, {
                                icon: 2,
                                time: 5000,
                                closeBtn: 0,
                            })
                        }
                    }
                });
            }
        });
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
 * 根据给定按钮的类型和返回内容，设置收藏按钮
 * @param  type 按钮类型，是进行获取、增加、删除中的一种
 * @param  data 返回收藏的id
 */
var modifyCollectBtn = function (type, data) {
    var $collectBtn = $("#collect-btn"),
        $collectLabel = $("#collect-btn i").eq(0);
    if (type == "GET") {
        $collectBtn.addClass(actionActiveClass);
        $collectLabel.removeClass(starEmpty);
        $collectLabel.addClass(starFull);
        $collectBtn.attr("data-i", data);
    } else if (type == "POST") {
        $collectBtn.addClass(actionActiveClass);
        $collectLabel.removeClass(starEmpty);
        $collectLabel.addClass(starFull);
        $collectBtn.attr("data-i", data);
    } else if (type == "DELETE") {
        $collectBtn.toggleClass(actionActiveClass);
        $collectLabel.toggleClass(starFull);
        $collectLabel.toggleClass(starEmpty);
        $collectBtn.attr("data-i", "");
    }
}