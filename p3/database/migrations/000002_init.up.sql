create table if not exists users (
    id serial primary key,
    name varchar(40) not null
);

insert into users (name) values ('Name 1'), ('Name 2');