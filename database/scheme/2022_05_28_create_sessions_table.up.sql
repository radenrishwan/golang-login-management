create table sessions
(
    id         varchar(255) primary key,
    user_id    varchar(255) unique not null,
    created_at varchar(255)        not null,
    expired_at varchar(255)        not null,
    foreign key (user_id) references users (id) on delete cascade
)