use `videowebsite`;
DROP TABLE IF EXISTS `vw_history`;
CREATE TABLE `vw_history` (
    id int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '观看记录编号',
    userid int(11) unsigned NOT NULL COMMENT '用户编号',
    videoid int(11) unsigned NOT NULL COMMENT '视频编号',
    `State` int NOT NULL COMMENT '状态(0:存在，1:不存在)',
    Pubtime timestamp NOT NULL COMMENT '发布时间',
    PRIMARY KEY(id),
    FOREIGN kEY(`userid`) REFERENCES `vw_user`(`id`),
    FOREIGN kEY(`videoid`) REFERENCES `vw_video`(`id`)
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT COMMENT="视频观看记录表";