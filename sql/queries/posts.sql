-- name: CreatePost :one
insert into
  posts (
    id,
    title,
    url,
    description,
    published_at,
    feed_id
  )
VALUES
  ($1, $2, $3, $4, $5, $6)
returning
  *;

-- name: GetPosts :many
select
  *
from
  posts
order by
  published_at desc
limit
  $1;
