create table sys_role_menu
(
    id               bigint auto_increment comment '编号',
    role_id          bigint                              null comment '角色ID',
    menu_id          bigint                              null comment '菜单ID',
    create_by        varchar(50)                         null comment '创建人',
    create_time      timestamp default CURRENT_TIMESTAMP null comment '创建时间',
    last_update_by   varchar(50)                         null comment '更新人',
    last_update_time datetime                            null comment '更新时间',
    UNIQUE KEY `id` (`id`) USING BTREE,
    PRIMARY KEY (`id`) USING BTREE
)
    comment '角色菜单';

INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (431, 8, 1, 'admin', '2018-09-23 19:55:08', null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (432, 8, 2, 'admin', '2018-09-23 19:55:08', null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (433, 8, 9, 'admin', '2018-09-23 19:55:08', null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (434, 8, 3, 'admin', '2018-09-23 19:55:08', null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (435, 8, 13, 'admin', '2018-09-23 19:55:08', null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (436, 8, 4, 'admin', '2018-09-23 19:55:08', null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (437, 8, 17, 'admin', '2018-09-23 19:55:08', null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (438, 8, 5, 'admin', '2018-09-23 19:55:08', null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (439, 8, 21, 'admin', '2018-09-23 19:55:08', null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (440, 8, 7, 'admin', '2018-09-23 19:55:08', null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (441, 8, 31, 'admin', '2018-09-23 19:55:08', null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (442, 8, 8, 'admin', '2018-09-23 19:55:08', null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (443, 8, 6, 'admin', '2018-09-23 19:55:08', null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (444, 8, 35, 'admin', '2018-09-23 19:55:08', null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (469, 2, 1, null, null, null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (470, 2, 2, null, null, null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (471, 2, 3, null, null, null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (472, 2, 4, null, null, null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (473, 2, 5, null, null, null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (474, 2, 6, null, null, null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (475, 2, 7, null, null, null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (476, 2, 8, null, null, null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (477, 2, 9, null, null, null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (478, 2, 10, null, null, null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (479, 2, 11, null, null, null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (480, 2, 12, null, null, null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (481, 2, 13, null, null, null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (482, 2, 14, null, null, null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (483, 2, 15, null, null, null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (484, 2, 16, null, null, null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (485, 2, 17, null, null, null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (486, 2, 18, null, null, null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (487, 2, 19, null, null, null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (488, 2, 20, null, null, null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (489, 2, 21, null, null, null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (490, 2, 22, null, null, null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (491, 2, 23, null, null, null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (492, 2, 24, null, null, null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (493, 2, 25, null, null, null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (494, 2, 26, null, null, null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (495, 2, 27, null, null, null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (496, 2, 28, null, null, null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (497, 2, 29, null, null, null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (498, 2, 30, null, null, null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (499, 2, 31, null, null, null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (500, 2, 32, null, null, null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (501, 2, 33, null, null, null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (502, 2, 34, null, null, null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (503, 2, 35, null, null, null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (504, 2, 36, null, null, null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (505, 2, 37, null, null, null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (506, 2, 43, null, null, null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (507, 2, 44, null, null, null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (508, 2, 45, null, null, null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (509, 2, 46, null, null, null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (510, 2, 47, null, null, null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (511, 2, 38, null, null, null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (512, 2, 39, null, null, null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (513, 2, 40, null, null, null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (514, 2, 41, null, null, null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (515, 2, 42, null, null, null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (516, 2, 48, null, null, null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (517, 2, 49, null, null, null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (518, 2, 50, null, null, null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (519, 2, 51, null, null, null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (520, 2, 52, null, null, null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (521, 2, 53, null, null, null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (522, 2, 54, null, null, null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (523, 2, 55, null, null, null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (524, 2, 56, null, null, null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (572, 3, 1, 'admin', '2019-01-22 14:45:28', null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (573, 3, 2, 'admin', '2019-01-22 14:45:28', null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (574, 3, 3, 'admin', '2019-01-22 14:45:28', null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (575, 3, 4, 'admin', '2019-01-22 14:45:28', null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (576, 3, 5, 'admin', '2019-01-22 14:45:28', null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (577, 3, 6, 'admin', '2019-01-22 14:45:28', null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (578, 3, 7, 'admin', '2019-01-22 14:45:28', null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (579, 3, 8, 'admin', '2019-01-22 14:45:28', null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (580, 3, 12, 'admin', '2019-01-22 14:45:28', null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (581, 3, 13, 'admin', '2019-01-22 14:45:28', null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (582, 3, 17, 'admin', '2019-01-22 14:45:28', null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (583, 3, 18, 'admin', '2019-01-22 14:45:28', null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (584, 3, 22, 'admin', '2019-01-22 14:45:28', null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (585, 3, 23, 'admin', '2019-01-22 14:45:28', null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (586, 3, 24, 'admin', '2019-01-22 14:45:28', null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (587, 3, 25, 'admin', '2019-01-22 14:45:28', null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (588, 3, 26, 'admin', '2019-01-22 14:45:28', null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (589, 3, 27, 'admin', '2019-01-22 14:45:28', null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (590, 3, 28, 'admin', '2019-01-22 14:45:28', null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (591, 3, 29, 'admin', '2019-01-22 14:45:28', null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (592, 3, 30, 'admin', '2019-01-22 14:45:28', null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (593, 3, 31, 'admin', '2019-01-22 14:45:28', null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (594, 3, 32, 'admin', '2019-01-22 14:45:28', null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (595, 3, 33, 'admin', '2019-01-22 14:45:28', null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (596, 3, 35, 'admin', '2019-01-22 14:45:28', null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (597, 3, 36, 'admin', '2019-01-22 14:45:28', null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (598, 3, 43, 'admin', '2019-01-22 14:45:28', null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (599, 3, 44, 'admin', '2019-01-22 14:45:28', null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (600, 3, 45, 'admin', '2019-01-22 14:45:28', null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (601, 3, 38, 'admin', '2019-01-22 14:45:28', null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (602, 3, 39, 'admin', '2019-01-22 14:45:28', null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (603, 3, 40, 'admin', '2019-01-22 14:45:28', null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (604, 3, 41, 'admin', '2019-01-22 14:45:28', null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (605, 3, 42, 'admin', '2019-01-22 14:45:28', null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (606, 3, 50, 'admin', '2019-01-22 14:45:28', null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (607, 3, 51, 'admin', '2019-01-22 14:45:28', null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (608, 4, 1, 'admin', '2019-01-22 14:46:44', null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (609, 4, 2, 'admin', '2019-01-22 14:46:44', null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (610, 4, 3, 'admin', '2019-01-22 14:46:44', null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (611, 4, 7, 'admin', '2019-01-22 14:46:44', null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (612, 4, 8, 'admin', '2019-01-22 14:46:44', null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (613, 4, 17, 'admin', '2019-01-22 14:46:44', null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (614, 4, 18, 'admin', '2019-01-22 14:46:44', null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (615, 4, 32, 'admin', '2019-01-22 14:46:44', null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (616, 4, 33, 'admin', '2019-01-22 14:46:44', null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (617, 4, 35, 'admin', '2019-01-22 14:46:44', null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (618, 4, 36, 'admin', '2019-01-22 14:46:44', null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (619, 4, 46, 'admin', '2019-01-22 14:46:44', null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (620, 4, 47, 'admin', '2019-01-22 14:46:44', null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (621, 4, 50, 'admin', '2019-01-22 14:46:44', null, null);
INSERT INTO gozero.sys_role_menu (id, role_id, menu_id, create_by, create_time, last_update_by, last_update_time) VALUES (622, 4, 51, 'admin', '2019-01-22 14:46:44', null, null);