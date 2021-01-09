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
    created_at      timestamp default current_timestamp not null,
    updated_at      timestamp default current_timestamp not null,
    deleted_at      timestamp
);

create unique index account_data_id_uindex
    on account_data (id);

alter table account_data
    add constraint account_data_pk
        primary key (id);


-- Task table holding the task details

create table tasks
(
    id           serial                              not null,
    project      int                                 not null,
    title        varchar                             not null,
    status       varchar,
    description  varchar,
    done         boolean   default false,
    deadline     timestamp,
    created_at   timestamp default CURRENT_TIMESTAMP not null,
    updated_at  timestamp default CURRENT_TIMESTAMP not null,
    deleted_at   timestamp,
    completed_at timestamp,
    creator      int                                 not null
        constraint tasks_account_data_id_fk_2
            references account_data,
    assignee     int
        constraint tasks_account_data_id_fk
            references account_data
);

create unique index tasks_id_uindex
    on tasks (id);

alter table tasks
    add constraint tasks_pk
        primary key (id);
