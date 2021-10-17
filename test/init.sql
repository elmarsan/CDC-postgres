CREATE DATABASE test;

CREATE TABLE users
(
    id    SERIAL,
    name  TEXT,
    email VARCHAR UNIQUE
);
