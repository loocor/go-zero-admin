create table pms_album
(
    id          bigint auto_increment,
    name        varchar(64)   null,
    cover_pic   varchar(1000) null,
    pic_count   int           null,
    sort        int           null,
    description varchar(1000) null,
    UNIQUE KEY `id` (`id`) USING BTREE,
    PRIMARY KEY (`id`) USING BTREE
)
    comment '相册表';
