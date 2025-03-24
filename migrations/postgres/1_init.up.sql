-- +goose Up
create table users
(
    id            serial primary key,
    firstname     varchar(50) unique      not null,
    lastname      varchar(50) unique      not null,
    email_address varchar(100) unique     not null,
    description   varchar(150) null,
    created_at    timestamp default now() not null,
    updated_at    timestamp default now() not null,
    deleted_at    timestamp null
);

create table participants
(
    id               serial primary key,
    firstname        varchar(50) unique      not null,
    lastname         varchar(50) unique      not null,
    contact_email    varchar(100) unique     not null,
    avatar_file_name varchar(150) unique     not null,
    user_id          int null,
    description      varchar(150) null,
    created_at       timestamp default now() not null,
    updated_at       timestamp default now() not null,
    deleted_at       timestamp null,
    constraint participants_users_user_id
        foreign key (user_id)
            references users (id)
            on update on delete cascade,
);
