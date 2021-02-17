create table ums_member_level
(
    id                     bigint auto_increment,
    name                   varchar(100)   null,
    growth_point           int            null,
    default_status         int(1)         null comment '是否为默认等级: 0->不是; 1->是',
    free_freight_point     decimal(10, 2) null comment '免运费标准',
    comment_growth_point   int            null comment '每次评价获取的成长值',
    privilege_free_freight int(1)         null comment '是否有免邮特权',
    privilege_sign_in      int(1)         null comment '是否有签到特权',
    privilege_comment      int(1)         null comment '是否有评论获奖励特权',
    privilege_promotion    int(1)         null comment '是否有专享活动特权',
    privilege_member_price int(1)         null comment '是否有会员价格特权',
    privilege_birthday     int(1)         null comment '是否有生日特权',
    note                   varchar(200)   null,
    UNIQUE KEY `id` (`id`) USING BTREE,
    PRIMARY KEY (`id`) USING BTREE
)
    comment '会员等级表';

INSERT INTO gozero.ums_member_level (id, name, growth_point, default_status, free_freight_point, comment_growth_point, privilege_free_freight, privilege_sign_in, privilege_comment, privilege_promotion, privilege_member_price, privilege_birthday, note) VALUES (1, '黄金会员', 1000, 0, 199.00, 5, 1, 1, 1, 1, 1, 1, null);
INSERT INTO gozero.ums_member_level (id, name, growth_point, default_status, free_freight_point, comment_growth_point, privilege_free_freight, privilege_sign_in, privilege_comment, privilege_promotion, privilege_member_price, privilege_birthday, note) VALUES (2, '白金会员', 5000, 0, 99.00, 10, 1, 1, 1, 1, 1, 1, null);
INSERT INTO gozero.ums_member_level (id, name, growth_point, default_status, free_freight_point, comment_growth_point, privilege_free_freight, privilege_sign_in, privilege_comment, privilege_promotion, privilege_member_price, privilege_birthday, note) VALUES (3, '钻石会员', 15000, 0, 69.00, 15, 1, 1, 1, 1, 1, 1, null);
INSERT INTO gozero.ums_member_level (id, name, growth_point, default_status, free_freight_point, comment_growth_point, privilege_free_freight, privilege_sign_in, privilege_comment, privilege_promotion, privilege_member_price, privilege_birthday, note) VALUES (4, '普通会员', 1, 1, 199.00, 20, 1, 1, 1, 1, 0, 0, null);