-- migrate:up
create table url_shorten
(
    id         bigserial primary key,
    short_path varchar(255) not null,
    real_url   text         not null
)

-- migrate:down

