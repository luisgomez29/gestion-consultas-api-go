/*
 * Author: Luis Guillermo Gómez Galeano
 *
 * Initial DML for the data base.
 */

--
-- INSERT INTO CONTENT_TYPE
--
INSERT INTO content_type(id, model)
VALUES (1, 'permission'),
       (2, 'group'),
       (3, 'users'),
       (4, 'appointments'),
       (5, 'rooms'),
       (6, 'messages');

SELECT setval('content_type_id_seq', 6);

--
-- TABLE AUTH_PERMISSION
--

INSERT INTO permissions(id, name, codename, content_type_id)
VALUES (1, 'Can add permission', 'add_permission', 1),
       (2, 'Can change permission', 'change_permission', 1),
       (3, 'Can delete permission', 'delete_permission', 1),
       (4, 'Can view permission', 'view_permission', 1),

       (5, 'Can add group', 'add_group', 2),
       (6, 'Can change group', 'change_group', 2),
       (7, 'Can delete group', 'delete_group', 2),
       (8, 'Can view group', 'view_group', 2),

       (9, 'Can add users', 'add_users', 3),
       (10, 'Can change users', 'change_users', 3),
       (11, 'Can delete users', 'delete_users', 3),
       (12, 'Can view users', 'view_users', 3),
       (13, 'Generate password reset link', 'password_reset', 3),

       (14, 'Can add appointments', 'add_appointments', 4),
       (15, 'Can change appointments', 'change_appointments', 4),
       (16, 'Can delete appointments', 'delete_appointments', 4),
       (17, 'Can view appointments', 'view_appointments', 4),
       (18, 'Can add my appointments', 'add_appointments_from_me', 4),
       (19, 'Can change my appointments', 'change_appointments_from_me', 4),
       (20, 'Can delete my appointments', 'delete_appointments_from_me', 4),
       (21, 'Can view my appointments', 'view_appointments_from_me', 4),

       (22, 'Can add rooms', 'add_rooms', 5),
       (23, 'Can change rooms', 'change_rooms', 5),
       (24, 'Can delete rooms', 'delete_rooms', 5),
       (25, 'Can view rooms', 'view_rooms', 5),
       (26, 'Can add my rooms', 'add_rooms_from_me', 5),
       (27, 'Can change my rooms', 'change_rooms_from_me', 5),
       (28, 'Can delete my rooms', 'delete_rooms_from_me', 5),
       (29, 'Can view my rooms', 'view_rooms_from_me', 5),

       (30, 'Can add messages', 'add_messages', 6),
       (31, 'Can change messages', 'change_messages', 6),
       (32, 'Can delete messages', 'delete_messages', 6),
       (33, 'Can view messages', 'view_messages', 6),
       (34, 'Can add my messages', 'add_messages_from_me', 6),
       (35, 'Can change my messages', 'change_messages_from_me', 6),
       (36, 'Can delete my messages', 'delete_messages_from_me', 6),
       (37, 'Can view my messages', 'view_messages_from_me', 6);

SELECT setval('permissions_id_seq', 37);

--
-- TABLE AUTH_GROUP
--

INSERT INTO groups(id, name)
VALUES (1, 'Users'),
       (2, 'Doctors');

SELECT setval('groups_id_seq', 2);

--
-- TABLE AUTH_GROUP_PERMISSIONS
--

INSERT INTO group_permissions(id, group_id, permission_id)
VALUES (1, 1, 18),
       (2, 1, 19),
       (3, 1, 21),
       (4, 1, 26),
       (5, 1, 29),
       (6, 1, 34),
       (7, 1, 35),
       (8, 1, 36),
       (9, 1, 37),

--     Doctors
       (10, 2, 12),
       (11, 2, 21),
       (12, 2, 26),
       (13, 2, 29),
       (14, 2, 34),
       (15, 2, 35),
       (16, 2, 36),
       (17, 2, 37);

SELECT setval('group_permissions_id_seq', 17);
