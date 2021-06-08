/*
 * Author: Luis Guillermo Gómez Galeano
 *
 * Delete initial database schema.
 */


--
-- DELETE INDEX
--


--
-- DELETE TABLES
--

DROP TABLE IF EXISTS content_type CASCADE;
DROP TABLE IF EXISTS permissions CASCADE;
DROP TABLE IF EXISTS groups CASCADE;
DROP TABLE IF EXISTS group_permissions CASCADE;
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
