layui.use(["form", "tree"], function () {
	var form = layui.form,
		layer = layui.layer,
		$ = layui.$;
	$.ajax({
		type: "post",
		url: "videotypelist.json?action=gettree",
		success: function (res) {
			layui.tree.render({
				elem: '#treeUl',
				data: res,
				click: function (node) { //点击节点回调
					var othis = $($(this)[0].elem).parents(".layui-form-select");
					othis.removeClass("layui-form-selected").find(".layui-select-title span").html(node.data.title).end().find("input:hidden[name='pid']").val(node.data.id);
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
	//监听提交
	form.on("submit(saveBtn)", function (data) {
		$.ajax({
			url: "videotypeadd.json?action=add",
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
							title: "添加用户失败",
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
}).use(['form', 'avatar'], function () {
	var $ = layui.jquery
		, form = layui.form
		, avatar = layui.avatar
		, layer = layui.layer;
	avatar.render({
		elem: "#vtypelogo"
		, success: function (base64, size) {
			var data = new Object()
			data.image = base64
			$.ajax({
				type: "post"
				, url: "../common/uploader?type=type-image"
				, data: data
				, success: function (res) {
					if (res.code == 0) {
						layer.msg('上传成功')
						document.getElementById("vtypelogoPath").value = res.data
						document.getElementById('logoimg').setAttribute("src", base64)
					} else {
						layer.msg(res.msg)
					}
				}
			});
		}
	});
});
