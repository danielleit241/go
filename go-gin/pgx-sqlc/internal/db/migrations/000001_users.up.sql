create extension if not exists "uuid-ossp";

create table if not exists users (
    id uuid primary key default uuid_generate_v4(),
    name varchar(50) not null,
    email varchar(100) not null unique,
    created_at_utc timestamptz not null default now()
);