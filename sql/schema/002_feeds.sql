-- +goose Up
create table feeds (
  id uuid primary key,
  created_at timestamp not null default now(),
  updated_at timestamp not null default now(),
  name text not null,
  url text not null unique,
  user_id uuid not null references users (id) on delete cascade
);

-- +goose Down
drop table feeds;
