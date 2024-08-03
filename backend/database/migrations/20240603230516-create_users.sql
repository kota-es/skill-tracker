
-- +migrate Up
create table if not exists users (
    id serial primary key,
    email varchar not null unique,
    password varchar not null,
    lastname varchar not null,
    firstname varchar not null,
    lastname_kana varchar,
    firstname_kana varchar,
    role varchar not null,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp
);

create table user_profiles (
    id serial primary key,
    user_id int not null unique,
    notes text,
    desires text,
    dislikes text,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp,
    foreign key (user_id) references users (id)
);

-- +migrate Down
drop table if exists user_profiles;
drop table if exists users;