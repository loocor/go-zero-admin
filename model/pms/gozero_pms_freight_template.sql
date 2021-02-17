create table pms_freight_template
(
    id              bigint auto_increment,
    name            varchar(64)    null,
    charge_type     int(1)         null comment '计费类型:0->按重量; 1->按件数',
    first_weight    decimal(10, 2) null comment '首重kg',
    first_fee       decimal(10, 2) null comment '首费 (元) ',
    continue_weight decimal(10, 2) null,
    continue_fee    decimal(10, 2) null,
    dest            varchar(255)   null comment '目的地 (省、市) ',
    UNIQUE KEY `id` (`id`) USING BTREE,
    PRIMARY KEY (`id`) USING BTREE
)
    comment '运费模版';
