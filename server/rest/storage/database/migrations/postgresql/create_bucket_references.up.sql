create table if not exists bucket_references
(
    id          char(36) primary key,
    application char(255) not null,
    name        varchar(255) not null,
    description text,
    is_public   boolean      not null,
    created_at  bigint       not null,
    updated_at  bigint,
    unique (application, name)
);
