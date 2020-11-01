-- Create database
CREATE DATABASE todo;

-- Connect to database
\c todo;

-- Create users table
CREATE TABLE registered_user(
    user_id         serial          PRIMARY KEY,
    email           varchar(50)     NOT NULL UNIQUE,
    password        varchar(100)    NOT NULL,
    first_name      varchar(100)    NOT NULL,
    last_name       varchar(100)    NOT NULL,
    profile_picture varchar(500) default 'https://static01.nyt.com/images/2019/04/02/science/28SCI-ZIMMER1/28SCI-ZIMMER1-articleLarge.jpg?quality=75&auto=webp&disable=upscale',
    created_at      timestamp       default current_timestamp
);

-- Create todos table
CREATE TABLE todo(
    todo_id     serial PRIMARY KEY,
    user_id     serial REFERENCES registered_user(id),
    title       varchar(100) NOT NULL,
    description text,
    created_at  timestamp default current_timestamp,
    deadline    timestamp,
    done        boolean default FALSE
);

-- Create the table to store the priority of the todos
CREATE TABLE todo_order(
    user_id serial references registered_user(id),
    todos text []
);

-- Insert test user values
INSERT INTO registered_user(email, first_name, last_name, password, profile_picture)
VALUES ('test1@mail.com', 'Test', 'Test', 'password', 'default'),
       ('test2@mail.com', 'Test', 'Test', 'password', 'default'),
       ('test3@mail.com', 'Test', 'Test', 'password', 'default');

-- Insert test todo values
INSERT INTO todo(user_id, title, description, deadline)
VALUES (1, 'Test todo 1', 'Description 1', '2016-06-22 19:10:25-07'),
       (2, 'Test todo 2', 'Description 2', '2016-06-22 19:10:25-07'),
       (3, 'Test todo 3', 'Description 3', '2016-06-22 19:10:25-07');

-- Insert sorted todos
INSERT INTO todo_order
VALUES (1 , ARRAY['1', '2', '3']);