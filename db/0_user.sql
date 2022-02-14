use videowebsite;
DROP TABLE IF EXISTS `vw_user`;
CREATE TABLE `vw_user` (
  `id`            int(11) unsigned  NOT NULL AUTO_INCREMENT       COMMENT 'ID',
  `username`      varchar(100)      NOT NULL UNIQUE DEFAULT ''    COMMENT '用户名',
  `password`      varchar(100)      NOT NULL DEFAULT ''           COMMENT '密码',
  `nickname`      varchar(100)      NOT NULL DEFAULT 'stranger'   COMMENT '昵称',
  `logoimage`     varchar(200)      DEFAULT NULL                  COMMENT '用户头像',
  `sex`           varchar(20)       DEFAULT '保密'                COMMENT '性别',
  `email`         varchar(100)      DEFAULT NULL                  COMMENT '电子邮箱',
  `birthday`      varchar(20)       DEFAULT NULL                  COMMENT '用户生日',
  `introduction`  varchar(400)      DEFAULT NULL                  COMMENT '简介',
  `status`        varchar(20)       NOT NULL DEFAULT '普通用户'    COMMENT '用户身份(普通用户,管理员)',
  `state`         int(11) unsigned  DEFAULT NULL                  COMMENT '状态信息',
  `remark`        varchar(200)      DEFAULT NULL                  COMMENT '备注',
  `create_at`     timestamp NULL    DEFAULT CURRENT_TIMESTAMP     COMMENT '创建时间',
  `update_at`     timestamp NULL    DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `delete_at`     timestamp NULL    DEFAULT NULL                  COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT COMMENT='系统用户表';

INSERT INTO vw_user(username, password, nickname, status, state) VALUES('admin', 'admin', 'superadmin', '超级管理员', '0');
INSERT INTO vw_user(username, password, nickname, status, state) VALUES('admin01', '123456', '管理员01', '管理员', '0');
INSERT INTO vw_user(username, password, nickname, status, state) VALUES('admin02', '123456', '管理员02', '管理员', '0');
INSERT INTO vw_user(username, password, nickname, status, state) VALUES('admin03', '123456', '管理员03', '管理员', '0');
INSERT INTO vw_user(username, password, nickname, status, state) VALUES('admin04', '123456', '管理员04', '管理员', '0');
INSERT INTO vw_user(username, password, nickname, status, state) VALUES('admin05', '123456', '管理员05', '管理员', '0');
INSERT INTO vw_user(username, password, nickname, status, state) VALUES('admin06', '123456', '管理员06', '管理员', '0');
INSERT INTO vw_user(username, password, nickname, status, state) VALUES('admin07', '123456', '管理员07', '管理员', '0');
INSERT INTO vw_user(username, password, nickname, status, state) VALUES('admin08', '123456', '管理员08', '管理员', '0');
INSERT INTO vw_user(username, password, nickname, status, state) VALUES('admin09', '123456', '管理员09', '管理员', '0');
INSERT INTO vw_user(username, password, nickname, status, state) VALUES('test10', '123456', '普通用户01', '普通用户', '0');
INSERT INTO vw_user(username, password, nickname, status, state) VALUES('test11', '123456', '普通用户02', '普通用户', '0');
INSERT INTO vw_user(username, password, nickname, status, state) VALUES('test12', '123456', '普通用户03', '普通用户', '0');
INSERT INTO vw_user(username, password, nickname, status, state) VALUES('test13', '123456', '普通用户04', '普通用户', '0');
INSERT INTO vw_user(username, password, nickname, status, state) VALUES('test14', '123456', '普通用户05', '普通用户', '0');
INSERT INTO vw_user(username, password, nickname, status, state) VALUES('test15', '123456', '普通用户06', '普通用户', '0');
INSERT INTO vw_user(username, password, nickname, status, state) VALUES('test16', '123456', '普通用户07', '普通用户', '0');
INSERT INTO vw_user(username, password, nickname, status, state) VALUES('test17', '123456', '普通用户08', '普通用户', '0');
INSERT INTO vw_user(username, password, nickname, status, state) VALUES('test18', '123456', '普通用户09', '普通用户', '0');
INSERT INTO vw_user(username, password, nickname, status, state) VALUES('test19', '123456', '普通用户10', '普通用户', '0');
INSERT INTO vw_user(username, password, nickname, status, state) VALUES('test20', '123456', '普通用户11', '普通用户', '0');
INSERT INTO vw_user(username, password, nickname, status, state) VALUES('test21', '123456', '普通用户12', '普通用户', '0');
INSERT INTO vw_user(username, password, nickname, status, state) VALUES('test22', '123456', '普通用户13', '普通用户', '0');
INSERT INTO vw_user(username, password, nickname, status, state) VALUES('test23', '123456', '普通用户14', '普通用户', '0');
INSERT INTO vw_user(username, password, nickname, status, state) VALUES('test24', '123456', '普通用户15', '普通用户', '0');
INSERT INTO vw_user(username, password, nickname, status, state) VALUES('test25', '123456', '普通用户16', '普通用户', '0');
INSERT INTO vw_user(username, password, nickname, status, state) VALUES('test26', '123456', '普通用户17', '普通用户', '0');
INSERT INTO vw_user(username, password, nickname, status, state) VALUES('test27', '123456', '普通用户18', '普通用户', '0');
INSERT INTO vw_user(username, password, nickname, status, state) VALUES('test28', '123456', '普通用户19', '普通用户', '0');
INSERT INTO vw_user(username, password, nickname, status, state) VALUES('test29', '123456', '普通用户20', '普通用户', '0');