-- Create database
CREATE DATABASE todanni;

-- Connect to database
\c todanni;

-- Account data table holding the public account info
create table account_data
(
    id serial not null,
    password varchar not null,
    email varchar not null,
    first_name varchar not null,
    last_name varchar not null,
    job_role varchar,
    profile_picture varchar,
    created_at timestamp default current_timestamp not null,
    updated_at timestamp default current_timestamp not null,
    deleted_at timestamp
);

create unique index account_data_id_uindex
    on account_data (id);

alter table account_data
    add constraint account_data_pk
        primary key (id);

