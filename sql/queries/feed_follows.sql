-- name: CreateFeedFollow :one
with
  inserted_feed_follow as (
    insert into
      feed_follows (user_id, feed_id, created_at, updated_at)
    values
      ($1, $2, $3, $4)
    returning
      *
  )
select
  inserted_feed_follow.*,
  feeds.name as feed_name,
  users.name as user_name
from
  inserted_feed_follow
  inner join users on users.id = inserted_feed_follow.user_id
  inner join feeds on feeds.id = inserted_feed_follow.feed_id;
