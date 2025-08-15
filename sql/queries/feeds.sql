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

-- name: GetFeedByURL :one
SELECT id, created_at, updated_at, name, url, user_id
FROM feeds
WHERE url = sqlc.arg(url)
LIMIT 1;

-- name: RemoveAllFeeds :exec
TRUNCATE TABLE feeds;

-- name: GetFeeds :many
SELECT id, created_at, updated_at, name, url, user_id FROM feeds;

-- name: GetAllFeeds :many
SELECT id, created_at, updated_at, name, url, user_id FROM feeds;

-- name: MarkFeedFetched :exec
UPDATE feeds SET last_fetched_at = sqlc.arg(last_fetched_at), updated_at = sqlc.arg(updated_at) WHERE id = sqlc.arg(id);

-- name: GetNextFeedToFetch :one
SELECT id, name, url, user_id, last_fetched_at, created_at, updated_at FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT 1;