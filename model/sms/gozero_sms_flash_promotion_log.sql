create table sms_flash_promotion_log
(
    id             int auto_increment,
    member_id      int          null,
    product_id     bigint       null,
    member_phone   varchar(64)  null,
    product_name   varchar(100) null,
    subscribe_time datetime     null comment '会员订阅时间',
    send_time      datetime     null,
    UNIQUE KEY `id` (`id`) USING BTREE,
    PRIMARY KEY (`id`) USING BTREE
)
    comment '限时购通知记录';
