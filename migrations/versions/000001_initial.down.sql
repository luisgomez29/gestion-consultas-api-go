/*
 * Author: Luis Guillermo Gómez Galeano
 * Delete initial database schema
 */


--
-- DELETE INDEX
--


--
-- DELETE TABLES
--

DROP TABLE IF EXISTS auth_permission CASCADE;
DROP TABLE IF EXISTS auth_group CASCADE;
DROP TABLE IF EXISTS auth_group_permissions CASCADE;
DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS user_permissions CASCADE;
DROP TABLE IF EXISTS user_groups CASCADE;
DROP TABLE IF EXISTS appointments CASCADE;
DROP TABLE IF EXISTS rooms CASCADE;
DROP TABLE IF EXISTS messages CASCADE;

--
-- DELETE ENUMS
--

DROP TYPE enum_users_type;
DROP TYPE enum_users_identification_type;
DROP TYPE enum_messages_type;
