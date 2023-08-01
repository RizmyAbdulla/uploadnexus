create table if not exists buckets
(
    id          char(36) primary key,
    name        varchar(255) not null,
    description text,
    is_public   boolean      not null,
    created_at  bigint       not null,
    updated_at  bigint,
    unique (name)
);
