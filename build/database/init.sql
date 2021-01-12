-- Create database
CREATE DATABASE todanni;

-- Connect to database
\c todanni;

-- Account data table holding the public account info
create table account_data
(
    id              serial                              not null,
    password        varchar                             not null,
    email           varchar                             not null,
    first_name      varchar                             not null,
    last_name       varchar                             not null,
    job_role        varchar,
    profile_picture varchar,
    created_at      timestamp default CURRENT_TIMESTAMP not null,
    updated_at      timestamp default CURRENT_TIMESTAMP not null,
    deleted_at      timestamp
);

create unique index account_data_id_uindex
    on account_data (id);

alter table account_data
    add constraint account_data_pk
        primary key (id);

-- Projects table holding the project details

create table projects
(
    id serial not null,
    title varchar not null,
    description varchar,
    logo varchar,
    created_at timestamp default CURRENT_TIMESTAMP,
    updated_at timestamp default CURRENT_TIMESTAMP,
    deleted_at timestamp,
    is_default boolean default false,
    creator int not null
        constraint projects_account_data_id_fk
            references account_data
);

create unique index projects_id_uindex
    on projects (id);

alter table projects
    add constraint projects_pk
        primary key (id);

-- Task table holding the task details

create table tasks
(
    id           serial                              not null
        constraint tasks_pk
            primary key,
    project      integer                             not null
        constraint tasks_projects_id_fk
            references projects,
    title        varchar                             not null,
    status       varchar,
    description  varchar,
    done         boolean   default false,
    deadline     timestamp,
    created_at   timestamp default CURRENT_TIMESTAMP not null,
    updated_at   timestamp default CURRENT_TIMESTAMP not null,
    deleted_at   timestamp,
    completed_at timestamp,
    creator      integer                             not null
        constraint tasks_account_data_id_fk_2
            references account_data,
    assignee     integer
        constraint tasks_account_data_id_fk
            references account_data
);

create unique index tasks_id_uindex
    on tasks (id);

alter table tasks
    add constraint tasks_pk
        primary key (id);
