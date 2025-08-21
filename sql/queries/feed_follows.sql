-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
  INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
    )
  RETURNING *
)
SELECT
  inserted_feed_follow.*,
  feeds.name AS feed_name,
  users.name AS user_name
FROM 
  inserted_feed_follow
  JOIN users ON inserted_feed_follow.user_id = users.id
  JOIN feeds ON inserted_feed_follow.feed_id = feeds.id
;

-- name: GetFeedFollowsForUser :many
SELECT
  feed_follows.*,
  feeds.name AS feed_name,
  users.name AS user_name
FROM
  feed_follows
  JOIN users ON feed_follows.user_id = users.id
  JOIN feeds ON feed_follows.feed_id = feeds.id
WHERE
  feed_follows.user_id = $1
;

-- name: UnfollowFeed :exec
DELETE FROM feed_follows 
USING users, feeds
WHERE  
  feed_follows.user_id = users.id AND
  feed_follows.feed_id = feeds.ID AND
  users.name = $1 AND
  feeds.url = $2
;
