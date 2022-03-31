use videowebsite;
DROP TABLE IF EXISTS `vw_video`;
CREATE TABLE `vw_video` (
    id              int(11) unsigned NOT NULL   AUTO_INCREMENT      COMMENT '视频编号',
    videoname       varchar(256) NOT NULL       DEFAULT 'untitled'  COMMENT '视频名称',
    typename        varchar(64) NOT NULL                            COMMENT '视频类型名称',
    classifiction    varchar(32) NOT NULL                            COMMENT '视频分类',
    introduction    varchar(512)                                    COMMENT '视频介绍',
    videologo       varchar(256)                                    COMMENT '视频图片地址',
    keywords        varchar(256)                                    COMMENT '视频关键字',
    videoresource   varchar(256)                NOT NULL            COMMENT '视频资源地址',
    copyright       varchar(8)                  DEFAULT '转载'      COMMENT '版权所有(原创,转载)',
    username        varchar(128)                NOT NULL            COMMENT '发布者用户名',
    pubtime         timestamp                   DEFAULT CURRENT_TIMESTAMP COMMENT '发布时间',
    viewnum         bigint(32)                  DEFAULT '0'         COMMENT '视频观看次数',
    scorenum        bigint(32)                  DEFAULT '0'         COMMENT '视频打分人数',
    remarknum       bigint(32)                  DEFAULT '0'         COMMENT '视频评论次数',
    averscore       decimal(32)                 DEFAULT '0'         COMMENT '用户评分平均分',
    totalscore      bigint(64)                  DEFAULT '0'         COMMENT '用户评分总分',
    passed          varchar(8)                  DEFAULT '待审核'    COMMENT '审核状态(待审核;通过审核)',
    recommand       tinyint(1)                  DEFAULT '0'         COMMENT '视频推荐(0:不推荐,1:推荐)',
    PRIMARY KEY(id)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT COMMENT='视频表';

INSERT INTO `vw_video`(videoname, classifiction, typename, videoresource, username) VALUES('test video', 'others', '剧情', 'static/store/video/testvideo.mp4', 'admin');
INSERT INTO `vw_video`(videoname, classifiction, typename, videoresource, username) VALUES('test video', 'others', '剧情', 'static/store/video/testvideo.mp4', 'admin');
INSERT INTO `vw_video`(videoname, classifiction, typename, videoresource, username) VALUES('test video', 'others', '剧情', 'static/store/video/testvideo.mp4', 'admin');
INSERT INTO `vw_video`(videoname, classifiction, typename, videoresource, username) VALUES('test video', 'others', '剧情', 'static/store/video/testvideo.mp4', 'admin');
INSERT INTO `vw_video`(videoname, classifiction, typename, videoresource, username) VALUES('test video', 'others', '剧情', 'static/store/video/testvideo.mp4', 'admin');
INSERT INTO `vw_video`(videoname, classifiction, typename, videoresource, username) VALUES('test video', 'others', '剧情', 'static/store/video/testvideo.mp4', 'admin');
INSERT INTO `vw_video`(videoname, classifiction, typename, videoresource, username) VALUES('test video', 'others', '剧情', 'static/store/video/testvideo.mp4', 'admin');
INSERT INTO `vw_video`(videoname, classifiction, typename, videoresource, username) VALUES('test video', 'others', '剧情', 'static/store/video/testvideo.mp4', 'admin');
INSERT INTO `vw_video`(videoname, classifiction, typename, videoresource, username) VALUES('test video', 'others', '剧情', 'static/store/video/testvideo.mp4', 'admin');
INSERT INTO `vw_video`(videoname, classifiction, typename, videoresource, username) VALUES('test video', 'others', '剧情', 'static/store/video/testvideo.mp4', 'admin');
INSERT INTO `vw_video`(videoname, classifiction, typename, videoresource, username) VALUES('test video', 'others', '剧情', 'static/store/video/testvideo.mp4', 'admin');
INSERT INTO `vw_video`(videoname, classifiction, typename, videoresource, username) VALUES('test video', 'others', '剧情', 'static/store/video/testvideo.mp4', 'admin');
INSERT INTO `vw_video`(videoname, classifiction, typename, videoresource, username) VALUES('test video', 'others', '剧情', 'static/store/video/testvideo.mp4', 'admin');
INSERT INTO `vw_video`(videoname, classifiction, typename, videoresource, username) VALUES('test video', 'others', '剧情', 'static/store/video/testvideo.mp4', 'admin');
INSERT INTO `vw_video`(videoname, classifiction, typename, videoresource, username) VALUES('test video', 'others', '剧情', 'static/store/video/testvideo.mp4', 'admin');