create table sys_user
(
    id               bigint auto_increment comment '编号',
    name             varchar(50)                         not null comment '用户名',
    nick_name        varchar(150)                        null comment '昵称',
    avatar           varchar(150)                        null comment '头像',
    password         varchar(100)                        null comment '密码',
    salt             varchar(40)                         null comment '加密盐',
    email            varchar(100)                        null comment '邮箱',
    mobile           varchar(100)                        null comment '手机号',
    status           tinyint                             null comment '状态  0: 禁用   1: 正常',
    dept_id          bigint                              null comment '机构ID',
    create_by        varchar(50)                         null comment '创建人',
    create_time      timestamp default CURRENT_TIMESTAMP null comment '创建时间',
    last_update_by   varchar(50)                         null comment '更新人',
    last_update_time datetime                            null comment '更新时间',
    del_flag         tinyint   default 0                 null comment '是否删除  -1: 已删除  0: 正常',
    CONSTRAINT name unique (name),
    UNIQUE KEY `id` (`id`) USING BTREE,
    PRIMARY KEY (`id`) USING BTREE
)
    comment '用户管理';

INSERT INTO gozero.sys_user (id, name, nick_name, avatar, password, salt, email, mobile, status, dept_id, create_by, create_time, last_update_by, last_update_time, del_flag) VALUES (1, 'admin', '超管', '', 'bd1718f058d8a02468134432b8656a86', 'YzcmCZNvbXocrsz9dm8e', 'admin@qq.com', '13612345678', 1, 4, 'admin', '2018-08-14 11:11:11', 'admin', '2018-08-14 11:11:11', 0);
INSERT INTO gozero.sys_user (id, name, nick_name, avatar, password, salt, email, mobile, status, dept_id, create_by, create_time, last_update_by, last_update_time, del_flag) VALUES (2, 'liubei', '刘备', '', 'fd80ebd493a655608dc893a9f897d845', 'YzcmCZNvbXocrsz9dm8e', 'test@qq.com', '13889700023', 1, 7, 'admin', '2018-09-23 19:43:00', 'admin', '2019-01-10 11:41:13', 0);