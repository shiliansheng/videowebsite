<div class="layui-header">
    <div class="website-logo">VIDEO WEBSITE</div>
    <ul class="layui-nav" lay-filter="nav-items">
        <li class="layui-nav-item" id="home"><a href="/">首页</a></li>
        <li class="layui-nav-item" id="movie"><a href="javascript:;" data-a="movie">电影</a></li>
        <li class="layui-nav-item" id="episode"><a href="javascript:;" data-a="episode">剧集</a></li>
        <li class="layui-nav-item" id="cartoon"><a href="javascript:;" data-a="cartoon">动漫</a></li>
        <li class="layui-nav-item" id="others"><a href="javascript:;" data-a="others">其他</a></li>
        <li class="layui-nav-item" id="library"><a href="javascript:;" data-a="library">片库</a></li>
        <div class="nav-right-box layui-layout-right">
            <li class="search-box ">
                <form action="/s/">
                    <input type="text" name="search" placeholder="请输入搜索内容" autocomplete="off" class="layui-input">
                    <button class="layui-btn search-btn" lay-submit lay-filter="search-btn"><i
                            class="fa fa-search"></i></button>
                </form>
            </li>
            <li class="user-box layui-nav-item">
            </li>
        </div>
    </ul>
</div>