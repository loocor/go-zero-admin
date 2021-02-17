create table ums_member_product_category_relation
(
    id                  bigint auto_increment,
    member_id           bigint null,
    product_category_id bigint null,
    UNIQUE KEY `id` (`id`) USING BTREE,
    PRIMARY KEY (`id`) USING BTREE
)
    comment '会员与商品分类关系表 (用户喜欢的分类) ';
