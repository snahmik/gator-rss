-- name: AddFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES ($1,$2,$3,$4,$5,(
    SELECT id FROM users WHERE users.name = $6))
RETURNING *;

-- name: GetFeeds :many
SELECT f.name,url,u.name FROM feeds f
INNER JOIN public.users u ON f.user_id = u.id;