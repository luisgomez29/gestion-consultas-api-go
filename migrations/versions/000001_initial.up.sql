-- Table users
CREATE TABLE users
(
    id         SERIAL PRIMARY KEY,
    first_name VARCHAR(40)        NOT NULL,
    last_name  VARCHAR(40)        NOT NULL,
    username   VARCHAR(60) UNIQUE NOT NULL,
    email      VARCHAR(60),
    phone      VARCHAR(12)        NOT NULL
);