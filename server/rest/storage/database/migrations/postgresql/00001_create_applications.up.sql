create table if not exists applications (
    id char(36) primary key,
    name varchar(255) not null,
    description text,
    created_at bigint not null,
    updated_at bigint,
    unique (name)
);