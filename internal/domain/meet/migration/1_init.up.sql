-- +goose Up
create table participants
(
    id               serial primary key,
    firstname        varchar(50) unique      not null,
    lastname         varchar(50) unique      not null,
    contact_email    varchar(100) unique     not null,
    user_id          int                     null,
    description      varchar(150)            null,
    avatar_file_name varchar(150)            null,
    created_at       timestamp default now() not null,
    updated_at       timestamp default now() not null,
    deleted_at       timestamp               null,
    constraint participants_users_user_id
        foreign key (user_id)
            references users (id)
            on update cascade
            on delete no action
);

-- +goose Up
create table meets
(
    id           serial primary key,
    title        varchar(150)            not null,
    status       varchar(50)             not null,
    started_at   timestamp               not null,
    finished_at  timestamp               not null,
    description  text                    not null,
    link         text                    null,
    organizer_id int                     not null,
    created_by   int                     not null,
    created_at   timestamp default now() not null,
    constraint meets_users_organizer_id
        foreign key (organizer_id)
            references users (id)
            on update cascade
            on delete restrict,
    constraint meets_users_created_by
        foreign key (created_by)
            references users (id)
            on update cascade
            on delete restrict
);
