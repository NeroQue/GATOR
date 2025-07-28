-- name: CreatePost :one
INSERT INTO posts (id, created_at, updated_at, title, url, description, published_at, feed_id)
VALUES (
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
--

-- name: GetPostsForUser :many
SELECT
    posts.*,
    feed_follows.user_id as user_id
from posts
inner join feeds on posts.feed_id = feeds.id
inner join feed_follows on feeds.id = feed_follows.feed_id
where feed_follows.user_id = $1
order by posts.created_at desc
limit $2;
