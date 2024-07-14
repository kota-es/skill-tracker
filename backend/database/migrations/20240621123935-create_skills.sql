
-- +migrate Up
create table skill_categories (
    id serial primary key,
    name varchar not null unique,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp
);

create table skills (
    id serial primary key,
    skill_category_id int not null,
    name varchar not null,
    description text,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp,
    unique (name, skill_category_id),
    foreign key (skill_category_id) references skill_categories(id)
);

create table skill_levels (
    id serial primary key,
    skill_id int not null,
    level int not null,
    explanation text,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp,
    unique (skill_id, level),
    foreign key (skill_id) references skills(id)
);

create table user_skills (
    id serial primary key,
    user_id int not null,
    skill_id int not null,
    level int not null,
    interested boolean not null,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp,
    unique (user_id, skill_id),
    foreign key (user_id) references users(id),
    foreign key (skill_id) references skills(id)
);

-- +migrate Down
drop table if exists user_skills;
drop table if exists skill_levels;
drop table if exists skills;
drop table if exists skill_categories;
