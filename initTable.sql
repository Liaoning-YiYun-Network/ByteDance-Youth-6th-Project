create schema skyline;
use skyline;
create table user
(
    userid         int auto_increment,
    username       varchar(16)   not null,
    passwd         varchar(128)  not null,
    avatar         varchar(64)   not null,
    background     varchar(64)   not null,
    signature      varchar(64)   null,
    follow_count   int default 0 not null,
    follower_count int default 0 not null,
    follow_db      varchar(128)  not null,
    follower_db    varchar(128)  not null,
    favorite_db    varchar(64)   not null,
    primary key (userid)
);
create table video
(
    id             int auto_increment,
    title          varchar(64)   not null,
    user_id        int           not null,
    play_url       varchar(128)  not null,
    cover_url      varchar(128)  not null,
    create_time    datetime      not null,
    favorite_count int default 0 not null,
    comment_count  int default 0 not null,
    comment_db     varchar(128)  not null,
    primary key (id)
);