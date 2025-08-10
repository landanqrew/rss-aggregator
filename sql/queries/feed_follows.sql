-- name: CreateFeedFollow :one
WITH insert_feed_follow as (
    INSERT INTO feed_follows (id, created_at, updated_at, feed_id, user_id)
    VALUES (
        $1,
        $2,
        $3,
        $4,
        $5
    )
    RETURNING id, created_at, updated_at, feed_id, user_id
)
SELECT
    ff.id, 
    ff.created_at, 
    ff.updated_at, 
    ff.feed_id,
    f.name as feed_name,
    f.url as feed_url,
    ff.user_id,
    u.name as user_name
FROM insert_feed_follow ff
    JOIN feeds f ON f.id = ff.feed_id
    JOIN users u ON u.id = ff.user_id;

-- name: GetAllFeedFollows :many
SELECT 
ff.id, 
ff.created_at, 
ff.updated_at, 
ff.feed_id,
f.name as feed_name,
f.url as feed_url,
ff.user_id,
u.name as user_name
FROM feed_follows ff
    JOIN feeds f ON f.id = ff.feed_id
    JOIN users u ON u.id = ff.user_id;

-- name: GetFeedFollowsForUser :many
SELECT 
ff.id, 
ff.created_at, 
ff.updated_at, 
ff.feed_id,
f.name as feed_name,
f.url as feed_url,
ff.user_id,
u.name as user_name
FROM feed_follows ff
    JOIN feeds f ON f.id = ff.feed_id
    JOIN users u ON u.id = ff.user_id
WHERE u.name = sqlc.arg(user_name);

-- name: RemoveFeedFollow :exec
DELETE FROM feed_follows WHERE id = sqlc.arg(id);