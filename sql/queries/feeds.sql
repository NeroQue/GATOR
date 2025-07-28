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
RETURNING *;

-- name: GetAllFeeds :many
SELECT * FROM feeds;

-- name: GetFeedByUUID :one
Select * from feeds where id = $1;

-- name: GetFeedByURL :one
select * from feeds where url = $1;


-- name: MarkFeedFetched :exec
update feeds
set last_fetched_at = now(), updated_at = now()
where id = $1;

-- name: GetNextFeedToFetch :one
SELECT * FROM feeds
ORDER BY last_fetched_at NULLS FIRST
LIMIT 1;
