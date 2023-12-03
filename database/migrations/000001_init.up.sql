create schema if not exists users;

create table users.post
(
    id          bigserial primary key,
    description text not null,
    code        text not null,
    created_dt  timestamp default now(),
    unique (code)
);

create table users.stream
(
    id          bigserial primary key,
    description text not null,
    code        text not null,
    created_dt  timestamp default now(),
    unique (code)
);

create table if not exists users.users
(
    id          uuid primary key   default gen_random_uuid(),
    first_name  text      not null,
    second_name text      not null,
    surname     text,
    avatar      text,
    post_id     bigint references users.post (id),
    stream_id   bigint references users.stream (id),
    Is_active   bool               default true,
    created_dt  timestamp not null default now(),
    updated_dt  timestamp not null
);