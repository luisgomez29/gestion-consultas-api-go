/*
 * Author: Luis Guillermo Gómez Galeano
 *
 * Create initial database schema.
 *
 * The created_at and updated_at fields have the default value CURRENT_TIMESTAMP
 */

--
-- TABLE CONTENT_TYPE
--

CREATE TABLE content_type
(
    id    SERIAL PRIMARY KEY,
    model VARCHAR(25) NOT NULL UNIQUE
);


--
-- TABLE AUTH_PERMISSION
--

CREATE TABLE auth_permission
(
    id           SERIAL PRIMARY KEY,
    name         VARCHAR(100) NOT NULL,
    content_type INTEGER      NOT NULL REFERENCES content_type (id) ON DELETE CASCADE,
    codename     VARCHAR(60)  NOT NULL UNIQUE
);

--
-- TABLE AUTH_GROUP
--

CREATE TABLE auth_group
(
    id   SERIAL PRIMARY KEY,
    name VARCHAR(40) NOT NULL UNIQUE
);


--
-- TABLE AUTH_GROUP_PERMISSIONS
--

CREATE TABLE auth_group_permissions
(
    id            SERIAL PRIMARY KEY,
    group_id      INTEGER NOT NULL REFERENCES auth_group (id) ON DELETE CASCADE,
    permission_id INTEGER NOT NULL REFERENCES auth_permission (id) ON DELETE CASCADE
);

--
--  TABLE USERS
--

-- Users type enum
CREATE TYPE enum_users_type AS ENUM ('ADMIN', 'DOC', 'USR');

-- Users identification type enum
CREATE TYPE enum_users_identification_type AS ENUM ('CC', 'CE');

-- Create users table
CREATE TABLE users
(
    id                    SERIAL PRIMARY KEY,
    role                  enum_users_type                NOT NULL,
    first_name            VARCHAR(40)                    NOT NULL,
    last_name             VARCHAR(40)                    NOT NULL,
    identification_type   enum_users_identification_type NOT NULL,
    identification_number VARCHAR(10)                    NOT NULL UNIQUE,
    username              VARCHAR(60)                    NOT NULL UNIQUE,
    email                 VARCHAR(60),
    password              VARCHAR(128)                   NOT NULL,
    phone                 VARCHAR(12)                    NOT NULL,
    picture               VARCHAR(100),
    city                  VARCHAR(40)                    NOT NULL,
    neighborhood          VARCHAR(40),
    address               VARCHAR(60),
    is_active             boolean                        NOT NULL,
    is_staff              boolean                        NOT NULL,
    is_superuser          boolean                        NOT NULL,
    last_login            TIMESTAMPTZ,
    created_at            TIMESTAMPTZ                    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at            TIMESTAMPTZ                    NOT NULL DEFAULT CURRENT_TIMESTAMP
);

--
-- TABLE USER_PERMISSIONS
--

CREATE TABLE user_permissions
(
    id            SERIAL PRIMARY KEY,
    user_id       INTEGER NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    permission_id INTEGER NOT NULL REFERENCES auth_permission (id) ON DELETE CASCADE UNIQUE
);

--
-- TABLE USER_GROUPS
--

CREATE TABLE user_groups
(
    id       SERIAL PRIMARY KEY,
    user_id  INTEGER NOT NULL REFERENCES users (id) ON DELETE CASCADE UNIQUE,
    group_id INTEGER NOT NULL REFERENCES auth_group (id) ON DELETE CASCADE UNIQUE
);

--
-- TABLE APPOINTMENTS
--

CREATE TABLE appointments
(
    id          SERIAL PRIMARY KEY,
    user_id     INTEGER     NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    doctor_id   INTEGER     NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    description VARCHAR(700),
    children    JSONB,
    aggressor   VARCHAR(500),
    audio       VARCHAR(255),
    start_time  DATE,
    end_time    DATE,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

--
-- TABLE ROOMS
--

CREATE TABLE rooms
(
    id            SERIAL PRIMARY KEY,
    name          VARCHAR(60) NOT NULL UNIQUE,
    user_owner    INTEGER     NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    user_receiver INTEGER     NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    created_at    TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

--
-- TABLE MESSAGES
--

-- Messages type enum
-- AD: audio
-- FL: file
-- IMG: image
-- TXT: text
-- VD: video
CREATE TYPE enum_messages_type AS ENUM ('AD', 'FL', 'IMG', 'TXT', 'VD');

CREATE TABLE messages
(
    id         SERIAL PRIMARY KEY,
    room_id    INTEGER            NOT NULL REFERENCES rooms (id) ON DELETE CASCADE,
    user_id    INTEGER            NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    type       enum_messages_type NOT NULL,
    content    TEXT               NOT NULL,
    created_at TIMESTAMPTZ        NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ        NOT NULL DEFAULT CURRENT_TIMESTAMP
);
