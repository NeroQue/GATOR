-- +goose Up
CREATE TABLE feeds (
    id uuid primary key,
    created_at TIMESTAMP NOT NULL default now(),
    updated_at TIMESTAMP NOT NULL default now(),
    name TEXT UNIQUE NOT NULL,
    url TEXT NOT NULL UNIQUE,
    user_id uuid NOT NULL references users(id)
                   ON DELETE CASCADE
);

-- +goose Down
DROP TABLE feeds;