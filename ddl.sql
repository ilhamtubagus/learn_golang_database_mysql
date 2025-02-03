create table if not exists learn_golang_database.comments
(
    id      int auto_increment
        primary key,
    email   varchar(100) not null,
    comment text         null
);

create table if not exists learn_golang_database.customer
(
    id         varchar(100)                         not null
        primary key,
    name       varchar(100)                         not null,
    email      varchar(100)                         null,
    balance    int        default 0                 null,
    rating     double     default 0                 null,
    created_at timestamp  default CURRENT_TIMESTAMP null,
    birth_date date                                 null,
    married    tinyint(1) default 0                 null
);

create table if not exists learn_golang_database.user
(
    username varchar(100) not null
        primary key,
    password varchar(100) not null
);

