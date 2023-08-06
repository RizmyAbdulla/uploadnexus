create table if not exists objects
(
    id            char(36) primary key,
    bucket        char(36)     not null,
    name          text         not null,
    mime_type     varchar(255) not null,
    size          bigint       not null,
    upload_status varchar(255) not null,
    metadata      jsonb,
    created_at    bigint       not null,
    updated_at    bigint,
    unique (bucket, name)
);