-- auto-generated definition
create table users
(
    id     bigint auto_increment comment 'id'
        primary key,
    name   varchar(50)   default ''  not null comment '名字',
    secret varchar(1024) default ''  not null comment '密码',
    age    int unsigned  default '0' not null comment '年龄',
    phone  varchar(50)   default ''  not null,
    constraint users_phone_uindex
        unique (phone)
)
    comment '用户表';

