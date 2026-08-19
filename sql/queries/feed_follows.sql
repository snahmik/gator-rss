-- name: CreateFeedFollow :one
WITH created_feed_follow AS (
    INSERT INTO feed_follows(id, created_at, updated_at, user_id, feed_id)
        VALUES ($1,$2,$3,
                (SELECT id FROM users WHERE users.name = $4),
                (SELECT id FROM feeds WHERE feeds.url = $5)
               )
           RETURNING *
)
SELECT cf.*, f.name AS feed_name,u.name AS user_name FROM created_feed_follow cf
INNER JOIN users u ON cf.user_id = u.id
INNER JOIN feeds f ON cf.feed_id = f.id;

-- name: GetFeedFollowsForUser :many
SELECT ff.*,f.name AS feed_name,u.name AS user_name  FROM feed_follows ff
INNER JOIN feeds f ON ff.feed_id = f.id
INNER JOIN users u ON ff.user_id = u.id
WHERE ff.user_id = (SELECT id FROM users WHERE users.name = $1);