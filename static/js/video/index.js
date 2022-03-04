layui.use(['element', 'layer', 'carousel'], function () {
    var carousel = layui.carousel
        , element = layui.element
        , layer = layui.layer
        , $ = layui.jquery;

    //轮播图
    carousel.render({
        elem: '#test1'
        , width: '730px' //设置容器宽度
        , height: '320px'
        , arrow: 'hover' //始终显示箭头
        , interval: 5000
        //,anim: 'updown' //切换动画方式
    });
    $("#home").click(function(){
        $("content-iframe").attr("src", "home.html")
    });
    $("#movie").click(function () {
        $("#content-iframe").attr("src", "movie.html");
    });
    $("#episode").click(function () {
        $("#content-iframe").attr("src", "episode.html");
    });
    $("#cartoon").click(function () {
        $("#content-iframe").attr("src", "cartoon.html");
    });
    $("#others").click(function () {
        $("#content-iframe").attr("src", "others.html");
    });
    $("#library").click(function () {
        $("#content-iframe").attr("src", "library.html");
    });
});

reHeight = function () {
    ifr = document.getElementById('content-iframe');
    ifr.height = 0;
    ifr.height = ifr.contentWindow.document.body.scrollHeight + 40;
};