create table pms_product_verify_record
(
    id          bigint auto_increment,
    product_id  bigint       null,
    create_time datetime     null,
    verify_man  varchar(64)  null comment '审核人',
    status      int(1)       null,
    detail      varchar(255) null comment '反馈详情',
    UNIQUE KEY `id` (`id`) USING BTREE,
    PRIMARY KEY (`id`) USING BTREE
)
    comment '商品审核记录';

INSERT INTO gozero.pms_product_verify_record (id, product_id, create_time, verify_man, status, detail) VALUES (1, 1, '2018-04-27 16:36:41', 'test', 1, '验证通过');
INSERT INTO gozero.pms_product_verify_record (id, product_id, create_time, verify_man, status, detail) VALUES (2, 2, '2018-04-27 16:36:41', 'test', 1, '验证通过');