$(function () {
    $("#home").addClass("layui-this");
    var userid;
    $.ajax({
        url: "../common/getlogininfo.json",
        type: "GET",
        success: function (res) {
            if (res.code == 0) {
                userid = res.data.id;
                $(".user-box").eq(0).html(LoadUserBox(res.data.logo, res.data.name))
                LoadUserFunc(res.data.id)
            } else {
                $(".user-box").eq(0).html(LoadLoginBox());
            }
        },
    });
    loadVideoBox('movie');
    loadVideoBox('cartoon');
    loadVideoBox('episode');
    setNavJump();
});

layui.use(['carousel', 'element'], function () {
    var carousel = layui.carousel,
        element = layui.element;
    //轮播图
    carousel.render({
        elem: '#home-carousel'
        , width: '730px' //设置容器宽度
        , height: '320px'
        , arrow: 'hover' //始终显示箭头
        , interval: 5000
    });
});

/******* 
 * 加载所属classification的内容，limit默认为14
 * @param  key 所属classification
 */
var loadVideoBox = function(key) {
    var $obj = $("#" + key + "-box .module-content").eq(0);
    $.ajax({
        url: "reloadvideo.json"
        , type: "GET"
        , data: {
            page: 1,
            limit: 14,
            sort: "",
            classification: key,
        }, success: function (res) {
            var data = res.data;
            if (res.code == 0) {
                for (var i = 0, len = data.length; i < len; i++) {
                    $obj.append(
                        '<li class="module-item">' +
                        '     <a href="/play?id=' + res.data[i].id + '" target="_blank" class="item-img-box">' +
                        '         <img src="' + res.data[i].videologo + '"></img>' +
                        '         <div class="num-info">' +
                        '             <span><i class="fa fa-eye" aria-hidden="true"></i> ' + res.data[i].viewnum + '</span >' +
                        '         </div >' +
                        '     </a>' +
                        '     <div class="item-info">' +
                        '         <a href="/play?id=' + res.data[i].id + '" target="_blank" class="item-title">' + res.data[i].videoname + '</a>' +
                        '         <span class="item-score">' + res.data[i].averscore + '</span>' +
                        '         <p class="item-tags">' + res.data[i].typename + '</p>' +
                        '     </div>' +
                        '</li>'
                    )
                }
            } else {
                console.log("get video information of " + key + " failed, but the reason maybe this information showing: ", res.msg)
            }
        }
    });
}