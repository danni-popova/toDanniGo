-- Create database
CREATE DATABASE todo;

-- Connect to database
\c todo;

-- Create users table
CREATE TABLE registered_user
(
    id              serial PRIMARY KEY,
    email           varchar(50)  NOT NULL UNIQUE,
    password        varchar(100) NOT NULL,
    first_name      varchar(100) NOT NULL,
    last_name       varchar(100) NOT NULL,
    profile_picture varchar(500) default 'https://static01.nyt.com/images/2019/04/02/science/28SCI-ZIMMER1/28SCI-ZIMMER1-articleLarge.jpg?quality=75&auto=webp&disable=upscale',
    created_at      timestamp    default current_timestamp,
    updated_at      timestamp    default current_timestamp,
    deleted_at      timestamp
);

-- Create todos table
CREATE TABLE todo
(
    id          serial PRIMARY KEY,
    user_id     serial REFERENCES registered_user (id),
    title       varchar(100) NOT NULL,
    description text,
    created_at  timestamp default current_timestamp,
    deadline    timestamp,
    done        boolean   default FALSE,
    updated_at  timestamp default current_timestamp,
    deleted_at  timestamp
);