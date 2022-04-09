use `videowebsite`;
DROP TABLE IF EXISTS `vw_score`;
CREATE TABLE `vw_score` (
    id int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '评分编号',
    userid int(11) unsigned NOT NULL COMMENT '用户编号',
    videoid int(11) unsigned NOT NULL COMMENT '视频编号',
    `Value` int NOT NULL COMMENT '评分',
    Pubtime timestamp NOT NULL COMMENT '发布时间',
    PRIMARY KEY(id),
    FOREIGN kEY(`userid`) REFERENCES `vw_user`(`id`),
    FOREIGN kEY(`videoid`) REFERENCES `vw_video`(`id`)
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT COMMENT="视频评分表";