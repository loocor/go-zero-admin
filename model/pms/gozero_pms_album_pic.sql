create table pms_album_pic
(
    id       bigint auto_increment,
    album_id bigint        null,
    pic      varchar(1000) null,
    UNIQUE KEY `id` (`id`) USING BTREE,
    PRIMARY KEY (`id`) USING BTREE
)
    comment '画册图片表';
