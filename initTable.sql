create schema skyline;
use
skyline;
create table user
(
    userid         int auto_increment,
    username       varchar(16)   not null,
    passwd         varchar(128)  not null,
    avatar         varchar(64)   not null,
    background     varchar(64)   not null,
    signature      varchar(64) null,
    follow_count   int default 0 not null,
    follower_count int default 0 not null,
    follow_db      varchar(128)  not null,
    follower_db    varchar(128)  not null,
    favorite_db    varchar(64)   not null,
    primary key (userid)
);

#视频数据库表
CREATE TABLE `video`
(
    `id`             int                                                           NOT NULL AUTO_INCREMENT COMMENT '视频唯一标识',
    `userid`         int                                                           NOT NULL COMMENT '视频作者id',
    `play_url`       varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '视频播放地址',
    `cover_url`      varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '视频封面地址',
    `favorite_count` int                                                           DEFAULT '0' COMMENT '视频的点赞总数',
    `comment_count`  int                                                           DEFAULT '0' COMMENT '视频的评论总数',
    `title`          varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '视频标题',
    `create_time`    datetime                                                      NOT NULL COMMENT '上传时间',
    `comment_db`     varchar(128)                                                  NOT NULL COMMENT '评论存储文件名',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=33 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

#用户详情表
CREATE TABLE `userdetail`
(
    `userid`           int       NOT NULL COMMENT '用户id',
    `name`             varchar(128) NOT NULL COMMENT '用户名称',
    `avatar`           varchar(255) DEFAULT NULL COMMENT '用户头像',
    `background_image` varchar(255) DEFAULT NULL COMMENT '用户个人页顶部大图',
    `signature`        varchar(255) DEFAULT NULL COMMENT '个人简介',
    `favorite_count`   int       DEFAULT NULL COMMENT '喜欢数',
    `follow_count`     int       DEFAULT NULL COMMENT '关注总数',
    `follower_count`   int       DEFAULT NULL COMMENT '粉丝总数',
    `total_favorited`  varchar(255) DEFAULT NULL COMMENT '获赞数量',
    `work_count`       int       DEFAULT NULL COMMENT '作品数',
    PRIMARY KEY (`userid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;