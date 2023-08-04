create table if not exists buckets
(
    id                 char(36) primary key,
    name               varchar(255) not null,
    description        text,
    allowed_mime_types text[]       not null,
    file_size_limit    bigint       not null,
    is_public          boolean      not null,
    created_at         bigint       not null,
    updated_at         bigint,
    unique (name)
);
