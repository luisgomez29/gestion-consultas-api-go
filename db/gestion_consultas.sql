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

--
-- TABLE AUTH_PERMISSION
--

INSERT INTO auth_permission(id, name, content_type, codename)
VALUES (1, 'Can add permission', 1, 'add_permission'),
       (2, 'Can change permission', 1, 'change_permission'),
       (3, 'Can delete permission', 1, 'delete_permission'),
       (4, 'Can view permission', 1, 'view_permission'),

       (5, 'Can add group', 2, 'add_group'),
       (6, 'Can change group', 2, 'change_group'),
       (7, 'Can delete group', 2, 'delete_group'),
       (8, 'Can view group', 2, 'view_group'),

       (9, 'Can add users', 3, 'add_users'),
       (10, 'Can change users', 3, 'change_users'),
       (11, 'Can delete users', 3, 'delete_users'),
       (12, 'Can view users', 3, 'view_users'),
       (13, 'Generate password reset link', 3, 'password_reset'),

       (14, 'Can add appointments', 4, 'add_appointments'),
       (15, 'Can change appointments', 4, 'change_appointments'),
       (16, 'Can delete appointments', 4, 'delete_appointments'),
       (17, 'Can view appointments', 4, 'view_appointments'),
       (18, 'Can add my appointments', 4, 'add_appointments_from_me'),
       (19, 'Can change my appointments', 4, 'change_appointments_from_me'),
       (20, 'Can delete my appointments', 4, 'delete_appointments_from_me'),
       (21, 'Can view my appointments', 4, 'view_appointments_from_me'),

       (22, 'Can add rooms', 5, 'add_rooms'),
       (23, 'Can change rooms', 5, 'change_rooms'),
       (24, 'Can delete rooms', 5, 'delete_rooms'),
       (25, 'Can view rooms', 5, 'view_rooms'),
       (26, 'Can add my rooms', 5, 'add_rooms_from_me'),
       (27, 'Can change my rooms', 5, 'change_rooms_from_me'),
       (28, 'Can delete my rooms', 5, 'delete_rooms_from_me'),
       (29, 'Can view my rooms', 5, 'view_rooms_from_me'),

       (30, 'Can add messages', 6, 'add_messages'),
       (31, 'Can change messages', 6, 'change_messages'),
       (32, 'Can delete messages', 6, 'delete_messages'),
       (33, 'Can view messages', 6, 'view_messages'),
       (34, 'Can add my messages', 6, 'add_messages_from_me'),
       (35, 'Can change my messages', 6, 'change_messages_from_me'),
       (36, 'Can delete my messages', 6, 'delete_messages_from_me'),
       (37, 'Can view my messages', 6, 'view_messages_from_me');

--
-- TABLE AUTH_GROUP
--

INSERT INTO auth_group(id, name)
VALUES (1, 'Users'),
       (2, 'Doctors');


--
-- TABLE AUTH_GROUP_PERMISSIONS
--

INSERT INTO auth_group_permissions(id, group_id, permission_id)
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
