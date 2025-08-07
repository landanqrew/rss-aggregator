-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING id, created_at, updated_at, name, url, user_id;


-- name: GetFeedByName :one
SELECT id, created_at, updated_at, name, url, user_id
FROM feeds
WHERE name = sqlc.arg(name)
LIMIT 1;

-- name: RemoveAllFeeds :exec
TRUNCATE TABLE feeds;

-- name: GetFeeds :many
SELECT id, created_at, updated_at, name, url, user_id FROM feeds;

-- name: GetAllFeeds :many
SELECT id, created_at, updated_at, name, url, user_id FROM feeds;
