-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table users (
    id     serial primary key not null,
    github varchar(39),
    name   varchar(64)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
drop table users;
