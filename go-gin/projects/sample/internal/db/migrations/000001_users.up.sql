create table if not exists users (
    id int primary key,
    name varchar(50) not null,
    email varchar(100) not null unique
);