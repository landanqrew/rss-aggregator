-- name: GetAllUserFeeds :many
SELECT
    f.id,
    f.created_at,
    f.updated_at,
    f.name,
    f.url,
    f.user_id,
    u.name as user_name
FROM feeds f
    JOIN users u
        ON f.user_id = u.id
ORDER BY user_name;
