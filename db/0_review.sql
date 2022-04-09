use videowebsite;
DROP TABLE IF EXISTS `vw_review`;
CREATE TABLE `vw_review` (
    id int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '评论编号',
    userid int(11) unsigned NOT NULL COMMENT '用户编号',
    videoid int(11) unsigned NOT NULL COMMENT '视频编号',
    Content text NOT NULL COMMENT '评论内容',
    `Status` varchar(8) NOT NULL COMMENT '评论状态',
    Pubtime timestamp NOT NULL COMMENT '发布时间',
    PRIMARY KEY(id),
    FOREIGN kEY(`userid`) REFERENCES `vw_user`(`id`),
    FOREIGN kEY(`videoid`) REFERENCES `vw_video`(`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT COMMENT="视频评论表";