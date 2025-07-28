-- +goose Up
Alter table feeds
    add last_fetched_at timestamp;

-- +goose Down
Alter table feeds
drop last_fetched_at;