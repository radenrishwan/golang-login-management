create table users
(
    id         varchar(255) primary key,
    email      varchar(255) unique   not null,
    username   varchar(255) unique   not null,
    password   varchar(255)          not null,
    role       enum ('admin','user') not null,
    created_at varchar(255)          not null,
    updated_at varchar(255)          not null
);