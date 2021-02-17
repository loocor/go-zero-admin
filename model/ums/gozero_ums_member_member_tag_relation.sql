create table ums_member_member_tag_relation
(
    id        bigint auto_increment,
    member_id bigint null,
    tag_id    bigint null,
    UNIQUE KEY `id` (`id`) USING BTREE,
    PRIMARY KEY (`id`) USING BTREE
)
    comment '用户和标签关系表';
