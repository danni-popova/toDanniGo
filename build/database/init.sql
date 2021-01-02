-- Create database
CREATE DATABASE todanni;

-- Connect to database
\c todanni;
--
-- -- Create users table
-- CREATE TABLE registered_user
-- (
--     id              serial PRIMARY KEY,
--     email           varchar(50)  NOT NULL UNIQUE,
--     password        varchar(100) NOT NULL,
--     first_name      varchar(100) NOT NULL,
--     last_name       varchar(100) NOT NULL,
--     profile_picture varchar(500) default 'https://static01.nyt.com/images/2019/04/02/science/28SCI-ZIMMER1/28SCI-ZIMMER1-articleLarge.jpg?quality=75&auto=webp&disable=upscale',
--     created_at      timestamp    default current_timestamp,
--     updated_at      timestamp    default current_timestamp,
--     deleted_at      timestamp
-- );
--
-- -- Create todos table
-- CREATE TABLE todo
-- (
--     id          serial PRIMARY KEY,
--     user_id     serial REFERENCES registered_user (id),
--     title       varchar(100) NOT NULL,
--     description text,
--     created_at  timestamp default current_timestamp,
--     deadline    timestamp,
--     done        boolean   default FALSE,
--     updated_at  timestamp default current_timestamp,
--     deleted_at  timestamp
-- );

-- create table accounts

create table accounts_data
(
    id              serial  not null,
    first_name      varchar not null,
    last_name       varchar not null,
    job_role        varchar,
    email           varchar not null,
    profile_picture varchar
);

create unique index accounts_id_uindex
    on accounts_data (id);

alter table accounts_data
    add constraint accounts_pk
        primary key (id);

create table tasks
(
    id           serial not null,
    description  varchar,
    done         boolean   default false,
    creator      int    not null,
    assignee     int    not null,
    created_at   timestamp default current_timestamp,
    updated_at   timestamp,
    completed_at timestamp,
    project      int    not null
);

create unique index tasks_id_uindex
    on tasks (id);

alter table tasks
    add constraint tasks_pk
        primary key (id);

create table projects
(
    id          serial  not null,
    title       varchar not null,
    description varchar,
    creator     int     not null,
    created_at  date    default current_timestamp,
    updated_at  timestamp,
    deleted_at  timestamp,
    is_default  boolean default false,
    logo        varchar
);

create unique index projects_id_uindex
    on projects (id);

alter table projects
    add constraint projects_pk
        primary key (id);

create table project_members
(
    project_id int,
    member_id  int,
    UNIQUE (project_id, member_id)
);

