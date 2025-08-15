-- name: CreatePost :one
INSERT INTO posts (id, created_at, updated_at, title, url, description, published_at, feed_id)
VALUES(
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8
)
RETURNING *;

-- name: CheckPostExists :one
SELECT COUNT(*) FROM posts WHERE url = sqlc.arg(url);

-- name: GetPostsForUser :many
SELECT p.* FROM posts p
JOIN feed_follows ff ON p.feed_id = ff.feed_id
WHERE ff.user_id = sqlc.arg(user_id)
ORDER BY p.published_at DESC
LIMIT sqlc.arg(max_results);