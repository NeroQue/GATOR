-- +goose Up
CREATE TABLE users (
    id uuid primary key,
    created_at TIMESTAMP NOT NULL default now(),
    updated_at TIMESTAMP NOT NULL default now(),
    name TEXT UNIQUE NOT NULL
);

-- +goose Down
DROP TABLE users;