layui.use(["form", "tree", "upload"], function () {
	var form = layui.form,
		layer = layui.layer,
		tree = layui.tree,
		upload = layui.upload,
		$ = layui.jquery;
	form.val ("videoinfo", {
		"copyright" : VideoCopyright,
		"recommand": VideoRecommand,
	});
	upload.render({
		elem: "#videologobtn",
		url: "../common/uploadfile?type=image-video",
		accept: "file",
		accept: "images",
		acceptMine: "image/*",
		exts: "jpg|png|gif|bmp|jpeg",
		size: 1024*3,
		field: "upfile",
		number: 1,
		choose: function(obj) {
		}, progress: function (n, elem, res, index) {
			var percent = n + '%' //获取进度百分比
			// element.progress('demo', percent); //可配合 layui 进度条元素使用
			console.log(percent); //得到当前上传文件的索引，多文件上传时的进度条控制，如：
			//element.progress('demo-' + index, n + '%'); //进度条
		},
		done: function(res, index, upload) {
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
		error: function() {
			layer.msg(
				"上传图片出错",
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
			url: "videoedit.json?action=update",
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