layui.use(["form", "tree", "upload", "element"], function () {
	var form = layui.form,
		layer = layui.layer,
		tree = layui.tree,
		upload = layui.upload,
		element = layui.element,
		$ = layui.jquery;
	form.val("videoinfo", {
	});
	upload.render({
		elem: "#videologobtn",
		url: "../common/uploadfile?type=image-video",
		accept: "file",
		accept: "images",
		acceptMine: "image/*",
		exts: "jpg|png|gif|bmp|jpeg",
		size: 1024 * 3,
		field: "upfile",
		number: 1,
		done: function (res, index, upload) {
			if (res.code == 0) {
				layer.msg('上传成功');
				document.getElementById("videologoPath").value = res.data;
				document.getElementById('logoimg').setAttribute("src", res.data);
			} else {
				layer.open({
					title: "错误",
					content: res.msg,
				});
			}
		},
		error: function () {
			layer.msg(
				"上传图片出错",
				{
					icon: 2,
					time: 5000,
				}
			);
		}
	});
	var beginTime = new Date().getTime();
	upload.render({
		elem: "#videoupBtn",
		url: "../common/uploadfile?type=video",
		accept: "file",
		accept: "video",
		acceptMine: "video/*",
		size: 1024 * 1024 * 3,
		field: "upfile",
		number: 1,
		choose: function(obj){
			beginTime = new Date().getTime()
			obj.preview(function(index, file){
				document.getElementById('videonameShow').innerHTML = file.name;
				document.getElementById('nowProgress').innerHTML = "上传文件";
				document.getElementsByName('videoname')[0].value = file.name.substring(0, file.name.lastIndexOf('.'))
			});
			layer.open({
				type: 1,
				title: "上传视频",
				area: ['640px', '220px'], //宽高
				shade: 0.1,
				resize: false,
				content: '<div class="layui-container" style="padding-top: 20px; width: 640px;">'
					+ '	<div class="layui-row layui-col-space15">'
					+ '		<div class="layui-col layui-col-sm2" style="text-align: right;">视频名称：</div>'
					+ '		<div class="layui-col layui-col-sm9" id="videonameShow">等待获取文件名</div>'
					+ '	</div>'
					+ '	<div class="layui-row layui-col-space15">'
					+ '		<div class="layui-col layui-col-sm2" style="text-align: right;">当前进程：</div>'
					+ '		<div class="layui-col layui-col-sm9" id="nowProgress">读取文件</div>'
					+ '	</div>'
					+ '	<br/>'
					+ '	<div class="layui-row layui-col-space15">'
					+ '		<div class="layui-col layui-col-sm2" style="text-align: right;">上传进度：</div>'
					+ '		<div class="layui-progress layui-col-sm9" lay-showPercent="yes" lay-filter="showupprogress" style="padding: 0; margin-left: 7px;">'
					+ '			<div class="layui-progress-bar layui-bg-red" lay-percent="0%"><span class="layui-progress-text" >0%</span></div>'
					+ '		</div>'
					+ '	</div>'
					+ '	<div class="layui-row layui-col-space15">'
					+ '		<div class="layui-col layui-col-sm2" style="text-align: right;">上传用时：</div>'
					+ '		<div class="layui-col layui-col-sm9" id="uptimer"></div>'
					+ '	</div>'
					+ '</div>'
			});
		},
		progress: function (n, elem, res, index) {
			var percent = n + '%' //获取进度百分比
			element.progress('showupprogress', percent); //可配合 layui 进度条元素使用
			var timeDif = new Date().getTime() - beginTime
			var minute = parseInt(timeDif / (1000 * 60))
			var second = parseInt(parseInt(timeDif / 1000)%60)
			document.getElementById("uptimer").innerHTML = minute + '分' + second + '秒'
		},
		done: function (res, index, upload) {
			if (res.code == 0) {
				layer.msg('上传成功');
				document.getElementById("videoresource").value = res.data;
			} else {
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
		}
	});
	// 校验两次密码是否一致
	form.verify({
	});
	//监听提交
	form.on("submit(saveBtn)", function (data) {
		console.log(data.field)
		$.ajax({
			url: "videoadd.json?action=add",
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
	$.ajax({
		type: "post",
		url: "videotypelist.json?action=gettree",
		success: function (res) {
			layui.tree.render({
				elem: '#treeUl',
				data: res,
				showLine: false,
				click: function (node) { //点击节点回调
					var othis = $($(this)[0].elem).parents(".layui-form-select");
					othis.removeClass("layui-form-selected").find(".layui-select-title span").html(node.data.title).end().find("input:hidden[name='typename']").val(node.data.title);
				}
			});
		}
	})

	$(".treeSelect").on("click", ".layui-select-title", function (e) {
		$(".layui-form-select").not($(this).parents(".layui-form-select")).removeClass("layui-form-selected");
		$(this).parents(".treeSelect").toggleClass("layui-form-selected");
		layui.stope(e);
	}).on("click", "dl i", function (e) {
		layui.stope(e);
	});
	$(document).on("click", function (e) {
		$(".layui-form-select").removeClass("layui-form-selected");
	});
});