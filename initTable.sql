create schema skyline;
use
skyline;
CREATE TABLE user
(
    userid   int          NOT NULL AUTO_INCREMENT COMMENT '用户唯一标识',
    username varchar(16)  NOT NULL COMMENT '用户名',
    password varchar(128) NOT NULL COMMENT '密码',
    state    int DEFAULT '1' COMMENT '用户状态（1-启用   0-封禁）',
    PRIMARY KEY (userid)
);

#视频数据库表
CREATE TABLE video
(
    id             int          NOT NULL AUTO_INCREMENT COMMENT '视频唯一标识',
    userid         int          NOT NULL COMMENT '视频作者id',
    play_url       varchar(128) NOT NULL COMMENT '视频播放地址',
    cover_url      varchar(128) DEFAULT 'https://tos.eyunnet.com/video_covers%2Fdefault.png' COMMENT '视频封面地址',
    favorite_count int          DEFAULT '0' COMMENT '视频的点赞总数',
    comment_count  int          DEFAULT '0' COMMENT '视频的评论总数',
    title          varchar(128) NOT NULL COMMENT '视频标题',
    create_time    datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT COMMENT '上传时间',
    comment_db     varchar(128) NOT NULL COMMENT '评论存储文件名',
    PRIMARY KEY (id)
);

#用户详情表
CREATE TABLE userdetail
(
    userid           int          NOT NULL AUTO_INCREMENT COMMENT '用户id',
    name             varchar(64)  NOT NULL COMMENT '用户名称',
    avatar           varchar(255) DEFAULT 'https://tos.eyunnet.com/avatars%2Fdefault.png' COMMENT '用户头像',
    background_image varchar(255) DEFAULT 'https://tos.eyunnet.com/backgrounds%2Fdefault.png' COMMENT '用户个人页顶部大图',
    signature        varchar(255) DEFAULT '这个人很懒没有写简介QAQ' COMMENT '个人简介',
    favorite_count   int          DEFAULT '0' COMMENT '喜欢数',
    follow_count     int          DEFAULT '0' COMMENT '关注总数',
    follower_count   int          DEFAULT '0' COMMENT '粉丝总数',
    total_favorited  int          DEFAULT '0' COMMENT '获赞数量',
    work_count       int          DEFAULT '0' COMMENT '作品数',
    follow_db        varchar(128) NOT NULL,
    follower_db      varchar(128) NOT NULL,
    favorite_db      varchar(128) NOT NULL,
    PRIMARY KEY (userid)
);
