var currentID;
$(function () {
    var currPageName = GetCurrentPageName();
    if (currPageName == "s") {
        $(".filter-box").hide();
        $(".sort-box").hide();
        $(".module-box:eq(0)").prepend('<div class="search-title">搜索结果<div class="search-error-info"></div></div>')
        loadSearchPage()
    } else {
        $("#" + currPageName).addClass("layui-this");
        reLoadVideo();
    }
    // 设置filter-item选择
    var $filterItems = $(".filter-row dd");
    $filterItems.click(function () {
        $filterItems.removeClass("on");
        $(this).addClass("on");
        reLoadVideo();
    });

    // 设置sort-item选择
    $(":radio[name='sort']").click(function() {
        reLoadVideo();
    });

    $.ajax({
        url: "../common/getlogininfo.json",
        type: "GET",
        success: function (res) {
            if (res.code == 0) {
                $(".user-box").eq(0).html(LoadUserBox(res.data.logo, res.data.name, res.data.id))
                LoadUserFunc()
            } else {
                $(".user-box").eq(0).html(LoadLoginBox());
            }
        },
    });

    setNavJump();
});

/******* 
 * 进行分页
 * @param  totalCount 总数据数
 * @param  page 当前页数
 */
var runPageSap = function (totalCount, page) {
    layui.use(['laypage'], function () {
        var laypage = layui.laypage;
        laypage.render({
            elem: 'library-page'
            , count: totalCount //数据总数，从服务端得到
            , limit: 42
            , curr: page
            , theme: "#00cec9"
            , jump: function (obj, first) {
                if (!first) {
                    if (GetCurrentPageName() == "s") {
                        loadSearchPage(obj.curr)
                    } else {
                        reLoadVideo(obj.curr, obj.limit)
                    }
                }
            }
        });
    });
}

// 先获取筛选信息，然后重新加载信息
var reLoadVideo = function(page, limit) {
    var $filterItems = $(".filter-row dd").filter(".on"),
        xmlBody = new Object(),
        $sortItem = $(":radio[name='sort']").filter(":checked"),
        vfilter = "";

    // 获取filter value
    for (var i = 0, len = $filterItems.length; i < len; i++) {
        vfilter += "+" + $filterItems.eq(i).attr("data-a");
    }
    if (vfilter[0] == '+' && vfilter.length > 0) {
        vfilter = vfilter.substring(1, vfilter.length);
    }

    // 设置page和limit
    if (limit == undefined) {
        limit = 42;
    }
    if (page == undefined) {
        page = 1;
    }

    xmlBody["filter"] = vfilter;
    xmlBody["sort"] = $sortItem.attr("value");
    xmlBody["page"] = page;
    xmlBody["limit"] = limit;
    xmlBody["classification"] = GetCurrentPageName(window.location.href);

    $.ajax({
        url: "reloadvideo.json",
        type: "GET",
        data: xmlBody,
        dataType: "json",
        success: function(res) {
            if (res.code == 0) {
                runPageSap(res.count, page)
                reLoadVideoPage(res.data)
            }
        },
        error: function() {
            console.log("reload video failed, and the reason is not found.")
        }
    });    
};

/******* 
 * 重新加载#video-module-box
 * @param  data 对象数组，对象为存储展示视频的信息
 */
var reLoadVideoPage = function(data) {
    var $container = $("#video-module-box");
    $container.empty();
    if (data.length == 0) {
        $(".search-error-info:eq(0)").html('当前结果为空')
    }
    for (var i = 0, len = data.length; i < len; i++) {
        $container.append(
            ' <li class="module-item">' +
            '     <a href="/play?id=' + data[i].id + '" target="_blank" class="item-img-box">' +
            '         <img src="' + data[i].videologo + '"></img>' +
            '     <div class="num-info">' +
            '         <span><i class="fa fa-eye" aria-hidden="true"></i> '+data[i].viewnum+'</span >' +
            '     </div >' +
            '     </a>' +
            '     <div class="item-info">' +
            '         <a href="/play?id=' + data[i].id + '" target="_blank" class="item-title">' + data[i].videoname + '</a>' +
            '         <span class="item-score">' + data[i].averscore + '</span>' +
            '         <p class="item-tags">' + data[i].typename + '</p>' +
            '     </div>' +
            ' </li>'
        )
    }
}

/******* 
 * 进行搜索内容的加载
 * @param page  当前页面
 */
var loadSearchPage = function(page) {
    if (page == undefined) {
        page = 1
    }
    var currHref = window.location.href
        , url = "/searcher" + currHref.substring(currHref.indexOf("?"), currHref.length);
    $.ajax({
        url: url
        , type: "GET"
        , dataType: "json"
        , data: {
            page: page
        }
        , success: function (res) {
            if (res.code == 0) {
                reLoadVideoPage(res.data);
                runPageSap(res.count, page);
            } else {
                $(".search-error-info:eq(0)").html('获取结果失败')
            }
        }
    })
}