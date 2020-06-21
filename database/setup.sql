-- -- Create database
-- CREATE DATABASE todo;

-- Create users table
CREATE TABLE registered_user(
    user_id     serial          PRIMARY KEY,
    email       varchar(50)     NOT NULL UNIQUE,
    first_name  varchar(100)    NOT NULL,
    last_name   varchar(100)    NOT NULL,
    created_at  timestamp       default current_timestamp
);

-- Create todos table
CREATE TABLE todo(
    todo_id     serial PRIMARY KEY,
    user_id     serial REFERENCES registered_user(user_id),
    title       varchar(100) NOT NULL,
    description text,
    created_at  timestamp default current_timestamp,
    deadline    timestamp,
    done        boolean default FALSE
);

-- Create the table to store the priority of the todos
CREATE TABLE todo_order(
    user_id serial references registered_user(user_id),
    todos text []
);

-- -- Insert test user values
-- INSERT INTO registered_user(email, first_name, last_name)
-- VALUES ('test1@mail.com', 'Test', 'Test'),
--        ('test2@mail.com', 'Test', 'Test'),
--        ('test3@mail.com', 'Test', 'Test');
--
-- -- Insert test todo values
-- INSERT INTO todo(user_id, title)
-- VALUES (1, 'Test todo 1'),
--        (1, 'Test todo 2'),
--        (1, 'Test todo 3');
--
-- -- Insert sorted todos
-- INSERT INTO todo_order
-- VALUES (1 , ARRAY['1', '2', '3']) ;