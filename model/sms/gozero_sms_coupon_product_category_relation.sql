create table sms_coupon_product_category_relation
(
    id                    bigint auto_increment,
    coupon_id             bigint       null,
    product_category_id   bigint       null,
    product_category_name varchar(200) null comment '商品分类名称',
    parent_category_name  varchar(200) null comment '父分类名称',
    UNIQUE KEY `id` (`id`) USING BTREE,
    PRIMARY KEY (`id`) USING BTREE
)
    comment '优惠券和商品分类关系表';

INSERT INTO gozero.sms_coupon_product_category_relation (id, coupon_id, product_category_id, product_category_name, parent_category_name) VALUES (4, 19, 30, '手机配件', '手机数码');