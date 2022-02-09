use videowebsite;
DROP TABLE IF EXISTS `vw_system_menu`;
CREATE TABLE `vw_system_menu` (
  `id`      int(11) unsigned  NOT NULL                 COMMENT 'ID',
  `pid`     int(11) unsigned  NOT NULL DEFAULT '0'     COMMENT '父ID',
  `title`   varchar(100)      NOT NULL DEFAULT ''      COMMENT '名称',
  `icon`    varchar(100)      NOT NULL DEFAULT ''      COMMENT '菜单图标',
  `href`    varchar(100)      NOT NULL DEFAULT ''      COMMENT '链接',
  `target`  varchar(20)       NOT NULL DEFAULT '_self' COMMENT '链接打开方式',
  `sort`    int(11)           DEFAULT '0'              COMMENT '菜单排序',
  `state`   tinyint(1) unsigned NOT NULL DEFAULT '1'   COMMENT '状态(0:禁用,1:启用)',
  `remark`  varchar(255)      DEFAULT NULL             COMMENT '备注信息',
  `create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `delete_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `title` (`title`),
  KEY `href` (`href`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT COMMENT='系统菜单表';

INSERT INTO vw_system_menu(id, pid, title, icon, href) VALUES(1, "0", "常规管理", "fa fa-address-book", "");
INSERT INTO vw_system_menu(id, pid, title, icon, href) VALUES(110, "1", "主页", "fa fa-home", "welcome.html");
INSERT INTO vw_system_menu(id, pid, title, icon, href) VALUES(120, "1", "系统设置", "fa fa-gears", "setting.html");
INSERT INTO vw_system_menu(id, pid, title, icon, href) VALUES(130, "1", "用户管理", "fa fa-users", "");
INSERT INTO vw_system_menu(id, pid, title, icon, href) VALUES(140, "1", "视频管理", "fa fa-file-video-o", "");
INSERT INTO vw_system_menu(id, pid, title, icon, href) VALUES(131, "130", "用户列表", "fa fa-user", "userlist.html");
INSERT INTO vw_system_menu(id, pid, title, icon, href) VALUES(141, "140", "视频列表", "fa fa-list-alt", "videolist.html");
-- INSERT INTO vw_system_menu(id, pid, title, icon, href) VALUES("1", "弹出层", "page/layer.html", "fa fa-shield");
-- INSERT INTO vw_system_menu(id, pid, title, icon, href) VALUES();
-- INSERT INTO vw_system_menu(id, pid, title, icon, href) VALUES();
