layui.use(["form", "miniTab", 'laydate', 'upload'], function () {
	var form = layui.form,
		layer = layui.layer,
		miniTab = layui.miniTab,
		laydate = layui.laydate,
		upload = layui.upload,
		$ = layui.$;
	// laydate实现
	laydate.render ({
		elem: '#dateselect',
		//type: 'datetime',
		//range: true
	});
	// 图片上传初始化
	upload.render({
		elem: "#logoimage",
		accpet: "images", // 指定允许上传时校验的文件类型
		acceptMime: 'image/jpg, image/png, image/jpeg', // 规定打开文件选择框时，筛选出的文件类型
		exts: "jpg|png|jpeg", // 允许上传的文件后缀
		number: 1, // 文件上传限制
		auto: false,
		bindAction: "#saveBtn",
		done: function (res, index, upload) {
			//获取当前触发上传的元素，一般用于 elem 绑定 class 的情况，注意：此乃 layui 2.1.0 新增
			var item = this.item;
		},
		choose: function(obj){
			obj.preview(function(index, file, result){
				document.getElementById("userlogo").setAttribute("src", result);
				var logo = document.getElementById("userlogo");
				logo.style.zoom = (150 / logo.width) * 1.0;
			});
		},
	});
	//监听提交
	form.on("submit(saveBtn)", function (data) {
		layer.confirm('确认修改?', function(){
			console.log(data.field);
			$.ajax({
				type : "post",
				url : "user_setting.json?action=changeSetting",
				data : data.field,
				success : function(res){
					if (res.code == 0) {
						layer.msg('修改信息成功');
						miniTab.deleteCurrentByIframe();
					} else {
						layer.msg(res.msg);
					}
				},
			});
		});
	});

	form.val("baseinfo", userInfo);

	form.verify({
		verifyEmail: function (value) {
			var ePattern =
				/^([A-Za-z0-9_\-\.])+\@([A-Za-z0-9_\-\.])+\.([A-Za-z]{2,4})$/;
			if (value != "" && !ePattern.test(value)) {
				return "邮箱格式不正确";
			}
		},
	});
});
