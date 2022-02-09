use videowebsite;
DROP TABLE IF EXISTS `vw_user`;
CREATE TABLE `vw_user` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `username` varchar(100) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(100) NOT NULL DEFAULT '' COMMENT '密码',
  `nickname` varchar(100) NOT NULL DEFAULT 'stranger' COMMENT '昵称',
  `sex` varchar(20) DEFAULT '保密' COMMENT '性别',
  `email` varchar(100) DEFAULT NULL COMMENT '电子邮箱',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '用户身份(1:普通用户,2:管理员)',
  `state` int(11) unsigned DEFAULT NULL COMMENT '状态信息',
  `remark` varchar(200) DEFAULT NULL COMMENT '备注',
  `create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `delete_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `username` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT COMMENT='系统用户表';

INSERT INTO vw_user(username, password, nickname, status) VALUES('admin', 'admin', 'admin', '2');