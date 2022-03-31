layui.use(["form", "layer", "upload", "element"], function () {
    var form = layui.form,
        layer = layui.layer,
        upload = layui.upload,
        element = layui.element,
        $ = layui.jquery;

    // video upload
    upload.render({
        elem: "#video-upload-btn",
        url: "../common/uploader?type=video",
        accept: "video",
        acceptMine: "video/*",
        // size:,
        field: "file",
        number: 1,
        choose: function (obj) {
            document.getElementById("upload-state").innerText = "等待上传";
        },
        before: function (obj) {
            obj.preview(function (index, file, result) {
                $("input[name=videoname]").val(file.name.substring(0, file.name.lastIndexOf('.')));
                var video = document.createElement("video"),
                    canvas = document.createElement("canvas"),
                    width = 1440,
                    height = 768,
                    output = document.getElementById("img-preview-box"),
                    frames = [1, 20, 100, 200, 500, 1000];
                output.innerHTML = "";
                video.width = width;
                video.height = height;
                video.setAttribute("crossOrigin", "anonymous");
                video.currentTime = frames[1];

                canvas.width = video.width;
                canvas.height = video.height;

                video.addEventListener("loadeddata", function() {
                    var dx = (video.width - video.videoWidth) / 2,
                        dy = (video.height - video.videoHeight)  / 2;
                    canvas.getContext('2d').drawImage(video, dx, dy, video.videoWidth, video.videoHeight);
                    var box = document.createElement("div"),
                        img = document.createElement("img");
                    box.id = "img-preview-show";
                    img.src = URL.createObjectURL(dataURLtoBlob(canvas.toDataURL('image/jpg')));
                    box.appendChild(img)
                    output.appendChild(box)
                });
                video.setAttribute("src", URL.createObjectURL(file));
            });
        },
        progress: function (n, elem, res, index) {
            var percent = n + "%";
            document.getElementById("upload-state").innerText = "正在上传";
            document.getElementById("upload-progress").innerText = percent;
            element.progress('video-upload-progress', percent);
        },
        done: function (res, index, upload) {
            if (res.code == 0) {
                layer.msg('上传成功');
                document.getElementById("upload-state").innerText = "上传成功"
                document.getElementById("videoresource").value = res.data.src;
            } else {
                document.getElementById("upload-state").innerText = "上传视频出错"
                layer.open({
                    title: "错误",
                    content: res.msg,
                });
            }
        },
        error: function () {
            layer.msg(
                "上传视频出错",
                {
                    icon: 2,
                    time: 5000,
                }
            );
            document.getElementById("upload-state").innerText = "上传视频出错"
        }
    });

    //监听提交
    form.on("submit(saveBtn)", function (data) {
        console.log(data.field)
        $.ajax({
            url: "videoadd.json",
            type: "post",
            data: data.field,
            success: function (res) {
                if (res.code == 0) {
                    layer.close(layer.index);
                    window.parent.location.reload();
                } else {
                    layer.alert(
                        res.msg,
                        {
                            title: "信息",
                            icon: 2,
                            time: 5000,
                        }
                    );
                }
            },
        });
        return false;
    });
});

layui.config({
    base: '/static/lib/cropper/' //layui自定义layui组件目录
}).use(['form', 'croppers'], function () {
    var $ = layui.jquery
        , croppers = layui.croppers;
    croppers.render({
        elem: '#img-button'
        , saveW: 193     //保存宽度
        , saveH: 289
        , mark: 193 / 289    //选取比例
        , area: '900px'  //弹窗宽度
        , url: "../common/uploader?type=image-video"  //图片上传接口返回和（layui 的upload 模块）返回的JOSN一样
        , done: function (url) { //上传完毕回调
            $("#img-input").val(url);
            $("#img-image").attr('src', url);
        }
    });
});

var dataURLtoBlob = function(dataurl) {
    var arr = dataurl.split(',');
    //注意base64的最后面中括号和引号是不转译的   
    var _arr = arr[1].substring(0, arr[1].length - 2);
    var mime = arr[0].match(/:(.*?);/)[1],
        bstr = atob(_arr),
        n = bstr.length,
        u8arr = new Uint8Array(n);
    while (n--) {
        u8arr[n] = bstr.charCodeAt(n);
    }
    return new Blob([u8arr], {
        type: mime
    });
};
