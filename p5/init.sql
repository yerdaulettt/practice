CREATE TABLE IF NOT EXISTS users (
    id serial primary key,
    name varchar(30),
    email varchar(100)
);