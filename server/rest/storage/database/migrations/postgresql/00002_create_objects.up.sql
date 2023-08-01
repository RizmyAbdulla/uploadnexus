create table if not exists objects
(
    id               char(36) primary key,
    bucket_reference char(36)     not null,
    file_key         text         not null,
    file_name        text         not null,
    file_type        varchar(255) not null,
    file_size        bigint       not null,
    upload_status    varchar(255) not null,
    metadata         jsonb,
    created_at       bigint       not null,
    updated_at       bigint,
    unique (bucket_reference, file_key)
);