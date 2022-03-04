use videowebsite;
DROP TABLE IF EXISTS `vw_tag`;
CREATE TABLE `vw_tag`(
    `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `name` varchar(32) NOT NULL COMMENT '标签名',
    `sequence` int(11) NOT NULL DEFAULT '0' COMMENT '顺序',
    PRIMARY KEY(`id`)
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT COMMENT='标签表';