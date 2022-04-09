use videowebsite;
DROP TABLE IF EXISTS `vw_videotype`;
CREATE TABLE `vw_videotype` (
    `id`            int(11) unsigned    NOT NULL AUTO_INCREMENT COMMENT '视频类型ID',
    `pid`           int(11) unsigned    NOT NULL DEFAULT '0'    COMMENT '视频分类父ID',
    `typename`      varchar(32)         NOT NULL UNIQUE         COMMENT '视频类型名称',
    `discription`   varchar(512)                 DEFAULT NULL   COMMENT '视频类型描述',
    `addid`         int(11) unsigned    NOT NULL                COMMENT '添加人ID',
    `createat`      timestamp NULL               DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `vtypelogo`     varchar(256)                 DEFAULT NULL   COMMENT '视频类型LOGO',
    `sequence`      int(11)                      DEFAULT '0'    COMMENT '显示顺序',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT COMMENT='视频类型表';

INSERT INTO `vw_videotype`(typename, addid) VALUES
('剧情', 1),
('喜剧', 1),
('动作', 1),
('爱情', 1),
('惊悚', 1),
('犯罪', 1),
('悬疑', 1),
('战争', 1),
('科幻', 1),
('动画', 1),
('恐怖', 1),
('家庭', 1),
('传记', 1),
('冒险', 1),
('奇幻', 1),
('武侠', 1),
('历史', 1),
('运动', 1),
('音乐', 1),
('记录', 1),
('伦理', 1);