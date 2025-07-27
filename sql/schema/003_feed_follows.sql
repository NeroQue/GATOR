-- +goose Up
CREATE TABLE feed_follows (
    id uuid NOT NULL PRIMARY KEY,
    created_at timestamp NOT NULL DEFAULT now(),
    updated_at timestamp NOT NULL DEFAULT now(),
    user_id uuid NOT NULL REFERENCES users(id) on delete cascade,
    feed_id uuid NOT NULL REFERENCES feeds(id) on delete cascade,
    constraint feed_follows_unique UNIQUE (user_id, feed_id)
);

-- +goose Down
DROP TABLE feed_follows;