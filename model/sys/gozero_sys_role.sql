create table sys_role
(
    id               bigint auto_increment comment '编号',
    name             varchar(100)                        null comment '角色名称',
    remark           varchar(100)                        null comment '备注',
    create_by        varchar(50)                         null comment '创建人',
    create_time      timestamp default CURRENT_TIMESTAMP null comment '创建时间',
    last_update_by   varchar(50)                         null comment '更新人',
    last_update_time datetime                            null comment '更新时间',
    del_flag         tinyint   default 0                 null comment '是否删除  -1: 已删除  0: 正常',
    UNIQUE KEY `id` (`id`) USING BTREE,
    PRIMARY KEY (`id`) USING BTREE
)
    comment '角色管理';

INSERT INTO gozero.sys_role (id, name, remark, create_by, create_time, last_update_by, last_update_time, del_flag) VALUES (1, 'admin', '超级管理员', 'admin', '2019-01-19 11:11:11', 'admin', '2019-01-19 19:07:18', 0);
INSERT INTO gozero.sys_role (id, name, remark, create_by, create_time, last_update_by, last_update_time, del_flag) VALUES (2, 'mng', '项目经理', 'admin', '2019-01-19 11:11:11', 'admin', '2019-01-19 11:39:28', 0);
INSERT INTO gozero.sys_role (id, name, remark, create_by, create_time, last_update_by, last_update_time, del_flag) VALUES (3, 'dev', '开发人员', 'admin', '2019-01-19 11:11:11', 'admin', '2019-01-19 11:39:28', 0);
INSERT INTO gozero.sys_role (id, name, remark, create_by, create_time, last_update_by, last_update_time, del_flag) VALUES (4, 'test', '测试人员', 'admin', '2019-01-19 11:11:11', 'admin', '2019-01-19 11:11:11', 0);
