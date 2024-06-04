
-- +migrate Up
create table users (
    id serial primary key,
    email varchar(255) not null,
    password varchar(255) not null,
    name varchar(255) not null,
    role varchar(255) not null,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp
);
-- +migrate Down
drop table if exists users;