-- +goose Up
-- +goose StatementBegin
create table users (
    id     serial primary key,
    github varchar(39) unique not null,
    name   varchar(64)
);

begin;
  update users
  set
    github = 'creasty',
    name = 'Yuki Iwanaga'
  where id = 1;

  insert into users (id, github, name)
  select 1, 'creasty', 'Yuki Iwanaga'
    where not exists (select 1 from users where id = 1);
commit;
-- +goose StatementEnd

-- +goose Down
drop table users;
