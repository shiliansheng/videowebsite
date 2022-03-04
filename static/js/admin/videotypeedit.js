layui.use(["form", "tree"], function () {
	var form = layui.form,
		layer = layui.layer,
		tree = layui.tree,
		$ = layui.jquery;
	// 校验两次密码是否一致
	form.verify({
	});
	//监听提交
	form.on("submit(saveBtn)", function (data) {
		$.ajax({
			url: "videotypeedit.html?action=update",
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
});
layui.config({
	base: '/static/lib/cropper/' //layui自定义layui组件目录
}).use(['form', 'croppers'], function () {
	var $ = layui.jquery
		, croppers = layui.croppers;
	croppers.render({
		elem: '#vtypelogo'
		, saveW: 150     //保存宽度
		, saveH: 150
		, mark: 1 / 1    //选取比例
		, area: '900px'  //弹窗宽度
		, url: "../common/uploadfile?type=image-type"  //图片上传接口返回和（layui 的upload 模块）返回的JOSN一样
		, done: function (url) { //上传完毕回调
			$("#vtypelogoPath").val(url);
			$("#logoimg").attr('src', url);
		}
	});
});