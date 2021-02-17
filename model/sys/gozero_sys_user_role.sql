create table sys_user_role
(
    id               bigint auto_increment comment '编号',
    user_id          bigint      null comment '用户ID',
    role_id          bigint      null comment '角色ID',
    create_by        varchar(50) null comment '创建人',
    create_time      datetime    null comment '创建时间',
    last_update_by   varchar(50) null comment '更新人',
    last_update_time datetime    null comment '更新时间',
    UNIQUE KEY `id` (`id`) USING BTREE,
    PRIMARY KEY (`id`) USING BTREE
)
    comment '用户角色';

INSERT INTO gozero.sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (1, 1, 1, null, null, null, null);
INSERT INTO gozero.sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (2, 2, 1, null, null, null, null);
INSERT INTO gozero.sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (3, 5, 3, null, null, null, null);
INSERT INTO gozero.sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (4, 6, 2, null, null, null, null);
INSERT INTO gozero.sys_user_role (id, user_id, role_id, create_by, create_time, last_update_by, last_update_time) VALUES (5, 4, 2, null, null, null, null);