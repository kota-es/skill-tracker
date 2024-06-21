
-- +migrate Up
create table skill_categories (
    id serial primary key,
    category_name varchar not null,
    created_at varchar not null,
    updated_at varchar not null
);

create table skills (
    id serial primary key,
    name varchar not null,
    skill_category_id int not null,
    created_at varchar not null,
    updated_at varchar not null,
    foreign key (skill_category_id) references skill_categories(id)
);

create table skill_levels (
    id serial primary key,
    skill_id int not null,
    level int not null,
    description text,
    created_at varchar not null,
    updated_at varchar not null,
    foreign key (skill_id) references skills(id)
);

create table user_skills (
    id serial primary key,
    user_id int not null,
    skill_id int not null,
    level int not null,
    interested boolean not null,
    created_at varchar not null,
    updated_at varchar not null,
    foreign key (user_id) references users(id),
    foreign key (skill_id) references skills(id)
);

-- +migrate Down
drop table if exists user_skills;
drop table if exists skill_levels;
drop table if exists skills;
drop table if exists skill_categories;
