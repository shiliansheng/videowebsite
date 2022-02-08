use videowebsite;

CREATE TABLE IF NOT EXISTS `vw_system_menu` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `pid` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '父ID',
  `title` varchar(100) NOT NULL DEFAULT '' COMMENT '名称',
  `icon` varchar(100) NOT NULL DEFAULT '' COMMENT '菜单图标',
  `href` varchar(100) NOT NULL DEFAULT '' COMMENT '链接',
  `target` varchar(20) NOT NULL DEFAULT '_self' COMMENT '链接打开方式',
  `sort` int(11) DEFAULT '0' COMMENT '菜单排序',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '状态(0:禁用,1:启用)',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注信息',
  `create_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `update_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `delete_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `title` (`title`),
  KEY `href` (`href`)
) ENGINE=InnoDB AUTO_INCREMENT=250 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT COMMENT='系统菜单表';


INSERT INTO vw_system_menu(pid, title, icon, href, target) VALUES("0", "常规管理", "fa fa-address-book", "", "_self");
INSERT INTO vw_system_menu(pid, title, icon, href, target) VALUES("1", "主页", "fa fa-home", "page/welcome.html", "_self");
INSERT INTO vw_system_menu(pid, title, icon, href, target) VALUES("1", "系统设置", "fa fa-gears", "page/setting.html", "_self");
INSERT INTO vw_system_menu(pid, title, icon, href, target) VALUES("1", "数据展示", "fa fa-file-text", "page/table.html", "_self");
INSERT INTO vw_system_menu(pid, title, icon, href, target) VALUES("2", "普通表单", "fa fa-list-alt", "page/form.html", "_self");
-- INSERT INTO vw_system_menu(pid, title, icon, href, target) VALUES("1", "弹出层", "page/layer.html", "fa fa-shield", "_self");
-- INSERT INTO vw_system_menu(pid, title, icon, href, target) VALUES();
-- INSERT INTO vw_system_menu(pid, title, icon, href, target) VALUES();
