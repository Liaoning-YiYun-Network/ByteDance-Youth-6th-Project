create schema skyline;
use skyline;
create table user
(
    userid         int auto_increment,
    username       varchar(16)  not null,
    passwd         varchar(128) not null,
    avatar         varchar(64)  not null,
    background     varchar(64)  not null,
    signature      varchar(64)  null,
    follow_count   int default 0 not null,
    follower_count int default 0 not null,
    primary key (userid)
);